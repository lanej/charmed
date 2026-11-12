package feedback_test

import (
	"image/color"
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func TestStatusCellUsesSemanticHealthColorsAndMonochromeGlyphs(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		for _, expected := range []struct {
			status feedback.Status
			fill   color.Color
			glyph  string
		}{
			{feedback.OK, palette.Success, "+"},
			{feedback.Warn, palette.Warning, "!"},
			{feedback.Error, palette.Danger, "x"},
			{feedback.Unknown, palette.Unknown, "?"},
			{feedback.Off, palette.Inactive, "-"},
			{feedback.Pending, palette.Pending, "p"},
			{feedback.Running, palette.Running, ">"},
			{feedback.Paused, palette.Paused, "="},
			{feedback.Disabled, palette.Inactive, "/"},
		} {
			frame := (feedback.StatusCell{Status: expected.status, Label: "v2.8"}).View(18, palette, false)
			if mode == theme.Mono {
				if !strings.HasPrefix(ansi.Strip(frame), " "+expected.glyph+" ") {
					t.Fatalf("%s lost monochrome health glyph: %q", expected.status, frame)
				}

				continue
			}

			cells := uv.NewStyledString(frame).Lines(ansi.GraphemeWidth)[0]
			if cells[2].Style.Bg == nil {
				t.Fatalf("%s lost semantic fill", expected.status)
			}

			r, g, b, _ := cells[2].Style.Bg.RGBA()
			wantR, wantG, wantB, _ := expected.fill.RGBA()

			if r != wantR || g != wantG || b != wantB {
				t.Fatalf("%s/%s fill is not its theme token", mode, expected.status)
			}
		}
	}
}

func TestStatusCellPreservesHealthFillAndSelectionGeometry(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)

		for _, status := range feedback.Statuses() {
			cell := feedback.StatusCell{Status: status, Label: "v2.8.1"}
			plain := cell.View(18, palette, false)

			selected := cell.View(18, palette, true)
			if ansi.StringWidth(plain) != 18 || ansi.StringWidth(selected) != 18 || !strings.Contains(ansi.Strip(selected), "v2.8.1") {
				t.Fatalf("%s/%s lost label or geometry: %q", mode, status, selected)
			}

			if !strings.HasPrefix(ansi.Strip(selected), palette.Glyph(theme.GlyphFocusStart)) || !strings.HasSuffix(ansi.Strip(selected), palette.Glyph(theme.GlyphFocusEnd)) {
				t.Fatalf("selection is not visible: %q", selected)
			}

			if mode != theme.Mono {
				a := uv.NewStyledString(plain).Lines(ansi.GraphemeWidth)[0]

				b := uv.NewStyledString(selected).Lines(ansi.GraphemeWidth)[0]
				if a[2].Style.Bg == nil || a[2].Style.Bg != b[2].Style.Bg {
					t.Fatalf("%s/%s selection changed health fill", mode, status)
				}
			}
		}
	}
}

func TestStatusCellHandlesUnknownAndUntrustedLabels(t *testing.T) {
	palette := theme.New(theme.Mono)
	for _, status := range []feedback.Status{"", "surprise", feedback.Unknown} {
		if label := ansi.Strip((feedback.StatusCell{Status: status}).View(18, palette, false)); !strings.Contains(label, "? UNKNOWN") {
			t.Fatalf("unrecognized status did not become UNKNOWN: %q", label)
		}
	}

	cell := feedback.StatusCell{Status: feedback.OK, Label: "\x1b[31m界\x1b[0m\n\ttext\x1b]0;title\a"}
	for width := -1; width <= 24; width++ {
		frame := cell.View(width, palette, true)
		if ansi.StringWidth(frame) != max(0, width) || strings.ContainsAny(ansi.Strip(frame), "\n\r\t") {
			t.Fatalf("unsafe or incorrectly sized cell at width %d: %q", width, frame)
		}
	}
}
