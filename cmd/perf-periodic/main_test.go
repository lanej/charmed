package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/internal/periodic"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func TestHeadlessPeriodicProbe(t *testing.T) {
	previous := os.Args
	defer func() { os.Args = previous }()

	for _, mode := range []string{"apply"} {
		dir := t.TempDir()
		file := filepath.Join(dir, "report.json")
		os.Args = []string{"perf-periodic", "--headless", "--mode", mode, "--warmup", "50ms", "--duration", "300ms", "--interval", "50ms", "--report", file}

		if err := run(); err != nil {
			t.Fatal(err)
		}

		directory, err := os.OpenRoot(dir)
		if err != nil {
			t.Fatal(err)
		}

		data, err := directory.ReadFile("report.json")

		closeErr := directory.Close()
		if closeErr != nil {
			t.Fatal(closeErr)
		}

		if err != nil {
			t.Fatal(err)
		}

		var got report
		if err = json.Unmarshal(data, &got); err != nil {
			t.Fatal(err)
		}

		if !got.Headless || len(got.Samples) < 1 || got.ElapsedNS <= 0 {
			t.Fatalf("incomplete measurement: %+v", got)
		}

		sawFrame := false

		for _, entry := range got.Samples {
			if mode == "apply" && entry.Status != unchanged {
				t.Fatalf("unchanged snapshot scheduled a frame: %+v", entry)
			}

			if entry.Status == presented {
				sawFrame = true
			}
		}

		if mode == "replace" && !sawFrame {
			t.Fatal("full refresh did not exercise presentation")
		}

		if mode == "apply" && got.PresentedFrames != 0 {
			t.Fatal("unchanged updates presented frames")
		}
	}
}

func TestStoppedProbeIgnoresQueuedMessages(t *testing.T) {
	config := periodic.Config{Records: 100, Points: 60, Scenario: periodic.Metric}

	source, err := periodic.NewSource(config)
	if err != nil {
		t.Fatal(err)
	}

	dashboard, err := periodic.NewDashboard(config, source.Snapshot())
	if err != nil {
		t.Fatal(err)
	}

	root := &runtime.Root{Scene: dashboard, Size: layout.Size{Width: 120, Height: 40}}
	root.Init()
	measured := &app{root: root, dashboard: dashboard, source: source, interval: 5 * time.Second}
	measured.startMeasurement()
	measured.Update(stopMsg{})

	before := source.Snapshot()
	frames := root.FrameRevision()

	for _, msg := range []tea.Msg{tickMsg{}, warmMsg{}, tea.KeyPressMsg{Code: 'h', Text: "h"}, struct{}{}} {
		measured.Update(msg)
	}

	if len(measured.report.Samples) != 0 || root.FrameRevision() != frames || !reflect.DeepEqual(before, source.Snapshot()) {
		t.Fatal("queued activity changed a completed measurement")
	}
}

func TestReplaceProbeTracksCompletedNativeFrame(t *testing.T) {
	config := periodic.Config{Records: 100, Points: 60, Scenario: periodic.Unchanged, Legacy: true}

	source, err := periodic.NewSource(config)
	if err != nil {
		t.Fatal(err)
	}

	dashboard, err := periodic.NewDashboard(config, source.Snapshot())
	if err != nil {
		t.Fatal(err)
	}

	root := &runtime.Root{Scene: dashboard, Size: layout.Size{Width: 120, Height: 40}}
	root.Init()
	measured := &app{root: root, dashboard: dashboard, source: source, interval: 5 * time.Second}
	measured.startMeasurement()

	_, command := measured.Update(tickMsg{})
	if command == nil || !root.PresentationPending() {
		t.Fatal("replacement did not request presentation")
	}

	batch, ok := command().(tea.BatchMsg)
	if !ok || len(batch) != 2 {
		t.Fatal("frame and next observation commands were not forwarded")
	}

	measured.Update(batch[0]())

	if len(measured.report.Samples) != 1 || measured.report.Samples[0].Status != presented {
		t.Fatal("frame completion was not measured")
	}
}
