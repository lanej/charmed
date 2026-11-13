package story

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/tui"
)

func TestNestedLayoutRecipesPreserveColumnStructure(t *testing.T) {
	for _, withGrid := range []bool{false, true} {
		root, ids := nestedLayout(withGrid)

		screen, err := tui.NewScreen(ids...)
		if err != nil {
			t.Fatal(err)
		}

		cases, err := screen.CheckLayouts(root, layout.Size{Width: 100, Height: 20}, layout.Size{Width: 40, Height: 24}, layout.Size{Width: 40, Height: 16})
		if err != nil {
			t.Fatal(err)
		}

		for i, name := range []string{"columns", "stacked", "inspection"} {
			plan := cases[i].Plan
			if plan.Variant() != name {
				t.Fatalf("expected %s, got %s", name, plan.Variant())
			}

			summary, _ := plan.Bounds("summary")

			details, _ := plan.Bounds("details")
			if summary.Min.X != details.Min.X || summary.Max.Y >= details.Min.Y {
				t.Fatal("summary and evidence lost their vertical column")
			}

			hosts, visible := plan.Bounds("hosts")

			switch name {
			case "columns":
				if !visible || hosts.Max.X >= summary.Min.X || hosts.Min.Y != summary.Min.Y {
					t.Fatal("inspection is not beside the host list")
				}
			case "stacked":
				if !visible || hosts.Max.Y >= summary.Min.Y || hosts.Min.X != summary.Min.X {
					t.Fatal("narrow panes did not stack")
				}
			case "inspection":
				if visible || len(plan.Hidden()) != 1 || plan.Hidden()[0].ID != "hosts" || plan.Hidden()[0].Reason == "" {
					t.Fatal("compact recipe lost its explicit omission")
				}
			}

			if withGrid {
				cpu, _ := plan.Bounds("cpu")

				memory, _ := plan.Bounds("memory")
				if cpu.Min.Y != memory.Min.Y || cpu.Max.X >= memory.Min.X || cpu.Min.X != summary.Min.X || cpu.Min.Y <= summary.Max.Y || cpu.Max.Y >= details.Min.Y {
					t.Fatal("metrics grid escaped its inspection column")
				}
			}
		}
	}
}

func TestNestedLayoutExamplesResizeFocusAndInspectRegions(t *testing.T) {
	for _, withGrid := range []bool{false, true} {
		e := Environment{Config: DefaultConfig()}

		m, err := nestedLayoutStory(withGrid, e)
		if err != nil {
			t.Fatal(err)
		}

		m.Resize(layout.Size{Width: 100, Height: 20})
		original := m.View()
		m.Update(tea.KeyPressMsg{Code: tea.KeyTab})

		if m.View() == original {
			t.Fatal("Tab did not move visible focus")
		}

		for _, size := range []layout.Size{{Width: 40, Height: 24}, {Width: 40, Height: 16}, {Width: 100, Height: 20}} {
			m.Resize(size)

			frame := m.View()
			if lipgloss.Width(frame) != size.Width || lipgloss.Height(frame) != size.Height {
				t.Fatal("nested frame escaped allocation")
			}

			plain := ansi.Strip(frame)
			for _, label := range []string{"Selected host", "Probe evidence", "HTTP 200"} {
				if !strings.Contains(plain, label) {
					t.Fatalf("%+v lost %s\n%s", size, label, plain)
				}
			}

			if withGrid && (!strings.Contains(plain, "43%") || !strings.Contains(plain, "61%")) {
				t.Fatal("nested metrics missing")
			}
		}

		m.Update(tea.KeyPressMsg{Code: 'l', Mod: tea.ModCtrl})

		if !strings.Contains(ansi.Strip(m.View()), "summary (") {
			t.Fatal("region inspection unavailable")
		}
	}
}
