package feedback_test

import (
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// themes covers every mode plus the explicit ASCII fallback of the colour modes.
func themes() map[string]theme.Theme {
	out := map[string]theme.Theme{}

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		out[string(mode)] = theme.New(mode)

		ascii := theme.New(mode)
		ascii.ASCII = true
		out[string(mode)+"-ascii"] = ascii
	}

	return out
}

func gaps() []feedback.Evidence {
	return []feedback.Evidence{
		feedback.EvidenceUnknown, feedback.EvidenceInsufficient, feedback.EvidenceAttributionGap,
		feedback.EvidenceStale, feedback.EvidenceTruncated, "",
	}
}

// TestEvidenceGapsNeverRenderAsHealthy guards the rule that missing, thin, aged, cut-short and unattributed evidence
// cannot pass as good: by status, glyph, word and view, in every theme including mono and ASCII.
func TestEvidenceGapsNeverRenderAsHealthy(t *testing.T) {
	for name, palette := range themes() {
		healthy := feedback.EvidenceHealthy.View(text.Normal, palette)
		okGlyph := palette.Glyph(theme.GlyphOK)

		for _, gap := range gaps() {
			if gap.Status() == feedback.OK {
				t.Fatalf("%s/%q maps to OK", name, gap)
			}

			footer := feedback.EvidenceFooter{Source: "connect", Scope: "dc1", State: gap}.View(80, palette)
			views := map[string]string{
				"pill":    gap.View(text.Normal, palette),
				"bright":  gap.View(text.Bright, palette),
				"summary": feedback.SummaryLine([]feedback.EvidenceCount{{State: gap, N: 3}}, 80, palette),
				"footer":  footer,
				"empty":   feedback.SummaryLine(nil, 80, palette),
			}

			for kind, view := range views {
				plain := strings.ToLower(ansi.Strip(view))
				if view == healthy || strings.Contains(plain, "healthy") || strings.Contains(plain, " ok ") {
					t.Fatalf("%s/%q/%s borrows the healthy word: %q", name, gap, kind, plain)
				}

				if strings.Contains(ansi.Strip(view), okGlyph) && okGlyph != "+" {
					t.Fatalf("%s/%q/%s borrows the healthy glyph: %q", name, gap, kind, plain)
				}
			}

			if !strings.Contains(ansi.Strip(views["pill"]), strings.ToUpper(gap.Label())) {
				t.Fatalf("%s/%q lost its word", name, gap)
			}
		}
	}
}

func TestEvidenceStatesAreDistinct(t *testing.T) {
	words := map[string]feedback.Evidence{}

	for _, e := range feedback.Evidences() {
		if other, dup := words[e.Label()]; dup {
			t.Fatalf("%q and %q share a word", e, other)
		}

		words[e.Label()] = e
	}

	// Unicode glyphs are unique per state; ASCII fallbacks may overlap because the word always accompanies them.
	unicode := theme.New(theme.Dark)
	seen := map[string]feedback.Evidence{}

	for _, e := range feedback.Evidences() {
		glyph := ansi.Strip(e.View(text.Normal, unicode, text.GlyphOnly))
		if other, dup := seen[glyph]; dup {
			t.Fatalf("%q and %q share glyph %q", e, other, glyph)
		}

		seen[glyph] = e
	}
}

func TestParseEvidence(t *testing.T) {
	if e, ok := feedback.ParseEvidence(" attribution gap "); !ok || e != feedback.EvidenceAttributionGap {
		t.Fatalf("parsed %q %v", e, ok)
	}

	if e, ok := feedback.ParseEvidence("fine"); ok || e != feedback.EvidenceUnknown {
		t.Fatalf("invalid input must be unknown, got %q %v", e, ok)
	}
}

func TestSummaryLineOrdersWorstFirstAndWraps(t *testing.T) {
	palette := theme.New(theme.Mono)
	got := ansi.Strip(feedback.SummaryLine([]feedback.EvidenceCount{
		{State: feedback.EvidenceHealthy, N: 2},
		{State: feedback.EvidenceUnknown, N: 1},
		{State: feedback.EvidenceStale, N: 1},
		{State: feedback.EvidenceDown, N: 1500},
		{State: "bogus", N: 1},
	}, 120, palette))

	if want := "x 1.5k down · ? 2 unknown · t 1 stale · + 2 healthy"; got != want {
		t.Fatalf("got %q want %q", got, want)
	}

	if ansi.Strip(feedback.SummaryLine(nil, 80, palette)) != "? no evidence" {
		t.Fatal("empty tally must say no evidence")
	}

	wrapped := feedback.SummaryLine([]feedback.EvidenceCount{
		{State: feedback.EvidenceDown, N: 1}, {State: feedback.EvidenceDegraded, N: 1}, {State: feedback.EvidenceHealthy, N: 1},
	}, 28, palette)
	for _, line := range strings.Split(wrapped, "\n") {
		if ansi.StringWidth(line) > 28 {
			t.Fatalf("line exceeds width: %q", line)
		}
	}
}

func TestEvidenceFooterStatesAgeAndOmission(t *testing.T) {
	palette := theme.New(theme.Mono)
	now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
	footer := feedback.EvidenceFooter{
		Source: "sentry", Scope: "prod", State: feedback.EvidenceStale, ObservedAt: now.Add(-18 * time.Minute), Now: now,
		Omission: &text.Omission{Kind: text.SourceTruncated, Count: 4, KnownCount: true},
	}

	got := ansi.Strip(footer.View(200, palette))
	for _, want := range []string{"source: sentry", "scope: prod", "STALE", "18m ago", "source truncated (4 omitted)"} {
		if !strings.Contains(got, want) {
			t.Fatalf("%q missing %q", got, want)
		}
	}

	missing := ansi.Strip(feedback.EvidenceFooter{State: feedback.EvidenceUnknown}.View(200, palette))
	for _, want := range []string{"source: ", "missing", "age unavailable"} {
		if !strings.Contains(missing, want) {
			t.Fatalf("%q missing %q", missing, want)
		}
	}

	future := feedback.EvidenceFooter{State: feedback.EvidenceDegraded, ObservedAt: now.Add(time.Minute), Now: now}
	if !strings.Contains(ansi.Strip(future.View(200, palette)), "in 1m (future)") {
		t.Fatal("future observation was not labeled")
	}
}

func TestEvidenceFooterNeverExceedsTheAllocation(t *testing.T) {
	for name, palette := range themes() {
		for _, width := range []int{20, 40, 55, 80} {
			for _, footer := range []feedback.EvidenceFooter{
				{Source: strings.Repeat("s", 60), Scope: "dc1", State: feedback.EvidenceStale},
				{Source: "connect", Scope: strings.Repeat("scope,", 12), State: feedback.EvidenceTruncated},
				{Source: strings.Repeat("s", 60), Scope: strings.Repeat("c", 60), State: feedback.EvidenceHealthy},
			} {
				view := footer.View(width, palette)

				for _, line := range strings.Split(view, "\n") {
					if ansi.StringWidth(line) > width {
						t.Fatalf("%s/%d: %d cells: %q", name, width, ansi.StringWidth(line), ansi.Strip(line))
					}
				}

				plain := strings.ReplaceAll(ansi.Strip(view), "\n", "")
				if !strings.Contains(plain, strings.Repeat("s", 10)) && footer.Source != "connect" {
					t.Fatalf("%s/%d lost the source: %q", name, width, plain)
				}
			}
		}
	}
}
