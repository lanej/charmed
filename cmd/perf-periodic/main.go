// Command perf-periodic measures bounded dashboards refreshed every five seconds.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	goruntime "runtime"
	"runtime/debug"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/internal/periodic"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

type (
	tickMsg struct{}
	warmMsg struct{}
	stopMsg struct{}
	status  string
)

const (
	presented status = "presented"
	unchanged status = "unchanged"
	cancelled status = "cancelled"
)

type counters struct {
	CPU          float64
	CPUAvailable bool
	Memory       goruntime.MemStats
}

func observe() counters {
	cpu, err := cpuSeconds()
	result := counters{CPU: cpu, CPUAvailable: err == nil}
	goruntime.ReadMemStats(&result.Memory)

	return result
}

type sample struct {
	AcceptedUnixNS int64            `json:"accepted_unix_ns"`
	Size           layout.Size      `json:"size"`
	Sequence       int              `json:"sequence"`
	Status         status           `json:"status"`
	Frame          uint64           `json:"frame"`
	ElapsedNS      int64            `json:"processing_ns"`
	CPUSeconds     float64          `json:"cpu_seconds"`
	CPUAvailable   bool             `json:"cpu_available"`
	AllocBytes     uint64           `json:"allocated_bytes"`
	Mallocs        uint64           `json:"allocations"`
	Changes        periodic.Changes `json:"changes"`
}

type report struct {
	InitialSize     layout.Size     `json:"initial_size"`
	FinalSize       layout.Size     `json:"final_size"`
	SchemaVersion   int             `json:"schema_version"`
	StartedUnixNS   int64           `json:"measurement_started_unix_ns"`
	EndedUnixNS     int64           `json:"measurement_ended_unix_ns"`
	Headless        bool            `json:"headless"`
	PresentedFrames uint64          `json:"presented_frames_in_window"`
	Config          periodic.Config `json:"config"`
	GoVersion       string          `json:"go_version"`
	OS, Arch        string
	Revision        string `json:"revision"`
	Dirty           bool   `json:"dirty"`
	TERM, ColorTERM string
	IntervalNS      int64    `json:"requested_interval_ns"`
	WarmupNS        int64    `json:"warmup_ns"`
	ElapsedNS       int64    `json:"measured_elapsed_ns"`
	CPUSeconds      float64  `json:"process_cpu_seconds"`
	CPUAvailable    bool     `json:"cpu_available"`
	OneCorePercent  float64  `json:"one_core_percent"`
	AllocBytes      uint64   `json:"allocated_bytes"`
	Mallocs         uint64   `json:"allocations"`
	GCCount         uint32   `json:"gc_count"`
	HeapStart       uint64   `json:"heap_start_after_gc_bytes"`
	HeapEnd         uint64   `json:"heap_end_after_gc_bytes"`
	Frames          uint64   `json:"frames"`
	Samples         []sample `json:"samples"`
}

type pending struct {
	start    time.Time
	before   counters
	measured bool
	sequence int
}

type app struct {
	ended                      time.Time
	end                        counters
	frameBaseline              uint64
	root                       *runtime.Root
	dashboard                  *periodic.Dashboard
	source                     *periodic.Source
	interval, warmup, duration time.Duration
	report                     report
	measuring                  bool
	started                    time.Time
	baseline                   counters
	pending                    *pending
	sequence                   int
	err                        error
}

func (a *app) Init() tea.Cmd {
	initial := a.root.Init()
	if a.warmup == 0 {
		a.startMeasurement()
	}

	return tea.Batch(initial, tea.Tick(a.interval, func(time.Time) tea.Msg { return tickMsg{} }), tea.Tick(a.warmup, func(time.Time) tea.Msg { return warmMsg{} }), tea.Tick(a.warmup+a.duration, func(time.Time) tea.Msg { return stopMsg{} }))
}

func (a *app) startMeasurement() {
	if a.measuring {
		return
	}

	goruntime.GC()

	a.baseline = observe()
	a.frameBaseline = a.root.FrameRevision()
	a.report.InitialSize = a.root.Size
	a.started = time.Now()
	a.measuring = true
}

func (a *app) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if !a.ended.IsZero() {
		return a, nil
	}

	switch message.(type) {
	case warmMsg:
		a.startMeasurement()
		return a, nil
	case stopMsg:
		a.finish(cancelled)
		a.finishMeasurement()
		a.err = errors.Join(a.err, a.root.Close(), a.root.Err(), a.dashboard.Err())

		return a, tea.Quit
	case tickMsg:
		a.sequence++
		if a.pending != nil {
			a.finish(cancelled)
		}

		a.pending = &pending{start: time.Now(), before: observe(), measured: a.measuring, sequence: a.sequence}
		before := a.root.FrameRevision()

		_, command := a.root.Update(periodic.SnapshotMsg{Snapshot: a.source.Next()})
		if a.root.FrameRevision() != before {
			a.finish(presented)
		} else if !a.root.PresentationPending() {
			a.finish(unchanged)
		}

		return a, tea.Batch(command, tea.Tick(a.interval, func(time.Time) tea.Msg { return tickMsg{} }))
	}

	before := a.root.FrameRevision()

	_, command := a.root.Update(message)
	if a.root.FrameRevision() != before {
		a.finish(presented)
	}

	if key, ok := message.(tea.KeyPressMsg); ok && key.String() == "ctrl+c" {
		a.finish(cancelled)
		a.finishMeasurement()
	}

	return a, command
}

