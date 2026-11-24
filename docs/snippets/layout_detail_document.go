package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutDetailDocument() string {
	palette := theme.New(theme.Dark)
	document := layout.DetailDocument{
		Header: layout.IdentityHeader{ID: "DELIVERY INSPECTION · ILLUSTRATIVE"},
		Summary: layout.RecordSummary{
			Title: "PR #42 · Route update", Status: text.Pill("Outcome uncertain", text.Warning, palette, text.Normal),
			Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}},
			Next:   &layout.Pair{Label: "Next action", Value: "Reconcile the saved merge operation"},
		},
		ActionDetail: "The request was recorded, but the remote response was unavailable.",
		Sections: []layout.DetailSection{
			{ID: "identity", Title: "Identity", Fields: []layout.Pair{
				{Label: "Repository", Value: "example/service"}, {Label: "Revision", Value: "a81f39d"}, {Label: "Owner", Value: "Worker 2"},
			}},
			{ID: "checks", Title: "Verification", Statuses: []layout.StatusItem{
				{Label: "Tests", Value: "Passed"}, {Label: "Review", Value: "Complete"}, {Label: "Merge", Value: "Unverified"},
			}},
			{ID: "feedback", Title: "Feedback delivery", FullWidth: true, Stages: []layout.StatusItem{
				{Label: "Transport", Value: "Delivered"}, {Label: "Receipt", Value: "Unconfirmed"}, {Label: "Disposition", Value: "Open"},
			}},
			{ID: "activity", Title: "Activity", FullWidth: true, Timeline: []layout.TimelineEvent{
				{When: "09:41 UTC", Title: "Merge requested", Detail: "Request persisted"},
				{When: "09:42 UTC", Title: "Remote response unavailable", Detail: "External outcome unknown"},
			}},
		},
	}

	return document.View(100, palette)
}
