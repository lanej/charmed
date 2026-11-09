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
	active := 1

	return &scene{
		update: func(msg tea.Msg) tea.Cmd {
			if key, ok := msg.(tea.KeyPressMsg); ok {
				switch key.String() {
				case "right":
					active = (active + 1) % len(graph.States)
				case "left":
					active = (active + len(graph.States) - 1) % len(graph.States)
				}
			}

			return nil
		},
		view: func(size layout.Size) string {
			graph.Active = graph.States[active].ID

			out, err := graph.View(size.Width, e.Config.Palette())
			if err != nil {
				return title("State transitions", e) + e.Config.Palette().Subtle().Render(ansi.Wrap(err.Error(), size.Width, ""))
			}

			if strings.Count(out, "\n")+5 > size.Height {
				return "State diagram needs more rows; enlarge this region."
			}

			return title("State transitions", e) + out + footer("Left/Right changes the active marker", e)
		},
	}, nil
}

func stateDiagramConfig() Config {
	c := DefaultConfig()
	c.Columns, c.Rows = 120, 40

	return c
}
