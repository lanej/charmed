package snippets

import (
	"time"

	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleFeedbackFreshness() string {
	now := time.Date(2026, 11, 10, 12, 0, 0, 0, time.UTC)
	freshness := feedback.Freshness{
		ObservedAt: now.Add(-8 * time.Minute), StaleAfter: time.Minute,
		Emphasis: text.Normal,
	}

	return freshness.View(now, theme.New(theme.Dark))
}
