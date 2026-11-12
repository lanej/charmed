// Package navigation contains keyboard-driven navigation components.
package navigation

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// Tab IDs should be unique and stable across item updates.
type Tab struct{ ID, Label string }

// Tabs activates on Left/Right or Home/End while focused. Tab and Shift+Tab
// belong to the parent's focus scope. Selection changes emit interaction.Selected.
type Tabs struct {
	items   []Tab
	index   int
	width   int
	focused bool
	Theme   theme.Theme
}

func NewTabs(items []Tab, t theme.Theme) Tabs {
	m := Tabs{Theme: t}
	m.SetItems(items)

	return m
}

// SetItems copies items and preserves the selected ID, falling back to the first
// item when that ID disappears. Width, focus and theme are retained.
func (m *Tabs) SetItems(items []Tab) {
	selected, hadSelection := m.Selected()

	m.items = append([]Tab(nil), items...)

	m.index = 0
	if hadSelection {
		m.Select(selected.ID)
	}
}

func (m *Tabs) Select(id string) bool {
	for i, item := range m.items {
		if item.ID == id {
			m.index = i
			return true
		}
	}

	return false
}

func (m *Tabs) Selected() (Tab, bool) {
	if len(m.items) == 0 {
		return Tab{}, false
	}

	return m.items[m.index], true
}

func (m *Tabs) SetWidth(width int) { m.width = max(0, width) }
func (m *Tabs) Focus()             { m.focused = true }
func (m *Tabs) Blur()              { m.focused = false }
func (m *Tabs) Focused() bool      { return m.focused }

func (m *Tabs) Update(msg tea.Msg) tea.Cmd {
	k, ok := msg.(tea.KeyPressMsg)
	if !ok || !m.focused || len(m.items) == 0 {
		return nil
	}

	previous := m.index

	switch k.String() {
	case "left":
		m.index = (m.index + len(m.items) - 1) % len(m.items)
	case "right":
		m.index = (m.index + 1) % len(m.items)
	case "home":
		m.index = 0
	case "end":
		m.index = len(m.items) - 1
	default:
		return nil
	}

	if m.index == previous {
		return nil
	}

	selected := interaction.Selected{Component: "tabs", ID: m.items[m.index].ID}

	return func() tea.Msg { return selected }
}

// View occupies exactly one row. Overflow keeps the active tab visible and marks
// hidden neighbors with < and >; very small allocations clip the active label.
func (m *Tabs) View() string {
	if len(m.items) == 0 {
		return layout.Fit("", layout.Size{Width: m.width, Height: 1})
	}

	if m.width < 5 {
		return layout.Fit(m.cell(m.index), layout.Size{Width: m.width, Height: 1})
	}

	available := m.width - 2
	start, end := m.index, m.index+1

	used := min(available, ansi.StringWidth(m.cell(m.index)))
	for start > 0 {
		width := ansi.StringWidth(m.cell(start - 1))
		if used+width > available {
			break
		}

		start--
		used += width
	}

	for end < len(m.items) {
		width := ansi.StringWidth(m.cell(end))
		if used+width > available {
			break
		}

		end++
		used += width
	}

	var cells []string
	for i := start; i < end; i++ {
		cells = append(cells, m.cell(i))
	}

	left, right := " ", " "
	if start > 0 {
		left = "<"
	}

	if end < len(m.items) {
		right = ">"
	}

	return left + layout.Fit(strings.Join(cells, ""), layout.Size{Width: available, Height: 1}) + right
}

func (m *Tabs) cell(index int) string {
	label := strings.SplitN(strings.ReplaceAll(m.items[index].Label, "\r", ""), "\n", 2)[0]
	if index != m.index {
		return m.Theme.Control().Render(label)
	}

	style := m.Theme.Selection()
	if m.focused {
		style = m.Theme.Focus()
	}

	return style.Render(label)
}
