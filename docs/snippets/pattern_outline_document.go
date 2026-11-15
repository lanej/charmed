package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func ExamplePatternOutlineDocument() (string, error) {
	document, err := pattern.NewOutlineDocument([]pattern.OutlineEntry{
		{ID: "proposal", Title: "Routing proposal", Children: []pattern.OutlineEntry{
			{ID: "intent", Title: "Intent", Body: func(int) string { return "Keep the last healthy route." }},
			{ID: "evidence", Title: "Evidence", Body: func(int) string { return "West: 42 ms\nEast: observation missing" }},
		}},
		{ID: "decision", Title: "Decision", Body: func(int) string { return "Wait for complete evidence." }},
	}, layout.Size{Width: 100, Height: 20}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	document.Select("evidence")

	return document.View(), nil
}
