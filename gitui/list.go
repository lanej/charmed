package gitui

import (
	"fmt"
	"slices"
	"strings"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// Item uses a stable caller identity. After contains nonselectable continuation rows.
type Item struct {
	ID    string
	Text  string
	Group string
	After []string
}

// List owns selection and filtering. Native Viewport styles and Input customization
// survive SetItems and Resize; content, dimensions and scroll position are managed here.
// History search retains every graph row and navigates matching commits.
type List struct {
	EmptyText     string
	Viewport      viewport.Model
	Input         textinput.Model
	Theme         theme.Theme
	Component     string
	items         []Item
	cursor        int
	width, height int
	focused       bool
	searching     bool
	preserveRows  bool
	minWidth      int
	query         string
}

func NewList(component string, items []Item, width, height int, t theme.Theme) (*List, error) {
	m := &List{Component: component, Theme: t, EmptyText: "No items", Viewport: collection.NewViewport("", width, height), Input: textinput.New(), focused: true}
	m.Input.Prompt = "/ "

	m.Input.CharLimit = 256
	if err := m.SetItems(items); err != nil {
		return nil, err
	}

	m.Resize(layout.Size{Width: width, Height: height})

	return m, nil
}

// SetItems validates and copies input before replacing the current collection.
func (m *List) SetItems(items []Item) error {
	seen := make(map[string]bool, len(items))

	cloned := slices.Clone(items)
	for i, item := range cloned {
		if item.ID == "" || seen[item.ID] {
			return fmt.Errorf("%s: empty or duplicate item ID %q", m.Component, item.ID)
		}

		seen[item.ID] = true
		cloned[i].After = slices.Clone(item.After)
	}

	selected := m.SelectedID()
	m.items = cloned

	m.cursor = min(m.cursor, max(0, len(items)-1))
	for i, item := range items {
		if item.ID == selected {
			m.cursor = i
			break
		}
	}

	m.reconcile()

	return nil
}

func (m *List) SelectedID() string {
	if len(m.items) == 0 || !m.matches(m.items[m.cursor]) {
		return ""
	}

	return m.items[m.cursor].ID
}

func (m *List) Select(id string) bool {
	for i, item := range m.items {
		if item.ID == id {
			m.cursor = i
			m.render()

			return true
		}
	}

	return false
}

// SetFocused changes presentation without following selection or moving the viewport.
func (m *List) SetFocused(focused bool) {
	if m.focused == focused {
		return
	}

	offset := m.Viewport.YOffset()
	m.focused = focused
	m.render()
	m.Viewport.SetYOffset(offset)
}
func (m *List) Init() tea.Cmd { return nil }
func (m *List) Close() error  { return nil }
func (m *List) Resize(size layout.Size) {
	m.width, m.height = max(1, size.Width), max(1, size.Height)
	m.Viewport.SetWidth(m.width)
	m.Viewport.SetHeight(max(1, m.height-2))
	m.Input.SetWidth(max(1, m.width-4))
	m.render()
}

func (m *List) matches(item Item) bool {
	return strings.Contains(strings.ToLower(ansi.Strip(item.Text)), strings.ToLower(m.query))
}

func (m *List) reconcile() {
	if len(m.items) > 0 && !m.matches(m.items[m.cursor]) {
		for i, item := range m.items {
			if m.matches(item) {
				m.cursor = i
				break
			}
		}
	}

	m.render()
}

func (m *List) move(delta int) {
	for i := m.cursor + delta; i >= 0 && i < len(m.items); i += delta {
		if m.matches(m.items[i]) {
			m.cursor = i
			return
		}
	}
}

func (m *List) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		if m.searching {
			return m.updateInput(msg)
		}

		return nil
	}

	if m.searching {
		switch key.String() {
		case "enter":
			m.searching = false
			m.Input.Blur()
		case "esc":
			m.searching = false
			m.Input.Blur()
			m.Input.SetValue("")
			m.query = ""
		default:
			return m.updateInput(msg)
		}
	} else {
		switch key.String() {
		case "/":
			m.searching = true
			return m.Input.Focus()
		case "esc":
			m.query = ""
			m.Input.SetValue("")
		case "ctrl+e":
			m.Viewport.ScrollDown(1)
			return nil
		case "ctrl+y":
			m.Viewport.ScrollUp(1)
			return nil
		case "down", "j":
			m.move(1)
		case "up", "k":
			m.move(-1)
		case "home", "g":
			m.cursor = 0
			m.reconcile()
		case "end", "G":
			m.cursor = max(0, len(m.items)-1)
			if len(m.items) > 0 && !m.matches(m.items[m.cursor]) {
				m.move(-1)
			}
		case "pgdown", "ctrl+d":
			for range max(1, m.height/2) {
				m.move(1)
			}
		case "pgup", "ctrl+u":
			for range max(1, m.height/2) {
				m.move(-1)
			}
		case "enter":
			id := m.SelectedID()

			component := m.Component
			if id != "" && m.matches(m.items[m.cursor]) {
				return func() tea.Msg { return interaction.Selected{Component: component, ID: id} }
			}
		}
	}

	m.render()

	return nil
}

func (m *List) render() {
	if m.width == 0 {
		return
	}

	if m.width < m.minWidth {
		m.Viewport.SetContent("More width needed for graph topology.")
		return
	}

	var lines []string

	selectedLine := 0
	matches := 0
	group := ""

	for i, item := range m.items {
		match := m.matches(item)
		if match {
			matches++
		}

		if !match && !m.preserveRows {
			continue
		}

		if item.Group != "" && item.Group != group {
			group = item.Group
			lines = append(lines, m.Theme.Heading().Render(clean(group)))
		}

		marker := " "

		if i == m.cursor && match {
			selectedLine = len(lines)

			if m.focused {
				marker = ">"
			}
		}

		row := marker + " " + item.Text

		lines = append(lines, ansi.Truncate(row, m.width, "…"))
		for _, line := range item.After {
			lines = append(lines, ansi.Truncate("  "+line, m.width, "…"))
		}
	}

	if matches == 0 {
		if m.preserveRows && len(lines) > 0 {
			lines = append([]string{"No matching commits"}, lines...)
		} else {
			message := "No matching items"
			if len(m.items) == 0 && m.query == "" {
				message = m.EmptyText
			}

			lines = []string{message}
		}
	}

	m.Viewport.SetContent(strings.Join(lines, "\n"))

	if selectedLine < m.Viewport.YOffset() {
		m.Viewport.SetYOffset(selectedLine)
	}

	if selectedLine >= m.Viewport.YOffset()+m.Viewport.Height() {
		m.Viewport.SetYOffset(selectedLine - m.Viewport.Height() + 1)
	}
}

func (m *List) View() string {
	footer := m.Theme.Subtle().Render("↑/↓ j/k · / search · Enter select")
	if m.Theme.ASCII {
		footer = m.Theme.Subtle().Render("up/down j/k · / search · Enter select")
	}

	if m.searching {
		footer = m.Input.View()
	} else if m.query != "" {
		footer = "search: " + clean(m.query) + " · Esc clear"
	}

	return portable(layout.Fit(m.Viewport.View()+"\n"+footer, layout.Size{Width: m.width, Height: m.height}), m.Theme)
}

func (m *List) updateInput(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	m.Input, cmd = m.Input.Update(msg)
	m.query = m.Input.Value()
	m.reconcile()

	return cmd
}
