package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutTextRow() string {
	palette := theme.New(theme.Dark)
	row := layout.TextRow{Gap: 2, Items: []layout.TextRowItem{
		{Text: palette.Heading().Render("API operations")},
		{Text: palette.Subtle().Render("Production / us-east-1")},
		{Flex: true},
		{Text: palette.Text().Render("Observed 00:16 UTC")},
	}}

	return row.View(64)
	// Use row.Stack = true for an explicitly vertical presentation.
	// Measure the returned height when allocating a parent panel or viewport.
}
