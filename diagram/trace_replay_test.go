package diagram_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestTraceReplayUsesOccurrencesAndBoundedNavigation(t *testing.T) {
	graph := diagram.StateDiagram{States: []diagram.State{{ID: "a", Initial: true}, {ID: "b"}}, Transitions: []diagram.Transition{{From: "a", To: "b"}, {From: "b", To: "a"}}, Trace: diagram.TransitionTrace{Steps: []int{0, 1, 0}}}

	replay, err := diagram.NewTraceReplay(graph, layout.Size{Width: 80, Height: 30}, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	graph.Trace.Steps[2] = 1

	replay.Update(tea.KeyPressMsg{Code: tea.KeyRight})

	selected, ok := replay.Selected()
	if !ok || selected.Position != 3 || selected.TransitionIndex != 0 {
		t.Fatal("occurrence identity or end clamp lost")
	}

	replay.Update(tea.KeyPressMsg{Code: tea.KeyHome})
	replay.Update(tea.KeyPressMsg{Code: tea.KeyLeft})

	if _, ok := replay.Selected(); ok || replay.Position() != 0 {
		t.Fatal("replay wrapped before beginning")
	}

	if _, err := replay.View(); err != nil {
		t.Fatal(err)
	}

	replay.Update(tea.KeyPressMsg{Code: tea.KeyEnd})
	replay.Resize(layout.Size{Width: 1, Height: 30})

	if _, err := replay.View(); err == nil {
		t.Fatal("insufficient width silently clipped")
	}

	replay.Resize(layout.Size{Width: 80, Height: 1})

	if _, err := replay.View(); err == nil {
		t.Fatal("insufficient height silently clipped")
	}

	graph.Trace.Steps = []int{1, 1}
	if err := replay.SetDiagram(graph); err == nil || replay.Position() != 3 {
		t.Fatal("invalid replacement accepted or mutated replay")
	}
}

func TestEmptyTraceReplayShowsSuppliedInitialState(t *testing.T) {
	replay, err := diagram.NewTraceReplay(diagram.StateDiagram{States: []diagram.State{{ID: "idle", Initial: true}}}, layout.Size{Width: 40, Height: 8}, theme.New(theme.Mono))
	if err != nil {
		t.Fatal(err)
	}

	out, err := replay.View()
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(out, "active") {
		t.Fatal("empty trace lost initial active state")
	}

	if _, ok := replay.Selected(); ok {
		t.Fatal("empty trace fabricated an occurrence")
	}
}
