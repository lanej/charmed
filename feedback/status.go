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
	glyph := theme.GlyphUnknown

	switch s {
	case OK:
		tone = text.Success
		glyph = theme.GlyphOK
	case Warn:
		tone = text.Warning
		glyph = theme.GlyphWarning
	case Error:
		tone = text.Danger
		glyph = theme.GlyphError
	case Unknown:
	case Off:
		glyph = theme.GlyphOff
	default:
		s = Unknown
	}

	return text.Pill(t.Glyph(glyph)+" "+string(s), tone, t, emphasis)
}
