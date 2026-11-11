package feedback

import (
	"time"

	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Freshness compares explicit times. A zero observation is unavailable, a future
// observation is labeled as such, and StaleAfter <= 0 disables stale classification.
// An observation is stale at age >= StaleAfter. Rendering never schedules ticks.
type Freshness struct {
	ObservedAt time.Time
	StaleAfter time.Duration
	Emphasis   text.Emphasis
}

func (f Freshness) Label(now time.Time) string {
	if f.ObservedAt.IsZero() {
		return "age unavailable"
	}

	if now.IsZero() {
		return "clock unavailable"
	}

	age := now.Sub(f.ObservedAt)
	if age < 0 {
		return "in " + text.ShortDuration(f.ObservedAt.Sub(now)) + " (future)"
	}

	label := text.ShortDuration(age) + " ago"
	if f.StaleAfter > 0 && age >= f.StaleAfter {
		label = "STALE " + label
	}

	return label
}

func (f Freshness) View(now time.Time, t theme.Theme) string {
	tone := text.Neutral
	if !now.IsZero() && !f.ObservedAt.IsZero() && (now.Before(f.ObservedAt) || (f.StaleAfter > 0 && now.Sub(f.ObservedAt) >= f.StaleAfter)) {
		tone = text.Warning
	}

	return text.ToneStyle(tone, f.Emphasis, t).Render(f.Label(now))
}
