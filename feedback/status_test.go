package feedback_test

import (
	"image/color"
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestStatusAtomsKeepUnavailableAndDisabledDistinct(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			for _, status := range feedback.Statuses() {
				label := ansi.Strip(status.View(emphasis, theme.New(mode)))
				if !strings.Contains(label, string(status)) {
					t.Fatalf("%s status lost its label", status)
				}

				if ansi.StringWidth(label) != len(status)+4 {
					t.Fatalf("%s %s %s lost icon, gap or padding: %q", mode, emphasis, status, label)
				}
			}
		}

		palette := theme.New(mode)
		if feedback.Attention(attention.Off, palette) != "" || !strings.Contains(ansi.Strip(feedback.Attention(attention.Unknown, palette)), "UNKNOWN") {
			t.Fatal("off and unknown attention states were conflated")
		}
	}
}

func TestLightAndContrastUnknownIsCoolAndOffIsAchromaticAtEveryEmphasis(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Light, theme.Contrast} {
		palette := theme.New(mode)

		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			colors := make([]color.Color, 0, 2)

			for _, status := range []feedback.Status{feedback.Unknown, feedback.Off} {
				cells := uv.NewStyledString(status.View(emphasis, palette)).Lines(ansi.GraphemeWidth)[0]
				if emphasis == text.Bright {
					colors = append(colors, cells[1].Style.Bg)
				} else {
					colors = append(colors, cells[1].Style.Fg)
				}
			}

			if colors[0] == nil || colors[1] == nil {
				t.Fatalf("%s %s status color disappeared", mode, emphasis)
			}

			r, _, b, _ := colors[0].RGBA()
			if b <= r {
				t.Fatalf("%s %s UNKNOWN lost its cool hue", mode, emphasis)
			}

			r, g, b, _ := colors[1].RGBA()
			if r != g || g != b {
				t.Fatalf("%s %s OFF must remain achromatic", mode, emphasis)
			}
		}
	}
}

func TestSageStatusColorsUseStateDimensions(t *testing.T) {
	palette := theme.New(theme.Dark)
	for _, example := range []struct {
		status feedback.Status
		want   color.Color
	}{
		{feedback.Unknown, palette.Muted},
		{feedback.Off, palette.Inactive},
		{feedback.Pending, palette.Muted},
		{feedback.Running, palette.Info},
		{feedback.Paused, palette.Muted},
		{feedback.Disabled, palette.Inactive},
		{feedback.Connected, palette.Connected},
		{feedback.Disconnected, palette.Inactive},
	} {
		rendered := example.status.View(text.Normal, palette, text.Compact)
		cell := uv.NewStyledString(rendered).Lines(ansi.GraphemeWidth)[0][0]

		if !sameStateColor(cell.Style.Fg, example.want) {
			t.Errorf("%s: got %v, want %v", example.status, cell.Style.Fg, example.want)
		}
	}
}

func TestSageInactiveStatesInheritCallerColors(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.Inactive = color.RGBA{R: 122, G: 200, B: 183, A: 255}
	palette.DimInactive = color.RGBA{R: 93, G: 166, B: 110, A: 255}

	for _, status := range []feedback.Status{feedback.Disabled, feedback.Disconnected} {
		for emphasis, want := range map[text.Emphasis]color.Color{text.Bright: palette.Inactive, text.Normal: palette.Inactive, text.Dim: palette.DimInactive} {
			cell := uv.NewStyledString(status.View(emphasis, palette, text.Compact)).Lines(ansi.GraphemeWidth)[0][0]
			if !sameStateColor(cell.Style.Fg, want) {
				t.Errorf("%s/%s: got %v, want inherited %v", status, emphasis, cell.Style.Fg, want)
			}
		}
	}

	palette.Disabled = palette.Warning
	cell := uv.NewStyledString(feedback.Disabled.View(text.Normal, palette, text.Compact)).Lines(ansi.GraphemeWidth)[0][0]

	if !sameStateColor(cell.Style.Fg, palette.Warning) {
		t.Errorf("disabled override: got %v, want %v", cell.Style.Fg, palette.Warning)
	}
}

func TestLifecycleStatusesKeepLabelsGlyphsAndReservedWidth(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)

		for _, state := range []struct {
			status         feedback.Status
			unicode, ascii string
		}{
			{feedback.Pending, "◷", "p"},
			{feedback.Running, "▶", ">"},
			{feedback.Paused, "‖", "="},
			{feedback.Disabled, "⊘", "/"},
			{feedback.Connected, "●", "+"},
			{feedback.Disconnected, "○", "-"},
		} {
			glyph := state.unicode
			if palette.ASCII {
				glyph = state.ascii
			}

			for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
				want := " " + glyph + " " + string(state.status) + " "
				if palette.ASCII {
					want = "[" + glyph + " " + string(state.status) + "]"
				}

				if got := ansi.Strip(state.status.View(emphasis, palette)); got != want {
					t.Fatalf("%s %s lifecycle status: got %q want %q", mode, emphasis, got, want)
				}
			}
		}

		if feedback.MaxStatusWidth(palette) != 16 {
			t.Fatal("status allocation must fit the sixteen-cell DISCONNECTED pill")
		}
	}
}

func TestAttentionWidthIncludesUnknownButExcludesLifecycleLabels(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		if feedback.MaxAttentionWidth(palette) != 11 {
			t.Fatal("attention must fit UNKNOWN without reserving lifecycle labels")
		}
	}
}
