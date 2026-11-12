package snippets

import (
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleChartCanvas() string {
	palette := theme.New(theme.Dark)
	canvasView := chart.NewCanvas(layout.Size{Width: 48, Height: 8}, palette)
	canvasView.SetStringWithStyle(canvas.Point{X: 2, Y: 2}, "CHARMED / ntcharts canvas", palette.Heading())
	canvasView.SetRuneWithStyle(canvas.Point{X: 2, Y: 4}, '*', palette.SeriesStyle(0))

	return canvasView.View()
}
