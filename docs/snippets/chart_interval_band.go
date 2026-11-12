package snippets

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/theme"
)

func ExampleChartIntervalBand() string {
	selected := 78.0
	band := chart.BandOptions{
		Min: 0, Max: 100, Center: 50, Selected: &selected,
		Intervals: []chart.Interval{
			{Low: 26, High: 74}, {Low: 34, High: 66}, {Low: 42, High: 58},
		},
	}

	return chart.IntervalBand(band, 64, theme.New(theme.Dark))
}
