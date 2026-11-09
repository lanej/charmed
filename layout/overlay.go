package layout

import (
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

// CenterOverlay preserves the background's cell geometry and clips both layers
// to size. Dimming replaces background styling without changing its text.
func CenterOverlay(background, foreground string, size Size, t theme.Theme, dim bool) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}

	base := Fit(background, size)

	if dim {
		style := t.Text()
		if t.DimForeground != nil {
			style = style.Foreground(t.DimForeground)
		} else if t.Muted != nil {
			style = style.Foreground(t.Muted)
		}

		base = style.Render(ansi.Strip(base))
	}

	panelSize := Size{Width: min(size.Width, lipgloss.Width(foreground)), Height: min(size.Height, lipgloss.Height(foreground))}
	panel := lipgloss.NewLayer(Fit(foreground, panelSize)).X((size.Width - panelSize.Width) / 2).Y((size.Height - panelSize.Height) / 2).Z(1)

	return Fit(lipgloss.NewCompositor(lipgloss.NewLayer(base), panel).Render(), size)
}
