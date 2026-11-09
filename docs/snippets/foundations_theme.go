package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFoundationsTheme() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		palette.Text().Render("Primary content"),
		palette.Subtle().Render("Supporting information"),
		palette.Heading().Render("Accent / interaction"),
		text.Pill("Success", text.Success, palette),
		text.Pill("Warning", text.Warning, palette),
	)
}
