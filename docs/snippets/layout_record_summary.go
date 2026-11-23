package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutRecordSummary() string {
	palette := theme.New(theme.Dark)
	summary := layout.RecordSummary{
		Title: "Gateway route update", Subtitle: "Production gateway",
		Status: text.Pill("Outcome unknown", text.Warning, palette, text.Normal),
		Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}, {Label: "Source", Value: "Controller journal"}},
		Next:   &layout.Pair{Label: "Next", Value: "Inspect the saved route request before changing traffic."},
	}

	return summary.View(72, palette, true)
}
