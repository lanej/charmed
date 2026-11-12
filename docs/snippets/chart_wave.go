package snippets

import (
	"math"

	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/wavelinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartWave() string {
	plot := chart.NewWave(layout.Size{Width: 48, Height: 12}, theme.New(theme.Dark),
		wavelinechart.WithXYRange(0, 31, 0, 100))
	for x := range 32 {
		plot.PlotDataSet("observed", canvas.Float64Point{
			X: float64(x), Y: 50 + 30*math.Sin(float64(x)*.2),
		})
	}

	plot.DrawAll()

	return plot.View()
}
