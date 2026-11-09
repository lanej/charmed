package collection

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	binding "charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/tree"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/log/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// LogTree retains the native tree API. SetEntries rebuilds nodes, preserving
// native styles/key maps and expansion/selection for retained stable identities.
type LogTree struct {
	// FocusedLeafGeometry returns [start,end) rendered rows in a zero-scroll native snapshot.
	// Vertically styled/wrapped nodes require this override; invalid geometry keeps the native view.
	FocusedLeafGeometry             func(tree.Model) (start, end int, ok bool)
	DetachedDetails, DetailsFocused bool
	Details                         viewport.Model
	height                          int
	detailID                        string
	Symbols                         text.Presentation
	ScrollbarGeometry               func(tree.Model) feedback.Scrollbar
	parentNodes                     map[*tree.Node]*tree.Node
	Model                           tree.Model
	Theme                           theme.Theme
	Limit, FieldDepth, FieldNodes   int
	Follow, GroupTraces             bool
	Folding, ShowScrollbar          bool
	width                           int
	records                         []LogRecord
	nodes                           map[string]*tree.Node
}

type logNode struct{ key, record, label string }

func (n logNode) String() string { return n.label }

func NewLogTree(width, height int, t theme.Theme) *LogTree {
	m := &LogTree{Theme: t, Limit: DefaultLogLimit, FieldDepth: DefaultFieldDepth, FieldNodes: DefaultFieldNodes, Follow: true, GroupTraces: true, Folding: true, ShowScrollbar: true, width: max(1, width)}
	m.Model = NewTree(tree.Root("Logs"), max(1, width-1), height, t)
	m.Model.SetShowHelp(false)
	m.height = max(1, height)
	m.Details = NewViewport("", width, height)
	m.nodes = map[string]*tree.Node{}

	return m
}

func (m *LogTree) Entries() []LogRecord { return slices.Clone(m.records) }
func (m *LogTree) Init() tea.Cmd        { return nil }
func (m *LogTree) Resize(width, height int) {
	m.width = max(1, width)
	m.height = max(1, height)

	inset := 0
	if m.ShowScrollbar && m.width > 1 {
		inset = 1
	}

	changed := m.Model.Width() != max(1, width-inset)
	m.Model.SetSize(max(1, width-inset), max(1, height))

	if changed {
		m.rebuild()
	}

	m.refreshDetached()
}

func (m *LogTree) View() string {
	if m.DetachedDetails {
		return m.detachedView()
	}

	return m.treeView()
}

func (m *LogTree) treeView() string {
	if !m.ShowScrollbar || m.width <= 1 {
		return m.Model.View()
	}

	bar := feedback.Scrollbar{Total: m.Model.Root().Height(), Visible: m.Model.Height(), Offset: m.Model.ViewportYOffset(), Height: m.Model.Height()}
	if m.ScrollbarGeometry != nil {
		bar = m.ScrollbarGeometry(m.Model)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, m.Model.View(), bar.View(m.Theme))
}

func (m *LogTree) SetSymbols(mode text.Presentation) { m.Symbols = mode; m.rebuild() }

func (m *LogTree) SetFolding(enabled bool) {
	m.Folding = enabled
	if enabled {
		m.Model.SetOpenCharacter(m.Theme.Glyph(theme.GlyphExpanded))
		m.Model.SetClosedCharacter(m.Theme.Glyph(theme.GlyphCollapsed))
	} else {
		m.Model.SetOpenCharacter("")
		m.Model.SetClosedCharacter("")
	}

	m.rebuild()
}

func (m *LogTree) Selected() string {
	for _, node := range m.Model.AllNodes() {
		if node.IsSelected() {
			if value, ok := node.GivenValue().(logNode); ok {
				return value.record
			}
		}
	}

	return ""
}

func (m *LogTree) Select(id string) bool {
	if node, ok := m.nodes[entryKey(id)]; ok {
		for parent := m.parentNodes[node]; parent != nil; parent = m.parentNodes[parent] {
			parent.Open()
		}

		m.Model.Root().Open()
		m.Model.SetNodes(m.Model.Root())
		m.Model.SetYOffset(node.YOffset())
		m.refreshDetached()

		return true
	}

	return false
}

func (m *LogTree) SetEntries(records []LogRecord) error {
	if err := validateRecordIDs(records, func(r LogRecord) string { return r.ID }); err != nil {
		return err
	}

	m.records = retainRecords(records, m.Limit)
	m.rebuild()

	return nil
}

