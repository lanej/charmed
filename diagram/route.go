package diagram

const (
	north uint8 = 1 << iota
	south
	east
	west
)

func route(edges []edgeLayout, width int) [][]uint8 {
	if len(edges) == 0 {
		return nil
	}

	rows := make([][]uint8, 3+2*(len(edges)-1))
	for i := range rows {
		rows[i] = make([]uint8, width)
	}

	for i, edge := range edges {
		lane := 2 + 2*i

		for _, x := range []int{edge.from, edge.to} {
			rows[0][x] |= north
			for y := 0; y < lane; y++ {
				rows[y][x] |= south
				rows[y+1][x] |= north
			}
		}

		for x := min(edge.from, edge.to); x < max(edge.from, edge.to); x++ {
			rows[lane][x] |= east
			rows[lane][x+1] |= west
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

		return '+'
	}

	glyphs := [16]rune{' ', '│', '│', '│', '─', '└', '┌', '├', '─', '┘', '┐', '┤', '─', '┴', '┬', '┼'}

	return glyphs[mask]
}
