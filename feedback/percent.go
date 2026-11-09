package feedback

import (
	"fmt"
	"math"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// PercentBar composes a native progress bar with a percentage label that preserves
// tiny nonzero values. Its domain is always 0–100%; small values may fall below the bar's terminal resolution.
func PercentBar(value float64, available bool, width int, t theme.Theme) string {
	label := "n/a"
	if available && !math.IsNaN(value) && !math.IsInf(value, 0) && value >= 0 && value <= 100 {
		label = percentLabel(value)
	} else {
		return layout.Fit(strings.Repeat(" ", max(0, width-len(label)))+t.Subtle().Render(label), layout.Size{Width: width, Height: 1})
	}

	const minBarWidth = 2

	labelWidth := max(ansi.StringWidth(percentLabel(100)), ansi.StringWidth(label))
	gap := strings.Repeat(" ", theme.SpaceSmall)

	barWidth := width - labelWidth - ansi.StringWidth(gap)
	if barWidth < minBarWidth {
		return layout.Fit(ansi.Truncate(label, max(0, width), t.Glyph(theme.GlyphLoading)), layout.Size{Width: width, Height: 1})
	}
	bar := NewProgress(barWidth, t)
	bar.EmptyColor = t.Border

	return layout.Fit(bar.ViewAs(value/100)+gap+t.Text().Render(fmt.Sprintf("%*s", labelWidth, label)), layout.Size{Width: width, Height: 1})
}

func percentLabel(value float64) string {
	const decimalPlaces = 3
	return strings.TrimPrefix((text.SignedNumber{Value: value, Available: true, Precision: decimalPlaces, Unit: "%"}).Label(), "+")
}