func (m *LogTree) Upsert(records ...LogRecord) error {
	if err := validateRecordIDs(records, func(r LogRecord) string { return r.ID }); err != nil {
		return err
	}

	return m.SetEntries(upsertRecords(m.records, records, func(r LogRecord) string { return r.ID }))
}

func (m *LogTree) SetGrouped(grouped bool) { m.GroupTraces = grouped; m.rebuild() }
func (m *LogTree) ResumeFollow()           { m.Follow = true; m.rebuild() }

func (m *LogTree) Update(msg tea.Msg) tea.Cmd {
	if m.DetachedDetails {
		available := m.CanFocusDetails()
		if !available {
			m.DetailsFocused = false
		}

		if pressed, ok := msg.(tea.KeyPressMsg); ok && pressed.String() == "D" {
			if available {
				m.DetailsFocused = !m.DetailsFocused
			}

			return nil
		}

		if m.DetailsFocused {
			var cmd tea.Cmd

			before := m.Details.YOffset()

			m.Details, cmd = m.Details.Update(msg)
			if before != m.Details.YOffset() {
				m.Follow = false
			}

			return cmd
		}
	}

	if pressed, ok := msg.(tea.KeyPressMsg); ok {
		keys := m.Model.KeyMap

		if binding.Matches(pressed, keys.GoToBottom) {
			m.ResumeFollow()
			return nil
		}

		if binding.Matches(pressed, keys.Up, keys.Down, keys.PageUp, keys.PageDown, keys.HalfPageUp, keys.HalfPageDown, keys.GoToTop) || (m.Folding && binding.Matches(pressed, keys.Open, keys.Close, keys.Toggle)) {
			m.Follow = false
		}

		switch pressed.String() {
		case "g":
			m.SetGrouped(!m.GroupTraces)
			return nil
		case "enter", "space", " ":
			if !m.Folding {
				break
			}

			m.Follow = false
			m.Model.ToggleCurrentNode()
			m.refreshDetached()

			return nil
		}
	}

	var cmd tea.Cmd

	keys := m.Model.KeyMap
	if !m.Folding {
		m.Model.KeyMap.Open.SetEnabled(false)
		m.Model.KeyMap.Close.SetEnabled(false)
		m.Model.KeyMap.Toggle.SetEnabled(false)
	}

	m.Model, cmd = m.Model.Update(msg)
	m.Model.KeyMap = keys
	m.refreshDetached()

	return cmd
}

func (m *LogTree) rebuild() {
	rootOpen := m.Model.Root().IsOpen()
	selected := ""
	rootSelected := false

	oldOffset := m.Model.ViewportYOffset()
	for _, node := range m.Model.AllNodes() {
		if node.IsSelected() {
			if value, ok := node.GivenValue().(logNode); ok {
				selected = value.key
			} else {
				rootSelected = true
			}

			break
		}
	}

	opened := map[string]bool{}
	for key, node := range m.nodes {
		opened[key] = node.IsOpen()
	}

	root := tree.Root(m.Theme.Subtle().Render(fmt.Sprintf("Logs · %d retained", len(m.records))))
	m.nodes = map[string]*tree.Node{}
	m.parentNodes = map[*tree.Node]*tree.Node{}

	parents, depths := m.traceNodes(root, opened)
	for _, record := range m.records {
		parent := root
		if m.GroupTraces {
			parent = parents[record.ID]
		}

		node := m.makeNode(entryKey(record.ID), record.ID, m.entryLabel(record, max(1, m.Model.Width()-3*(depths[record.ID]+1)-6)), false, opened)
		parent.Child(node)
		m.parentNodes[node] = parent

		if m.DetachedDetails {
			continue
		}

		available := max(1, m.Model.Width()-3*(depths[record.ID]+1)-6)

		m.recordFields(node, record, available, opened)
	}

	if m.Folding && !rootOpen && !m.Follow {
		root.Close()
	}

	m.Model.SetNodes(root)

	if node, ok := m.nodes[selected]; ok && !m.Follow {
		m.Model.SetYOffset(node.YOffset())

		if node.LineOffset() >= oldOffset && node.LineOffset() < oldOffset+m.Model.Height() {
			m.Model.SetViewportYOffset(oldOffset)
		}
	} else if m.Follow && len(m.records) > 0 {
		m.Select(m.records[len(m.records)-1].ID)
	} else if rootSelected {
		m.Model.SetYOffset(0)
	} else if len(m.records) > 0 {
		m.Select(m.records[0].ID)
	}

	m.refreshDetached()
}

