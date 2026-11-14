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
	"github.com/lanej/charmed/internal/rendercache"
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

type listViewKey struct {
	revision                uint64
	width, height, x, y     int
	outerWidth, outerHeight int
}

// List owns selection and filtering. Native Viewport styles and Input customization
// survive SetItems and Resize; content, dimensions and scroll position are managed here.
// History search retains every graph row and navigates matching commits.
type List struct {
	rowBuilds                           uint64
	viewCache                           rendercache.Cache[listViewKey]
	revision                            uint64
	updating, rowsDirty, preserveOffset bool

	ShowHelp      bool // NewList enables help; compositions can opt out.
	EmptyText     string
	Viewport      viewport.Model
	Input         textinput.Model
	Theme         theme.Theme
	Component     string
	items         []Item
	searchText    []string
	foldedQuery   string
	preparedRows  [][]string
	preparedWidth int
	cursor        int
	selectedLine  int
	width, height int
	focused       bool
	searching     bool
	preserveRows  bool
	minWidth      int
	query         string
}

func NewList(component string, items []Item, width, height int, t theme.Theme) (*List, error) {
	m := &List{ShowHelp: true, Component: component, Theme: t, EmptyText: "No items", Viewport: collection.NewViewport("", width, height), Input: textinput.New(), focused: true}
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

	m.searchText = make([]string, len(cloned))
	for i, item := range cloned {
		m.searchText[i] = strings.ToLower(ansi.Strip(item.Text))
	}

	m.preparedRows = nil

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
	if len(m.items) == 0 || !m.matches(m.cursor) {
		return ""
	}

	return m.items[m.cursor].ID
}

func (m *List) Select(id string) bool {
	for i, item := range m.items {
		if item.ID == id {
			m.cursor = i
			m.preserveOffset = false
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
	m.preserveOffset = m.updating
	m.render()
	m.Viewport.SetYOffset(offset)
}
func (m *List) Init() tea.Cmd { return nil }
func (m *List) Close() error  { return nil }
func (m *List) Resize(size layout.Size) {
	m.resize(size, true)
}

// ResizeViewport changes geometry without refreshing externally mutated row
// presentation. Use Resize to reapply native theme or presentation changes.
func (m *List) ResizeViewport(size layout.Size) { m.resize(size, false) }

func (m *List) resize(size layout.Size, refresh bool) {
	widthChanged := m.width != max(1, size.Width)
	m.width, m.height = max(1, size.Width), max(1, size.Height)
	m.Viewport.SetWidth(m.width)
	m.Viewport.SetHeight(max(1, m.height-m.footerHeight()))
	m.Input.SetWidth(max(1, m.width-4))

	if refresh || widthChanged {
		m.render()
	} else {
		m.followSelection()
	}
}

func (m *List) matches(index int) bool {
	if m.query == "" {
		return true
	}

	return strings.Contains(m.searchText[index], m.foldedQuery)
}

func (m *List) reconcile() {
	m.preserveOffset = false
	if len(m.items) > 0 && !m.matches(m.cursor) {
		for i := range m.items {
			if m.matches(i) {
				m.cursor = i
				break
			}
		}
	}

	m.render()
}

func (m *List) move(delta int) {
	for i := m.cursor + delta; i >= 0 && i < len(m.items); i += delta {
		if m.matches(i) {
			m.cursor = i
			return
		}
	}
}

func (m *List) Update(msg tea.Msg) tea.Cmd {
	m.updating = true
	defer func() { m.updating = false }()

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
			m.render()

			return m.Input.Focus()
		case "esc":
			m.query = ""
			m.Input.SetValue("")
		case "ctrl+e":
			m.flushRows()
			before := m.Viewport.YOffset()
			m.Viewport.ScrollDown(1)

			if before != m.Viewport.YOffset() {
				m.revision++
			}

			return nil
		case "ctrl+y":
			m.flushRows()
			before := m.Viewport.YOffset()
			m.Viewport.ScrollUp(1)

			if before != m.Viewport.YOffset() {
				m.revision++
			}

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
			if len(m.items) > 0 && !m.matches(m.cursor) {
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
			if id != "" && m.matches(m.cursor) {
				return func() tea.Msg { return interaction.Selected{Component: component, ID: id} }
			}
		default:
			return nil
		}
	}

	m.render()

	return nil
}

func (m *List) footerHeight() int {
	if m.ShowHelp || m.searching || m.query != "" {
		return 2
	}

	return 0
}

func (m *List) render() {
	m.revision++
	if m.updating {
		m.preserveOffset = false
		m.rowsDirty = true

		return
	}

	m.rowsDirty = false
	m.rowBuilds++

	m.Viewport.SetHeight(max(1, m.height-m.footerHeight()))

	if m.width == 0 {
		return
	}

	if m.width < m.minWidth {
		m.Viewport.SetContent("More width needed for graph topology.")
		return
	}

	if m.preparedRows == nil || m.preparedWidth != m.width {
		m.preparedRows = make([][]string, len(m.items))

		m.preparedWidth = m.width
		for i, item := range m.items {
			rows := []string{strings.Clone(ansi.Truncate("  "+item.Text, m.width, "…"))}
			for _, line := range item.After {
				rows = append(rows, strings.Clone(ansi.Truncate("  "+line, m.width, "…")))
			}

			m.preparedRows[i] = rows
		}
	}

	var lines []string

	selectedLine := 0
	matches := 0
	group := ""

	for i, item := range m.items {
		match := m.matches(i)
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
				marker = m.Theme.Glyph(theme.GlyphCollapsed)
			}
		}

		row := m.preparedRows[i][0]
		if marker != " " {
			row = ansi.Truncate(marker+" "+item.Text, m.width, "…")
		}

		lines = append(lines, row)
		lines = append(lines, m.preparedRows[i][1:]...)
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

	m.selectedLine = selectedLine
	if !m.preserveOffset {
		m.followSelection()
	}

	m.preserveOffset = false
}

