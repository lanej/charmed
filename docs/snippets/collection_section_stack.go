package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionSectionStack() (string, error) {
	stack, err := collection.NewSectionStack([]collection.DocumentSection{
		{ID: "intent", Title: "Intent", Body: func(int) string { return "Retain the last healthy route while observations are incomplete." }},
		{ID: "evidence", Title: "Evidence", Summary: "3 observations", Body: func(int) string { return "West: 42 ms\nCentral: 61 ms\nEast: not captured" }},
	}, layout.Size{Width: 72, Height: 12}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	return stack.View(), nil
}
