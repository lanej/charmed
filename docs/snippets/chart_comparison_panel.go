package snippets

import (
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartComparisonPanel() string {
	palette := theme.New(theme.Dark)
	start := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	series := []chart.TimeSeriesData{
		{ID: "api", Label: "API", Unit: "%"},
		{ID: "worker", Label: "Worker", Unit: "%"},
	}

	samples := [][]float64{{72, 80, 68, 54, 66, 92, 90, 74}, {60, 66, 62, 48, 56, 66, 82, 68}}
	for i, values := range samples {
		for minute, value := range values {
			series[i].Points = append(series[i].Points, timeserieslinechart.TimePoint{
				Time: start.Add(time.Duration(minute) * time.Minute), Value: value,
			})
		}
	}

	distribution := chart.DistributionComparison{
		Sigma: 3, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp",
		Series: []chart.DistributionSeries{
			{Label: "API", Mean: 72, StdDev: 6, Members: 64, HasSpread: true},
			{Label: "Worker", Mean: 62, StdDev: 8, Members: 32, HasSpread: true},
		},
	}
	history := chart.NewTimeSeries(series, layout.Size{Width: 60, Height: 12}, palette)
	panel := chart.ComparisonPanel{Title: "Response share", Distribution: distribution, History: history}

	return panel.View(layout.Size{Width: 64, Height: 20}, palette, false)
}