func (m *List) followSelection() {
	selectedLine := m.selectedLine
	if selectedLine < m.Viewport.YOffset() {
		m.Viewport.SetYOffset(selectedLine)
	}

	if selectedLine >= m.Viewport.YOffset()+m.Viewport.Height() {
		m.Viewport.SetYOffset(selectedLine - m.Viewport.Height() + 1)
	}
}

// Invalidate reapplies caller-supplied presentation after direct native edits.
func (m *List) Invalidate() {
	m.viewCache.Clear()

	if !m.rowsDirty {
		m.preserveOffset = true
	}

	m.rowsDirty = true
	m.revision++
}

// PresentationRevision identifies changes from Charmed's own update path.
func (m *List) PresentationRevision() uint64 { return m.revision }

func (m *List) View() string {
	if m.rowsDirty {
		m.render()
	}

	key := listViewKey{
		revision: m.revision, width: m.Viewport.Width(), height: m.Viewport.Height(),
		x: m.Viewport.XOffset(), y: m.Viewport.YOffset(), outerWidth: m.width, outerHeight: m.height,
	}

	return m.viewCache.Get(key, m.view)
}

func (m *List) view() string {
	if m.footerHeight() == 0 {
		return portable(layout.Fit(m.Viewport.View(), layout.Size{Width: m.width, Height: m.height}), m.Theme)
	}

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
	m.revision++

	var cmd tea.Cmd

	m.Input, cmd = m.Input.Update(msg)
	if m.query == m.Input.Value() {
		return cmd
	}

	m.query = m.Input.Value()
	m.foldedQuery = strings.ToLower(m.query)
	m.reconcile()

	return cmd
}

func (m *List) flushRows() {
	if m.rowsDirty {
		updating := m.updating
		m.updating = false
		m.render()
		m.updating = updating
	}
}
