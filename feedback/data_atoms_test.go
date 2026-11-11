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
