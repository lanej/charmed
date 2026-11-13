package story

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func stateDiagramStory(e Environment) (runtime.Scene, error) {
	graph := diagram.StateDiagram{
		States:      []diagram.State{{ID: "idle", Label: "Idle", Initial: true}, {ID: "running", Label: "Running"}, {ID: "done", Label: "Done", Final: true, Outcome: diagram.OutcomeSuccess}, {ID: "failed", Label: "Failed", Final: true, Outcome: diagram.OutcomeError}},
		Transitions: []diagram.Transition{{From: "idle", To: "running", Label: "start"}, {From: "running", To: "running", Label: "retry"}, {From: "running", To: "done", Label: "finish"}, {From: "running", To: "idle", Label: "cancel"}, {From: "running", To: "failed", Label: "fail"}},
	}
	graph.Trace = diagram.TransitionTrace{Steps: []int{0, 1, 3, 0, 2}}

	replay, err := diagram.NewTraceReplay(graph, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	mode := e.Config.TraceWindow
	if mode == "" {
		mode = "last"
	}

	group, err := input.NewRadioGroup([]input.RadioOption{{ID: "last", Label: "Last transition"}, {ID: "last-3", Label: "Last three"}, {ID: "all", Label: "Complete trace"}}, mode, e.Config.Palette())
	if err != nil {
		return nil, err
	}

	toolbar := false

	return &scene{
		update: func(msg tea.Msg) tea.Cmd {
			if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "tab" {
				toolbar = !toolbar
				if toolbar {
					group.Focus()
				} else {
					group.Blur()
				}

				return nil
			}

			if toolbar {
				return group.Update(msg)
			}

			return replay.Update(msg)
		},
		view: func(size layout.Size) string {
			mode, _ := group.Selected()
			replay.Last, replay.All = 0, false

			switch mode {
			case "last-3":
				replay.Last = 3
			case "all":
				replay.All = true
			}

			replay.Resize(layout.Size{Width: size.Width, Height: max(0, size.Height-8)})

			out, err := replay.View()
			if err != nil {
				return e.Config.Palette().Subtle().Render(ansi.Wrap(err.Error(), size.Width, ""))
			}

			return title("State transitions", e) + out + "\n" + group.View() + footer("Left/Right · Home/End · Tab trace window", e)
		},
	}, nil
}

func stateDiagramConfig() Config {
	c := DefaultConfig()
	c.Columns, c.Rows = 120, 40
	c.TraceWindow = "last"

	return c
}
