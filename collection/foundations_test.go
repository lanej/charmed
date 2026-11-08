package collection_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/table"

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

func TestNativeTableHeaderPaddingMatchesAllocationHelpers(t *testing.T) {
	columns := []table.Column{{Title: "NAME", Width: 6}, {Title: "STATE", Width: 8}}
	palette := theme.New(theme.Dark)

	for _, specimen := range []struct {
		model   table.Model
		padding int
	}{
		{collection.NewTable(columns, nil, 40, 4, palette), collection.TablePaddingWidth(len(columns))},
		{collection.NewDenseTable(columns, nil, 40, 4, palette), collection.DenseTablePaddingWidth(len(columns))},
	} {
		header := strings.Split(ansi.Strip(specimen.model.View()), "\n")[0]
		if ansi.StringWidth(header) != 14+specimen.padding {
			t.Fatalf("native cell insets diverged from allocation: %q", header)
		}
	}
}
