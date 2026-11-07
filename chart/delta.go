package chart

import (
	"fmt"
	"math"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// DeltaProperty cancels the shared reference/candidate amount and renders only
// their signed difference as a zero-centred interval and exact numeric value.
// Extent is supplied by the parent so peer properties share the same scale.
type DeltaProperty struct {
	Value, Extent float64
	Available     bool
	Style         lipgloss.Style
}

func (d DeltaProperty) View(width int, t theme.Theme) string {
	if !d.Available || !finite(d.Value) || !finite(d.Extent) || d.Extent <= 0 {
		return layout.Fit("n/a", layout.Size{Width: width, Height: 1})
	}
	label := fmt.Sprintf("%+.3fpp", d.Value)
	if d.Value != 0 && math.Abs(d.Value) < 0.0005 {
		label = fmt.Sprintf("%+.3gpp", d.Value)
	}
	labelWidth := max(10, len(label))
	barWidth := width - labelWidth - 1
	if barWidth%2 == 0 {
		barWidth--
	}
	if barWidth < 3 {
		return layout.Fit(d.Style.Render(label), layout.Size{Width: width, Height: 1})
	}
	style := d.Style
	bar := IntervalBand(BandOptions{Min: -d.Extent, Max: d.Extent, Center: 0, Intervals: []Interval{{Low: min(0, d.Value), High: max(0, d.Value)}}, Style: &style}, barWidth, t)
	return layout.Fit(bar+" "+d.Style.Render(fmt.Sprintf("%*s", labelWidth, label)), layout.Size{Width: width, Height: 1})
}

// SymmetricExtent rounds an observed magnitude up to a readable 1/2/5 domain.
func SymmetricExtent(magnitude float64) float64 {
	if !finite(magnitude) || magnitude <= 0 {
		return 1
	}
	power := math.Pow(10, math.Floor(math.Log10(magnitude)))
	for _, step := range []float64{1, 2, 5, 10} {
		if magnitude <= step*power {
			return step * power
		}
	}
	return magnitude
}
