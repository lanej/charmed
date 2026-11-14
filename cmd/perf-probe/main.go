// Command perf-probe measures accepted updates through completed runtime frames.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/story"
)

type pendingSample struct {
	first, last time.Time
	count       int
}
type probe struct {
	root     *runtime.Root
	trace    *json.Encoder
	pending  [3]pendingSample
	traceErr error
}

type frameSample struct {
	Kind           string `json:"kind"`
	Count          int    `json:"count"`
	Frame          uint64 `json:"frame"`
	Status         string `json:"status"`
	Duration       int64  `json:"duration_ns"`
	LatestDuration int64  `json:"latest_duration_ns"`
}

func (m *probe) Init() tea.Cmd { return m.root.Init() }
func (m *probe) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	started := time.Now()
	before := m.root.FrameRevision()

	_, command := m.root.Update(message)
	if m.root.FrameRevision() != before {
		m.flush(time.Now(), "presented")

		if m.traceErr != nil {
			return m, tea.Quit
		}

		return m, command
	}

	index := 2

	switch event := message.(type) {
	case tea.KeyPressMsg:
		if event.String() == "ctrl+c" {
			m.flush(time.Now(), "cancelled")
			return m, command
		}

		index = 0
	case tea.WindowSizeMsg:
		index = 1
	}

	sample := &m.pending[index]
	if sample.count == 0 {
		sample.first = started
	}

	sample.last = started
	sample.count++

	if !m.root.PresentationPending() {
		m.flush(time.Now(), "unchanged")

		if m.traceErr != nil {
			return m, tea.Quit
		}
	}

	return m, command
}

func (m *probe) flush(now time.Time, status string) {
	kinds := [3]string{"key", "resize", "message"}

	for i, sample := range m.pending {
		if sample.count == 0 {
			continue
		}

		record := frameSample{Kind: kinds[i], Count: sample.count, Frame: m.root.FrameRevision(), Status: status}
		if status == "presented" {
			record.Duration = now.Sub(sample.first).Nanoseconds()
			record.LatestDuration = now.Sub(sample.last).Nanoseconds()
		}

		if err := m.trace.Encode(record); err != nil {
			m.traceErr = errors.Join(m.traceErr, err)
		}

		m.pending[i] = pendingSample{}
	}
}
func (m *probe) View() tea.View { return m.root.View() }

func run() (err error) {
	id := flag.String("story", "git/browser", "Native fixture")
	path := flag.String("trace", "", "Trace JSONL destination")
	live := flag.Bool("live", false, "Enable fixture animation")

	flag.Parse()

	if *path == "" {
		return errors.New("--trace is required")
	}

	file, err := os.Create(*path)
	if err != nil {
		return err
	}

	defer func() { err = errors.Join(err, file.Close()) }()

	definition, err := story.Find(*id)
	if err != nil {
		return err
	}

	config := definition.Defaults
	config.Live = *live

	model, err := story.New(*id, config, nil)
	if err != nil {
		return err
	}

	defer func() { err = errors.Join(err, model.Close()) }()

	measured := &probe{root: model, trace: json.NewEncoder(file)}
	_, programErr := tea.NewProgram(measured).Run()
	measured.flush(time.Now(), "cancelled")

	return errors.Join(programErr, measured.traceErr)
}

func main() {
	if err := run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(2)
		}

		os.Exit(1)
	}
}
