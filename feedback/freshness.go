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

	switch {
	case f.ObservedAt.IsZero() || now.IsZero():
		t = statusTheme(t, t.Missing, t.DimMissing)
	case now.Before(f.ObservedAt):
		tone = text.Warning
	case f.StaleAfter > 0 && now.Sub(f.ObservedAt) >= f.StaleAfter:
		tone = text.Warning
		if t.Stale != nil {
			t.Warning = t.Stale
		}
	case f.StaleAfter > 0:
		t = statusTheme(t, t.Fresh, t.DimFresh)
	}

	label := f.Label(now)

	if !t.ASCII {
		glyph := theme.GlyphClock
		if f.ObservedAt.IsZero() || now.IsZero() {
			glyph = theme.GlyphUnavailable
		}

		label = t.Glyph(glyph) + " " + label
	}

	return text.ToneStyle(tone, f.Emphasis, t).Render(label)
}
