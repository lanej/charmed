package pattern

import (
	"fmt"
	"strings"
	"unicode"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// InspectorRecord supplies a stable identity with structured presentation.
// Summary renders the list row and Document supplies sectioned inspection.
// Heading is optional plain text above the record, not a selectable item.
//
// Row and Detail are optional pure custom renderers that override Summary and
// Document respectively. Row receives the width after the selection gutter;
// Detail receives its full content width. With neither Document nor Detail,
// inspection is disabled. The caller owns ordering, evidence and actions.
type InspectorRecord struct {
	ID, Heading string
	Summary     *layout.RecordSummary
	Document    *layout.DetailDocument
	Row         func(width int, selected bool) string
	Detail      func(width int) string
}

// ListInspectorKeys act on records while browsing. During inspection, native
// Detail bindings handle scrolling and Back returns to the retained list.
type ListInspectorKeys struct{ Previous, Next, Inspect, Back key.Binding }

type inspectorRowSpan struct {
	id         string
	start, end int
}

type inspectorListAnchor struct {
	id            string
	row, fallback int
}

// ListInspector composes a selectable multiline list and optional detail preview.
// Wide browsing shows both panes; narrow browsing shows the list. Inspection
// uses the full allocation. The parent owns application navigation and effects.
//
// List and Detail retain their native models, styles and key maps across record
// replacement and resize. Content, dimensions, SoftWrap and row anchoring belong
// to this wrapper. Hidden panes retain their native size and offset until shown;
// visible reflow can clamp offsets. Structured documents retain their leading
// visible section and relative row across reflow; side-by-side sections anchor
// to the first section in reading order. Custom details retain their row offset.
// Changing selection resets the detail position.
//
// SetRecords copies typed data, including nested slices, but not callback
// captures. Renderers run during mutation or reflow, never View, and must remain
// pure on the owning UI loop. Call SetRecords to replace typed observations;
// Refresh after changing captured data, theme or native presentation. Public
// presentation options also take effect on Refresh or Resize.
type ListInspector struct {
	List, Detail               viewport.Model
	Keys                       ListInspectorKeys
	Theme                      theme.Theme
	Preview                    bool
	ListTitle, DetailTitle     string
	Empty                      string
	ListWidth, MinDetailWidth  int
	records                    []InspectorRecord
	indices                    map[string]int
	selected                   string
	spans                      []inspectorRowSpan
	detailSections             []layout.DetailSectionRange
	size, listSize, detailSize layout.Size
	focused, inspecting, split bool
	pendingReveal              bool
}

func NewListInspector(records []InspectorRecord, size layout.Size, t theme.Theme) (*ListInspector, error) {
	m := &ListInspector{
		List: collection.NewViewport("", 1, 1), Detail: collection.NewViewport("", 1, 1),
		Theme: t, Preview: true, ListTitle: "Items", DetailTitle: "Details", Empty: "No records",
		ListWidth: 40, MinDetailWidth: 44, size: size, focused: true,
		Keys: ListInspectorKeys{
			Previous: key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("k/↑", "previous record")),
			Next:     key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("j/↓", "next record")),
			Inspect:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "inspect")),
			Back:     key.NewBinding(key.WithKeys("esc", "backspace"), key.WithHelp("esc", "back")),
		},
	}
	for _, model := range []*viewport.Model{&m.List, &m.Detail} {
		model.KeyMap.PageUp = key.NewBinding(key.WithKeys("pgup", "["), key.WithHelp("[/pgup", "page up"))
		model.KeyMap.PageDown = key.NewBinding(key.WithKeys("pgdown", "]"), key.WithHelp("]/pgdn", "page down"))
	}

	if err := m.SetRecords(records); err != nil {
		return nil, err
	}

	return m, nil
}

// SetRecords validates the whole replacement before changing either pane. IDs
// must be unique, nonempty and free of control characters; Summary or Row must
// be present. Structured section IDs have the same rules within each document.
// A retained ID keeps selection. Removing it selects the first remaining record
// and closes inspection, so a different identity cannot silently take its place.
func (m *ListInspector) SetRecords(records []InspectorRecord) error {
	indices := make(map[string]int, len(records))
	for index, record := range records {
		if strings.TrimSpace(record.ID) == "" || strings.IndexFunc(record.ID, unicode.IsControl) >= 0 {
			return fmt.Errorf("empty or unsafe inspector record ID %q", record.ID)
		}

		if _, exists := indices[record.ID]; exists {
			return fmt.Errorf("duplicate inspector record ID %q", record.ID)
		}

		if record.Row == nil && record.Summary == nil {
			return fmt.Errorf("inspector record %q has no summary or row renderer", record.ID)
		}

		if err := validateInspectorDocument(record); err != nil {
			return err
		}

		indices[record.ID] = index
	}

	previous := m.selected
	reveal := m.selectionVisible()
	m.records, m.indices = cloneInspectorRecords(records), indices

	if _, retained := indices[m.selected]; !retained {
		m.selected, m.inspecting = "", false
		if len(records) > 0 {
			m.selected = records[0].ID
		}
	}

	if record, present := m.current(); !present || !record.hasDetail() {
		m.inspecting = false
	}

	if m.selected != previous {
		m.Detail.GotoTop()
		m.detailSections = nil
		m.pendingReveal = true
		reveal = true
	}

	m.reflow(reveal)

	return nil
}

