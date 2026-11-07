package feedback

import (
	"fmt"
	"math"
	"strings"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// PercentBar composes a native progress bar with an exact percentage. Its domain
// is always 0–100%; small values may fall below the bar's terminal resolution.
func PercentBar(value float64, available bool, width int, t theme.Theme) string {
	label := "n/a"
	if available && !math.IsNaN(value) && !math.IsInf(value, 0) && value >= 0 && value <= 100 {
		label = fmt.Sprintf("%.3f%%", value)
	} else {
		return layout.Fit(strings.Repeat(" ", max(0, width-len(label)))+t.Subtle().Render(label), layout.Size{Width: width, Height: 1})
	}
	barWidth := width - 9
	if barWidth < 2 {
		return layout.Fit(label, layout.Size{Width: width, Height: 1})
	}
	bar := NewProgress(barWidth, t)
	bar.EmptyColor = t.Border
	return layout.Fit(bar.ViewAs(value/100)+" "+t.Text().Render(fmt.Sprintf("%8s", label)), layout.Size{Width: width, Height: 1})
}
