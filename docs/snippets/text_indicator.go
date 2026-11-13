package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextIndicator() string {
	palette := theme.New(theme.Dark)
	glyph := palette.Glyph(theme.GlyphUp)

	return layout.Column(
		text.Indicator(glyph, "Up", text.Success, palette, text.Normal),
		text.Indicator(glyph, "Up", text.Success, palette, text.Normal, text.GlyphOnly),
		text.Indicator(glyph, "Up", text.Success, palette, text.Normal, text.LabelOnly),
		text.Indicator(glyph, "Up", text.Success, palette, text.Normal, text.Compact),
	)
}