func (m *LogTree) makeNode(key, record, label string, open bool, previous map[string]bool) *tree.Node {
	node := tree.Root(logNode{key, record, label})
	if old, ok := previous[key]; ok {
		open = old
	}

	if !open && m.Folding {
		node.Close()
	}

	m.nodes[key] = node

	return node
}

func (m *LogTree) entryLabel(record LogRecord, available int) string {
	stamp := text.TimeOfDay(record.Time)

	var style lipgloss.Style

	switch record.Level {
	case log.DebugLevel:
		style = logStyle(m.Theme, m.Theme.Muted)
	case log.InfoLevel:
		style = logStyle(m.Theme, m.Theme.Info)
	case log.WarnLevel:
		style = logStyle(m.Theme, m.Theme.Warning)
	case log.ErrorLevel, log.FatalLevel:
		style = logStyle(m.Theme, m.Theme.Danger)
	default:
		style = m.Theme.Text()
	}

	mark := text.Log("", m.Theme, m.Symbols)
	if mark != "" {
		mark += " "
	}

	prefix := mark + style.Width(5).Render(strings.ToUpper(record.Level.String())) + " "
	if m.Model.Width() >= 64 {
		prefix = m.Theme.Subtle().Render(stamp) + " " + prefix
	}

	if m.DetachedDetails {
		return ansi.Truncate(prefix+m.Theme.Text().Render(text.Inline(record.Message)), available, "…")
	}

	return hangingWrap(prefix, m.Theme.Text().Render(text.Inline(record.Message)), available)
}

func (m *LogTree) traceNodes(root *tree.Node, opened map[string]bool) (map[string]*tree.Node, map[string]int) {
	result := map[string]*tree.Node{}
	depths := map[string]int{}

	if !m.GroupTraces {
		for _, r := range m.records {
			result[r.ID] = root
		}

		return result, depths
	}

	attached := map[string]bool{}
	traces := map[string]*tree.Node{}
	spans := map[string]*tree.Node{}
	parentIDs := map[string]string{}
	conflicts := map[string]bool{}
	spanTrace := map[string]string{}
	spanNames := map[string]string{}
	traceKey := func(r LogRecord) string { return strconv.Quote(r.TraceID) }
	spanKey := func(trace, span string) string { return trace + "/" + strconv.Quote(span) }

	for _, r := range m.records {
		if r.TraceID == "" {
			continue
		}

		tk := traceKey(r)
		if traces[tk] == nil {
			traces[tk] = m.makeNode("trace/"+tk, "", text.Trace("Trace", m.Theme, m.Symbols)+" "+text.Inline(r.TraceID), true, opened)
			root.Child(traces[tk])
			m.parentNodes[traces[tk]] = root
		}

		if r.SpanID != "" {
			sk := spanKey(tk, r.SpanID)
			if spans[sk] == nil {
				spans[sk] = m.makeNode("span/"+sk, "", text.Span("Span", m.Theme, m.Symbols)+" "+text.Inline(r.SpanID), true, opened)
				parentIDs[sk] = r.ParentSpanID
				spanTrace[sk] = tk
				spanNames[sk] = r.SpanID
			} else if parentIDs[sk] != r.ParentSpanID {
				conflicts[sk] = true
			}
		}
	}

	invalidParents := map[string]bool{}

	for sk, parentID := range parentIDs {
		seen := map[string]bool{sk: true}
		invalid := conflicts[sk]

		ancestor := parentID
		for ancestor != "" {
			ak := spanKey(spanTrace[sk], ancestor)
			if seen[ak] || spans[ak] == nil || conflicts[ak] {
				invalid = true
				break
			}

			seen[ak] = true
			ancestor = parentIDs[ak]
		}

		invalidParents[sk] = invalid
		if invalid {
			spans[sk] = m.makeNode("span/"+sk, "", text.Span("Span", m.Theme, m.Symbols)+" "+text.Inline(spanNames[sk])+" [unresolved parent]", true, opened)
		}
	}

	for _, r := range m.records {
		if !m.GroupTraces {
			result[r.ID] = root
			continue
		}

		if r.TraceID == "" {
			result[r.ID] = root
			continue
		}

		tk := traceKey(r)
		if r.SpanID == "" {
			result[r.ID] = traces[tk]
			depths[r.ID] = 1

			continue
		}

		sk := spanKey(tk, r.SpanID)

		node := spans[sk]
		if !attached[sk] {
			parent := traces[tk]
			if !invalidParents[sk] {
				if p := spans[spanKey(tk, parentIDs[sk])]; p != nil {
					parent = p
				}
			}

			parent.Child(node)
			m.parentNodes[node] = parent

			attached[sk] = true
		}

		result[r.ID] = node
		depths[r.ID] = 2

		if !invalidParents[sk] {
			for ancestor := parentIDs[sk]; ancestor != ""; ancestor = parentIDs[spanKey(tk, ancestor)] {
				depths[r.ID]++
			}
		}
	}

	return result, depths
}

