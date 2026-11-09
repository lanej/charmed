package layout

import (
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// IdentityHeader makes an entity's identifier the headline. Kind, ID and Context
// are escaped data; Actions accepts styled display text. Identifiers move below
// the kind before wrapping, and are never abbreviated. Measure the rendered
// height when allocating the remaining screen. Actions appear only when they fit.
type IdentityHeader struct{ Kind, ID, Context, Actions string }

func (h IdentityHeader) View(width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	id := text.Inline(h.ID)
	if id == "" {
		id = "unavailable"
	}

	heading := t.Heading().Render(id)
	if h.Kind != "" {
		kind := t.Subtle().Render(text.Inline(h.Kind))

		separator := " "
		if ansi.StringWidth(kind)+1+ansi.StringWidth(heading) > width {
			separator = "\n"
		}

		heading = kind + separator + heading
	}

	heading = ansi.Hardwrap(heading, width, true)
	if lipgloss.Height(heading) == 1 {
		heading = chromeRow(heading, singleLine(h.Actions), width)
	}

	if h.Context != "" {
		heading += "\n" + ansi.Wrap(t.Subtle().Render(text.Inline(h.Context)), width, "")
	}

	return Fit(heading, Size{Width: width, Height: lipgloss.Height(heading)})
}
