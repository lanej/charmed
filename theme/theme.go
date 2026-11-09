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
	Mode                                                        Mode
	ASCII                                                       bool
	ReducedMotion                                               bool
	Background, Surface, Foreground, Muted, Border, Accent      color.Color
	FocusBorder, SelectionForeground, SelectionBackground       color.Color
	Success, Info, Warning, Danger                              color.Color
	Unknown, Inactive, DimUnknown, DimInactive                  color.Color
	Pending, Running, Paused, DimPending, DimRunning, DimPaused color.Color
	DimForeground, DimSuccess, DimInfo, DimWarning, DimDanger   color.Color
	Series, Sequential                                          []color.Color
	Permissions                                                 PermissionColors
	HTTP                                                        HTTPColors
	Syntax                                                      SyntaxColors
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
		t.SelectionBackground = hex("#d1e8e3")
		t.Unknown, t.Inactive = hex("#475e86"), hex("#626262")
		t.DimUnknown, t.DimInactive = hex("#5c6e8a"), hex("#666666")
		t.Pending, t.Running, t.Paused = hex("#345a9c"), hex("#146375"), hex("#756044")
		t.DimPending, t.DimRunning, t.DimPaused = hex("#526fa0"), hex("#426f7c"), hex("#7f6b51")
		t.DimForeground, t.DimSuccess, t.DimInfo, t.DimWarning, t.DimDanger = hex("#526176"), hex("#477451"), hex("#526fa8"), hex("#8a6841"), hex("#a05b5b")
		t.Series = colors("#6654ad", "#3568b0", "#a3497f", "#8c663d", "#575ba6")
		t.Sequential = colors("#e5ecf7", "#b6cbea", "#80a4d3", "#4e79b0", "#1d3e6e")
	case Contrast:
		t.Background, t.Foreground = hex("#000000"), hex("#ffffff")
		t.Surface = hex("#151515")
		t.Muted, t.Border, t.Accent = hex("#d0d0d0"), hex("#ffffff"), hex("#67e8f9")
		t.Success, t.Warning, t.Danger = hex("#86efac"), hex("#fde047"), hex("#fca5a5")
		t.Info = hex("#93c5fd")
		t.SelectionBackground = hex("#1c4245")
		t.Unknown, t.Inactive = hex("#b7c9ed"), hex("#b0b0b0")
		t.DimUnknown, t.DimInactive = hex("#8da4cb"), hex("#929292")
		t.Pending, t.Running, t.Paused = hex("#abc1ff"), hex("#88ddeb"), hex("#e1ccb1")
		t.DimPending, t.DimRunning, t.DimPaused = hex("#8ba2da"), hex("#65aab7"), hex("#b5a28a")
		t.DimForeground, t.DimSuccess, t.DimInfo, t.DimWarning, t.DimDanger = hex("#c0c0c0"), hex("#6cbe80"), hex("#739ac7"), hex("#c5b75a"), hex("#d48282")
		t.Series = colors("#c8b4ff", "#9ac2ff", "#f4b0d8", "#e8c49a", "#abb6ff")
		t.Sequential = colors("#18283f", "#354f76", "#5b81b0", "#8db5e1", "#d9ebff")
	case Mono:
		t.ASCII, t.ReducedMotion = true, true
		t.Background, t.Surface = hex("#000000"), hex("#151515")
		t.Foreground, t.Muted, t.Border, t.Accent = hex("#ffffff"), hex("#bdbdbd"), hex("#777777"), hex("#ffffff")
		t.Sequential = colors("#202020", "#505050", "#808080", "#b0b0b0", "#eeeeee")
	default:
		t.Mode = Dark
		t.Background, t.Foreground = hex("#101722"), hex("#eeeeee")
		t.Surface = hex("#1b2734")
		t.Muted, t.Border, t.Accent = hex("#a8a8a8"), hex("#475569"), hex("#67d4c0")
		t.Success, t.Warning, t.Danger = hex("#86cf9a"), hex("#f2c879"), hex("#f28b82")
		t.Info = hex("#78b8ed")
		t.SelectionBackground = hex("#24433f")
		t.Unknown, t.Inactive = hex("#a7bce0"), hex("#9a9a9a")
		t.DimUnknown, t.DimInactive = hex("#829bc3"), hex("#898989")
		t.Pending, t.Running, t.Paused = hex("#8ca8e8"), hex("#70c6d8"), hex("#c5b49b")
		t.DimPending, t.DimRunning, t.DimPaused = hex("#758fc4"), hex("#519bab"), hex("#a59682")
		t.DimForeground, t.DimSuccess, t.DimInfo, t.DimWarning, t.DimDanger = hex("#a8a8a8"), hex("#629c73"), hex("#5a8fb9"), hex("#b5945c"), hex("#c17670")
		t.Series = colors("#b9a7f5", "#87b5f9", "#e49ac2", "#d7b58b", "#8f9de8")
		t.Sequential = colors("#202844", "#354b72", "#4e70a1", "#729bd0", "#afcbef")
	}

	t.FocusBorder, t.SelectionForeground = t.Accent, t.Foreground
	t.Permissions = permissionColors(t.Mode)
	t.HTTP, t.Syntax = httpColors(t.Mode)

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
const InlinePadding = SpaceSmall

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

// Selection styles a padded selection surface. Nested foreground styles retain
// their colors, including native table cells and semantic pills.
func (t Theme) Selection() lipgloss.Style {
	if t.Mode == Mono {
		return t.Focus()
	}

	if t.SelectionBackground == nil {
		// WHY: Custom themes need a visible cue without assuming a fill contrast.
		return t.Control().Bold(true).Underline(true)
	}

	style := t.Control().Background(t.SelectionBackground).Bold(true)
	if t.SelectionForeground != nil {
		style = style.Foreground(t.SelectionForeground)
	}

	return style
}
func (t Theme) IsDark() bool { return t.Mode != Light }

func (t Theme) SeriesColor(index int) color.Color {
	if len(t.Series) == 0 {
		return nil
	}
	return t.Series[index%len(t.Series)]
}

// SeriesStyle keeps plotted data independent of interaction emphasis.
func (t Theme) SeriesStyle(index int) lipgloss.Style {
	style := t.Text()
	if color := t.SeriesColor(index); color != nil {
		style = style.Foreground(color)
	}

	return style
}
