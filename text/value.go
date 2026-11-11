package text

import (
	"fmt"

	"github.com/lanej/charmed/theme"
)

// ValueState describes presence independently of the value's contents.
// Zero, false and an empty string are all present values.
type ValueState string

const (
	ValuePresent  ValueState = ""
	ValueMissing  ValueState = "missing"
	ValueRedacted ValueState = "redacted"
	ValueInvalid  ValueState = "invalid"
)

func (s ValueState) Label(value string) string {
	switch s {
	case ValuePresent:
		if value == "" {
			return `""`
		}

		return Inline(value)
	case ValueMissing:
		return "missing"
	case ValueRedacted:
		return "redacted"
	case ValueInvalid:
		return "invalid"
	default:
		return fmt.Sprintf("unknown value state %q", Inline(string(s)))
	}
}

func (s ValueState) View(value string, emphasis Emphasis, t theme.Theme) string {
	tone := Neutral
	glyph := theme.GlyphUnknown

	switch s {
	case ValuePresent:
	case ValueMissing:
		glyph = theme.GlyphMissing
	case ValueRedacted:
		glyph = theme.GlyphRedacted
	case ValueInvalid:
		tone = Warning
		glyph = theme.GlyphWarning
	default:
		tone = Warning
	}

	label := s.Label(value)
	if s != ValuePresent && !t.ASCII {
		label = t.Glyph(glyph) + " " + label
	}

	return ToneStyle(tone, emphasis, t).Render(label)
}
