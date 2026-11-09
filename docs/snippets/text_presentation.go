package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextPresentation() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		feedback.Connected.View(text.Normal, palette, text.GlyphAndLabel),
		feedback.Connected.View(text.Normal, palette, text.GlyphOnly),
		feedback.Connected.View(text.Normal, palette, text.LabelOnly),
	)
}
