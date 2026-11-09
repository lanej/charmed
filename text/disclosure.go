package text

import "github.com/lanej/charmed/theme"

// Disclosure labels a fold without owning its expansion state or key handling.
func Disclosure(label string, open bool, t theme.Theme, options ...Presentation) string {
	glyph := theme.GlyphCollapsed
	if open {
		glyph = theme.GlyphExpanded
	}

	return t.Heading().Render(presentation(options).Format(t.Glyph(glyph), label))
}
