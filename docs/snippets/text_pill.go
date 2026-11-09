package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextPill() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		text.Pill("Ready", text.Success, palette, text.Bright),
		text.Pill("Warning", text.Warning, palette, text.Normal),
		text.Pill("Unavailable", text.Neutral, palette, text.Dim),
	)
}
