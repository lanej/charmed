package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutRecordSummary() string {
	palette := theme.New(theme.Dark)
	summary := layout.RecordSummary{
		Title: "PR #42 · Route update", Status: text.Pill("Outcome uncertain", text.Warning, palette, text.Normal),
		Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}},
		Next:   &layout.Pair{Label: "Next action", Value: "Reconcile the saved merge operation"},
	}

	return summary.View(72, palette, true)
}
