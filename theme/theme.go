// Package theme defines the semantic colors and terminal policies shared by components.
package theme

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

type Mode string

const (
	Dark     Mode = "dark"
	Light    Mode = "light"
	Contrast Mode = "contrast"
	Mono     Mode = "mono"
)

// Theme is a value: a preview can change its theme without mutating other sessions.
type Theme struct {
	Mode                                                   Mode
	ASCII                                                  bool
	ReducedMotion                                          bool
	Background, Surface, Foreground, Muted, Border, Accent color.Color
	Success, Info, Warning, Danger                         color.Color
	Series                                                 []color.Color
}

func New(mode Mode) Theme {
	t := Theme{Mode: mode}
	switch mode {
	case Light:
		t.Background, t.Foreground = hex("#f8fafc"), hex("#172033")
		t.Surface = hex("#edf2f7")
		t.Muted, t.Border, t.Accent = hex("#526176"), hex("#94a3b8"), hex("#075985")
		t.Success, t.Warning, t.Danger = hex("#166534"), hex("#92400e"), hex("#b91c1c")
		t.Info = hex("#1d4ed8")
	case Contrast:
		t.Background, t.Foreground = hex("#000000"), hex("#ffffff")
		t.Surface = hex("#151515")
		t.Muted, t.Border, t.Accent = hex("#d0d0d0"), hex("#ffffff"), hex("#67e8f9")
		t.Success, t.Warning, t.Danger = hex("#86efac"), hex("#fde047"), hex("#fca5a5")
		t.Info = hex("#93c5fd")
	case Mono:
		t.ASCII, t.ReducedMotion = true, true
		t.Background, t.Surface = hex("#000000"), hex("#151515")
		t.Foreground, t.Muted, t.Border, t.Accent = hex("#ffffff"), hex("#bdbdbd"), hex("#777777"), hex("#ffffff")
	default:
		t.Mode = Dark
		t.Background, t.Foreground = hex("#101722"), hex("#eeeeee")
		t.Surface = hex("#1b2734")
		t.Muted, t.Border, t.Accent = hex("#a8a8a8"), hex("#475569"), hex("#67d4c0")
		t.Success, t.Warning, t.Danger = hex("#86cf9a"), hex("#f2c879"), hex("#f28b82")
		t.Info = hex("#78b8ed")
	}
	if mode != Mono {
		t.Series = []color.Color{t.Accent, hex("#c4a7e7"), t.Warning, hex("#78b8ed"), t.Success}
	}
	return t
}

func hex(s string) color.Color { return lipgloss.Color(s) }

func (t Theme) Text() lipgloss.Style    { return lipgloss.NewStyle().Foreground(t.Foreground) }
func (t Theme) Subtle() lipgloss.Style  { return lipgloss.NewStyle().Foreground(t.Muted) }
func (t Theme) Heading() lipgloss.Style { return lipgloss.NewStyle().Foreground(t.Accent).Bold(true) }
func (t Theme) Focus() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.Accent).Bold(true).Reverse(true)
}
func (t Theme) IsDark() bool { return t.Mode != Light }

func (t Theme) SeriesColor(index int) color.Color {
	if len(t.Series) == 0 {
		return nil
	}
	return t.Series[index%len(t.Series)]
}
