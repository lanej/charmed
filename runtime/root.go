// Package runtime adapts a composed scene to one Bubble Tea terminal session.
package runtime

import (
	"image/color"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
)

// Scene is a workbench/application boundary, not a required widget interface.
type Scene interface {
	Init() tea.Cmd
	Update(tea.Msg) tea.Cmd
	Resize(layout.Size)
	View() string
}

type StepMsg struct{}

type Root struct {
	Scene      Scene
	Background color.Color
	Size       layout.Size
	AltScreen  bool
	ASCII      bool
}

func (m *Root) Close() error {
	if c, ok := m.Scene.(interface{ Close() error }); ok {
		return c.Close()
	}
	return nil
}
func (m *Root) Init() tea.Cmd { m.Scene.Resize(m.Size); return m.Scene.Init() }
func (m *Root) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if k, ok := msg.(tea.KeyPressMsg); ok {
		if k.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if k.String() == "ctrl+n" {
			return m, m.Scene.Update(StepMsg{})
		}
	}
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.Size = layout.Size{Width: size.Width, Height: size.Height}
		m.Scene.Resize(m.Size)
		return m, nil
	}
	return m, m.Scene.Update(msg)
}
func (m *Root) View() tea.View {
	content := m.Scene.View()
	if m.ASCII {
		content = layout.ASCII(content)
	}
	v := tea.NewView(layout.Fit(content, m.Size))
	v.AltScreen = m.AltScreen
	v.BackgroundColor = m.Background
	return v
}
