package layout

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

// Header places a title opposite optional actions, followed by an optional
// context row. Actions are omitted when they cannot fit beside the full title.
// Actions are display text; applications own their bindings and effects.
type Header struct{ Title, Context, Actions string }

func (h Header) View(width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	top := chromeRow(t.Heading().Render(singleLine(h.Title)), singleLine(h.Actions), width)
	if h.Context == "" {
		return top
	}

	return Column(top, Fit(t.Subtle().Render(singleLine(h.Context)), Size{Width: width, Height: 1}))
}

// Footer places help opposite status on one row. Help takes precedence when
// both cannot fit. Fields accept pre-styled text such as native help output.
type Footer struct{ Help, Status string }

func (f Footer) View(width int, t theme.Theme) string {
	return chromeRow(t.Subtle().Render(singleLine(f.Help)), t.Subtle().Render(singleLine(f.Status)), width)
}

func chromeRow(left, right string, width int) string {
	if width <= 0 {
		return ""
	}

	gap := width - ansi.StringWidth(left) - ansi.StringWidth(right)
	if right == "" || gap < 1 {
		return Fit(left, Size{Width: width, Height: 1})
	}

	return left + strings.Repeat(" ", gap) + right
}

func singleLine(value string) string {
	return strings.SplitN(strings.ReplaceAll(value, "\r", ""), "\n", 2)[0]
}