func (m *LogTree) field(parent *tree.Node, key, record, name string, value any, depth, available int, budget *int, opened map[string]bool) {
	if *budget <= 0 {
		return
	}

	*budget--

	if depth >= max(1, m.FieldDepth) {
		parent.Child(text.Inline(name) + ": [depth limit]")
		return
	}

	label := logStyle(m.Theme, m.Theme.Syntax.Name).Render(name)

	switch v := value.(type) {
	case map[string]any:
		node := m.makeNode(key, record, label+fmt.Sprintf(" {%d}", len(v)), false, opened)
		parent.Child(node)
		m.parentNodes[node] = parent

		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}

		slices.Sort(keys)

		for _, name := range keys {
			m.field(node, key+"/"+strconv.Quote(name), record, text.Inline(name), v[name], depth+1, available-3, budget, opened)

			if *budget <= 0 {
				node.Child("[field node limit]")
				break
			}
		}
	case []any:
		node := m.makeNode(key, record, label+fmt.Sprintf(" [%d]", len(v)), false, opened)
		parent.Child(node)
		m.parentNodes[node] = parent

		for i, item := range v {
			m.field(node, key+"/"+strconv.Itoa(i), record, fmt.Sprintf("[%d]", i), item, depth+1, available-3, budget, opened)

			if *budget <= 0 {
				node.Child("[field node limit]")
				break
			}
		}
	default:
		value := fieldLabel(v)

		node := m.makeNode(key, record, hangingWrap(label+": ", logStyle(m.Theme, m.Theme.Syntax.String).Render(value), available), true, opened)
		parent.Child(node)
		m.parentNodes[node] = parent
	}
}

func entryKey(id string) string { return "entry/" + strconv.Quote(id) }

func fieldLabel(value any) string {
	var result string

	switch v := value.(type) {
	case nil:
		result = "null"
	case string:
		if len(v) > DefaultFieldValueBytes {
			v = v[:DefaultFieldValueBytes] + " [value truncated]"
		}

		return text.Inline(strconv.Quote(v))
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, json.Number:
		result = fmt.Sprint(v)
	case time.Time:
		result = v.Format(time.RFC3339Nano)
	default:
		result = fmt.Sprintf("<%T>", v)
	}

	return text.Inline(result)
}

func fieldName(value any) string {
	if name, ok := value.(string); ok {
		return text.Inline(name)
	}

	return fieldLabel(value)
}

func fieldIdentity(value any) string {
	if raw, ok := value.(string); ok {
		return "string:" + strconv.Quote(raw)
	}

	return fmt.Sprintf("%T:%s", value, fieldLabel(value))
}

func (m *LogTree) recordFields(node *tree.Node, record LogRecord, available int, opened map[string]bool) {
	budget := max(1, m.FieldNodes)
	if !record.Time.IsZero() {
		m.field(node, entryKey(record.ID)+"/metadata/time", record.ID, "time", record.Time, 0, available-3, &budget, opened)
	}

	occurrences := map[string]int{}

	for i := 0; i < len(record.Fields); i += 2 {
		name := fieldName(record.Fields[i])

		var value any = "[unpaired field]"
		if i+1 < len(record.Fields) {
			value = record.Fields[i+1]
		}

		identity := fieldIdentity(record.Fields[i])
		occurrence := occurrences[identity]
		occurrences[identity]++
		m.field(node, entryKey(record.ID)+"/field/"+strconv.Quote(identity)+"/"+strconv.Itoa(occurrence), record.ID, name, value, 0, available-3, &budget, opened)

		if budget <= 0 {
			node.Child("[field node limit]")
			break
		}
	}
}
