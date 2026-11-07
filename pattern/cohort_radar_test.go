package pattern

import (
	"math"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestRadarUsesReferenceHostSpreadAndStableWindowDomain(t *testing.T) {
	d, _ := fixture.Data("isolated-hosts")
	m, err := NewCohortComparison(CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark), MinRequests: 1000, Radar: true})
	if err != nil {
		t.Fatal(err)
	}
	o, span, missing := m.radarOptions()
	if len(missing) != 0 || len(o.Series) != 3 {
		t.Fatal("complete profiles lost observations")
	}
	for i, summaries := range m.summaries {
		want := (summaries[1].Mean - summaries[0].Mean) / summaries[0].StdDev
		if o.Series[0].Values[i].Value != 0 || math.Abs(o.Series[1].Values[i].Value-want) > 1e-12 {
			t.Fatal("profile used fleet weighting or candidate normalization")
		}
	}
	selected := m.Selected()
	m.SetSigma(1)
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	after, afterSpan, _ := m.radarOptions()
	if span != afterSpan || o.Min != after.Min || o.Max != after.Max || len(after.Series[1].Values[3].Intervals) != 1 {
		t.Fatal("selection or visible sigma limit rescaled the window")
	}
	if selected == m.Selected() {
		t.Fatal("host navigation did not update the profile")
	}
	m.SetWindow(10)
	m.SetMetric("3xx")
	m.SetSmooth(false)
	selected = m.Selected()
	m.SetRadar(false)
	m.SetRadar(true)
	for _, size := range []layout.Size{{Width: 40, Height: 16}, {Width: 132, Height: 48}} {
		m.Resize(size)
		if m.Selected() != selected || m.WindowIndex() != 10 || m.MetricID() != "3xx" || m.Sigma() != 1 || m.Smooth() || !m.RadarEnabled() {
			t.Fatal("view switch or resize lost shared selection state")
		}
	}
}

func TestRadarOmitsZeroOrUnavailableReferenceSpread(t *testing.T) {
	d, _ := fixture.Data("isolated-hosts")
	for wi := range d.Windows {
		for si := range d.Windows[wi].Samples {
			s := &d.Windows[wi].Samples[si]
			s.Counts["3xx"] = 0
		}
	}
	m, err := NewCohortComparison(CohortComparisonOptions{Data: d, Theme: theme.New(theme.Dark), MinRequests: 1000, Radar: true})
	if err != nil {
		t.Fatal(err)
	}
	o, _, missing := m.radarOptions()
	if len(missing) != 1 || missing[0] != "3xx" {
		t.Fatal("zero reference spread was not reported")
	}
	for _, s := range o.Series {
		if s.Values[1].Valid {
			t.Fatal("unavailable normalization became a finite profile value")
		}
	}
	d, _ = fixture.Data("insufficient-samples")
	if err := m.SetData(d); err != nil {
		t.Fatal(err)
	}
	m.Resize(layout.Size{Width: 132, Height: 48})
	frame := ansi.Strip(m.View())
	if !strings.Contains(frame, "Reference spread unavailable") || strings.Contains(frame, "NaN") || strings.Contains(frame, "Inf") {
		t.Fatal("sparse data manufactured a normalized profile")
	}
}
