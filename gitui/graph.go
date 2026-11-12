package gitui

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lanej/charmed/theme"
)

type graphRow struct {
	prefix string
	after  []string
}

type graphEdge struct{ from, to int }

// graph retains lanes for parents outside the window until explicit boundary rows.
func graph(history History, t theme.Theme) ([]graphRow, int, error) {
	positions := make(map[string]int, len(history.Commits))
	for i, commit := range history.Commits {
		if commit.ID == "" {
			return nil, 0, fmt.Errorf("commit %d has no ID", i)
		}

		if _, exists := positions[commit.ID]; exists {
			return nil, 0, fmt.Errorf("duplicate commit %q", commit.ID)
		}

		positions[commit.ID] = i
	}

	for i, commit := range history.Commits {
		seen := map[string]bool{}
		for _, parent := range commit.ParentIDs {
			if parent == "" || seen[parent] {
				return nil, 0, fmt.Errorf("commit %q has empty or duplicate parent", commit.ID)
			}

			seen[parent] = true
			if position, exists := positions[parent]; exists {
				if position <= i {
					return nil, 0, fmt.Errorf("parent %q must follow child %q", parent, commit.ID)
				}
			} else if !history.Truncated {
				return nil, 0, fmt.Errorf("parent %q outside complete history", parent)
			}
		}
	}

	var lanes []string

	rows := make([]graphRow, len(history.Commits))
	width := 1

	for i, commit := range history.Commits {
		lane := slices.Index(lanes, commit.ID)
		if lane < 0 {
			lanes, lane = assignLane(lanes, commit.ID)
		}

		marks := make([]string, len(lanes))
		for j, id := range lanes {
			marks[j] = "  "
			if id != "" {
				marks[j] = "│ "
			}
		}

		marks[lane] = "● "

		if t.ASCII {
			for j, id := range lanes {
				marks[j] = "  "
				if id != "" {
					marks[j] = "| "
				}
			}

			marks[lane] = "* "
		}

		rows[i].prefix = colorLanes(marks, 1, t)
		next := slices.Clone(lanes)

		next[lane] = ""
		for p, parent := range commit.ParentIDs {
			if slices.Contains(next, parent) {
				continue
			}

			if p == 0 {
				next[lane] = parent
			} else {
				next, _ = assignLane(next, parent)
			}
		}

		for len(next) > 0 && next[len(next)-1] == "" {
			next = next[:len(next)-1]
		}

		var edges []graphEdge

		for j, id := range lanes {
			if j == lane {
				for _, parent := range commit.ParentIDs {
					edges = append(edges, graphEdge{j * 2, slices.Index(next, parent) * 2})
				}
			} else if id != "" {
				edges = append(edges, graphEdge{j * 2, slices.Index(next, id) * 2})
			}
		}

		columns := max(len(lanes), len(next)) * 2
		width = max(width, columns)
		rows[i].after = routeEdges(edges, columns, t)
		lanes = next
	}

	if len(rows) > 0 {
		for lane, id := range lanes {
			if id == "" {
				continue
			}

			marks := make([]string, len(lanes))
			for j, pending := range lanes {
				marks[j] = "  "
				if pending != "" {
					marks[j] = "│ "
					if t.ASCII {
						marks[j] = "| "
					}
				}
			}

			marks[lane] = "~ "
			rows[len(rows)-1].after = append(rows[len(rows)-1].after, colorLanes(marks, 1, t)+clean(id)+" (history continues)")
			lanes[lane] = ""
		}
	}

	return rows, width, nil
}

func colorLanes(marks []string, cellsPerLane int, t theme.Theme) string {
	var out []string

	for i, mark := range marks {
		style := t.Text()
		if t.Mode != theme.Mono && len(t.Series) > 0 {
			style = style.Foreground(t.Series[(i/cellsPerLane)%len(t.Series)])
		}

		out = append(out, style.Render(mark))
	}

	return strings.Join(out, "")
}

// Stable lanes keep unrelated vertical paths separate. A crossing has no junction:
// ':' in ASCII and '╪' in Unicode distinguish it from connected '+' / '┼'.
func routeEdges(edges []graphEdge, columns int, t theme.Theme) []string {
	moves := 0

	for _, edge := range edges {
		if edge.from != edge.to {
			moves++
		}
	}

	if moves == 0 {
		return nil
	}

	height := moves + 2

	grid := make([][]uint8, height)
	for i := range grid {
		grid[i] = make([]uint8, columns)
	}

	route := 1
	endpoints := make(map[[2]int]bool)

	for _, edge := range edges {
		if edge.from == edge.to {
			for y := range height {
				grid[y][edge.from] |= 3
			}

			continue
		}

		for y := 0; y < route; y++ {
			grid[y][edge.from] |= 3
		}

		for y := route + 1; y < height; y++ {
			grid[y][edge.to] |= 3
		}

		left, right := min(edge.from, edge.to), max(edge.from, edge.to)
		for x := left; x < right; x++ {
			grid[route][x] |= 8
			grid[route][x+1] |= 4
		}

		endpoints[[2]int{route, edge.from}] = true
		endpoints[[2]int{route, edge.to}] = true
		grid[route][edge.from] |= 1
		grid[route][edge.to] |= 2
		route++
	}

	rows := make([]string, height)
	for y, cells := range grid {
		marks := make([]string, columns)

		for x, mask := range cells {
			glyph := " "

			if t.ASCII {
				switch mask {
				case 0:
				case 1, 2, 3:
					glyph = "|"
				case 4, 8, 12:
					glyph = "-"
				default:
					glyph = "+"
				}
			} else {
				glyphs := [...]string{" ", "│", "│", "│", "─", "┘", "┐", "┤", "─", "└", "┌", "├", "─", "┴", "┬", "┼"}
				glyph = glyphs[mask]
			}

			if mask == 15 && !endpoints[[2]int{y, x}] {
				glyph = "╪"
				if t.ASCII {
					glyph = ":"
				}
			}

			marks[x] = glyph
		}

		rows[y] = colorLanes(marks, 2, t)
	}

	return rows
}

// assignLane fills holes without moving pending parents; only trailing empty lanes collapse.
func assignLane(lanes []string, id string) ([]string, int) {
	if lane := slices.Index(lanes, ""); lane >= 0 {
		lanes[lane] = id
		return lanes, lane
	}

	return append(lanes, id), len(lanes)
}
