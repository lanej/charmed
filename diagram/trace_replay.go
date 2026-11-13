package diagram

import (
	"fmt"
	"slices"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// TraceOccurrence identifies an occurrence, not merely an edge which may repeat.
type TraceOccurrence struct {
	Position, TransitionIndex int
	Transition                Transition
}

// TraceReplay inspects a complete supplied trace without executing transitions or clocks.
// Position zero precedes the first transition. Last and All control visible history.
type TraceReplay struct {
	Theme    theme.Theme
	Last     int
	All      bool
	graph    StateDiagram
	position int
	size     layout.Size
}

func NewTraceReplay(graph StateDiagram, size layout.Size, t theme.Theme) (*TraceReplay, error) {
	m := &TraceReplay{Theme: t, size: size, Last: graph.Trace.Last, All: graph.Trace.All}
	if err := m.SetDiagram(graph); err != nil {
		return nil, err
	}

	return m, nil
}

// SetDiagram replaces graph data and selects its final occurrence, retaining Last and All.
func (m *TraceReplay) SetDiagram(graph StateDiagram) error {
	if _, _, err := graph.validate(); err != nil {
		return err
	}

	if _, err := graph.traceSteps(); err != nil {
		return err
	}

	graph.States = slices.Clone(graph.States)
	graph.Transitions = slices.Clone(graph.Transitions)
	graph.Trace.Steps = slices.Clone(graph.Trace.Steps)
	m.graph, m.position = graph, len(graph.Trace.Steps)

	return nil
}

func (m *TraceReplay) SetPosition(position int) {
	m.position = max(0, min(position, len(m.graph.Trace.Steps)))
}
func (m *TraceReplay) Position() int           { return m.position }
func (m *TraceReplay) Resize(size layout.Size) { m.size = size }

func (m *TraceReplay) Selected() (TraceOccurrence, bool) {
	if m.position == 0 {
		return TraceOccurrence{}, false
	}

	index := m.graph.Trace.Steps[m.position-1]

	return TraceOccurrence{Position: m.position, TransitionIndex: index, Transition: m.graph.Transitions[index]}, true
}

func (m *TraceReplay) Update(msg tea.Msg) tea.Cmd {
	if k, ok := msg.(tea.KeyPressMsg); ok {
		switch k.String() {
		case "left", "h":
			m.SetPosition(m.position - 1)
		case "right", "l":
			m.SetPosition(m.position + 1)
		case "home":
			m.SetPosition(0)
		case "end":
			m.SetPosition(len(m.graph.Trace.Steps))
		}
	}

	return nil
}

func (m *TraceReplay) View() (string, error) {
	d := m.graph

	d.Trace = TransitionTrace{Steps: d.Trace.Steps[:m.position], Last: m.Last, All: m.All}
	if len(m.graph.Trace.Steps) > 0 {
		d.Active = d.Transitions[m.graph.Trace.Steps[0]].From
		if occurrence, ok := m.Selected(); ok {
			d.Active = occurrence.Transition.To
		}
	}

	if len(m.graph.Trace.Steps) == 0 && d.Active == "" {
		for _, state := range d.States {
			if state.Initial {
				d.Active = state.ID
				break
			}
		}
	}

	out, err := d.View(m.size.Width, m.Theme)
	if err != nil {
		return "", err
	}

	if lipgloss.Height(out) > m.size.Height {
		return "", fmt.Errorf("trace replay needs %d rows; available %d", lipgloss.Height(out), m.size.Height)
	}

	return out, nil
}
