package story

import (
	"fmt"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
)

func deploymentStory(e Environment) (runtime.Scene, error) {
	return comparisonStory(e, false)
}

func deploymentRadarStory(e Environment) (runtime.Scene, error) {
	return comparisonStory(e, true)
}

func comparisonStory(e Environment, radar bool) (runtime.Scene, error) {
	data, err := fixture.DataWithVersions(e.Config.Scenario, fixture.Versions{Reference: e.Config.ReferenceID, Candidate: e.Config.CandidateID})
	if err != nil {
		return nil, err
	}
	if e.Config.State == "empty" {
		data.Windows = nil
	}
	t := e.Config.Palette()
	m, err := pattern.NewCohortComparison(pattern.CohortComparisonOptions{Title: "Deployment comparison", Scope: "method=blue/green / SYNTHETIC / " + e.Config.Scenario,
		Data: data, Theme: t, MinRequests: 1000, Sigma: e.Config.Sigma, Emit: e.Emit, MetricTones: fixture.StatusTones(), Radar: radar})
	if err != nil {
		return nil, err
	}
	if e.Config.Curve == "straight" {
		m.SetSmooth(false)
	}
	return m, nil
}

func radarStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		if e.Config.State == "empty" {
			return layout.Fit("No profile observations", size)
		}
		o := chart.RadarOptions{Axes: []string{"2xx", "3xx", "4xx", "5xx"}, Min: -4, Max: 4, Reference: 0,
			Rings: []float64{-3, -2, -1, 1, 2, 3}, Smooth: e.Config.Curve != "straight", SelectedAxis: -1}
		for c, values := range [][]float64{{0, 0, 0, 0}, {-.8, .1, .4, 1.2}, {-1.1, .2, .4, 1.8}} {
			s := chart.RadarSeries{Label: []string{"reference", "candidate", "host"}[c], Marker: []rune{'R', 'C', 'X'}[c]}
			for _, value := range values {
				p := chart.RadarValue{Value: value, Valid: true}
				if c < 2 {
					sd := 1.0
					if c == 1 {
						sd = .4
					}
					for k := e.Config.Sigma; k >= 1; k-- {
						p.Intervals = append(p.Intervals, chart.Interval{Low: value - float64(k)*sd, High: value + float64(k)*sd})
					}
				}
				s.Values = append(s.Values, p)
			}
			o.Series = append(o.Series, s)
		}
		return layout.Fit("R reference / C candidate / X host; normalized units\n"+chart.Radar(o, layout.Size{Width: size.Width, Height: size.Height - 2}, e.Config.Palette())+
			fmt.Sprintf("\n0=reference mean; outward=higher; +/- %ds whiskers", e.Config.Sigma), size)
	}), nil
}

func bandStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		if e.Config.State == "empty" {
			return "No interval estimates"
		}
		var intervals []chart.Interval
		for k := e.Config.Sigma; k >= 1; k-- {
			intervals = append(intervals, chart.Interval{Low: 50 - float64(k)*8, High: 50 + float64(k)*8})
		}
		selected := 78.0
		return fmt.Sprintf("Nested intervals / +/- %ds\n\n", e.Config.Sigma) + chart.IntervalBand(chart.BandOptions{Min: 0, Max: 100, Center: 50, Intervals: intervals, Selected: &selected}, size.Width, e.Config.Palette()) +
			"\n0                                              100\n\n| center / X selected observation\nOuter to inner: 3s / 2s / 1s\nIntervals are supplied by the caller; this atom makes no statistical claim."
	}), nil
}
