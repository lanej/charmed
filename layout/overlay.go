package layout

import (
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/theme"
)

// CenterOverlay retains the original boolean entrypoint. For independently
// configurable dim, monochrome and unchanged backgrounds, use CenterOverlayWithBackdrop.
func CenterOverlay(background, foreground string, size Size, t theme.Theme, dim bool) string {
	mode := BackdropNone
	if dim {
		mode = BackdropDim
	}

	return CenterOverlayWithBackdrop(background, foreground, size, t, mode)
}

func CenterOverlayWithBackdrop(background, foreground string, size Size, t theme.Theme, mode BackdropMode) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}

	base := renderBackdrop(Fit(background, size), mode, t)
	panelSize := Size{Width: min(size.Width, lipgloss.Width(foreground)), Height: min(size.Height, lipgloss.Height(foreground))}
	panel := lipgloss.NewLayer(Fit(foreground, panelSize)).X((size.Width - panelSize.Width) / 2).Y((size.Height - panelSize.Height) / 2).Z(1)

	return Fit(lipgloss.NewCompositor(lipgloss.NewLayer(base), panel).Render(), size)
}
