package snippets

import (
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleChartWaffleHeader() string {
	palette := theme.New(theme.Dark)
	fleet := chart.Waffle{Parts: []chart.WafflePart{
		{Label: "ready", Count: 92, Tone: text.Success},
		{Label: "pending", Count: 3, Tone: text.Info},
		{Label: "failed", Count: 5, Tone: text.Danger},
	}}

	return fleet.StatusHeader("PRODUCTION", layout.Size{Width: 64, Height: 2}, palette)
}
