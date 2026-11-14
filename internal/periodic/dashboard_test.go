package periodic

import (
	"strings"
	"testing"

	"github.com/lanej/charmed/collection"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func TestDashboardSnapshotsMatchFullReplacement(t *testing.T) {
	for _, scenario := range []Scenario{Unchanged, Metric, Request, Append, Mixed, Hidden} {
		t.Run(string(scenario), func(t *testing.T) {
			config := Config{Records: 100, Points: 60, Scenario: scenario}

			source, err := NewSource(config)
			if err != nil {
				t.Fatal(err)
			}

			optimized, err := NewDashboard(config, source.Snapshot())
			if err != nil {
				t.Fatal(err)
			}

			config.Legacy = true

			legacy, err := NewDashboard(config, source.Snapshot())
			if err != nil {
				t.Fatal(err)
			}

			for range 3 {
				snapshot := source.Next()

				changes, applyErr := optimized.Apply(snapshot)
				if applyErr != nil {
					t.Fatal(applyErr)
				}

				if _, applyErr = legacy.Apply(snapshot); applyErr != nil {
					t.Fatal(applyErr)
				}

				if scenario == "metric" && (!changes.Charts[0] || changes.Charts[1] || changes.Inspector || changes.Health) {
					t.Fatalf("unrelated preparation: %+v", changes)
				}

				if scenario == "hidden" && changes.Inspector {
					t.Fatal("hidden inspector prepared")
				}

				if optimized.View() != legacy.View() {
					t.Fatalf("snapshot differs from replacement in %s", scenario)
				}

				if optimized.Err() != nil || legacy.Err() != nil {
					t.Fatalf("render failed: %v %v", optimized.Err(), legacy.Err())
				}
			}

			if scenario == "hidden" {
				key := tea.KeyPressMsg{Code: 'h', Text: "h"}
				optimized.Update(key)
				legacy.Update(key)

				if optimized.View() != legacy.View() {
					t.Fatal("reopened inspector missed latest observations")
				}
			}
		})
	}
}

func TestUnchangedAndHiddenUpdatesDoNotScheduleFrames(t *testing.T) {
	for _, scenario := range []Scenario{Unchanged, Hidden} {
		config := Config{Records: 100, Points: 60, Scenario: scenario}

		source, err := NewSource(config)
		if err != nil {
			t.Fatal(err)
		}

		dashboard, err := NewDashboard(config, source.Snapshot())
		if err != nil {
			t.Fatal(err)
		}

		root := &runtime.Root{Scene: dashboard, Size: layout.Size{Width: 120, Height: 40}}
		root.Init()

		before := root.FrameRevision()
		for range 5 {
			_, command := root.Update(SnapshotMsg{Snapshot: source.Next()})
			if command != nil || root.PresentationPending() || root.FrameRevision() != before {
				t.Fatalf("%s update scheduled unnecessary frame", scenario)
			}
		}
	}
}

func BenchmarkPeriodicSnapshot(b *testing.B) {
	for _, legacy := range []bool{true, false} {
		for _, scenario := range []Scenario{Unchanged, Metric, Request, Append, Mixed, Hidden} {
			label := "apply/"
			if legacy {
				label = "replace/"
			}

			b.Run(label+string(scenario), func(b *testing.B) {
				config := Config{Records: 1000, Points: 300, Scenario: scenario, Legacy: legacy}

				source, err := NewSource(config)
				if err != nil {
					b.Fatal(err)
				}

				dashboard, err := NewDashboard(config, source.Snapshot())
				if err != nil {
					b.Fatal(err)
				}

				root := &runtime.Root{Scene: dashboard, Size: layout.Size{Width: 120, Height: 40}}
				root.Init()

				snapshots := [2]Snapshot{source.Next(), source.Next()}
				i := 0

				b.ReportAllocs()

				for b.Loop() {
					root.Update(SnapshotMsg{Snapshot: snapshots[i%2]})

					if root.PresentationPending() {
						if err := root.PresentNow(); err != nil {
							b.Fatal(err)
						}
					}

					i++

					if root.Err() != nil || dashboard.Err() != nil {
						b.Fatalf("frame failed: %v %v", root.Err(), dashboard.Err())
					}
				}
			})
		}
	}
}

func TestDataUpdateDoesNotMaskInvalidGeometry(t *testing.T) {
	config := Config{Records: 100, Points: 60, Scenario: "unchanged"}

	source, err := NewSource(config)
	if err != nil {
		t.Fatal(err)
	}

	dashboard, err := NewDashboard(config, source.Snapshot())
	if err != nil {
		t.Fatal(err)
	}

	dashboard.Resize(layout.Size{Width: 20, Height: 10})
	dashboard.Update(SnapshotMsg{Snapshot: source.Next()})

	if dashboard.Err() == nil {
		t.Fatal("valid data masked invalid geometry")
	}

	dashboard.Resize(layout.Size{Width: 120, Height: 40})

	if dashboard.Err() != nil {
		t.Fatal(dashboard.Err())
	}
}

func TestDashboardRecoversFromErrorThroughRuntime(t *testing.T) {
	config := Config{Records: 100, Points: 60, Scenario: Unchanged}

	source, err := NewSource(config)
	if err != nil {
		t.Fatal(err)
	}

	dashboard, err := NewDashboard(config, source.Snapshot())
	if err != nil {
		t.Fatal(err)
	}

	root := &runtime.Root{Scene: dashboard, Size: layout.Size{Width: 120, Height: 40}}
	root.Init()

	invalid := source.Snapshot()
	invalid.Logs = []collection.LogRecord{{ID: "duplicate"}, {ID: "duplicate"}}

	_, command := root.Update(SnapshotMsg{Snapshot: invalid})
	if command == nil {
		t.Fatal("error did not schedule a frame")
	}

	root.Update(command())

	if !strings.Contains(root.View().Content, "duplicate") {
		t.Fatal("error frame missing")
	}

	_, command = root.Update(SnapshotMsg{Snapshot: source.Snapshot()})
	if command == nil {
		t.Fatal("valid unchanged data did not replace error frame")
	}

	root.Update(command())

	if dashboard.Err() != nil || strings.Contains(root.View().Content, "duplicate") {
		t.Fatal("error frame persisted after recovery")
	}
}

func TestDashboardForwardsNativeCommandResults(t *testing.T) {
	config := Config{Records: 100, Points: 60, Scenario: Unchanged}

	source, err := NewSource(config)
	if err != nil {
		t.Fatal(err)
	}

	dashboard, err := NewDashboard(config, source.Snapshot())
	if err != nil {
		t.Fatal(err)
	}

	dashboard.Update(collection.CommandSelectedMsg{ID: "help/toggle"})

	if !dashboard.inspector.HelpVisible() {
		t.Fatal("native command result was dropped")
	}

	frame := dashboard.View()
	if dashboard.Err() != nil || frame == "" {
		t.Fatal("command did not render")
	}
}

func TestResizeRecoversFromFailedRenderWithoutNewData(t *testing.T) {
	config := Config{Records: 100, Points: 60, Scenario: Unchanged}

	source, err := NewSource(config)
	if err != nil {
		t.Fatal(err)
	}

	dashboard, err := NewDashboard(config, source.Snapshot())
	if err != nil {
		t.Fatal(err)
	}

	root := &runtime.Root{Scene: dashboard, Size: layout.Size{Width: 120, Height: 40}}
	root.Init()

	if err = dashboard.screen.ArrangeCompiled(nil, root.Size); err == nil {
		t.Fatal("invalid arrangement unexpectedly succeeded")
	}

	if err = root.PresentNow(); err != nil {
		t.Fatal(err)
	}

	if dashboard.Err() == nil {
		t.Fatal("failed render was not observable")
	}

	for _, size := range []layout.Size{{Width: 40, Height: 40}, {Width: 120, Height: 40}} {
		_, command := root.Update(tea.WindowSizeMsg{Width: size.Width, Height: size.Height})
		if command == nil {
			t.Fatal("resize did not schedule recovery")
		}

		root.Update(command())

		if dashboard.Err() != nil {
			t.Fatal(dashboard.Err())
		}

		if strings.Contains(root.View().Content, "Periodic probe:") {
			t.Fatal("resize retained stale error frame")
		}
	}
}
