package layout_test

import (
	"image/color"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestBackdropModesPreserveTextWithDistinctColorPolicies(t *testing.T) {
	tokens := theme.New(theme.Dark)
	tokens.Background = color.NRGBA{A: 255}

	background := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Bold(true).Render("R") + lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render("G")
	for _, mode := range []layout.BackdropMode{layout.BackdropDim, layout.BackdropMonochrome, layout.BackdropNone} {
		view := layout.CenterOverlayWithBackdrop(background, "", layout.Size{Width: 2, Height: 1}, tokens, mode)
		if ansi.Strip(view) != "RG" {
			t.Fatal("backdrop changed text", mode, view)
		}

		cells := uv.NewStyledString(view).Lines(ansi.GraphemeWidth)[0]
		if len(cells) != 2 {
			t.Fatal("backdrop changed cells")
		}

		a, b := color.NRGBAModel.Convert(cells[0].Style.Fg), color.NRGBAModel.Convert(cells[1].Style.Fg)

		switch mode {
		case layout.BackdropDim:
			if a != (color.NRGBA{R: 128, A: 255}) || b != (color.NRGBA{G: 128, A: 255}) || cells[0].Style.Attrs&uv.AttrBold == 0 {
				t.Fatal("dim lost hue or attributes", a, b)
			}
		case layout.BackdropMonochrome:
			if a != b {
				t.Fatal("monochrome retained category colors")
			}
		case layout.BackdropNone:
			if a != (color.NRGBA{R: 255, A: 255}) || b != (color.NRGBA{G: 255, A: 255}) {
				t.Fatal("unchanged backdrop altered colors")
			}
		case layout.DefaultBackdrop:
		default:
			t.Fatal("unexpected mode")
		}
	}

	for _, mode := range []layout.BackdropMode{layout.BackdropDim, layout.BackdropMonochrome, layout.BackdropNone} {
		view := layout.CenterOverlayWithBackdrop(background, "X", layout.Size{Width: 4, Height: 3}, tokens, mode)
		if !strings.Contains(ansi.Strip(view), " X ") {
			t.Fatal("backdrop affected foreground position")
		}
	}
}
