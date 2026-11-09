package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutHeader() string {
	palette := theme.New(theme.Dark)
	header := layout.Header{
		Title:   "Service workbench",
		Context: "Production / us-west-2",
		Actions: text.KeyHint("r", "refresh", palette),
	}

	return header.View(64, palette)
	// Bind r in the application; Header only renders the hint.
}
