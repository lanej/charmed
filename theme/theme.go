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
	Mode                                                      Mode
	ASCII                                                     bool
	ReducedMotion                                             bool
	Background, Foreground, Muted, Border, Accent             color.Color
	SelectionBackground                                       color.Color
	Success, Info, Warning, Danger                            color.Color
	DimForeground, DimSuccess, DimInfo, DimWarning, DimDanger color.Color
	Series                                                    []color.Color
}

func New(mode Mode) Theme {
	t := Theme{Mode: mode}
	switch mode {
	case Light:
		t.Background, t.Foreground = hex("#f8fafc"), hex("#172033")
		t.Muted, t.Border, t.Accent = hex("#526176"), hex("#94a3b8"), hex("#075985")
		t.Success, t.Warning, t.Danger = hex("#166534"), hex("#92400e"), hex("#b91c1c")
		t.Info = hex("#1d4ed8")
		t.SelectionBackground = hex("#d1e8e3")
		t.DimForeground, t.DimSuccess, t.DimInfo, t.DimWarning, t.DimDanger = hex("#526176"), hex("#477451"), hex("#526fa8"), hex("#8a6841"), hex("#a05b5b")
		t.Series = colors("#6654ad", "#3568b0", "#a3497f", "#8c663d", "#575ba6")
	case Contrast:
		t.Background, t.Foreground = hex("#000000"), hex("#ffffff")
		t.Muted, t.Border, t.Accent = hex("#d0d0d0"), hex("#ffffff"), hex("#67e8f9")
		t.Success, t.Warning, t.Danger = hex("#86efac"), hex("#fde047"), hex("#fca5a5")
		t.Info = hex("#93c5fd")
		t.SelectionBackground = hex("#1c4245")
		t.DimForeground, t.DimSuccess, t.DimInfo, t.DimWarning, t.DimDanger = hex("#c0c0c0"), hex("#6cbe80"), hex("#739ac7"), hex("#c5b75a"), hex("#d48282")
		t.Series = colors("#c8b4ff", "#9ac2ff", "#f4b0d8", "#e8c49a", "#abb6ff")
	case Mono:
		t.ASCII, t.ReducedMotion = true, true
	default:
		t.Mode = Dark
		t.Background, t.Foreground = hex("#101722"), hex("#eeeeee")
		t.Muted, t.Border, t.Accent = hex("#a8a8a8"), hex("#475569"), hex("#67d4c0")
		t.Success, t.Warning, t.Danger = hex("#86cf9a"), hex("#f2c879"), hex("#f28b82")
		t.Info = hex("#78b8ed")
		t.SelectionBackground = hex("#24433f")
		t.DimForeground, t.DimSuccess, t.DimInfo, t.DimWarning, t.DimDanger = hex("#a8a8a8"), hex("#629c73"), hex("#5a8fb9"), hex("#b5945c"), hex("#c17670")
		t.Series = colors("#b9a7f5", "#87b5f9", "#e49ac2", "#d7b58b", "#8f9de8")
	}

	return t
}

func hex(s string) color.Color { return lipgloss.Color(s) }

func colors(values ...string) []color.Color {
	result := make([]color.Color, len(values))
	for i, value := range values {
		result[i] = hex(value)
	}

	return result
}

// InlinePadding is the horizontal inset shared by controls and panel content, in cells.
const InlinePadding = 1

func (t Theme) Text() lipgloss.Style    { return lipgloss.NewStyle().Foreground(t.Foreground) }
func (t Theme) Subtle() lipgloss.Style  { return lipgloss.NewStyle().Foreground(t.Muted) }
func (t Theme) Heading() lipgloss.Style { return lipgloss.NewStyle().Foreground(t.Accent).Bold(true) }

// Control reserves the same inset as Focus so selection does not move its label.
func (t Theme) Control() lipgloss.Style { return t.Text().Padding(0, InlinePadding) }

func (t Theme) Focus() lipgloss.Style {
	if t.Mode == Mono {
		return t.Control().Bold(true).Reverse(true)
	}

	return t.Control().Foreground(t.Background).Background(t.Accent).Bold(true)
}

// Selection keeps nested semantic cells readable on a highlighted collection row.
func (t Theme) Selection() lipgloss.Style {
	if t.Mode == Mono {
		return t.Focus()
	}

	if t.SelectionBackground == nil {
		// WHY: Custom themes need a visible cue without assuming a fill contrast.
		return t.Control().Bold(true).Underline(true)
	}

	return t.Control().Background(t.SelectionBackground).Bold(true)
}
func (t Theme) IsDark() bool { return t.Mode != Light }

func (t Theme) SeriesColor(index int) color.Color {
	if len(t.Series) == 0 {
		return nil
	}
	return t.Series[index%len(t.Series)]
}
