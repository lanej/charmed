package snippets

import (
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartLine() string {
	palette := theme.New(theme.Dark)
	values := []float64{72, 80, 68, 54, 66, 92, 90, 74}
	plot := chart.NewLine(layout.Size{Width: 64, Height: 14}, 0, 7, 0, 100, palette)
	plot.DrawXYAxisAndLabel()

	var previous canvas.Float64Point

	for i, value := range values {
		point := canvas.Float64Point{X: float64(i), Y: value}
		if i > 0 {
			plot.DrawBrailleLine(previous, point)
		}

		previous = point
	}

	return plot.View()
}
