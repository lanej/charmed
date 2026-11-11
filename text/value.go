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

	switch s {
	case ValuePresent, ValueMissing, ValueRedacted:
	case ValueInvalid:
		tone = Warning
	default:
		tone = Warning
	}

	return ToneStyle(tone, emphasis, t).Render(s.Label(value))
}
