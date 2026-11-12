package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextPending() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		text.Pending(palette),
		text.PendingWithPresentation(text.GlyphAndLabel, palette),
		text.Pending(palette, 2), // Optional animation frame supplied by your clock.
	)
}
