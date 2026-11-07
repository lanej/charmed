package story

import (
	"fmt"

	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/examples/deployment/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
)

func deploymentStory(e Environment) (runtime.Scene, error) {
	return comparisonStory(e, false)
}

func deploymentBandsStory(e Environment) (runtime.Scene, error) {
	return comparisonStory(e, true)
}

func comparisonStory(e Environment, overview bool) (runtime.Scene, error) {
	data, err := fixture.DataWithVersions(e.Config.Scenario, fixture.Versions{Reference: e.Config.ReferenceID, Candidate: e.Config.CandidateID})
	if err != nil {
		return nil, err
	}
	if e.Config.State == "empty" {
		data.Windows = nil
	}
	t := e.Config.Palette()
	var boundaries map[string]cohort.ComparisonPolicy
	if e.Config.Boundaries {
		boundaries = fixture.Boundaries()
	}
	return pattern.NewCohortComparison(pattern.CohortComparisonOptions{Title: "Deployment comparison", Scope: "method=blue/green / SYNTHETIC / " + e.Config.Scenario,
		Data: data, Theme: t, MinRequests: 1000, Sigma: e.Config.Sigma, Emit: e.Emit, MetricTones: fixture.StatusTones(), Boundaries: boundaries, BandOverview: overview})
}

func bandMatrixStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		matrix := chart.BandMatrix{Title: fmt.Sprintf("Host mean +/- %ds", e.Config.Sigma)}
		if e.Config.State != "empty" {
			for _, metric := range []struct {
				label    string
				mean, sd float64
			}{{"Success", 96, .4}, {"Server error", .2, .03}} {
				selected := metric.mean + 5*metric.sd
				d := chart.DistributionComparison{Sigma: e.Config.Sigma, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp",
					Series: []chart.DistributionSeries{
						{Label: "a13c9f2", Mean: metric.mean, StdDev: metric.sd, Members: 64, HasSpread: true},
						{Label: "f4d7e81", Mean: metric.mean + metric.sd, StdDev: metric.sd * 1.5, Members: 32, HasSpread: true, Selected: &selected},
					}}
				matrix.Rows = append(matrix.Rows, chart.BandRow{Title: metric.label, Distribution: d})
			}
		}
		return matrix.View(size, e.Config.Palette(), false)
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
