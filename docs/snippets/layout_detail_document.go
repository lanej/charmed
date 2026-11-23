package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutDetailDocument() string {
	palette := theme.New(theme.Dark)
	document := layout.DetailDocument{
		Header: layout.IdentityHeader{Kind: "Operation", ID: "route-0042"},
		Summary: layout.RecordSummary{
			Title: "Gateway route update", Status: text.Pill("Outcome unknown", text.Warning, palette, text.Normal),
			Next: &layout.Pair{Label: "Next", Value: "Read the remote route before changing traffic."},
		},
		Sections: []layout.DetailSection{
			{ID: "observation", Title: "Observation", Fields: []layout.Pair{
				{Label: "Source", Value: "Controller journal"}, {Label: "Observed", Value: "09:42 UTC"}, {Label: "Receipt", State: text.ValueMissing},
			}},
			{ID: "verification", Title: "Verification", Statuses: []layout.StatusItem{
				{Label: "Submission", Value: "Recorded"}, {Label: "Remote route", Value: "Not observed", Detail: "No independent read follows the request."},
			}},
			{ID: "history", Title: "Activity", FullWidth: true, Timeline: []layout.TimelineEvent{
				{When: "09:41 UTC", Actor: "Controller", Title: "Request recorded"},
				{When: "09:42 UTC", Actor: "Controller", Title: "Response timed out", Detail: "A lost response does not establish the remote result."},
			}},
		},
	}

	return document.View(100, palette)
}
