// Package diagram renders themed terminal diagrams from explicit graph data.
package diagram

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
)

// StateOutcome optionally gives a state semantic success or error styling.
// Outcomes do not imply Final or execute transitions.
type StateOutcome uint8

const (
	OutcomeNone StateOutcome = iota
	OutcomeSuccess
	OutcomeError
)

// State names a node. Empty Label uses ID. Initial and Final are visible markers;
// at most one state may be initial, and final states cannot have outgoing edges.
type State struct {
	ID, Label      string
	Initial, Final bool
	Outcome        StateOutcome
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

// View renders labeled nodes and directed edges without a legend. It prefers
// a horizontal layout, then stacks nodes vertically when that fits the width.
// The complete graph is returned or an error reports the minimum required width.
// Labels are single-line plain text; height depends on the chosen layout.
func (d StateDiagram) View(width int, t theme.Theme) (string, error) {
	nodes, edges, err := d.validate()
	if err != nil {
		return "", err
	}

	if len(nodes) == 0 {
		return "", nil
	}

	horizontal := planDiagram(nodes, edges, false)

	drawing := horizontal
	if drawing.width > width {
		vertical := planDiagram(nodes, edges, true)
		if vertical.width > width {
			return "", fmt.Errorf("state diagram needs %d columns; available %d", min(horizontal.width, vertical.width), width)
		}

		drawing = vertical
	}

	out := drawing.render(d.Active, t)
	if t.Mode == theme.Mono {
		out = ansi.Strip(out)
	}

	return out, nil
}

func (d StateDiagram) validate() ([]nodeBox, []link, error) {
	nodes := make([]nodeBox, len(d.States))
	index := make(map[string]int, len(nodes))
	initial := false

	for i, state := range d.States {
		if strings.TrimSpace(state.ID) == "" || !singleLine(state.ID) {
			return nil, nil, fmt.Errorf("state %d has an invalid ID", i)
		}

		if _, exists := index[state.ID]; exists {
			return nil, nil, fmt.Errorf("duplicate state ID %q", state.ID)
		}

		if !singleLine(state.Label) {
			return nil, nil, fmt.Errorf("state %q label must be single-line text", state.ID)
		}

		if state.Initial && initial {
			return nil, nil, fmt.Errorf("multiple initial states")
		}

		initial = initial || state.Initial
		index[state.ID] = i

		label := state.Label
		if label == "" {
			label = state.ID
		}

		switch state.Outcome {
		case OutcomeNone, OutcomeSuccess, OutcomeError:
		default:
			return nil, nil, fmt.Errorf("state %q has invalid outcome %d", state.ID, state.Outcome)
		}

		nodes[i] = nodeBox{state: state, label: label, markers: state.markers(state.ID == d.Active)}
	}

	if d.Active != "" {
		if _, exists := index[d.Active]; !exists {
			return nil, nil, fmt.Errorf("unknown active state %q", d.Active)
		}
	}

	for i, edge := range d.Transitions {
		from, fromOK := index[edge.From]

		_, toOK := index[edge.To]
		if !fromOK || !toOK {
			return nil, nil, fmt.Errorf("transition %d references unknown state: %q -> %q", i, edge.From, edge.To)
		}

		if nodes[from].state.Final {
			return nil, nil, fmt.Errorf("final state %q has outgoing transition %d", edge.From, i)
		}

		if !singleLine(edge.Label) {
			return nil, nil, fmt.Errorf("transition %d label must be single-line text", i)
		}
	}

	edges := make([]link, len(d.Transitions))
	for i, edge := range d.Transitions {
		edges[i] = link{from: index[edge.From], to: index[edge.To], label: edge.Label}
	}

	return nodes, edges, nil
}

func singleLine(value string) bool {
	for _, r := range value {
		if unicode.IsControl(r) || r == '\u2028' || r == '\u2029' {
			return false
		}
	}

	return true
}

func (s State) markers(active bool) string {
	var parts []string
	if s.Initial {
		parts = append(parts, "initial")
	}

	if s.Final && s.Outcome == OutcomeNone {
		parts = append(parts, "final")
	}

	switch s.Outcome {
	case OutcomeNone:
	case OutcomeSuccess:
		parts = append(parts, "success")
	case OutcomeError:
		parts = append(parts, "error")
	default:
	}

	if active {
		parts = append(parts, "active")
	}

	return strings.Join(parts, ", ")
}
