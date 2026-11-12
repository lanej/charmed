package text

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/x/ansi"
)

// CellLabel converts display text to one plain terminal line without changing
// caller data. Escape sequences cannot override surrounding semantic styling.
func CellLabel(value string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsControl(r) || r == '\u2028' || r == '\u2029' {
			return ' '
		}

		return r
	}, ansi.Strip(value))
}
