package diagram

import (
	"image/color"
	"strings"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

func (d drawing) render(active string, t theme.Theme) string {
	screen := uv.NewScreenBuffer(d.width, d.height)
	screen.Method = ansi.GraphemeWidth
	masks := d.routes()
	ascii := t.ASCII || t.Mode == theme.Mono
	put := func(x, y int, glyph rune, style uv.Style) {
		cell := uv.NewCell(ansi.GraphemeWidth, string(glyph))
		cell.Style = style
		screen.SetCell(x, y, cell)
	}
	edgeStyle := uv.Style{Fg: t.Muted}

	for y, row := range masks {
		for x, mask := range row {
			if mask != 0 {
				put(x, y, lineGlyph(mask, ascii), edgeStyle)
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

						put(x, neighbor, glyph, edgeStyle)
					}
				}
			}
		}
	}

	for _, node := range d.nodes {
		semantic := nodeColor(node.state, t)

		border := uv.Style{Fg: semantic}
		if node.state.ID == active {
			border.Fg = t.Accent
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
					put(node.x+x, node.y+y, nodeGlyph(mask|masks[node.y+y][node.x+x], node.state, ascii), border)
				}
			}
		}

		label := t.Text()
		if node.state.Outcome != OutcomeNone {
			label = label.Foreground(semantic)
		}

		if node.state.ID == active {
			label = label.Bold(true)
			if node.state.Outcome == OutcomeNone {
				label = label.Foreground(t.Accent)
			}
		}

		drawText(screen, node.x+2, node.y+1, label.Render(node.label))
		drawText(screen, node.x+2, node.y+2, t.Text().Foreground(semantic).Render(node.markers))
	}

	for _, edge := range d.edges {
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

		put(at.x, at.y, arrow, edgeStyle)

		if edge.label != "" {
			drawText(screen, edge.labelAt.x-1, edge.labelAt.y, t.Text().Render(" "+edge.label+" "))
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

func nodeColor(s State, t theme.Theme) color.Color {
	switch s.Outcome {
	case OutcomeNone:
		return t.Border
	case OutcomeSuccess:
		return t.Success
	case OutcomeError:
		return t.Danger
	default:
		return t.Foreground
	}
}

func nodeGlyph(mask uint8, s State, ascii bool) rune {
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

	return lineGlyph(mask, ascii)
}
