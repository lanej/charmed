package diagram

const (
	north uint8 = 1 << iota
	south
	east
	west
)

func (d drawing) routes() [][]uint8 {
	rows := make([][]uint8, d.height)
	for i := range rows {
		rows[i] = make([]uint8, d.width)
	}

	for _, edge := range d.edges {
		for i := 1; i < len(edge.points); i++ {
			a, b := edge.points[i-1], edge.points[i]

			dx, dy := 0, 0
			if b.x > a.x {
				dx = 1
			} else if b.x < a.x {
				dx = -1
			}

			if b.y > a.y {
				dy = 1
			} else if b.y < a.y {
				dy = -1
			}

			for a != b {
				next := point{a.x + dx, a.y + dy}
				if dx == 1 {
					rows[a.y][a.x] |= east
					rows[next.y][next.x] |= west
				}

				if dx == -1 {
					rows[a.y][a.x] |= west
					rows[next.y][next.x] |= east
				}

				if dy == 1 {
					rows[a.y][a.x] |= south
					rows[next.y][next.x] |= north
				}

				if dy == -1 {
					rows[a.y][a.x] |= north
					rows[next.y][next.x] |= south
				}

				a = next
			}
		}
	}

	return rows
}

func lineGlyph(mask uint8, ascii bool) rune {
	if mask == 0 {
		return ' '
	}

	if ascii {
		if mask&(north|south) == 0 {
			return '-'
		}

		if mask&(east|west) == 0 {
			return '|'
		}

		if mask == north|south|east|west {
			return '-'
		}

		return '+'
	}

	glyphs := [16]rune{' ', '│', '│', '│', '─', '└', '┌', '├', '─', '┘', '┐', '┤', '─', '┴', '┬', '─'}

	return glyphs[mask]
}
