package layout

import (
	"image/color"
	"math"
	"strings"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

type BackdropMode string

const (
	DefaultBackdrop    BackdropMode = ""
	BackdropDim        BackdropMode = "dim"
	BackdropMonochrome BackdropMode = "monochrome"
	BackdropNone       BackdropMode = "none"
)

func (mode BackdropMode) Valid() bool {
	switch mode {
	case DefaultBackdrop, BackdropDim, BackdropMonochrome, BackdropNone:
		return true
	default:
		return false
	}
}

const backdropColorRetention = .5

func dimColor(value, background color.Color) color.Color {
	if value == nil || background == nil {
		return value
	}

	a, ok := color.NRGBAModel.Convert(value).(color.NRGBA)
	if !ok {
		return value
	}

	b, ok := color.NRGBAModel.Convert(background).(color.NRGBA)
	if !ok {
		return value
	}

	mix := func(v, base uint8) uint8 {
		return uint8(math.Round(float64(v)*backdropColorRetention + float64(base)*(1-backdropColorRetention)))
	}

	return color.NRGBA{R: mix(a.R, b.R), G: mix(a.G, b.G), B: mix(a.B, b.B), A: a.A}
}

func renderBackdrop(source string, mode BackdropMode, t theme.Theme) string {
	switch mode {
	case BackdropNone:
		return source
	case BackdropMonochrome:
		style := t.Text()
		if t.DimForeground != nil {
			style = style.Foreground(t.DimForeground)
		} else if t.Muted != nil {
			style = style.Foreground(t.Muted)
		}

		return style.Render(ansi.Strip(source))
	case DefaultBackdrop, BackdropDim:
	default:
		return source
	}

	var out strings.Builder

	for y, line := range uv.NewStyledString(source).Lines(ansi.GraphemeWidth) {
		if y > 0 {
			if err := out.WriteByte('\n'); err != nil {
				return source
			}
		}

		for _, cell := range line {
			if cell.Width == 0 {
				continue
			}

			style := cell.Style
			if style.Fg == nil {
				style.Fg = t.Foreground
			}

			style.Fg = dimColor(style.Fg, t.Background)
			style.Bg = dimColor(style.Bg, t.Background)
			style.UnderlineColor = dimColor(style.UnderlineColor, t.Background)

			content := style.Styled(cell.Content)
			if cell.Link.URL != "" {
				content = ansi.SetHyperlink(cell.Link.URL, cell.Link.Params) + content + ansi.ResetHyperlink()
			}

			if _, err := out.WriteString(content); err != nil {
				return source
			}
		}
	}

	return out.String()
}
