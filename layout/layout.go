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
	borderSpan := 2 * theme.BorderCells
	return Size{Width: max(0, size.Width-borderSpan-2*panelGutter(size.Width)), Height: max(0, size.Height-borderSpan)}
}

func panelGutter(width int) int {
	innerWidth := max(0, width-2*theme.BorderCells)
	// WHY: Tiny panels reduce gutters so at least one content cell can remain.
	return min(theme.InlinePadding, max(0, (innerWidth-1)/2))
}

func Panel(title, content string, size Size, t theme.Theme, focused bool) string {
	const minPanelWidth = 2*theme.BorderCells + 2

	const minPanelHeight = 2*theme.BorderCells + 1

	if size.Width < minPanelWidth || size.Height < minPanelHeight {
		return Fit(content, size)
	}

	innerWidth := size.Width - 2*theme.BorderCells
	frame := t.FrameBorder()
	border := t.BorderStyle(focused)
	label := ansi.Truncate(" "+title+" ", innerWidth, "")

	labelStyle := t.Subtle()
	if focused {
		labelStyle = t.Heading()
	}

	top := border.Render(frame.TopLeft) + labelStyle.Render(label) + border.Render(strings.Repeat(frame.Top, max(0, innerWidth-lipgloss.Width(label)))+frame.TopRight)
	padding := strings.Repeat(" ", panelGutter(size.Width))

	inside := strings.Split(Fit(content, PanelContentSize(size)), "\n")
	for i := range inside {
		inside[i] = border.Render(frame.Left) + padding + inside[i] + padding + border.Render(frame.Right)
	}

	bottom := border.Render(frame.BottomLeft + strings.Repeat(frame.Bottom, innerWidth) + frame.BottomRight)

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
