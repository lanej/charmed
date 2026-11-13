package collection

import (
	"strings"

	"charm.land/bubbles/v2/tree"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
)

// SetDetachedDetails moves log fields from the hierarchy into a full-width
// selected-record viewport in place of the selected leaf. Native detail styles/key maps survive data changes;
// content is regenerated on selection/data/size changes. D switches scroll focus.
func (m *LogTree) SetDetachedDetails(enabled bool) {
	selected := m.Selected()
	m.DetachedDetails = enabled
	m.DetailsFocused = false
	m.Model.SetSize(max(1, m.Model.Width()), m.height)
	m.rebuild()

	if !m.Follow && selected != "" {
		m.Select(selected)
	}
}

func (m *LogTree) refreshDetached() {
	if !m.DetachedDetails {
		return
	}

	m.Model.SetHeight(m.height)
	m.Model.SetYOffset(m.Model.YOffset())
	size := layout.PanelContentSize(layout.Size{Width: m.breakoutWidth(), Height: m.height})

	inset := 0
	if m.ShowScrollbar && size.Width > 1 {
		inset = 1
	}

	m.Details.SetWidth(max(1, size.Width-inset))
	id := m.Selected()

	content := m.Theme.Subtle().Render("Select a log entry")
	for _, record := range m.records {
		if record.ID == id {
			content = m.detachedContent(record)
			break
		}
	}

	m.Details.SetContent(content)
	context := min(max(0, m.Model.Root().Height()-1), max(0, m.height/2))
	bodyHeight := min(m.Details.TotalLineCount(), max(1, m.height-context-2))
	m.Details.SetHeight(max(1, bodyHeight))

	if id != m.detailID {
		m.Details.GotoTop()
	}

	m.detailID = id
	if id == "" {
		m.DetailsFocused = false
	}
}

func (m *LogTree) detachedContent(record LogRecord) string {
	// WHY: Reuse bounded field construction, independent of trace/span indentation.
	fields := &LogTree{Theme: m.Theme, FieldDepth: m.FieldDepth, FieldNodes: m.FieldNodes, nodes: map[string]*tree.Node{}, parentNodes: map[*tree.Node]*tree.Node{}}
	root := tree.Root("")
	fields.recordFields(root, record, m.Details.Width(), nil)
	model := NewTree(root, m.Details.Width(), 1, m.Theme)
	model.SetShowHelp(false)
	model.SetCursorCharacter("")
	model.SetOpenCharacter("")
	model.SetClosedCharacter("")

	styles := tree.DefaultStyles(m.Theme.IsDark())
	styles.NodeStyle = m.Theme.Text()
	styles.ParentNodeStyle = m.Theme.Text()
	styles.EnumeratorStyle = m.Theme.Subtle()
	styles.IndenterStyle = m.Theme.Subtle()
	styles.SelectedNodeStyle = m.Theme.Text()
	styles.RootNodeStyle = m.Theme.Text()
	model.SetStyles(styles)

	for _, node := range fields.nodes {
		node.Open()
	}

	model.SetNodes(root)
	model.SetHeight(root.Height())
	// The empty synthetic root contributes one line; the fields retain their order.
	lines := strings.Split(model.View(), "\n")
	fieldView := strings.Join(lines[min(1, len(lines)):], "\n")
	labelModel := *m
	labelModel.DetachedDetails = false

	return labelModel.entryLabel(record, m.Details.Width()) + "\n" + fieldView
}

func (m *LogTree) breakoutWidth() int {
	if m.ShowScrollbar && m.width > 1 {
		return m.width - 1
	}

	return m.width
}

