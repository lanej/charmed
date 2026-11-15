package pattern_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/cohort"
	regions "github.com/lanej/charmed/examples/cohorts/fixture"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func TestComparisonVariantsPreserveSelectionAndOptionalAttention(t *testing.T) {
	d, _ := fixture.Data("isolated-hosts")

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: d.Cohorts[0].ID, Theme: theme.New(theme.Dark), MinRequests: 1000, Variant: pattern.Bands, Boundaries: fixture.Boundaries()})
	if err != nil {
		t.Fatal(err)
	}
	if m.Variant() != pattern.Bands {
		t.Fatal("explicit band variant ignored")
	}
	m.Init()
	m.Resize(layout.Size{Width: 132, Height: 39})
	frame := ansi.Strip(m.Panels.View())
	if !strings.Contains(frame, "Disconnected bands") || !strings.Contains(frame, fixture.ReferenceVersionID) || !strings.Contains(frame, "ERROR") {
		t.Fatal("overview lost direct version identity or policy attention")
	}
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	m.SetMetric("2xx")
	m.SetWindow(10)
	m.SetSigma(1)
	selected := m.Selected()
	if err := m.SetVariant("unknown"); err == nil || m.Variant() != pattern.Bands {
		t.Fatal("invalid variant changed presentation")
	}
	m.SetVariant(pattern.Stacked)
	m.SetVariant(pattern.Bands)
	for _, size := range []layout.Size{{Width: 40, Height: 16}, {Width: 132, Height: 39}} {
		m.Resize(size)
		if m.Selected() != selected || m.MetricID() != "2xx" || m.WindowIndex() != 10 || m.Sigma() != 1 || m.Variant() != pattern.Bands {
			t.Fatal("view or resize lost shared state")
		}
	}
	if err := m.SetBoundaries(nil); err != nil {
		t.Fatal(err)
	}
	if strings.Contains(ansi.Strip(m.Panels.View()), "ERROR") {
		t.Fatal("removed attention left stale overview state")
	}
	if strings.Contains(ansi.Strip(m.Panels.View()), "History %:") {
		t.Fatal("overview retained a series legend")
	}
	d, _ = fixture.Data("insufficient-samples")
	if err := m.SetData(d); err != nil {
		t.Fatal(err)
	}
	frame = ansi.Strip(m.Panels.View())
	if !strings.Contains(frame, "n/a") || strings.Contains(frame, "NaN") || strings.Contains(frame, "Inf") {
		t.Fatal("sparse bands invented observations")
	}
}

func TestDefaultComparisonRetainsStackedPresentation(t *testing.T) {
	d, _ := fixture.Data("isolated-hosts")

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: d.Cohorts[0].ID, Theme: theme.New(theme.Dark), MinRequests: 1000})
	if err != nil {
		t.Fatal(err)
	}
	if m.Variant() != pattern.Stacked || strings.Contains(ansi.Strip(m.Panels.View()), "Disconnected bands") {
		t.Fatal("the band variant replaced the default stacked presentation")
	}
}

func TestBandMatrixKeepsComparisonColumnsAcrossMetricsFocusAndWindows(t *testing.T) {
	d := regions.Data()
	for i := range d.Windows[len(d.Windows)-1].Samples {
		sample := &d.Windows[len(d.Windows)-1].Samples[i]
		if sample.CohortID == "west" {
			sample.Measurements["energy"] = cohort.Measurement{Value: 100000000}
		}
	}

	d.Metrics[0].Unit = "µs"

	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, ReferenceID: "north", Theme: theme.New(theme.Mono), Variant: pattern.Bands})
	if err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 150, Height: 80})

	baseline := -1

	for _, window := range []int{0, len(d.Windows) - 1} {
		if err := m.SetWindow(window); err != nil {
			t.Fatal(err)
		}

		for _, metric := range d.Metrics {
			if err := m.SetMetric(metric.ID); err != nil {
				t.Fatal(err)
			}

			rows := 0

			for _, line := range strings.Split(ansi.Strip(m.Panels.View()), "\n") {
				if !strings.Contains(line, "vs North") {
					continue
				}

				index := strings.Index(line[1:], "|") + 1
				if index <= 0 {
					t.Fatal("comparison bar missing")
				}

				column := ansi.StringWidth(line[:index])
				if baseline < 0 {
					baseline = column
				} else if column != baseline {
					t.Fatalf("zero column moved: %d to %d for %s window %d", baseline, column, metric.ID, window)
				}

				rows++
			}

			if rows != 8 {
				t.Fatalf("got %d peer comparisons, want 8", rows)
			}
		}
	}
}

func TestCohortHeaderReflowsWithoutOverlappingPanels(t *testing.T) {
	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Title: "Regional measurements", Scope: "SYNTHETIC sensor observations", Data: regions.Data(), ReferenceID: "north", Theme: theme.New(theme.Dark), Variant: pattern.Bands})
	if err != nil {
		t.Fatal(err)
	}

	for _, width := range []int{40, 80, 150} {
		m.Resize(layout.Size{Width: width, Height: 60})

		lines := strings.Split(ansi.Strip(m.View()), "\n")
		if len(lines) != 60 {
			t.Fatalf("frame height %d at width %d", len(lines), width)
		}

		panel := -1

		for i, line := range lines {
			if ansi.StringWidth(line) != width {
				t.Fatalf("line width at %d: %q", width, line)
			}

			if strings.Contains(line, "Disconnected bands") {
				panel = i
				break
			}
		}

		if panel < 0 {
			t.Fatal("header hid the primary panel")
		}

		header := strings.Join(lines[:panel], "\n")
		for _, value := range []string{"Regional measurements", "SYNTHETIC sensor observations", "Reference North", "South", "West", "eligible"} {
			if !strings.Contains(header, value) {
				t.Fatalf("header lost %q at width %d: %s", value, width, header)
			}
		}
	}
}

func TestManyCohortsKeepPrimaryPanelReachableInShortTerminal(t *testing.T) {
	d := regions.Data()
	for i := 0; i < 9; i++ {
		group := d.Cohorts[0]
		group.ID = "extra-" + strings.Repeat("x", i+1)
		group.Label = group.ID
		d.Cohorts = append(d.Cohorts, group)
	}

	for _, variant := range []pattern.ComparisonVariant{pattern.Stacked, pattern.Bands} {
		m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Title: "Measurements", Scope: "Synthetic", Data: d, ReferenceID: "north", Theme: theme.New(theme.Dark), Variant: variant})
		if err != nil {
			t.Fatal(err)
		}

		m.Resize(layout.Size{Width: 40, Height: 16})

		frame := ansi.Strip(m.View())
		if !strings.Contains(frame, "12 cohorts") || !strings.Contains(frame, "temp") {
			t.Fatalf("header hid primary evidence for %s:\n%s", variant, frame)
		}
	}
}
