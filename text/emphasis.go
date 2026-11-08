package text

import (
	"image/color"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/theme"
)

// Emphasis controls prominence independently of a semantic tone.
type Emphasis string

const (
	Bright Emphasis = "bright"
	Normal Emphasis = "normal"
	Dim    Emphasis = "dim"
)

func toneColor(tone Tone, t theme.Theme) color.Color {
	switch tone {
	case Success:
		return t.Success
	case Info:
		return t.Info
	case Warning:
		return t.Warning
	case Danger:
		return t.Danger
	}

	return t.Foreground
}

func dimColor(tone Tone, t theme.Theme) color.Color {
	var dim color.Color

	switch tone {
	case Neutral:
		dim = t.DimForeground
	case Success:
		dim = t.DimSuccess
	case Info:
		dim = t.DimInfo
	case Warning:
		dim = t.DimWarning
	case Danger:
		dim = t.DimDanger
	default:
		dim = t.DimForeground
	}

	if dim == nil {
		return toneColor(tone, t)
	}

	return dim
}

func blend(a, b color.Color, amount float64) color.Color {
	if a == nil || b == nil {
		return b
	}
	x := color.NRGBAModel.Convert(a).(color.NRGBA)
	y := color.NRGBAModel.Convert(b).(color.NRGBA)
	channel := func(a, b uint8) uint8 { return uint8(float64(a)*amount + float64(b)*(1-amount) + .5) }

	return color.NRGBA{R: channel(x.R, y.R), G: channel(x.G, y.G), B: channel(x.B, y.B), A: 255}
}

// ToneStyle is shared by bare numbers and pill content. Dim uses explicit
// hue-preserving theme tokens rather than the terminal's faint-text attribute.
func ToneStyle(tone Tone, emphasis Emphasis, t theme.Theme) lipgloss.Style {
	style := t.Text()
	if !t.ASCII {
		fg := toneColor(tone, t)
		if emphasis == Dim {
			fg = dimColor(tone, t)
		}

		style = style.Foreground(fg)
	}

	return style.Bold(emphasis == Bright)
}
