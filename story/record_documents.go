package story

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/examples/inspection/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func recordDefinitions() []Definition {
	summaryDefaults := DefaultConfig()
	summaryDefaults.Columns, summaryDefaults.Rows = 72, 10
	documentDefaults := DefaultConfig()
	documentDefaults.Columns, documentDefaults.Rows = 112, 34

	return []Definition{
		{ID: "layout/record-summary", Title: "Molecules/Layout/Record summary", Level: "Molecules", Family: "layout", Description: "Title, state, metadata and a next step form a reusable record summary. Styled text wraps without dropping evidence, and selection keeps the same geometry.", Source: "layout/record_summary.go", Symbol: "RecordSummary", Composes: []string{"layout/text-row", "layout/key-value"}, Defaults: summaryDefaults, Factory: recordSummaryStory},
		{ID: "layout/detail-document", Title: "Organisms/Layout/Detail document", Level: "Organisms", Family: "layout", Description: "Lead with a next action and its evidence, then group identity, verification, delivery stages and event history. Sections pair when space permits, stack when narrow and preserve their complete content.", Source: "layout/detail_document.go", Symbol: "DetailDocument", Composes: []string{"layout/identity-header", "layout/record-summary", "layout/key-value"}, Defaults: documentDefaults, Factory: detailDocumentStory},
	}
}

func recordSummaryStory(e Environment) (runtime.Scene, error) {
	summary := fixture.Records("ready", e.Config.Palette())[0].Summary

	return static(func(size layout.Size) string { return summary.View(size.Width, e.Config.Palette(), true) }), nil
}

func detailDocumentStory(e Environment) (runtime.Scene, error) {
	document := fixture.Records("ready", e.Config.Palette())[0].Document
	viewport := collection.NewViewport("", 1, 1)
	pane := collection.NewScrollPane(&viewport, e.Config.Size(), e.Config.Palette())
	resize := func(size layout.Size) {
		pane.Resize(size)
		pane.SetContent(document.View(viewport.Width(), e.Config.Palette()))
	}
	resize(e.Config.Size())

	return &scene{update: pane.Update, resize: resize, view: func(layout.Size) string { return pane.View() }}, nil
}
