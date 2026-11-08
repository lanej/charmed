// Package layout composes ANSI text using terminal cells rather than bytes or pixels.
package layout

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

// Size includes a component's border and padding.
type Size struct{ Width, Height int }

// Fit clips and pads to exactly the requested cell rectangle. It never wraps.
func Fit(content string, size Size) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	out := make([]string, size.Height)
	for y := range out {
		line := ""
		if y < len(lines) {
			line = ansi.Truncate(lines[y], size.Width, "")
		}
		out[y] = line + strings.Repeat(" ", max(0, size.Width-lipgloss.Width(line)))
	}
	return strings.Join(out, "\n")
}

// ASCII replaces terminal decoration when Unicode is unavailable. Application
// state is untouched; unsupported text glyphs get an explicit '?' fallback.
func ASCII(content string) string {
	return strings.Map(func(r rune) rune {
		if r < 128 {
			return r
		}
		switch r {
		case '─', '━', '═', '–', '—':
			return '-'
		case '│', '┃', '║':
			return '|'
		case '┌', '┐', '└', '┘', '├', '┤', '┬', '┴', '┼', '╭', '╮', '╰', '╯':
			return '+'
		case '●', '•', '·', '▪':
			return '*'
		case '○':
			return 'o'
		case '↑', '▲':
			return '^'
		case '↓', '▼', '▾':
			return 'v'
		case '→', '▶', '›', '▸':
			return '>'
		case '←', '◀', '‹':
			return '<'
		case '…':
			return '.'
		case '✓', '✔':
			return '+'
		case '×', '✗':
			return 'x'
		}

		return '?'
	}, content)
}

// PanelContentSize excludes the border and horizontal content gutters.
func PanelContentSize(size Size) Size {
	padding := min(theme.InlinePadding, max(0, (size.Width-3)/2))
	return Size{Width: max(0, size.Width-2*theme.BorderCells-2*padding), Height: max(0, size.Height-2*theme.BorderCells)}
}

func Panel(title, content string, size Size, t theme.Theme, focused bool) string {
	if size.Width < 4 || size.Height < 3 {
		return Fit(content, size)
	}

	frame := t.FrameBorder()
	border := t.BorderStyle(focused)
	label := ansi.Truncate(" "+title+" ", size.Width-2, "")

	labelStyle := t.Subtle()
	if focused {
		labelStyle = t.Heading()
	}

	top := border.Render(frame.TopLeft) + labelStyle.Render(label) + border.Render(strings.Repeat(frame.Top, max(0, size.Width-2-lipgloss.Width(label)))+frame.TopRight)
	padding := strings.Repeat(" ", min(theme.InlinePadding, max(0, (size.Width-3)/2)))

	inside := strings.Split(Fit(content, PanelContentSize(size)), "\n")
	for i := range inside {
		inside[i] = border.Render(frame.Left) + padding + inside[i] + padding + border.Render(frame.Right)
	}

	bottom := border.Render(frame.BottomLeft + strings.Repeat(frame.Bottom, size.Width-2) + frame.BottomRight)

	return strings.Join(append(append([]string{top}, inside...), bottom), "\n")
}

func Row(parts ...string) string    { return lipgloss.JoinHorizontal(lipgloss.Top, parts...) }
func Column(parts ...string) string { return lipgloss.JoinVertical(lipgloss.Left, parts...) }

// Split allocates a horizontal pair only when both minimums fit.
func Split(width, gap, leftMin, rightMin int, ratio float64) (left, right int, sideBySide bool) {
	if width < leftMin+rightMin+gap {
		return width, width, false
	}
	left = max(leftMin, min(int(float64(width-gap)*ratio), width-gap-rightMin))
	return left, width - gap - left, true
}