func (m *ListInspector) SelectedID() (string, bool) {
	_, present := m.indices[m.selected]
	return m.selected, present
}

// Select leaves state unchanged for unknown IDs. Selecting the current record
// only reveals it; changing identity also resets the detail offset.
func (m *ListInspector) Select(id string) bool {
	if _, present := m.indices[id]; !present {
		return false
	}

	if id != m.selected {
		m.selected = id
		m.pendingReveal = true
		m.Detail.GotoTop()
		m.detailSections = nil

		if record, _ := m.current(); !record.hasDetail() {
			m.inspecting = false
		}

		m.reflow(false)
	}

	m.revealSelection(true)

	return true
}

// Move follows the supplied order and clamps at either end; it never wraps.
func (m *ListInspector) Move(delta int) {
	index, present := m.indices[m.selected]
	if !present {
		return
	}

	index += max(-index, min(delta, len(m.records)-1-index))
	m.Select(m.records[index].ID)
}

func (m *ListInspector) Inspect() bool {
	record, present := m.current()
	if !present || !record.hasDetail() {
		return false
	}

	if !m.inspecting {
		m.inspecting = true
		m.reflow(false)
	}

	return true
}

func (m *ListInspector) Back() {
	if m.inspecting {
		m.inspecting = false
		m.reflow(false)
	}
}

func (m *ListInspector) Inspecting() bool { return m.inspecting }
func (m *ListInspector) Split() bool      { return m.split }
func (m *ListInspector) Focus()           { m.focused = true }
func (m *ListInspector) Blur()            { m.focused = false }

func (m *ListInspector) Resize(size layout.Size) {
	reveal := m.selectionVisible()
	m.size = size
	m.reflow(reveal)
}

func (m *ListInspector) Refresh() { m.reflow(m.selectionVisible()) }

func (m *ListInspector) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	if press, ok := msg.(tea.KeyPressMsg); ok {
		switch {
		case m.inspecting && key.Matches(press, m.Keys.Back):
			m.Back()
			return nil
		case !m.inspecting && key.Matches(press, m.Keys.Previous):
			m.Move(-1)
			return nil
		case !m.inspecting && key.Matches(press, m.Keys.Next):
			m.Move(1)
			return nil
		case !m.inspecting && key.Matches(press, m.Keys.Inspect):
			m.Inspect()
			return nil
		}
	}

	active := &m.List
	if m.inspecting {
		active = &m.Detail
	}

	var cmd tea.Cmd

	*active, cmd = active.Update(msg)

	return cmd
}

func (m *ListInspector) View() string {
	if m.size.Width <= 0 || m.size.Height <= 0 {
		return ""
	}

	if m.inspecting {
		return layout.Panel(text.Inline(m.DetailTitle), m.Detail.View(), m.detailSize, m.Theme, m.focused)
	}

	list := layout.Panel(text.Inline(m.ListTitle), m.List.View(), m.listSize, m.Theme, m.focused)
	if !m.split {
		return list
	}

	detail := layout.Panel(text.Inline(m.DetailTitle), m.Detail.View(), m.detailSize, m.Theme, false)
	gap := layout.Fit("", layout.Size{Width: theme.SpaceMedium, Height: m.size.Height})

	return layout.Row(list, gap, detail)
}

func (m *ListInspector) current() (InspectorRecord, bool) {
	index, present := m.indices[m.selected]
	if !present {
		return InspectorRecord{}, false
	}

	return m.records[index], true
}

func (m *ListInspector) reflow(reveal bool) {
	m.size.Width, m.size.Height = max(0, m.size.Width), max(0, m.size.Height)
	m.split = false

	m.listSize, m.detailSize = layout.Size{}, layout.Size{}
	if m.size.Width == 0 || m.size.Height == 0 {
		return
	}

	if m.inspecting {
		m.detailSize = m.size
	} else {
		m.listSize = m.size
		listWidth, detailMinimum := max(1, m.ListWidth), max(1, m.MinDetailWidth)
		remaining := m.size.Width - min(listWidth, m.size.Width) - theme.SpaceMedium

		if m.Preview && remaining >= detailMinimum {
			m.split = true
			m.listSize.Width = listWidth
			m.detailSize = layout.Size{Width: remaining, Height: m.size.Height}
		}
	}

	if m.listSize.Width > 0 {
		anchor := m.listAnchor()
		inspectorResizeViewport(&m.List, m.listSize)
		m.renderRows()
		m.restoreList(anchor)

		if reveal || m.pendingReveal {
			m.revealSelection(m.pendingReveal)
			m.pendingReveal = false
		}
	}

	if m.detailSize.Width > 0 {
		anchor := m.detailAnchor()
		inspectorResizeViewport(&m.Detail, m.detailSize)
		width := max(1, inspectorViewportSize(&m.Detail).Width)
		content := m.Theme.Subtle().Render("No details supplied")
		m.detailSections = nil

		if record, present := m.current(); present {
			switch {
			case record.Detail != nil:
				content = record.Detail(width)
			case record.Document != nil:
				rendered := record.Document.Render(width, m.Theme)
				content, m.detailSections = rendered.Content, rendered.Sections
			}
		}

		m.Detail.SetContent(ansi.Hardwrap(content, width, true))
		m.restoreDetail(anchor)
	}
}

