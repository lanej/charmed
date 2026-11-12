package diagram

import (
	"image/color"
	"strings"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

func (d drawing) render(active string, steps []int, t theme.Theme) string {
	screen := uv.NewScreenBuffer(d.width, d.height)
	screen.Method = ansi.GraphemeWidth
	masks := d.routes()
	ascii := t.ASCII || t.Mode == theme.Mono
	put := func(x, y int, glyph rune, style uv.Style) {
		cell := uv.NewCell(ansi.GraphemeWidth, string(glyph))
		cell.Style = style
		screen.SetCell(x, y, cell)
	}
	edgeStyle := uv.Style{Fg: t.DimForeground}
	traced := make(map[int]bool, len(steps))
	visited := make(map[string]bool)
	latest := -1

	for _, step := range steps {
		traced[step] = true
		latest = step
	}

	for y, row := range masks {
		for x, mask := range row {
			if mask != 0 {
				put(x, y, lineGlyph(mask, ascii), edgeStyle)
			}
		}
	}

	for _, step := range steps {
		edge := d.edges[step]

		style := uv.Style{Fg: t.Foreground}
		if step == latest {
			style.Fg = t.Accent
			style.Attrs = uv.AttrBold
		}

		edge.walk(func(at point, horizontal bool) {
			// WHY: At a crossing the visible stroke belongs to the horizontal edge.
			if masks[at.y][at.x] == north|south|east|west && !horizontal {
				return
			}

			glyph := lineGlyph(masks[at.y][at.x], ascii)
			if step == latest {
				glyph = traceGlyph(masks[at.y][at.x], ascii)
			}

			put(at.x, at.y, glyph, style)
		})

		for _, node := range d.nodes {
			first, last := edge.points[0], edge.points[len(edge.points)-1]
			if onBox(node, first) || onBox(node, last) {
				visited[node.state.ID] = true
			}
		}
	}

	// WHY: A crossing is not a state. Interrupt the underpassing vertical edge
	// with a dotted continuation instead of drawing a junction outside a node.
	for y, row := range masks {
		for x, mask := range row {
			if mask == north|south|east|west {
				for _, neighbor := range []int{y - 1, y + 1} {
					if neighbor >= 0 && neighbor < len(masks) && masks[neighbor][x] == north|south {
						glyph := '┊'
						if ascii {
							glyph = ':'
						}

						style := edgeStyle
						if cell := screen.CellAt(x, neighbor); cell != nil {
							style = cell.Style
						}

						put(x, neighbor, glyph, style)
					}
				}
			}
		}
	}

	for _, node := range d.nodes {
		isActive := node.state.ID == active

		semantic := nodeColor(node.state, isActive, t)
		if !isActive && visited[node.state.ID] {
			semantic = visitedColor(node.state, t)
		}

		border := uv.Style{Fg: semantic}
		if isActive {
			border.Attrs = uv.AttrBold
		}

		for y := 0; y < node.h; y++ {
			for x := 0; x < node.w; x++ {
				mask := uint8(0)

				if y == 0 || y == node.h-1 {
					if x > 0 {
						mask |= west
					}

					if x < node.w-1 {
						mask |= east
					}
				}

				if x == 0 || x == node.w-1 {
					if y > 0 {
						mask |= north
					}

					if y < node.h-1 {
						mask |= south
					}
				}

				if mask != 0 {
					put(node.x+x, node.y+y, nodeGlyph(mask|masks[node.y+y][node.x+x], node.state, isActive, ascii), border)
				}
			}
		}

		label := t.Text().Foreground(semantic)
		if isActive {
			label = label.Bold(true)
		}

		drawText(screen, node.x+2, node.y+1, label.Render(node.label))
		drawText(screen, node.x+2, node.y+2, t.Text().Foreground(semantic).Render(node.markers))
	}

	for i, edge := range d.edges {
		style := edgeStyle
		if traced[i] {
			style.Fg = t.Foreground
		}

		if i == latest {
			style.Fg = t.Accent
			style.Attrs = uv.AttrBold
		}

		end, previous := edge.points[len(edge.points)-1], edge.points[len(edge.points)-2]

		arrow, at := '▶', end
		if end.x > previous.x {
			at.x--

			if ascii {
				arrow = '>'
			}
		}

		if end.x < previous.x {
			at.x++

			arrow = '◀'
			if ascii {
				arrow = '<'
			}
		}

		if end.y > previous.y {
			at.y--

			arrow = '▼'
			if ascii {
				arrow = 'v'
			}
		}

		if end.y < previous.y {
			at.y++

			arrow = '▲'
			if ascii {
				arrow = '^'
			}
		}

		put(at.x, at.y, arrow, style)

		if edge.label != "" {
			drawText(screen, edge.labelAt.x-1, edge.labelAt.y, t.Text().Foreground(style.Fg).Bold(i == latest).Render(" "+edge.label+" "))
		}
	}

	return strings.TrimRight(screen.Render(), "\n")
}

func drawText(screen uv.ScreenBuffer, x, y int, value string) {
	if value == "" {
		return
	}

	uv.NewStyledString(value).Draw(screen, uv.Rect(x, y, ansi.StringWidth(value), 1))
}

func nodeColor(s State, active bool, t theme.Theme) color.Color {
	switch s.Outcome {
	case OutcomeNone:
		if active {
			return t.Accent
		}

		return t.DimForeground
	case OutcomeSuccess:
		if active {
			return t.Success
		}

		return t.DimSuccess
	case OutcomeError:
		if active {
			return t.Danger
		}

		return t.DimDanger
	default:
		return t.Foreground
	}
}

func nodeGlyph(mask uint8, s State, active, ascii bool) rune {
	if s.Final {
		if ascii {
			if mask == east|west {
				return '='
			}

			if mask&(north|south) != 0 && mask&(east|west) != 0 {
				return '#'
			}

			return lineGlyph(mask, true)
		}

		glyphs := [16]rune{' ', '║', '║', '║', '═', '╚', '╔', '╟', '═', '╝', '╗', '╢', '═', '╧', '╤', '╪'}

		return glyphs[mask]
	}

	if s.Initial {
		switch mask {
		case south | east:
			if ascii {
				return '('
			}

			return '╭'
		case south | west:
			if ascii {
				return ')'
			}

			return '╮'
		case north | east:
			if ascii {
				return '('
			}

			return '╰'
		case north | west:
			if ascii {
				return ')'
			}

			return '╯'
		}
	}

	if active && !ascii {
		glyphs := [16]rune{' ', '┃', '┃', '┃', '━', '┗', '┏', '┣', '━', '┛', '┓', '┫', '━', '┻', '┳', '╋'}
		return glyphs[mask]
	}

	return lineGlyph(mask, ascii)
}

func (e edgePath) walk(visit func(point, bool)) {
	for i := 1; i < len(e.points); i++ {
		a, b := e.points[i-1], e.points[i]

		dx, dy := 0, 0
		if b.x > a.x {
			dx = 1
		}

		if b.x < a.x {
			dx = -1
		}

		if b.y > a.y {
			dy = 1
		}

		if b.y < a.y {
			dy = -1
		}

		visit(a, dy == 0)

		for a != b {
			a.x += dx
			a.y += dy
			visit(a, dy == 0)
		}
	}
}

func onBox(n nodeBox, p point) bool {
	return p.x >= n.x && p.x < n.x+n.w && p.y >= n.y && p.y < n.y+n.h
}

func visitedColor(s State, t theme.Theme) color.Color {
	switch s.Outcome {
	case OutcomeNone:
		return t.Foreground
	case OutcomeSuccess:
		return t.Success
	case OutcomeError:
		return t.Danger
	default:
		return t.Foreground
	}
}

func traceGlyph(mask uint8, ascii bool) rune {
	if ascii {
		if mask&(north|south) == 0 || mask == north|south|east|west {
			return '='
		}

		if mask&(east|west) == 0 {
			return '#'
		}

		return '+'
	}

	glyphs := [16]rune{' ', '┃', '┃', '┃', '━', '┗', '┏', '┣', '━', '┛', '┓', '┫', '━', '┻', '┳', '━'}

	return glyphs[mask]
}
