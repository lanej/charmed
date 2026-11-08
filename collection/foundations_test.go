package collection_test

import (
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestListAppliesGlyphsToItsNativePaginator(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		palette := theme.New(mode)
		model := collection.NewList(nil, 40, 16, palette)

		if ansi.Strip(model.Paginator.ActiveDot) != palette.Glyph(theme.GlyphSelected) || ansi.Strip(model.Paginator.InactiveDot) != palette.Glyph(theme.GlyphUnselected) {
			t.Fatalf("%s list paginator retained upstream glyphs", mode)
		}

		if model.Paginator.ActiveDot == model.Paginator.InactiveDot {
			t.Fatalf("%s list pagination conflated selected and unselected", mode)
		}
	}
}
