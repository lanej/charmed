package story

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/diagram"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func stateDiagramStory(e Environment) (runtime.Scene, error) {
	graph := diagram.StateDiagram{
		States:      []diagram.State{{ID: "idle", Label: "Idle", Initial: true}, {ID: "running", Label: "Running"}, {ID: "done", Label: "Done", Final: true, Outcome: diagram.OutcomeSuccess}, {ID: "failed", Label: "Failed", Final: true, Outcome: diagram.OutcomeError}},
		Transitions: []diagram.Transition{{From: "idle", To: "running", Label: "start"}, {From: "running", To: "running", Label: "retry"}, {From: "running", To: "done", Label: "finish"}, {From: "running", To: "idle", Label: "cancel"}, {From: "running", To: "failed", Label: "fail"}},
	}
	path := []int{0, 1, 3, 0, 2}
	position, traceMode := len(path), 0

	switch e.Config.TraceWindow {
	case "last-3":
		traceMode = 1
	case "all":
		traceMode = 2
	}

	return &scene{
		update: func(msg tea.Msg) tea.Cmd {
			if key, ok := msg.(tea.KeyPressMsg); ok {
				switch key.String() {
				case "right":
					position = (position + 1) % (len(path) + 1)
				case "left":
					position = (position + len(path)) % (len(path) + 1)
				case "h":
					traceMode = (traceMode + 1) % 3
				}
			}

			return nil
		},
		view: func(size layout.Size) string {
			graph.Active = "idle"
			if position > 0 {
				graph.Active = graph.Transitions[path[position-1]].To
			}

			graph.Trace = diagram.TransitionTrace{Steps: path[:position]}
			mode := "last"

			switch traceMode {
			case 0:
			case 1:
				graph.Trace.Last = 3
				mode = "last 3"
			case 2:
				graph.Trace.All = true
				mode = "all"
			}

			out, err := graph.View(size.Width, e.Config.Palette())
			if err != nil {
				return title("State transitions", e) + e.Config.Palette().Subtle().Render(ansi.Wrap(err.Error(), size.Width, ""))
			}

			if strings.Count(out, "\n")+5 > size.Height {
				return "State diagram needs more rows; enlarge this region."
			}

			return title("State transitions", e) + out + footer("Left/Right replay; H: "+mode, e)
		},
	}, nil
}

func stateDiagramConfig() Config {
	c := DefaultConfig()
	c.Columns, c.Rows = 120, 40
	c.TraceWindow = "last"

	return c
}
