package snippets

import (
	"time"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackEvidenceFooter() string {
	now := time.Date(2026, 11, 11, 12, 0, 0, 0, time.UTC)
	footer := feedback.EvidenceFooter{
		Source: "sentry", Scope: "prod", State: feedback.EvidenceStale,
		ObservedAt: now.Add(-18 * time.Minute), Now: now, // Explicit clocks keep rendering deterministic.
	}

	return footer.View(80, theme.New(theme.Dark))
}
