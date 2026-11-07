package chart

import (
	"math"
	"sort"

	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

type Interval struct{ Low, High float64 }
type BandMarker struct {
	Value float64
	Glyph rune
	Style lipgloss.Style
}
type BandOptions struct {
	Min, Max float64
	// Intervals are nested from outermost to innermost. They are clipped to
	// the domain. Center is the mean; Selected is an optional observed value.
	Intervals []Interval
	Center    float64
	Selected  *float64
	Series    int
	Markers   []BandMarker
}

// IntervalBand is a one-row ntcharts canvas primitive. Density encodes nested
// intervals independently of color; a selected observation gets a distinct mark.
// Invalid geometry produces an explicit fallback instead of a misleading band.
func IntervalBand(o BandOptions, width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}
	valid := finite(o.Min) && finite(o.Max) && o.Max > o.Min && finite(o.Center)
	for _, i := range o.Intervals {
		valid = valid && finite(i.Low) && finite(i.High) && i.Low <= i.High
	}
	if !valid {
		return layout.Fit("Invalid interval", layout.Size{Width: width, Height: 1})
	}
	plot := NewCanvas(layout.Size{Width: width, Height: 1}, t)
	style := lipgloss.NewStyle().Foreground(t.SeriesColor(o.Series))
	for x := range width {
		plot.SetRuneWithStyle(canvas.Point{X: x}, '·', t.Subtle())
		if t.ASCII {
			plot.SetRuneWithStyle(canvas.Point{X: x}, '.', t.Subtle())
		}
	}
	project := func(v float64) int {
		return int(math.Round((max(o.Min, min(o.Max, v)) - o.Min) / (o.Max - o.Min) * float64(width-1)))
	}
	intervals := append([]Interval(nil), o.Intervals...)
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i].High-intervals[i].Low > intervals[j].High-intervals[j].Low })
	glyphs := []rune{'░', '▒', '▓'}
	if t.ASCII {
		glyphs = []rune{'.', '=', '#'}
	}
	for level, i := range intervals {
		if i.High < o.Min || i.Low > o.Max {
			continue
		}
		for x := project(i.Low); x <= project(i.High); x++ {
			plot.SetRuneWithStyle(canvas.Point{X: x}, glyphs[min(level+max(0, 3-len(intervals)), len(glyphs)-1)], style)
		}
	}
	if o.Center >= o.Min && o.Center <= o.Max {
		plot.SetRuneWithStyle(canvas.Point{X: project(o.Center)}, '|', style.Bold(true))
	}
	for _, marker := range o.Markers {
		if finite(marker.Value) && marker.Value >= o.Min && marker.Value <= o.Max {
			plot.SetRuneWithStyle(canvas.Point{X: project(marker.Value)}, marker.Glyph, marker.Style)
		}
	}
	if o.Selected != nil && finite(*o.Selected) && *o.Selected >= o.Min && *o.Selected <= o.Max {
		plot.SetRuneWithStyle(canvas.Point{X: project(*o.Selected)}, 'X', t.Text().Bold(true))
	}
	return layout.Fit(plot.View(), layout.Size{Width: width, Height: 1})
}
