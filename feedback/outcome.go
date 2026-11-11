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
	label, tone := "? UNKNOWN", text.Neutral

	switch o {
	case OutcomeUnknown:
	case Succeeded:
		label, tone = "+ SUCCEEDED", text.Success
	case Failed:
		label, tone = "x FAILED", text.Danger
	case Canceled:
		label = "- CANCELED"
	case TimedOut:
		label, tone = "! TIMED OUT", text.Warning
	default:
		label = "? UNKNOWN OUTCOME"
	}

	return text.Pill(label, tone, t, emphasis)
}
