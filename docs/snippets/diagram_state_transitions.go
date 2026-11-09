package snippets

import (
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/theme"
)

func ExampleDiagramStateTransitions() (string, error) {
	palette := theme.New(theme.Dark)
	graph := diagram.StateDiagram{
		States: []diagram.State{
			{ID: "idle", Label: "Idle", Initial: true},
			{ID: "running", Label: "Running"},
			{ID: "done", Label: "Done", Final: true, Outcome: diagram.OutcomeSuccess},
			{ID: "failed", Label: "Failed", Final: true, Outcome: diagram.OutcomeError},
		},
		Transitions: []diagram.Transition{
			{From: "idle", To: "running", Label: "start"},
			{From: "running", To: "running", Label: "retry"},
			{From: "running", To: "done", Label: "finish"},
			{From: "running", To: "idle", Label: "cancel"},
			{From: "running", To: "failed", Label: "fail"},
		},
		Active: "done",
		Trace: diagram.TransitionTrace{
			Steps: []int{0, 1, 3, 0, 2}, // start, retry, cancel, start, finish.
			Last:  1,                    // Last N steps; zero defaults to one. All: true shows the whole path.
		},
	}

	rendered, err := graph.View(64, palette)
	if err != nil {
		return "", err
	}
	// Update Active and Trace.Steps when application state changes, then render again.
	return rendered, nil
}
