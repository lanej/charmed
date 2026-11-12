package feedback

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// StatusCell combines health with an application-owned label such as a version
// or a refined state. A zero Status means UNKNOWN, never an absent mapping.
type StatusCell struct {
	Status Status
	Label  string
}

// View renders one terminal row. Selection replaces the side gutters with
// focus edges (brackets in ASCII mode), retaining both the status fill and the cell's allocation.
func (c StatusCell) View(width int, t theme.Theme, selected bool) string {
	if width <= 0 {
		return ""
	}

	status, glyph, tone, palette := statusAppearance(c.Status, t)

	label := CellLabel(c.Label)
	if label == "" {
		label = string(status)
	}

	inside := layout.Fit(glyph+" "+label, layout.Size{Width: max(0, width-2), Height: 1})

	left, right := " ", " "
	if selected {
		left, right = t.Glyph(theme.GlyphFocusStart), t.Glyph(theme.GlyphFocusEnd)
	}

	style := text.ToneStyle(tone, text.Bright, palette)
	if t.Mode != theme.Mono {
		style = style.Background(style.GetForeground()).Foreground(t.Background)
	}

	return style.Render(layout.Fit(left+inside+right, layout.Size{Width: width, Height: 1}))
}

// CellLabel converts display text to one plain terminal line. Data held by the
// caller remains unchanged; escape sequences cannot override semantic styling.
func CellLabel(value string) string {
	return text.CellLabel(value)
}
