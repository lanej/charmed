package collection_test

import (
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestPaletteMatchEmphasisIsOptionalAndActionRemainsDistinct(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		m, err := collection.NewCommandPalette([]collection.Command{{ID: "timeline", Label: "Layout: timeline", Description: "Arrange observations", Shortcut: "F5"}}, 72, 24, theme.New(mode))
		if err != nil {
			t.Fatal(err)
		}

		m.Open()
		m.SetQuery("laytm")

		if m.MatchHighlighting() {
			t.Fatal("highlighting enabled by default")
		}

		for _, enabled := range []bool{false, true, false} {
			m.SetMatchHighlighting(enabled)
			m.Resize(layout.Size{Width: 72, Height: 24})

			underlined := false
			action, description, shortcut := false, false, false

			for _, line := range uv.NewStyledString(m.Results.View()).Lines(ansi.GraphemeWidth) {
				for _, cell := range line {
					underlined = underlined || cell.Style.Underline != uv.UnderlineNone
					switch cell.Content {
					case "L":
						action = cell.Style.Attrs&uv.AttrBold != 0
					case "A":
						description = cell.Style.Attrs&uv.AttrBold == 0
					case "F":
						shortcut = cell.Style.Attrs&uv.AttrBold == 0
					}
				}
			}

			if underlined != enabled || !action || !description || !shortcut {
				t.Fatal("match or action emphasis incorrect", mode, enabled)
			}

			if m.Search.Value() != "laytm" || m.SelectedID() != "timeline" || len(m.Results.Items()) != 1 {
				t.Fatal("display option changed search or selection")
			}
		}
	}
}
