package collection

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (m *JSONViewer) Update(msg tea.Msg) tea.Cmd {
	if pressed, ok := msg.(tea.KeyPressMsg); ok && m.Interactive() {
		keys := m.Model.KeyMap
		switch {
		case key.Matches(pressed, m.Toggle):
			m.SetCollapsed(m.selected, !m.Collapsed(m.selected))
		case key.Matches(pressed, keys.Up):
			m.move(-1)
		case key.Matches(pressed, keys.Down):
			m.move(1)
		case key.Matches(pressed, keys.Right):
			if m.Collapsed(m.selected) {
				m.SetCollapsed(m.selected, false)
			} else if children := m.nodes[m.selected].children; len(children) > 0 {
				m.Select(children[0])
			}
		case key.Matches(pressed, keys.Left):
			if len(m.nodes[m.selected].children) > 0 && !m.Collapsed(m.selected) {
				m.SetCollapsed(m.selected, true)
			} else if parent := m.nodes[m.selected].parent; parent >= 0 {
				m.Select(parent)
			}
		default:
			return m.updateViewport(msg)
		}

		return nil
	}

	return m.updateViewport(msg)
}

// Paging and wheel input retain native line scrolling so long wrapped scalar
// values remain readable. Structural input reveals the selected field again.
func (m *JSONViewer) updateViewport(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	*m.Model, cmd = m.Model.Update(msg)

	return cmd
}
