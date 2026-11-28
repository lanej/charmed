package feedback

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Outcome is a reported operation result, not a prediction about outstanding work.
// Zero deliberately means unknown. Cancellation and timeout do not imply rollback.
type Outcome string

const (
	OutcomeUnknown Outcome = ""
	Succeeded      Outcome = "succeeded"
	Failed         Outcome = "failed"
	Canceled       Outcome = "canceled"
	TimedOut       Outcome = "timed-out"
)

func (o Outcome) View(emphasis text.Emphasis, t theme.Theme) string {
	label, tone, glyph := "UNKNOWN", text.Neutral, theme.GlyphUnknown

	switch o {
	case OutcomeUnknown:
		t = statusTheme(t, t.UnknownOutcome, t.DimUnknownOutcome)
	case Succeeded:
		label, tone, glyph = "SUCCEEDED", text.Success, theme.GlyphOK
	case Failed:
		label, tone, glyph = "FAILED", text.Danger, theme.GlyphError
	case Canceled:
		label, glyph = "CANCELED", theme.GlyphDisabled
		t = statusTheme(t, t.Canceled, t.DimCanceled)
	case TimedOut:
		label, tone, glyph = "TIMED OUT", text.Warning, theme.GlyphClock
	default:
		label = "UNKNOWN OUTCOME"
		t = statusTheme(t, t.UnknownOutcome, t.DimUnknownOutcome)
	}

	return text.Pill(t.Glyph(glyph)+" "+label, tone, t, emphasis)
}
