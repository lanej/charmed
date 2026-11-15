package pattern

import (
	"fmt"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/tree"
	tea "charm.land/bubbletea/v2"
	ltree "charm.land/lipgloss/v2/tree"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// OutlineEntry describes an arbitrary item and its document section. Root-level
// leaves produce a flat list. Children produce groups, with no path semantics.
type OutlineEntry = collection.DocumentSection

type outlineLabel struct{ id, label string }

func (n outlineLabel) String() string { return n.label }

type outlineRect struct{ x, y, width, height int }

func (r outlineRect) contains(x, y int) bool {
	return x >= r.x && x < r.x+r.width && y >= r.y && y < r.y+r.height
}

func (r outlineRect) size() layout.Size {
	return layout.Size{Width: r.width, Height: r.height}
}

func (r outlineRect) content() outlineRect {
	if r.width < 4 || r.height < 3 {
		return r
	}

	size := layout.PanelContentSize(r.size())

	return outlineRect{x: r.x + (r.width-size.Width)/2, y: r.y + 1, width: size.Width, height: size.Height}
}

// OutlineDocument links a native tree and one scrolling section stack.
// It owns nodes, content, geometry and fold synchronization. Native tree styles,
// enumerators and bindings survive SetEntries and Resize. Nodes are rebuilt;
// node-specific customizations must be reapplied after SetEntries or folding.
// Keep node labels one row for hit testing. Body renderers may be multiline.
// Parent applications route Tab and call SetFocus; mouse coordinates are screen
// coordinates translated by SetOrigin. View never changes selection or scroll.
type OutlineDocument struct {
	Outline  tree.Model
	Document *collection.SectionStack
	Theme    theme.Theme
	AppKeys  []key.Binding
	// IndentDepth caps visible ancestor rails; hierarchy and folding are retained.
	// SetIndentDepth refreshes existing nodes. Native Indenter overrides remain available.
	IndentDepth int
	nodes       map[string]*tree.Node
	parents     map[string]string
	depths      map[string]int
	selected    string
	size        layout.Size
	docFocus    bool
	originX     int
	originY     int
}

func NewOutlineDocument(entries []OutlineEntry, size layout.Size, t theme.Theme) (*OutlineDocument, error) {
	stack, err := collection.NewSectionStack(entries, size, t)
	if err != nil {
		return nil, err
	}

	m := &OutlineDocument{Document: stack, Theme: t, size: size, IndentDepth: 2}
	m.Outline = collection.NewTree(tree.Root("Contents"), size.Width, size.Height, t)
	m.Outline.Enumerator(func(children ltree.Children, i int) string {
		if i == children.Length()-1 {
			return "└"
		}

		return "├"
	})
	m.Outline.Indenter(m.outlineIndent)
	m.Outline.SetShowHelp(false)
	m.rebuildOutline()
	m.Resize(size)

	if err := m.SetKeyLayout(interaction.VimKeys); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *OutlineDocument) SetEntries(entries []OutlineEntry) error {
	if err := m.Document.SetSections(entries); err != nil {
		return err
	}

	m.rebuildOutline()

	return nil
}

func (m *OutlineDocument) SelectedID() string     { return m.selected }
func (m *OutlineDocument) DocumentFocused() bool  { return m.docFocus }
func (m *OutlineDocument) SetFocus(document bool) { m.docFocus = document }

// SetOrigin is the top-left screen cell of View, excluding any parent chrome.
func (m *OutlineDocument) SetOrigin(x, y int) { m.originX, m.originY = x, y }

// Select reveals the item or its collapsed ancestor without changing folds.
func (m *OutlineDocument) Select(id string) bool {
	if _, exists := m.nodes[id]; !exists {
		return false
	}

	m.selectVisible(id)
	m.Document.ScrollTo(id)

	return true
}

func (m *OutlineDocument) Collapsed(id string) bool { return m.Document.Collapsed(id) }

func (m *OutlineDocument) SetCollapsed(id string, collapsed bool) bool {
	if !m.Document.SetCollapsed(id, collapsed) {
		return false
	}

	m.rebuildOutline()

	return true
}

// SetKeyLayout applies a preset once, before native binding overrides. Resize
// and data replacement never reapply it. Enter folds the active section.
func (m *OutlineDocument) SetKeyLayout(preset interaction.KeyLayout) error {
	nav, ok := interaction.Navigation(preset)
	if !ok {
		return fmt.Errorf("unknown outline navigation preset %q", preset)
	}

	km := &m.Outline.KeyMap
	km.Up, km.Down, km.Open, km.Close = nav.Up, nav.Down, nav.Right, nav.Left
	km.PageUp, km.PageDown = nav.PageUp, nav.PageDown
	km.HalfPageUp, km.HalfPageDown = nav.HalfPageUp, nav.HalfPageDown
	km.GoToTop, km.GoToBottom = nav.Top, nav.Bottom
	vp := &m.Document.Viewport.KeyMap
	vp.Up, vp.Down = nav.Up, nav.Down
	vp.PageUp, vp.PageDown = nav.PageUp, nav.PageDown
	vp.HalfPageUp, vp.HalfPageDown = nav.HalfPageUp, nav.HalfPageDown

	return nil
}

func (m *OutlineDocument) rebuildOutline() {
	oldParents := m.parents
	m.parents = make(map[string]string)
	m.nodes = make(map[string]*tree.Node)
	m.depths = make(map[string]int)
	root := tree.Root("Contents")

	var appendNodes func(*tree.Node, []OutlineEntry, string, int)

	appendNodes = func(parent *tree.Node, entries []OutlineEntry, parentID string, depth int) {
		for _, entry := range entries {
			label := entry.Title
			if len(entry.Children) == 0 {
				glyph := theme.GlyphExpanded
				if m.Collapsed(entry.ID) {
					glyph = theme.GlyphCollapsed
				}

				label = m.Theme.Glyph(glyph) + " " + label
			}

			if entry.Summary != "" {
				label += "  " + m.Theme.Subtle().Render(entry.Summary)
			}

			label = ansi.Truncate(label, max(1, m.Outline.Width()-min(depth, m.indentDepth())-7), "…")
			value := outlineLabel{id: entry.ID, label: label}

			var node *tree.Node
			if len(entry.Children) > 0 {
				node = tree.Root(value)
				appendNodes(node, entry.Children, entry.ID, depth+1)

				if m.Collapsed(entry.ID) {
					node.Close()
				}
			} else {
				parent.Child(value)
				children := parent.ChildNodes()
				node = children[len(children)-1]
			}

			m.parents[entry.ID] = parentID
			m.depths[entry.ID] = depth

			m.nodes[entry.ID] = node
			if len(entry.Children) > 0 {
				parent.Child(node)
			}
		}
	}

	appendNodes(root, m.Document.Sections(), "", 0)
	m.Outline.SetNodes(root)

	id := m.selected
	for id != "" && m.nodes[id] == nil {
		id = oldParents[id]
	}

	if id == "" {
		id = m.Document.TopID()
	}

	m.selectVisible(id)
}

// SetIndentDepth changes visible indentation without flattening the data.
// Zero suppresses ancestor rails; negative depths are rejected atomically.
func (m *OutlineDocument) SetIndentDepth(depth int) error {
	if depth < 0 {
		return fmt.Errorf("outline indentation depth must not be negative")
	}

	m.IndentDepth = depth
	m.rebuildOutline()

	return nil
}

func (m *OutlineDocument) indentDepth() int {
	return min(max(0, m.IndentDepth), max(0, m.Outline.Width()-16))
}

func (m *OutlineDocument) outlineIndent(children ltree.Children, i int) string {
	node, ok := children.At(i).(*tree.Node)
	if !ok {
		return ""
	}

	label, ok := node.GivenValue().(outlineLabel)
	if !ok || m.depths[label.id] >= m.indentDepth() {
		return ""
	}

	if i == children.Length()-1 {
		return " "
	}

	return "│"
}

func (m *OutlineDocument) selectVisible(id string) {
	visible := m.Document.VisibleIDs()
	for id != "" {
		for _, candidate := range visible {
			if candidate == id {
				m.selected = id
				m.Outline.SetYOffset(m.nodes[id].YOffset())

				return
			}
		}

		id = m.parents[id]
	}

	m.selected = ""
}

func (m *OutlineDocument) paneRects() (outline, document outlineRect) {
	w, h := max(0, m.size.Width), max(0, m.size.Height-1)
	if w >= 80 {
		left := min(36, w/3)
		return outlineRect{width: left, height: h}, outlineRect{x: left, width: w - left, height: h}
	}

	top := min(8, h/3)

	return outlineRect{width: w, height: top}, outlineRect{y: top, width: w, height: h - top}
}

func (m *OutlineDocument) Resize(size layout.Size) {
	m.size = size
	outline, document := m.paneRects()
	a, b := outline.content(), document.content()
	widthChanged := m.Outline.Width() != max(1, a.width)
	m.Outline.SetSize(max(1, a.width), max(1, a.height))

	if widthChanged {
		m.rebuildOutline()
	}

	m.Document.Resize(b.size())
}

func (m *OutlineDocument) Init() tea.Cmd { return nil }
func (m *OutlineDocument) Close() error  { return nil }

func (m *OutlineDocument) Update(msg tea.Msg) tea.Cmd {
	if mouse, ok := msg.(tea.MouseMsg); ok {
		return m.updateMouse(msg, mouse.Mouse())
	}

	if pressed, ok := msg.(tea.KeyPressMsg); ok {
		if m.foldKey(pressed) {
			return nil
		}

		if m.docFocus {
			before := m.Document.Viewport.YOffset()
			switch {
			case key.Matches(pressed, m.Outline.KeyMap.GoToTop):
				m.Document.Viewport.GotoTop()
			case key.Matches(pressed, m.Outline.KeyMap.GoToBottom):
				m.Document.Viewport.GotoBottom()
			default:
				cmd := m.Document.Update(msg)
				m.trackScroll(before)

				return cmd
			}

			m.trackScroll(before)

			return nil
		}
	}

	if m.docFocus {
		return m.Document.Update(msg)
	}

	var cmd tea.Cmd

	m.Outline, cmd = m.Outline.Update(msg)
	m.syncOutlineSelection()

	return cmd
}

func (m *OutlineDocument) foldKey(pressed tea.KeyPressMsg) bool {
	id := m.selected

	km := m.Outline.KeyMap
	switch {
	case key.Matches(pressed, km.Toggle):
		m.SetCollapsed(id, !m.Collapsed(id))
	case key.Matches(pressed, km.Open):
		m.SetCollapsed(id, false)
	case key.Matches(pressed, km.Close):
		m.SetCollapsed(id, true)
	default:
		return false
	}

	return true
}

func (m *OutlineDocument) syncOutlineSelection() {
	node := m.Outline.NodeAtCurrentOffset()
	if node != nil {
		if label, ok := node.GivenValue().(outlineLabel); ok {
			if label.id != m.selected {
				m.selected = label.id
				m.Document.ScrollTo(label.id)
			}

			return
		}
	}

	ids := m.Document.VisibleIDs()
	if len(ids) > 0 {
		m.Select(ids[0])
	}
}

func (m *OutlineDocument) trackScroll(before int) {
	if m.Document.Viewport.YOffset() != before {
		m.selectVisible(m.Document.TopID())
	}
}

func (m *OutlineDocument) updateMouse(msg tea.Msg, mouse tea.Mouse) tea.Cmd {
	x, y := mouse.X-m.originX, mouse.Y-m.originY
	outline, document := m.paneRects()
	a, b := outline.content(), document.content()

	switch typed := msg.(type) {
	case tea.MouseClickMsg:
		if typed.Button != tea.MouseLeft {
			return nil
		}

		switch {
		case a.contains(x, y):
			row := y - a.y + m.Outline.ViewportYOffset()
			if node := m.Outline.Node(row); node != nil {
				if label, ok := node.GivenValue().(outlineLabel); ok {
					m.SetFocus(false)
					m.Select(label.id)

					if len(node.ChildNodes()) > 0 {
						m.SetCollapsed(label.id, !m.Collapsed(label.id))
					}
				}
			}
		case b.contains(x, y):
			m.SetFocus(true)

			if id := m.Document.HeaderAt(y - b.y); id != "" {
				m.selectVisible(id)
				m.SetCollapsed(id, !m.Collapsed(id))
			}
		}
	case tea.MouseWheelMsg:
		switch {
		case b.contains(x, y):
			before := m.Document.Viewport.YOffset()
			cmd := m.Document.Update(msg)
			m.trackScroll(before)

			return cmd
		case a.contains(x, y):
			for range 3 {
				switch typed.Button {
				case tea.MouseWheelUp:
					m.Outline.Up()
				case tea.MouseWheelDown:
					m.Outline.Down()
				}
			}

			m.syncOutlineSelection()
		}
	}

	return nil
}

func (m *OutlineDocument) View() string {
	outline, document := m.paneRects()
	a := layout.Panel("Outline", m.Outline.View(), outline.size(), m.Theme, !m.docFocus)
	b := layout.Panel("Document", m.Document.View(), document.size(), m.Theme, m.docFocus)

	body := layout.Row(a, b)
	if m.size.Width < 80 {
		body = layout.Column(a, b)
	}

	help := "↑/↓ navigate · Enter fold · Tab pane"
	if m.docFocus {
		help = "↑/↓ scroll · Enter fold · Tab pane"
	}

	if m.size.Width < 60 {
		help = "↑/↓ · Enter fold · Tab"
	}

	for _, binding := range m.AppKeys {
		if binding.Enabled() {
			help += " · " + binding.Help().Key + " " + binding.Help().Desc
		}
	}

	if m.Theme.ASCII {
		body, help = layout.ASCII(body), layout.ASCII(help)
	}

	help = ansi.Truncate(m.Theme.Subtle().Render(help), max(0, m.size.Width), "")

	return layout.Fit(body+"\n"+help, m.size)
}

// SyncScroll updates the active outline entry after direct native scrolling.
func (m *OutlineDocument) SyncScroll() { m.selectVisible(m.Document.TopID()) }
