package snippets

import (
	"github.com/NimbleMarkets/ntcharts/v2/barchart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartBars() string {
	palette := theme.New(theme.Dark)

	plot := chart.NewBars(layout.Size{Width: 48, Height: 12}, palette)
	for i, name := range []string{"API", "Events", "Jobs"} {
		plot.Push(barchart.BarData{Label: name, Values: []barchart.BarValue{
			{Name: "Successful", Value: float64(30 + i*10), Style: palette.SeriesStyle(0)},
			{Name: "Retried", Value: float64(12 - i*3), Style: palette.SeriesStyle(1)},
		}})
	}

	plot.Draw()

	return plot.View()
}
