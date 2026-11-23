package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExamplePatternListInspector() (string, error) {
	palette := theme.New(theme.Dark)
	summary := layout.RecordSummary{
		Title: "Gateway route update", Status: text.Pill("Outcome unknown", text.Warning, palette, text.Normal),
		Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}},
		Next:   &layout.Pair{Label: "Next", Value: "Inspect the saved route request."},
	}
	document := layout.DetailDocument{
		Header: layout.IdentityHeader{Kind: "Operation", ID: "route-0042"}, Summary: summary,
		Sections: []layout.DetailSection{
			{ID: "source", Title: "Observation", Fields: []layout.Pair{{Label: "Source", Value: "Controller journal"}, {Label: "Receipt", State: text.ValueMissing}}},
			{ID: "checks", Title: "Verification", Statuses: []layout.StatusItem{{Label: "Submission", Value: "Recorded"}, {Label: "Remote route", Value: "Not observed"}}},
			{ID: "history", Title: "Activity", FullWidth: true, Timeline: []layout.TimelineEvent{
				{When: "09:42 UTC", Actor: "Controller", Title: "Response timed out", Detail: "Submission is recorded; the external result remains unknown."},
			}},
		},
	}
	records := []pattern.InspectorRecord{
		{ID: "route-0042", Heading: "Network", Summary: &summary, Document: &document},
		{ID: "pool-0007", Heading: "Compute", Summary: &layout.RecordSummary{
			Title: "Worker capacity", Status: text.Pill("Outcome confirmed", text.Success, palette, text.Normal),
			Fields: []layout.Pair{{Label: "Capacity", Value: "12 workers observed"}},
			Next:   &layout.Pair{Label: "Next", Value: "Check queue latency."},
		}},
	}

	inspector, err := pattern.NewListInspector(records, layout.Size{Width: 104, Height: 20}, palette)
	if err != nil {
		return "", err
	}

	inspector.ListTitle, inspector.DetailTitle = "Operations", "Evidence"
	inspector.Refresh()

	return layout.Column(inspector.View(), text.KeyHint("j/k", "select", palette)+"  "+text.KeyHint("enter", "inspect", palette)+"  "+text.KeyHint("esc", "back", palette)), nil
}
