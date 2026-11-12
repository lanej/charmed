package collection

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/theme"
)

func logStyle(t theme.Theme, foreground color.Color) lipgloss.Style {
	style := t.Text()
	if t.Mode != theme.Mono && foreground != nil {
		style = style.Foreground(foreground)
	}

	return style
}
