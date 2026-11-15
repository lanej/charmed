// Command operations is a copyable dashboard with deterministic observations.
package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/examples/operations/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

type model struct {
	dashboard *pattern.OperationalDashboard
	quit      key.Binding
}

func newModel() *model {
	t := theme.New(theme.Dark)
	m := &model{
		dashboard: pattern.NewOperationalDashboard(fixture.Snapshot("ready", t), layout.Size{Width: 100, Height: 30}, t),
		quit:      key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
	}
	m.dashboard.AppKeys = []key.Binding{m.quit}
	m.dashboard.Resize(layout.Size{Width: 100, Height: 30})

	return m
}

func (m *model) Init() tea.Cmd { return nil }
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if press, ok := msg.(tea.KeyPressMsg); ok && key.Matches(press, m.quit) {
		return m, tea.Quit
	}

	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.dashboard.Resize(layout.Size{Width: size.Width, Height: size.Height})
		return m, nil
	}

	return m, m.dashboard.Update(msg)
}

func (m *model) View() tea.View {
	view := tea.NewView(m.dashboard.View())
	view.AltScreen = true
	view.BackgroundColor = m.dashboard.Theme.Background

	return view
}

func main() {
	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}
}
