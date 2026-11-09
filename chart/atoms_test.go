package chart_test

import (
	"image/color"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/heatmap"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/streamlinechart"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/wavelinechart"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestSingleSeriesChartsUseDataColor(t *testing.T) {
	palette := theme.New(theme.Dark)
	size := layout.Size{Width: 40, Height: 10}
	spark := chart.NewSparkline([]float64{2, 4, 3, 5}, size, palette)
	stream := chart.NewStream(size, palette, streamlinechart.WithYRange(0, 10))
	wave := chart.NewWave(size, palette, wavelinechart.WithXYRange(0, 4, 0, 10))

	for i, v := range []float64{2, 4, 3, 5} {
		stream.Push(v)
		wave.Plot(canvas.Float64Point{X: float64(i), Y: v})
	}

	stream.DrawAll()
	wave.Draw()

	for name, view := range map[string]string{"spark": spark.View(), "stream": stream.View(), "wave": wave.View()} {
		found := false

		for _, line := range uv.NewStyledString(view).Lines(ansi.GraphemeWidth) {
			for _, cell := range line {
				if cell.Content != " " && cell.Style.Fg != nil && color.NRGBAModel.Convert(cell.Style.Fg) == color.NRGBAModel.Convert(palette.Series[0]) {
					found = true
				}
			}
		}

		if !found {
			t.Fatalf("%s data did not use its series color", name)
		}
	}

	line := chart.NewLine(size, 0, 4, 0, 10, palette)
	if line.Style.GetForeground() != palette.Series[0] {
		t.Fatal("native line default used interaction color")
	}
}

func TestHeatmapScaleRetainsNativeOverridesAndDoesNotAliasTheme(t *testing.T) {
	palette := theme.New(theme.Dark)
	size := layout.Size{Width: 10, Height: 6}
	plot := chart.NewHeatmap(size, palette)
	original := palette.Sequential[0]
	plot.ColorScale[0] = lipgloss.Color("#ff0000")

	if palette.Sequential[0] != original {
		t.Fatal("native scale mutation changed the theme")
	}

	custom := []color.Color{lipgloss.Color("#010101"), lipgloss.Color("#fefefe")}

	overridden := chart.NewHeatmap(size, palette, heatmap.WithColorScale(custom))
	if overridden.ColorScale[0] != custom[0] {
		t.Fatal("native heatmap option was overwritten")
	}

	fallback := chart.NewHeatmap(size, theme.Theme{})
	if len(fallback.ColorScale) != 2 {
		t.Fatal("partial custom theme lost neutral heatmap fallback")
	}
}
