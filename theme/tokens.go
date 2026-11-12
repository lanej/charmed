package theme

import "charm.land/lipgloss/v2"

// Spacing is measured in terminal cells, not pixels. Horizontal and vertical
// application remain explicit because character cells are taller than wide.
const (
	SpaceNone   = 0
	SpaceSmall  = 1
	SpaceMedium = 2
	SpaceLarge  = 4

	// BorderCells is the thickness of each side of a frame.
	BorderCells = 1
)

// FrameBorder exposes the native border shape with an explicit ASCII fallback.
func (t Theme) FrameBorder() lipgloss.Border {
	if t.ASCII {
		return lipgloss.ASCIIBorder()
	}

	return lipgloss.NormalBorder()
}

// BorderStyle distinguishes focus without changing the frame allocation.
func (t Theme) BorderStyle(focused bool) lipgloss.Style {
	style := lipgloss.NewStyle().Foreground(t.Border)
	if !focused {
		return style
	}

	color := t.FocusBorder
	if color == nil {
		color = t.Accent
	}

	return style.Foreground(color).Bold(true)
}

// LinkStyle keeps links recognizable independently of color. Partial custom
// themes fall back to Accent, then their ordinary text foreground.
func (t Theme) LinkStyle() lipgloss.Style {
	foreground := t.Link
	if foreground == nil {
		foreground = t.Accent
	}

	style := t.Text().Underline(true)
	if foreground != nil {
		style = style.Foreground(foreground)
	}

	return style
}
