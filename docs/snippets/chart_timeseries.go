package snippets

import (
	"time"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartTimeseries() string {
	palette := theme.New(theme.Dark)
	start := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	series := []chart.TimeSeriesData{
		{ID: "api", Label: "API", Unit: "req/min"},
		{ID: "worker", Label: "Worker", Unit: "req/min"},
	}

	samples := [][]float64{{72, 80, 68, 54, 66, 92, 90, 74}, {60, 66, 62, 48, 56, 66, 82, 68}}
	for i, values := range samples {
		for minute, value := range values {
			series[i].Points = append(series[i].Points, timeserieslinechart.TimePoint{
				Time: start.Add(time.Duration(minute) * time.Minute), Value: value,
			})
		}
	}

	plot := chart.NewTimeSeries(series, layout.Size{Width: 64, Height: 14}, palette)

	return plot.View()
}