func (m *LogTree) detachedView() string {
	if m.detailID == "" {
		return m.treeView()
	}

	width := m.Details.Width()
	if m.ShowScrollbar && width > 1 {
		width++
	}

	pane := &ScrollPane{Model: &m.Details, Theme: m.Theme, ShowScrollbar: m.ShowScrollbar}
	pane.Resize(layout.Size{Width: width, Height: m.Details.Height()})
	details := pane.View()

	panelHeight := min(m.height, m.Details.Height()+2)
	panel := layout.Panel("Log · "+text.Inline(m.detailID), details, layout.Size{Width: m.breakoutWidth(), Height: panelHeight}, m.Theme, true)
	// WHY: Render the native hierarchy before replacing its selected leaf, so
	// siblings keep their original order and indentation on either side.
	lines, selected, end, ok := m.focusedLeafLines()
	if !ok {
		return m.treeView()
	}

	combined := append([]string{}, lines[:selected]...)
	combined = append(combined, strings.Split(panel, "\n")...)
	combined = append(combined, lines[end:]...)
	above := max(0, (m.height-panelHeight)/2)
	offset := min(max(0, selected-above), max(0, len(combined)-m.height))

	content := layout.Fit(strings.Join(combined[offset:], "\n"), layout.Size{Width: m.breakoutWidth(), Height: m.height})
	if m.ShowScrollbar && m.width > 1 {
		bar := feedback.Scrollbar{Total: len(combined), Visible: m.height, Offset: offset, Height: m.height}
		content = lipgloss.JoinHorizontal(lipgloss.Top, content, bar.View(m.Theme))
	}

	return layout.Fit(content, layout.Size{Width: m.width, Height: m.height})
}

// Native node offsets exclude the TreeStyle frame. A single-node probe measures
// that frame without exposing or reconstructing the caller's native styles.
func (m *LogTree) focusedLeafLines() ([]string, int, int, bool) {
	probe := m.Model
	root := tree.Root("@")

	probe.SetShowHelp(false)
	probe.SetCursorCharacter("")
	probe.SetNodes(root)
	probe.SetHeight(m.height)
	probe.SetViewportYOffset(0)

	frameTop := -1

	for i, line := range strings.Split(probe.View(), "\n") {
		if strings.Contains(ansi.Strip(line), "@") {
			frameTop = i
			break
		}
	}

	for i, line := range strings.Split(root.String(), "\n") {
		if strings.Contains(ansi.Strip(line), "@") {
			frameTop -= i
			break
		}
	}

	if frameTop < 0 {
		return nil, 0, 0, false
	}

	snapshot := m.Model
	snapshot.SetShowHelp(false)
	snapshot.SetHeight(m.Model.Root().Height() + frameTop + m.height)
	snapshot.SetViewportYOffset(0)

	lines := strings.Split(snapshot.View(), "\n")
	for len(lines) > 1 && strings.TrimSpace(ansi.Strip(lines[len(lines)-1])) == "" {
		lines = lines[:len(lines)-1]
	}

	if m.FocusedLeafGeometry != nil {
		start, end, ok := m.FocusedLeafGeometry(snapshot)
		return lines, start, end, ok && start >= 0 && end > start && end <= len(lines)
	}
	// Native offsets assume single-line nodes. Keep the native view when a
	// customization violates that assumption rather than replacing a wrong row.
	nodes := m.Model.AllNodes()
	if m.Model.Root().Height() != len(nodes) {
		return nil, 0, 0, false
	}

	for _, node := range nodes {
		if lipgloss.Height(node.Value()) != 1 {
			return nil, 0, 0, false
		}
	}

	node := m.Model.NodeAtCurrentOffset()
	selected := node.LineOffset() + frameTop

	end := min(len(lines), selected+max(1, lipgloss.Height(node.Value())))
	if selected < 0 || selected >= len(lines) {
		return nil, 0, 0, false
	}

	return lines, selected, end, true
}

// CanFocusDetails reports whether the selected leaf has a visible focused container.
func (m *LogTree) CanFocusDetails() bool {
	if !m.DetachedDetails || m.detailID == "" {
		return false
	}

	_, _, _, ok := m.focusedLeafLines()

	return ok
}
