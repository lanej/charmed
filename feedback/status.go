package feedback

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Status is display state, independent of attention-rule evaluation.
type Status string

const (
	OK      Status = "OK"
	Warn    Status = "WARN"
	Error   Status = "ERROR"
	Unknown Status = "UNKNOWN"
	Off     Status = "OFF"
)

func (s Status) View(emphasis text.Emphasis, t theme.Theme) string {
	tone := text.Neutral
	switch s {
	case OK:
		tone = text.Success
	case Warn:
		tone = text.Warning
	case Error:
		tone = text.Danger
	case Unknown, Off:
	default:
		s = Unknown
	}
	if emphasis == text.Bright {
		return text.Pill(string(s), tone, t, emphasis)
	}
	return text.ToneStyle(tone, emphasis, t).Render(string(s))
}
