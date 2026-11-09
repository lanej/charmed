package text_test

import (
	"image/color"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestPendingFramesHaveStableStatusCellWidth(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		t.Run(string(mode), func(t *testing.T) {
			palette := theme.New(mode)
			for _, value := range []string{text.Pending(palette), text.Pending(palette, 0), text.Pending(palette, 1), text.Pending(palette, 2), text.Pending(palette, 3)} {
				if ansi.StringWidth(value) != 5 {
					t.Fatalf("pending changed status cell width: %q", value)
				}
			}

			if !palette.ReducedMotion {
				for frame, want := range []string{".", "..", "..."} {
					if got := strings.TrimSpace(ansi.Strip(text.Pending(palette, uint(frame)))); got != want {
						t.Fatalf("frame %d: got %q, want %q", frame, got, want)
					}
				}
			}

			palette.ReducedMotion = true
			if text.Pending(palette, 2) != text.Pending(palette) {
				t.Fatal("reduced motion animated pending")
			}

			if ansi.Strip(text.Pending(palette, 0)) != ansi.Strip(text.Pending(palette, 3)) {
				t.Fatal("ellipsis does not cycle")
			}
		})
	}
}

func TestObservationAtomsEscapeLabelsAndKeepASCIIFallbacks(t *testing.T) {
	palette := theme.New(theme.Mono)
	for _, specimen := range []struct {
		view   func(string, theme.Theme, ...text.Presentation) string
		prefix string
	}{
		{text.Log, "L"}, {text.Trace, "T"}, {text.Span, "S"}, {text.Upstream, "^"}, {text.Downstream, "v"},
	} {
		got := ansi.Strip(specimen.view("example\x1b[31m", palette))
		if got != specimen.prefix+` example\u001b[31m` {
			t.Fatal("lost identity or escaped label", got)
		}
	}
}

func TestAbbreviatedTimesHaveExplicitIndependentExpectations(t *testing.T) {
	for _, example := range []struct {
		value time.Duration
		want  string
	}{
		{0, "0s"}, {420 * time.Nanosecond, "420ns"}, {842 * time.Microsecond, "842µs"}, {842 * time.Millisecond, "842ms"}, {1234567890, "1.23s"}, {62 * time.Second, "1m2s"}, {-1234567890, "-1.23s"},
	} {
		if got := text.ShortDuration(example.value); got != example.want {
			t.Fatalf("%s: got %s, want %s", example.value, got, example.want)
		}
	}

	at := time.Date(2026, 11, 8, 12, 4, 1, 234567000, time.FixedZone("local", 3600))
	if got := text.TimeOfDay(at); got != "12:04:01.234" {
		t.Fatal("clock changed supplied location or precision", got)
	}

	if !strings.Contains(text.TimeOfDay(time.Time{}), "--:--:--") {
		t.Fatal("missing time became an observation")
	}
}

func TestPendingKeepsCustomForegroundWithoutRole(t *testing.T) {
	palette := theme.Theme{Foreground: color.RGBA{R: 220, G: 80, B: 60, A: 255}}
	want := palette.Control().Width(5).AlignHorizontal(0.5).Render(palette.Glyph(theme.GlyphPending))

	if got := text.Pending(palette); got != want {
		t.Fatalf("pending lost custom foreground: got %q, want %q", got, want)
	}
}