func (a *app) finish(state status) {
	if a.pending == nil {
		return
	}

	current := a.pending
	a.pending = nil

	if !current.measured {
		return
	}

	after := observe()
	a.report.Samples = append(a.report.Samples, sample{AcceptedUnixNS: current.start.UnixNano(), Size: a.root.Size, Sequence: current.sequence, Status: state, Frame: a.root.FrameRevision(), ElapsedNS: time.Since(current.start).Nanoseconds(), CPUSeconds: after.CPU - current.before.CPU, CPUAvailable: after.CPUAvailable && current.before.CPUAvailable, AllocBytes: after.Memory.TotalAlloc - current.before.Memory.TotalAlloc, Mallocs: after.Memory.Mallocs - current.before.Memory.Mallocs, Changes: a.dashboard.LastChanges})
}

func (a *app) View() tea.View { return a.root.View() }

func (a *app) finishMeasurement() {
	if !a.ended.IsZero() {
		return
	}

	a.end = observe()
	a.ended = time.Now()
}

func (a *app) completedReport() report {
	result := a.report
	a.finishMeasurement()

	end := a.end
	if a.measuring {
		result.StartedUnixNS = a.started.UnixNano()
		result.EndedUnixNS = a.ended.UnixNano()
		result.ElapsedNS = a.ended.Sub(a.started).Nanoseconds()
		result.CPUAvailable = end.CPUAvailable && a.baseline.CPUAvailable
		result.CPUSeconds = end.CPU - a.baseline.CPU
		result.OneCorePercent = 100 * result.CPUSeconds / (float64(result.ElapsedNS) / float64(time.Second))
		result.AllocBytes = end.Memory.TotalAlloc - a.baseline.Memory.TotalAlloc
		result.Mallocs = end.Memory.Mallocs - a.baseline.Memory.Mallocs
		result.GCCount = end.Memory.NumGC - a.baseline.Memory.NumGC
		result.HeapStart = a.baseline.Memory.HeapAlloc
	}

	goruntime.GC()

	var memory goruntime.MemStats
	goruntime.ReadMemStats(&memory)
	result.HeapEnd = memory.HeapAlloc
	result.FinalSize = a.root.Size
	result.Frames = a.root.FrameRevision()
	result.PresentedFrames = result.Frames - a.frameBaseline

	return result
}

func run() (err error) {
	flags := flag.NewFlagSet("perf-periodic", flag.ContinueOnError)
	scenario := flags.String("scenario", "unchanged", "unchanged, metric, request, append, mixed or hidden")
	mode := flags.String("mode", "apply", "apply skips unchanged data; replace forces full refresh")
	records := flags.Int("records", 100, "retained records: 100 or 1000")
	points := flags.Int("points", 60, "points per chart: 60 or 300")
	interval := flags.Duration("interval", 5*time.Second, "requested source-update interval")
	duration := flags.Duration("duration", 2*time.Minute, "measurement duration after warmup")
	warmup := flags.Duration("warmup", 6*time.Second, "warmup excluded from summary counters")
	history := flags.Bool("history", false, "inspect oldest request during updates")
	headless := flags.Bool("headless", false, "run the native event loop without terminal rendering; lower-bound CPU only")

	output := flags.String("report", "", "write JSON to this file; default stderr")
	if err = flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	if *interval <= 0 || *duration <= 0 || *warmup < 0 {
		return errors.New("interval/duration must be positive and warmup nonnegative")
	}

	if *mode != "apply" && *mode != "replace" {
		return errors.New("mode must be apply or replace")
	}

	config := periodic.Config{Records: *records, Points: *points, Scenario: periodic.Scenario(*scenario), Legacy: *mode == "replace", History: *history}

	source, err := periodic.NewSource(config)
	if err != nil {
		return err
	}

	dashboard, err := periodic.NewDashboard(config, source.Snapshot())
	if err != nil {
		return err
	}

	root := &runtime.Root{Scene: dashboard, Size: layout.Size{Width: 120, Height: 40}, AltScreen: true, Background: nil}
	defer func() { err = errors.Join(err, root.Close()) }()

	measured := &app{root: root, dashboard: dashboard, source: source, interval: *interval, duration: *duration, warmup: *warmup, report: report{SchemaVersion: 1, Config: config, Headless: *headless, GoVersion: goruntime.Version(), OS: goruntime.GOOS, Arch: goruntime.GOARCH, TERM: os.Getenv("TERM"), ColorTERM: os.Getenv("COLORTERM"), IntervalNS: int64(*interval), WarmupNS: int64(*warmup)}}

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				measured.report.Revision = setting.Value
			case "vcs.modified":
				measured.report.Dirty = setting.Value == "true"
			}
		}
	}

	options := []tea.ProgramOption{}
	if *headless {
		options = append(options, tea.WithInput(nil), tea.WithOutput(io.Discard), tea.WithoutRenderer(), tea.WithWindowSize(120, 40))
	}

	if _, err = tea.NewProgram(measured, options...).Run(); err != nil {
		return err
	}

	if measured.err != nil {
		return measured.err
	}

	var writer io.Writer = os.Stderr

	if *output != "" {
		file, openErr := os.Create(*output)
		if openErr != nil {
			return openErr
		}
		defer func() { err = errors.Join(err, file.Close()) }()

		writer = file
	}

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	return encoder.Encode(measured.completedReport())
}

func main() {
	if err := run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}
}
