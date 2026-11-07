package pattern_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func TestComparisonVariantsPreserveSelectionAndOptionalAttention(t *testing.T) {
	d, _ := fixture.Data("isolated-hosts")
	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark), MinRequests: 1000, Variant: pattern.Bands, Boundaries: fixture.Boundaries()})
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
	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark), MinRequests: 1000})
	if err != nil {
		t.Fatal(err)
	}
	if m.Variant() != pattern.Stacked || strings.Contains(ansi.Strip(m.Panels.View()), "Disconnected bands") {
		t.Fatal("the band variant replaced the default stacked presentation")
	}
}
