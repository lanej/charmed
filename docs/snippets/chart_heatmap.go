package snippets

import (
	"math"

	"github.com/NimbleMarkets/ntcharts/v2/heatmap"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartHeatmap() string {
	plot := chart.NewHeatmap(layout.Size{Width: 48, Height: 12}, theme.New(theme.Dark),
		heatmap.WithValueRange(0, 100), heatmap.WithCellSize(1, 1))
	plot.SetXYRange(0, 12, 0, 8)
	plot.SetViewXYRange(0, 12, 0, 8)

	for x := range 12 {
		for y := range 8 {
			value := 50 + 45*math.Sin(float64(x+y)*.4)
			plot.Push(heatmap.NewHeatPoint(float64(x), float64(y), value))
		}
	}

	plot.Draw()

	return plot.View()
}
