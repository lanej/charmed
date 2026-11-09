// Package diagram renders themed terminal diagrams from explicit graph data.
package diagram

import (
	"fmt"
	"strings"
	"unicode"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

// State names a node. Empty Label uses ID. Initial and Final are visible markers;
// at most one state may be initial, and final states cannot have outgoing edges.
type State struct {
	ID, Label      string
	Initial, Final bool
}

// Transition is a directed edge. Parallel transitions and self-loops are retained.
type Transition struct{ From, To, Label string }

// StateDiagram displays a graph; it does not execute transitions. Slice order
// controls node and routing order. Active is an optional state ID.
type StateDiagram struct {
	States      []State
	Transitions []Transition
	Active      string
}

type nodeLayout struct {
	state          State
	label, markers string
	x, width       int
	ports          []int
}

type edgeLayout struct {
	from, to int
	label    string
}

// View renders the complete graph or returns an error, never a clipped graph.
// Width is a maximum cell allocation. Height grows with the transition count.
// Nodes follow declaration order; every transition has separate ports and a
// routing row. Crossings are not junctions. Labels must be single-line text.
func (d StateDiagram) View(width int, t theme.Theme) (string, error) {
	nodes, edges, graphWidth, err := d.plan()
	if err != nil {
		return "", err
	}

	if len(nodes) == 0 {
		return "", nil
	}

	required := graphWidth

	for _, edge := range edges {
		if edge.label != "" {
			required = max(required, graphWidth+2+ansi.StringWidth(edge.label))
		}
	}

	if width < required {
		return "", fmt.Errorf("state diagram needs %d columns; available %d", required, width)
	}

	masks := route(edges, graphWidth)
	crossing := false

	targets := make([]bool, graphWidth)
	for _, edge := range edges {
		targets[edge.to] = true
	}

	for _, row := range masks {
		for _, mask := range row {
			if mask == north|south|east|west {
				crossing = true
			}
		}
	}

	const crossingNote = "Crossings do not join transitions."
	if crossing {
		required = max(required, len(crossingNote))
	}

	if width < required {
		return "", fmt.Errorf("state diagram needs %d columns; available %d", required, width)
	}

	ascii := t.ASCII || t.Mode == theme.Mono

	panels := make([]string, 0, len(nodes))
	for i, node := range nodes {
		if i > 0 {
			panels = append(panels, strings.Repeat(" ", 3))
		}

		panels = append(panels, renderNode(node, d.Active, t, ascii))
	}

	lines := []string{lipgloss.JoinHorizontal(lipgloss.Top, panels...)}

	for y, row := range masks {
		line := make([]rune, len(row))
		for x, mask := range row {
			glyph := lineGlyph(mask, ascii)
			if y == 0 && targets[x] {
				glyph = '↑'
				if ascii {
					glyph = '^'
				}
			}

			line[x] = glyph
		}

		rendered := t.Subtle().Render(string(line))

		if y >= 2 && y%2 == 0 {
			if edge := edges[(y-2)/2]; edge.label != "" {
				rendered += "  " + t.Text().Render(edge.label)
			}
		}

		lines = append(lines, rendered)
	}

	if crossing {
		lines = append(lines, t.Subtle().Render(crossingNote))
	}

	out := strings.Join(lines, "\n")
	if t.Mode == theme.Mono {
		out = ansi.Strip(out)
	}

	return out, nil
}

func (d StateDiagram) plan() ([]nodeLayout, []edgeLayout, int, error) {
	nodes := make([]nodeLayout, len(d.States))
	index := make(map[string]int, len(nodes))
	degree := make([]int, len(nodes))
	initial := false

	for i, state := range d.States {
		if strings.TrimSpace(state.ID) == "" || !singleLine(state.ID) {
			return nil, nil, 0, fmt.Errorf("state %d has an invalid ID", i)
		}

		if _, exists := index[state.ID]; exists {
			return nil, nil, 0, fmt.Errorf("duplicate state ID %q", state.ID)
		}

		if !singleLine(state.Label) {
			return nil, nil, 0, fmt.Errorf("state %q label must be single-line text", state.ID)
		}

		if state.Initial && initial {
			return nil, nil, 0, fmt.Errorf("multiple initial states")
		}

		initial = initial || state.Initial
		index[state.ID] = i

		label := state.Label
		if label == "" {
			label = state.ID
		}

		var markers []string
		if state.Initial {
			markers = append(markers, "initial")
		}

		if state.Final {
			markers = append(markers, "final")
		}

		if state.ID == d.Active {
			markers = append(markers, "active")
		}

		nodes[i] = nodeLayout{state: state, label: label, markers: strings.Join(markers, ", ")}
	}

	if d.Active != "" {
		if _, exists := index[d.Active]; !exists {
			return nil, nil, 0, fmt.Errorf("unknown active state %q", d.Active)
		}
	}

	for i, edge := range d.Transitions {
		from, fromOK := index[edge.From]

		to, toOK := index[edge.To]
		if !fromOK || !toOK {
			return nil, nil, 0, fmt.Errorf("transition %d references unknown state: %q -> %q", i, edge.From, edge.To)
		}

		if nodes[from].state.Final {
			return nil, nil, 0, fmt.Errorf("final state %q has outgoing transition %d", edge.From, i)
		}

		if !singleLine(edge.Label) {
			return nil, nil, 0, fmt.Errorf("transition %d label must be single-line text", i)
		}

		degree[from]++
		degree[to]++
	}

	total := 0

	for i := range nodes {
		node := &nodes[i]
		markerWidth := ansi.StringWidth(node.markers)

		if d.Active != "" {
			possible := []string{}
			if node.state.Initial {
				possible = append(possible, "initial")
			}

			if node.state.Final {
				possible = append(possible, "final")
			}

			possible = append(possible, "active")
			markerWidth = max(markerWidth, ansi.StringWidth(strings.Join(possible, ", ")))
		}

		node.width = max(7, ansi.StringWidth(node.label)+4, markerWidth+4, 2*degree[i]+3)
		node.x = total
		total += node.width + 3
	}

	edges := make([]edgeLayout, len(d.Transitions))
	for i, edge := range d.Transitions {
		from, to := index[edge.From], index[edge.To]
		edges[i] = edgeLayout{from: nodes[from].port(), to: nodes[to].port(), label: edge.Label}
	}

	if len(nodes) > 0 {
		total -= 3
	}

	return nodes, edges, total, nil
}

func (n *nodeLayout) port() int {
	x := n.x + 2 + 2*len(n.ports)
	n.ports = append(n.ports, x-n.x)

	return x
}

func singleLine(value string) bool {
	for _, r := range value {
		if unicode.IsControl(r) || r == '\u2028' || r == '\u2029' {
			return false
		}
	}

	return true
}

func renderNode(node nodeLayout, active string, t theme.Theme, ascii bool) string {
	h, v, tl, tr, bl, br, port := '─', '│', '┌', '┐', '└', '┘', '┬'
	if ascii {
		h, v, tl, tr, bl, br, port = '-', '|', '+', '+', '+', '+', '+'
	}

	border := t.Text().Foreground(t.Border)

	label := t.Text()
	if node.state.ID == active {
		border = border.Foreground(t.Accent)
		label = label.Foreground(t.Accent).Bold(true)
	}

	bottom := []rune(string(bl) + strings.Repeat(string(h), node.width-2) + string(br))
	for _, x := range node.ports {
		bottom[x] = port
	}

	row := func(value string, style lipgloss.Style) string {
		return border.Render(string(v)) + style.Render(" "+value+strings.Repeat(" ", node.width-3-ansi.StringWidth(value))) + border.Render(string(v))
	}

	return strings.Join([]string{
		border.Render(string(tl) + strings.Repeat(string(h), node.width-2) + string(tr)),
		row(node.label, label), row(node.markers, t.Subtle()), border.Render(string(bottom)),
	}, "\n")
}
