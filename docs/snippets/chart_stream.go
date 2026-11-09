package snippets

import (
	"math"

	"github.com/NimbleMarkets/ntcharts/v2/linechart/streamlinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartStream() string {
	plot := chart.NewStream(layout.Size{Width: 48, Height: 12}, theme.New(theme.Dark),
		streamlinechart.WithYRange(0, 100))
	for i := range 48 {
		plot.Push(50 + 30*math.Sin(float64(i)*.2))
	}

	plot.DrawAll()

	return plot.View()
}
