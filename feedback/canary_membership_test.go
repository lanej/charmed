package feedback_test

import (
	"image/color"
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestCanaryMembershipKeepsRoutingSeparateFromHealth(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		for _, ascii := range []bool{false, true} {
			palette := theme.New(mode)
			palette.ASCII = ascii

			for _, emphasis := range []text.Emphasis{text.Normal, text.Bright, text.Dim} {
				for _, state := range []struct {
					value               feedback.CanaryMembership
					label, glyph, ascii string
					color, dim          color.Color
				}{
					{feedback.CanaryOn, "Canary on", "◉", "+", palette.Info, palette.DimInfo},
					{feedback.CanaryOff, "Canary off", "◌", "-", palette.Inactive, palette.DimInactive},
				} {
					glyph := state.glyph
					if ascii {
						glyph = state.ascii
					}

					full := ansi.Strip(state.value.View(emphasis, palette))
					if !strings.Contains(full, glyph+" "+state.label) {
						t.Fatalf("membership label lost: %q", full)
					}

					compact := state.value.View(emphasis, palette, text.Compact)
					if ansi.Strip(compact) != glyph || ansi.StringWidth(compact) != 1 {
						t.Fatalf("membership compact geometry: %q", compact)
					}

					cell := uv.NewStyledString(compact).Lines(ansi.GraphemeWidth)[0][0]
					if cell.Style.Bg != nil {
						t.Fatal("compact membership retained fill")
					}

					if !ascii {
						expected := state.color
						if emphasis == text.Dim {
							expected = state.dim
						}

						if mode == theme.Mono && state.value == feedback.CanaryOff {
							expected = palette.Foreground
						}

						if !sameStateColor(cell.Style.Fg, expected) {
							t.Fatalf("%s %s %s lost routing color: got %#v want %#v", mode, state.label, emphasis, cell.Style.Fg, expected)
						}
					}
				}
			}

			if got := ansi.Strip(feedback.Canary.View(text.Normal, palette)); !strings.Contains(got, "Canary") || strings.Contains(got, "Canary on") || strings.Contains(got, "Canary off") {
				t.Fatal("cohort inferred membership", got)
			}
		}
	}
}

func TestUnknownCanaryMembershipDoesNotImplyOutOfPool(t *testing.T) {
	palette := theme.New(theme.Dark)
	for _, state := range []feedback.CanaryMembership{feedback.CanaryMembershipUnknown, feedback.CanaryMembership("bad")} {
		if got := ansi.Strip(state.View(text.Normal, palette)); got != " ? Unknown " {
			t.Fatal("unknown membership changed", got)
		}

		if got := ansi.Strip(state.View(text.Normal, palette, text.Compact)); got != "?" {
			t.Fatal("unknown compact membership changed", got)
		}
	}
}
