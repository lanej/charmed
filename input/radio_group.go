package input

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/theme"
)

// RadioOption carries stable identity and an optional explanation of unavailability.
type RadioOption struct {
	ID, Label      string
	Disabled       bool
	DisabledReason string
}

// RadioGroup owns exclusivity and internal focus. Arrows move focus; Enter or
// Space commits selection. Parents own traversal into and out of the group.
type RadioGroup struct {
	radios   []Radio
	cursor   int
	selected string
	focused  bool
}

func NewRadioGroup(options []RadioOption, selected string, t theme.Theme) (*RadioGroup, error) {
	m := &RadioGroup{}

	seen := make(map[string]bool, len(options))
	for _, o := range options {
		if o.ID == "" || seen[o.ID] {
			return nil, fmt.Errorf("empty or duplicate radio ID %q", o.ID)
		}

		seen[o.ID] = true
		r := NewRadio(o.ID, o.Label, false, t)
		r.Disabled, r.DisabledReason = o.Disabled, o.DisabledReason
		m.radios = append(m.radios, r)
	}

	if selected != "" && !m.Select(selected) {
		return nil, fmt.Errorf("unknown or disabled radio %q", selected)
	}

	if selected == "" {
		m.cursor = -1
		m.move(1)
	}

	return m, nil
}

// Select updates selection without emitting an activation command.
func (m *RadioGroup) Select(id string) bool {
	for i, r := range m.radios {
		if r.ID == id && !r.unavailable() {
			m.selected, m.cursor = id, i
			for j := range m.radios {
				m.radios[j].Selected = j == i
			}

			m.syncFocus()

			return true
		}
	}

	return false
}

func (m *RadioGroup) Selected() (string, bool) { return m.selected, m.selected != "" }
func (m *RadioGroup) Focus()                   { m.focused = true; m.syncFocus() }
func (m *RadioGroup) Blur()                    { m.focused = false; m.syncFocus() }
func (m *RadioGroup) Focused() bool            { return m.focused }

func (m *RadioGroup) syncFocus() {
	for i := range m.radios {
		m.radios[i].Blur()

		if m.focused && i == m.cursor {
			m.radios[i].Focus()
		}
	}
}

func (m *RadioGroup) move(delta int) {
	for i := m.cursor + delta; i >= 0 && i < len(m.radios); i += delta {
		if !m.radios[i].unavailable() {
			m.cursor = i
			break
		}
	}

	m.syncFocus()
}

func (m *RadioGroup) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	if k, ok := msg.(tea.KeyPressMsg); ok {
		switch k.String() {
		case "left", "up", "h", "k":
			m.move(-1)
			return nil
		case "right", "down", "l", "j":
			m.move(1)
			return nil
		}
	}

	if m.cursor < 0 || m.cursor >= len(m.radios) {
		return nil
	}

	var cmd tea.Cmd

	m.radios[m.cursor], cmd = m.radios[m.cursor].Update(msg)
	if cmd != nil {
		m.Select(m.radios[m.cursor].ID)
	}

	return cmd
}

func (m *RadioGroup) View() string {
	views := make([]string, len(m.radios))
	for i, r := range m.radios {
		views[i] = r.View()
	}

	return strings.Join(views, "\n")
}
