package chart

import (
	"image/color"
	"slices"

	"charm.land/lipgloss/v2"

	"github.com/NimbleMarkets/ntcharts/v2/barchart"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/canvas/runes"
	"github.com/NimbleMarkets/ntcharts/v2/heatmap"
	"github.com/NimbleMarkets/ntcharts/v2/linechart"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/streamlinechart"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/wavelinechart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// Native constructors return the concrete upstream model. Further options apply
// after design-system defaults, allowing targeted customization.
func NewLine(size layout.Size, minX, maxX, minY, maxY float64, t theme.Theme, opts ...linechart.Option) linechart.Model {
	defaults := []linechart.Option{linechart.WithStyles(t.Subtle(), t.Subtle(), t.SeriesStyle(0))}
	return linechart.New(max(20, size.Width), max(6, size.Height), minX, maxX, minY, maxY, append(defaults, opts...)...)
}

func NewBars(size layout.Size, t theme.Theme, opts ...barchart.Option) barchart.Model {
	defaults := []barchart.Option{barchart.WithStyles(t.Subtle(), t.Subtle())}
	return barchart.New(max(1, size.Width), max(1, size.Height), append(defaults, opts...)...)
}

func NewHeatmap(size layout.Size, t theme.Theme, opts ...heatmap.Option) heatmap.Model {
	defaults := []heatmap.Option{}

	scale := t.Sequential
	if len(scale) == 0 {
		// WHY: Custom themes without a scale use a neutral ramp, never categorical colors.
		scale = []color.Color{color.Black, color.White}
	}

	defaults = append(defaults, heatmap.WithColorScale(slices.Clone(scale)))
	return heatmap.New(max(1, size.Width), max(1, size.Height), append(defaults, opts...)...)
}

func NewStream(size layout.Size, t theme.Theme, opts ...streamlinechart.Option) streamlinechart.Model {
	defaults := []streamlinechart.Option{streamlinechart.WithStyles(runes.ThinLineStyle, t.SeriesStyle(0)), streamlinechart.WithAxesStyles(t.Subtle(), t.Subtle())}
	return streamlinechart.New(max(20, size.Width), max(6, size.Height), append(defaults, opts...)...)
}

func NewWave(size layout.Size, t theme.Theme, opts ...wavelinechart.Option) wavelinechart.Model {
	defaults := []wavelinechart.Option{wavelinechart.WithStyles(runes.ThinLineStyle, t.SeriesStyle(0)), wavelinechart.WithAxesStyles(t.Subtle(), t.Subtle())}
	return wavelinechart.New(max(20, size.Width), max(6, size.Height), append(defaults, opts...)...)
}

func NewCanvas(size layout.Size, t theme.Theme) canvas.Model {
	m := canvas.New(max(1, size.Width), max(1, size.Height))
	m.Style = lipgloss.NewStyle().Foreground(t.Foreground)
	return m
}
