package feedback_test

import (
	"image/color"
	"strings"
	"testing"
	"time"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestFreshnessUsesExplicitClockAndExactStaleBoundary(t *testing.T) {
	now := time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)
	for _, tc := range []struct {
		f    feedback.Freshness
		want string
	}{
		{feedback.Freshness{}, "age unavailable"},
		{feedback.Freshness{ObservedAt: now}, "0s ago"},
		{feedback.Freshness{ObservedAt: now.Add(-12 * time.Second), StaleAfter: time.Minute}, "12s ago"},
		{feedback.Freshness{ObservedAt: now.Add(-time.Minute), StaleAfter: time.Minute}, "STALE 1m0s ago"},
		{feedback.Freshness{ObservedAt: now.Add(-8 * time.Minute)}, "8m0s ago"},
		{feedback.Freshness{ObservedAt: now.Add(time.Second)}, "in 1s (future)"},
	} {
		if got := tc.f.Label(now); got != tc.want {
			t.Fatalf("freshness got %q want %q", got, tc.want)
		}
	}

	if got := (feedback.Freshness{ObservedAt: now}).Label(time.Time{}); got != "clock unavailable" {
		t.Fatal("missing clock looked current")
	}
}

func TestSageFreshnessColorsRequireAnObservedWindow(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.Warning = color.RGBA{R: 230, G: 171, B: 75, A: 255}
	now := time.Unix(60, 0)

	for _, example := range []struct {
		freshness feedback.Freshness
		want      color.Color
	}{
		{feedback.Freshness{}, palette.Muted},
		{feedback.Freshness{ObservedAt: now}, palette.Foreground},
		{feedback.Freshness{ObservedAt: now, StaleAfter: time.Second}, palette.Fresh},
		{feedback.Freshness{ObservedAt: now.Add(-time.Second), StaleAfter: time.Second}, palette.Warning},
		{feedback.Freshness{ObservedAt: now.Add(time.Second), StaleAfter: time.Second}, palette.Warning},
	} {
		rendered := example.freshness.View(now, palette)
		cell := uv.NewStyledString(rendered).Lines(ansi.GraphemeWidth)[0][0]

		if !sameStateColor(cell.Style.Fg, example.want) {
			t.Errorf("%s: got %v, want %v", example.freshness.Label(now), cell.Style.Fg, example.want)
		}
	}
}

func TestUnknownOutcomePreservesOrdinaryTextUnlessOverridden(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Light, theme.Contrast} {
		palette := theme.New(mode)
		palette.Unknown = color.RGBA{R: 133, G: 213, B: 115, A: 255}

		for _, outcome := range []feedback.Outcome{feedback.OutcomeUnknown, "other"} {
			cell := uv.NewStyledString(outcome.View(text.Normal, palette)).Lines(ansi.GraphemeWidth)[0][1]
			if !sameStateColor(cell.Style.Fg, palette.Foreground) {
				t.Errorf("%s/%q: got %v, want normal %v", mode, outcome, cell.Style.Fg, palette.Foreground)
			}
		}

		palette.UnknownOutcome = palette.Warning
		cell := uv.NewStyledString(feedback.OutcomeUnknown.View(text.Normal, palette)).Lines(ansi.GraphemeWidth)[0][1]

		if !sameStateColor(cell.Style.Fg, palette.Warning) {
			t.Errorf("%s outcome override: got %v, want %v", mode, cell.Style.Fg, palette.Warning)
		}
	}
}

func TestOperationOutcomesRemainDistinctWithoutColor(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			for _, tc := range []struct {
				outcome feedback.Outcome
				label   string
			}{
				{feedback.OutcomeUnknown, "UNKNOWN"}, {feedback.Succeeded, "SUCCEEDED"}, {feedback.Failed, "FAILED"}, {feedback.Canceled, "CANCELED"}, {feedback.TimedOut, "TIMED OUT"}, {feedback.Outcome("other"), "UNKNOWN OUTCOME"},
			} {
				if got := ansi.Strip(tc.outcome.View(emphasis, theme.New(mode))); !strings.Contains(got, tc.label) {
					t.Fatalf("outcome lost: %q", got)
				}
			}
		}
	}
}
