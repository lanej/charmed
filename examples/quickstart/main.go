// Command quickstart runs an editable Charmed text field in Bubble Tea.
package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

type model struct{ field input.Text }

func (m *model) Init() tea.Cmd { return m.field.Focus() }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "ctrl+c" {
		return m, tea.Quit
	}

	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.field.Resize(max(1, size.Width-4))
		return m, nil
	}

	var cmd tea.Cmd

	m.field, cmd = m.field.Update(msg)

	return m, cmd
}

func (m *model) View() tea.View { return tea.NewView("Environment\n" + m.field.View()) }

func main() {
	field := input.NewText(input.TextOptions{Value: "production", Width: 36, Theme: theme.New(theme.Dark)})
	app := &model{field: field}
	// The field model and command stay owned by the application's event loop.
	if _, err := tea.NewProgram(app).Run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}
}
