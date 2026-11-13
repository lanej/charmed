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

func TestSemanticStateAtomsKeepMeaningAndCompactGeometry(t *testing.T) {
	states := []struct {
		label, glyph, ascii string
		view                func(text.Emphasis, theme.Theme, ...text.Presentation) string
		tone                text.Tone
	}{
		{"Up", "↑", "^", feedback.Up.View, text.Success},
		{"Down", "↓", "v", feedback.Down.View, text.Danger},
		{"Healthy", "♥", "H", feedback.Healthy.View, text.Success},
		{"Unhealthy", "×", "x", feedback.Unhealthy.View, text.Danger},
		{"Ready", "✓", "r", feedback.Ready.View, text.Success},
		{"Not ready", "◇", "-", feedback.NotReady.View, text.Neutral},
		{"Canary", "◔", "c", feedback.Canary.View, text.Info},
		{"Live", "●", "l", feedback.Live.View, text.Info},
	}
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		for _, ascii := range []bool{false, true} {
			palette := theme.New(mode)
			palette.ASCII = ascii

			for _, emphasis := range []text.Emphasis{text.Normal, text.Bright, text.Dim} {
				for _, state := range states {
					glyph := state.glyph
					if ascii {
						glyph = state.ascii
					}

					full := ansi.Strip(state.view(emphasis, palette))
					if !strings.Contains(full, glyph+" "+state.label) {
						t.Fatalf("%s lost its label or glyph: %q", state.label, full)
					}

					compact := state.view(emphasis, palette, text.Compact)
					if ansi.Strip(compact) != glyph || ansi.StringWidth(compact) != 1 {
						t.Fatalf("%s compact is not a single bare glyph: %q", state.label, compact)
					}

					glyphPill := ansi.Strip(state.view(emphasis, palette, text.GlyphOnly))
					if ansi.StringWidth(glyphPill) != 3 {
						t.Fatalf("glyph-only pill lost its shell: %q", glyphPill)
					}

					cells := uv.NewStyledString(compact).Lines(ansi.GraphemeWidth)[0]
					if cells[0].Style.Bg != nil {
						t.Fatalf("compact %s retained a fill", state.label)
					}

					if !ascii {
						expected := uv.NewStyledString(text.ToneStyle(state.tone, emphasis, palette).Render(glyph)).Lines(ansi.GraphemeWidth)[0][0].Style.Fg
						if !sameStateColor(cells[0].Style.Fg, expected) {
							t.Fatalf("%s lost its tone", state.label)
						}
					}
				}
			}
		}
	}
}

func sameStateColor(a, b color.Color) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()

	return ar == br && ag == bg && ab == bb && aa == ba
}

func TestSemanticStateUnknownsNeverBecomePositiveObservations(t *testing.T) {
	palette := theme.New(theme.Dark)

	states := []func(text.Emphasis, theme.Theme, ...text.Presentation) string{
		feedback.AvailabilityUnknown.View, feedback.Availability("bad").View,
		feedback.HealthUnknown.View, feedback.Health("bad").View,
		feedback.ReadinessUnknown.View, feedback.Readiness("bad").View,
		feedback.RolloutUnknown.View, feedback.Rollout("bad").View,
	}
	for _, view := range states {
		if got := ansi.Strip(view(text.Normal, palette)); got != " ? Unknown " {
			t.Fatal("unknown observation changed", got)
		}

		if got := ansi.Strip(view(text.Normal, palette, text.Compact)); got != "?" {
			t.Fatal("compact unknown changed", got)
		}
	}

	for _, status := range feedback.Statuses() {
		compact := status.View(text.Normal, palette, text.Compact)
		if ansi.StringWidth(compact) != 1 {
			t.Fatal("legacy status compact retained its shell", status, compact)
		}
	}

	if got := ansi.Strip(feedback.Follow(true, text.Normal, palette, text.Compact)); got != "▶" {
		t.Fatal("follow compact", got)
	}

	if got := ansi.Strip(feedback.EvidenceHealthy.View(text.Normal, palette, text.Compact)); got != "✓" {
		t.Fatal("evidence compact", got)
	}
}