func inspectorResizeViewport(model *viewport.Model, size layout.Size) {
	inner := size
	if size.Width >= 2*theme.BorderCells+2 && size.Height >= 2*theme.BorderCells+1 {
		inner = layout.PanelContentSize(size)
	}

	model.SetWidth(max(1, inner.Width))
	model.SetHeight(max(1, inner.Height))
	model.SoftWrap = false
}

func inspectorViewportSize(model *viewport.Model) layout.Size {
	width, height := model.Width(), model.Height()
	if styled := model.Style.GetWidth(); styled > 0 {
		width = min(width, styled)
	}

	if styled := model.Style.GetHeight(); styled > 0 {
		height = min(height, styled)
	}

	gutter := 0
	if model.LeftGutterFunc != nil {
		gutter = ansi.StringWidth(model.LeftGutterFunc(viewport.GutterContext{}))
	}

	return layout.Size{Width: max(0, width-model.Style.GetHorizontalFrameSize()-gutter), Height: max(0, height-model.Style.GetVerticalFrameSize())}
}

func (m *ListInspector) renderRows() {
	width := max(1, inspectorViewportSize(&m.List).Width)
	gutter := min(2, width-1)
	padding := strings.Repeat(" ", gutter)
	lines := []string{}
	m.spans = nil

	for _, record := range m.records {
		if len(lines) > 0 {
			lines = append(lines, "")
		}

		start := len(lines)

		if record.Heading != "" {
			heading := ansi.Hardwrap(m.Theme.Subtle().Render(text.Inline(record.Heading)), width, true)
			lines = append(lines, strings.Split(heading, "\n")...)
		}

		selected := record.ID == m.selected

		body := record.row(width-gutter, selected, m.Theme)

		body = ansi.Hardwrap(body, width-gutter, true)
		for index, line := range strings.Split(body, "\n") {
			prefix := padding
			if selected && index == 0 && gutter > 0 {
				prefix = m.Theme.Heading().Render(m.Theme.Glyph(theme.GlyphRight)) + strings.Repeat(" ", gutter-1)
			}

			lines = append(lines, prefix+line)
		}

		m.spans = append(m.spans, inspectorRowSpan{id: record.ID, start: start, end: len(lines)})
	}

	if len(lines) == 0 {
		lines = strings.Split(ansi.Hardwrap(m.Theme.Subtle().Render(text.Escape(m.Empty)), width, true), "\n")
	}

	m.List.SetContentLines(lines)
}

func (m *ListInspector) listAnchor() inspectorListAnchor {
	offset := m.List.YOffset()
	for _, span := range m.spans {
		if offset < span.end {
			return inspectorListAnchor{id: span.id, row: max(0, offset-span.start), fallback: offset}
		}
	}

	return inspectorListAnchor{fallback: offset}
}

func (m *ListInspector) restoreList(anchor inspectorListAnchor) {
	for _, span := range m.spans {
		if span.id == anchor.id {
			m.List.SetYOffset(span.start + min(anchor.row, span.end-span.start-1))
			return
		}
	}

	m.List.SetYOffset(anchor.fallback)
}

func (m *ListInspector) selectionVisible() bool {
	height := inspectorViewportSize(&m.List).Height
	for _, span := range m.spans {
		if span.id == m.selected {
			return span.end > m.List.YOffset() && span.start < m.List.YOffset()+height
		}
	}

	return false
}

func (m *ListInspector) revealSelection(topOfTall bool) {
	height := inspectorViewportSize(&m.List).Height
	if height == 0 || m.inspecting {
		return
	}

	for _, span := range m.spans {
		if span.id != m.selected {
			continue
		}

		offset := m.List.YOffset()
		switch {
		case span.end-span.start > height:
			if topOfTall || offset < span.start || offset >= span.end {
				m.List.SetYOffset(span.start)
			}
		case span.start < offset:
			m.List.SetYOffset(span.start)
		case span.end > offset+height:
			m.List.SetYOffset(span.end - height)
		}

		return
	}
}
