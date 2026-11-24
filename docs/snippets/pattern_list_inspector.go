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
		Title: "PR #42 · Route update", Status: text.Pill("Outcome uncertain", text.Warning, palette, text.Normal),
		Fields: []layout.Pair{{Label: "Observed", Value: "09:42 UTC"}},
		Next:   &layout.Pair{Label: "Next action", Value: "Reconcile the saved merge operation"},
	}
	document := layout.DetailDocument{
		Header: layout.IdentityHeader{ID: "DELIVERY INSPECTION · ILLUSTRATIVE"}, Summary: summary,
		ActionDetail: "The request was recorded, but the remote response was unavailable.",
		Sections: []layout.DetailSection{
			{ID: "identity", Title: "Identity", Fields: []layout.Pair{{Label: "Repository", Value: "example/service"}, {Label: "Owner", Value: "Worker 2"}}},
			{ID: "checks", Title: "Verification", Statuses: []layout.StatusItem{{Label: "Tests", Value: "Passed"}, {Label: "Merge", Value: "Unverified"}}},
		},
	}
	records := []pattern.InspectorRecord{
		{ID: "example/service#42", Heading: "Needs attention", Summary: &summary, Document: &document},
		{ID: "example/service#41", Heading: "Completed", Summary: &layout.RecordSummary{
			Title: "PR #41 · Request parsing", Status: text.Pill("Merged", text.Success, palette, text.Normal),
			Fields: []layout.Pair{{Label: "Observed", Value: "09:38 UTC"}},
		}},
	}

	inspector, err := pattern.NewListInspector(records, layout.Size{Width: 104, Height: 20}, palette)
	if err != nil {
		return "", err
	}

	inspector.ListTitle, inspector.DetailTitle = "Deliveries", ""
	inspector.Refresh()

	return layout.Column(inspector.View(), text.KeyHint("j/k", "select", palette)+"  "+text.KeyHint("enter", "inspect", palette)+"  "+text.KeyHint("esc", "back", palette)), nil
}
