package snippets

import (
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleDiagramTraceReplay() (string, error) {
	graph := diagram.StateDiagram{
		States:      []diagram.State{{ID: "idle", Initial: true}, {ID: "running"}, {ID: "done", Final: true}},
		Transitions: []diagram.Transition{{From: "idle", To: "running", Label: "start"}, {From: "running", To: "idle", Label: "cancel"}, {From: "running", To: "done", Label: "finish"}},
		Trace:       diagram.TransitionTrace{Steps: []int{0, 1, 0, 2}},
	}

	replay, err := diagram.NewTraceReplay(graph, layout.Size{Width: 80, Height: 24}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	replay.SetPosition(3)
	// Repeated edge zero is selected by occurrence three, independently of edge index.
	// Left/Right stop at either end. Home/End jump; no clock runs.
	return replay.View()
}
