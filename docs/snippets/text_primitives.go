package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextPrimitives() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		text.Label("production", palette),
		text.Badge("READY", palette)+" "+text.Badge("3 SELECTED", palette),
		text.Pill("Success", text.Success, palette),
		text.KeyHint("Enter", "select", palette),
	)
}
