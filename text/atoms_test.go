package text_test

import (
	"math"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestSignedNumberPreservesDirectionAndMissingness(t *testing.T) {
	for _, tc := range []struct {
		value     float64
		available bool
		want      string
	}{
		{.25, true, "+0.250pp"}, {-.25, true, "-0.250pp"}, {0, true, "+0.000pp"}, {math.Copysign(0, -1), true, "+0.000pp"},
		{-.000001, true, "-1e-06pp"}, {0, false, "n/a"}, {math.NaN(), true, "n/a"}, {math.Inf(1), true, "n/a"},
	} {
		n := text.SignedNumber{Value: tc.value, Available: tc.available, Precision: 3, Unit: "pp"}
		for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
			if got := ansi.Strip(n.View(theme.New(mode))); got != tc.want {
				t.Fatalf("%v %g: got %q want %q", mode, tc.value, got, tc.want)
			}
		}
	}
	palette := theme.New(theme.Dark)
	n := text.SignedNumber{Value: .25, Available: true, Precision: 3, Tones: text.SignTones{Positive: text.Danger, Negative: text.Success}}
	if n.View(palette) != text.ToneStyle(text.Danger, text.Normal, palette).Render("+0.250") {
		t.Fatal("caller-supplied increasing-is-bad tone was ignored")
	}
	n.Available = false
	n.Tones.Zero = text.Success
	if n.View(palette) != text.ToneStyle(text.Neutral, text.Normal, palette).Render("n/a") {
		t.Fatal("missing observation inherited a positive tone")
	}
}

func TestPillEmphasisKeepsToneAndLabel(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		for _, tone := range []text.Tone{text.Neutral, text.Success, text.Info, text.Warning, text.Danger} {
			seen := map[string]bool{}
			for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
				view := text.Pill("5xx", tone, palette, emphasis)
				if !strings.Contains(ansi.Strip(view), "5xx") {
					t.Fatal("pill emphasis removed its semantic label")
				}
				seen[view] = true
				if mode == theme.Mono {
					for _, r := range ansi.Strip(view) {
						if r > 127 {
							t.Fatal("ASCII pill contains non-ASCII decoration")
						}
					}
				}
			}
			if mode != theme.Mono && len(seen) != 3 {
				t.Fatalf("%s %s emphasis variants are indistinguishable", mode, tone)
			}
		}
	}
}
