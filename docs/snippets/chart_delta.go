package snippets

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/theme"
)

func ExampleChartDelta() string {
	palette := theme.New(theme.Dark)
	reference, candidate := 20.0, 20.18 // Percentage points, not fractions.
	property := chart.DeltaProperty{
		Value: candidate - reference, Extent: .5,
		Available: true, Style: palette.Text(),
	}

	return property.View(40, palette)
	// Reuse Extent across peer properties to keep their zero-centred scales comparable.
}
