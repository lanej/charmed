package chart

import (
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// DeltaProperty cancels the shared reference/candidate amount and renders only
// their signed difference as a zero-centred interval and compact numeric label.
// Extent is supplied by the parent so peer properties share the same scale.
type DeltaProperty struct {
	Value, Extent float64
	Available     bool
	Style         lipgloss.Style
	Tones         text.SignTones
	// Unit defaults to percentage points for existing callers.
	Unit     string
	Unitless bool
	// LabelWidth is a shared minimum numeric column width, measured in cells.
	LabelWidth int
}

func (d DeltaProperty) number() text.SignedNumber {
	unit := d.Unit
	if unit == "" && !d.Unitless {
		unit = "pp"
	}

	return text.SignedNumber{Value: d.Value, Available: d.Available && finite(d.Extent) && d.Extent > 0, Precision: 3, Unit: unit, Tones: d.Tones, Abbreviate: true}
}

// Label exposes the formatted value for shared cell allocation by a composition.
func (d DeltaProperty) Label() string { return d.number().Label() }

func (d DeltaProperty) View(width int, t theme.Theme) string {
	number := d.number()
	label := number.Label()

	labelWidth := max(d.LabelWidth, lipgloss.Width(label))
	if d.LabelWidth <= 0 {
		labelWidth = max(10, labelWidth)
	}

	barWidth := width - labelWidth - 1
	if barWidth%2 == 0 {
		barWidth--
	}
	if barWidth < 3 {
		return layout.Fit(number.View(t), layout.Size{Width: width, Height: 1})
	}

	if label == "n/a" {
		if d.LabelWidth <= 0 {
			return layout.Fit(label, layout.Size{Width: width, Height: 1})
		}

		return layout.Fit(strings.Repeat(" ", barWidth+1+max(0, labelWidth-lipgloss.Width(label)))+number.View(t), layout.Size{Width: width, Height: 1})
	}

	style := d.Style
	bar := IntervalBand(BandOptions{Min: -d.Extent, Max: d.Extent, Center: 0, Intervals: []Interval{{Low: min(0, d.Value), High: max(0, d.Value)}}, Style: &style}, barWidth, t)

	return layout.Fit(bar+" "+strings.Repeat(" ", max(0, labelWidth-lipgloss.Width(label)))+number.View(t), layout.Size{Width: width, Height: 1})
}

// Plot renders only the numeric track so a composition can place labels in shared columns.
func (d DeltaProperty) Plot(width int, t theme.Theme) string {
	if d.Label() == "n/a" {
		return layout.Fit("", layout.Size{Width: width, Height: 1})
	}
	style := d.Style

	return IntervalBand(BandOptions{Min: -d.Extent, Max: d.Extent, Center: 0, Intervals: []Interval{{Low: min(0, d.Value), High: max(0, d.Value)}}, Style: &style}, width, t)
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
