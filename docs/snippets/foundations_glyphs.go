package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleFoundationsGlyphs() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		palette.Glyph(theme.GlyphSearch)+" Search",
		palette.Glyph(theme.GlyphExpanded)+" Expanded",
		palette.Glyph(theme.GlyphCollapsed)+" Collapsed",
		palette.Glyph(theme.GlyphSelected)+" Selected",
	)
	// Use theme.New(theme.Mono) for the matching ASCII glyphs.
}
