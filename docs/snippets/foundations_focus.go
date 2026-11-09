package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleFoundationsFocus() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		palette.Control().Render("Normal control"),
		palette.Focus().Render("Focused control"),
		palette.Selection().Render("Selected row"),
	)
}
