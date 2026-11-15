package collection

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// DocumentSection is content, not a file. Sibling order is preserved. Labels
// accept styled text; Body receives its available cell width and may return ANSI.
// Collapsed is an initial value, used only for IDs absent from the previous data.
type DocumentSection struct {
	ID, Title, Summary string
	Body               func(width int) string
	Children           []DocumentSection
	Collapsed          bool
}

type sectionRange struct {
	id         string
	start, end int
}

// SectionStack owns projected content, folds, dimensions and scroll anchoring.
// Native viewport bindings and styles survive SetSections and Resize. SoftWrap
// is managed here: bodies are hard-wrapped before offsets are calculated.
// Render callbacks are cached until SetSections, Refresh or a width change.
type SectionStack struct {
	Viewport viewport.Model
	Theme    theme.Theme
	sections []DocumentSection
	index    map[string]DocumentSection
	parents  map[string]string
	folded   map[string]bool
	bodies   map[string]string
	ranges   []sectionRange
	size     layout.Size
}

func NewSectionStack(sections []DocumentSection, size layout.Size, t theme.Theme) (*SectionStack, error) {
	m := &SectionStack{Viewport: NewViewport("", size.Width, size.Height), Theme: t, size: size}
	if err := m.SetSections(sections); err != nil {
		return nil, err
	}

	return m, nil
}

// SetSections validates and copies the complete hierarchy before replacing it.
// Retained IDs keep their folds; removed IDs lose state. Renderers remain
// caller-owned functions and must be deterministic until the next refresh.
func (m *SectionStack) SetSections(sections []DocumentSection) error {
	index := make(map[string]DocumentSection)
	parents := make(map[string]string)
	folded := make(map[string]bool)

	var copySections func([]DocumentSection, string) ([]DocumentSection, error)

	copySections = func(source []DocumentSection, parent string) ([]DocumentSection, error) {
		result := make([]DocumentSection, len(source))
		for i, section := range source {
			if _, exists := index[section.ID]; exists || section.ID == "" {
				return nil, fmt.Errorf("empty or duplicate section ID %q", section.ID)
			}

			index[section.ID] = section
			parents[section.ID] = parent

			folded[section.ID] = section.Collapsed
			if value, exists := m.folded[section.ID]; exists {
				folded[section.ID] = value
			}

			children, err := copySections(section.Children, section.ID)
			if err != nil {
				return nil, err
			}

			section.Children = children
			result[i] = section
			index[section.ID] = section
		}

		return result, nil
	}

	cloned, err := copySections(sections, "")
	if err != nil {
		return err
	}

	anchor := m.anchor()
	m.sections, m.index, m.parents, m.folded = cloned, index, parents, folded
	m.bodies = make(map[string]string)
	m.render()
	m.restore(anchor)

	return nil
}

// Sections returns a structural copy; callback closures are shared.
func (m *SectionStack) Sections() []DocumentSection {
	var clone func([]DocumentSection) []DocumentSection

	clone = func(source []DocumentSection) []DocumentSection {
		result := make([]DocumentSection, len(source))
		for i, entry := range source {
			entry.Children = clone(entry.Children)
			result[i] = entry
		}

		return result
	}

	return clone(m.sections)
}

func (m *SectionStack) Collapsed(id string) bool { return m.folded[id] }

func (m *SectionStack) SetCollapsed(id string, collapsed bool) bool {
	if _, exists := m.index[id]; !exists {
		return false
	}

	if m.folded[id] == collapsed {
		return true
	}

	anchor := m.anchor()
	m.folded[id] = collapsed
	m.render()
	m.restore(anchor)

	return true
}

// VisibleIDs returns depth-first headings, including collapsed summaries.
func (m *SectionStack) VisibleIDs() []string {
	ids := make([]string, len(m.ranges))
	for i, span := range m.ranges {
		ids[i] = span.id
	}

	return ids
}

// ScrollTo reveals a heading without opening it. Hidden descendants resolve to
// the nearest visible ancestor. Native viewport clamping applies near the end.
func (m *SectionStack) ScrollTo(id string) bool {
	for id != "" {
		for _, span := range m.ranges {
			if span.id == id {
				m.Viewport.SetYOffset(span.start)
				return true
			}
		}

		id = m.parents[id]
	}

	return false
}

func (m *SectionStack) TopID() string { return m.IDAt(m.Viewport.YOffset()) }

// IDAt maps an absolute rendered row to a section, including its body.
func (m *SectionStack) IDAt(row int) string {
	for _, span := range m.ranges {
		if row >= span.start && row < span.end {
			return span.id
		}
	}

	return ""
}

// HeaderAt maps a viewport-local row to a heading; body rows return an empty ID.
func (m *SectionStack) HeaderAt(row int) string {
	for _, span := range m.ranges {
		if span.start == row+m.Viewport.YOffset() {
			return span.id
		}
	}

	return ""
}

func (m *SectionStack) Resize(size layout.Size) {
	anchor := m.anchor()
	widthChanged := size.Width != m.size.Width
	m.size = size
	m.Viewport.SetWidth(max(1, size.Width))
	m.Viewport.SetHeight(max(1, size.Height))

	if widthChanged {
		m.bodies = make(map[string]string)
		m.render()
	}

	m.restore(anchor)
}

// Refresh invalidates rendered bodies after caller-owned presentation changes.
func (m *SectionStack) Refresh() {
	anchor := m.anchor()
	m.bodies = make(map[string]string)
	m.render()
	m.restore(anchor)
}

func (m *SectionStack) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	m.Viewport, cmd = m.Viewport.Update(msg)

	return cmd
}

func (m *SectionStack) View() string {
	return layout.Fit(m.Viewport.View(), m.size)
}

func (m *SectionStack) render() {
	width := max(1, m.size.Width)

	var lines []string

	m.ranges = nil

	var visit func([]DocumentSection)

	visit = func(sections []DocumentSection) {
		for _, section := range sections {
			start := len(lines)

			heading := text.Disclosure(section.Title, !m.folded[section.ID], m.Theme)
			if section.Summary != "" {
				heading += "  " + m.Theme.Subtle().Render(section.Summary)
			}

			lines = append(lines, ansi.Truncate(m.Theme.Heading().Render(heading), width, "…"))
			if !m.folded[section.ID] && section.Body != nil {
				body, cached := m.bodies[section.ID]
				if !cached {
					body = ansi.Hardwrap(section.Body(width), width, true)
					m.bodies[section.ID] = body
				}

				if body != "" {
					lines = append(lines, strings.Split(body, "\n")...)
				}
			}

			if len(section.Children) == 0 && !m.folded[section.ID] {
				lines = append(lines, "")
			}

			m.ranges = append(m.ranges, sectionRange{id: section.ID, start: start, end: len(lines)})
			if !m.folded[section.ID] {
				visit(section.Children)
			}
		}
	}

	visit(m.sections)

	if len(lines) == 0 {
		lines = []string{m.Theme.Subtle().Render("No sections")}
	}

	content := strings.Join(lines, "\n")
	if m.Theme.ASCII {
		content = layout.ASCII(content)
	}

	m.Viewport.SoftWrap = false
	m.Viewport.SetContent(content)
}

type stackAnchor struct {
	id         string
	row        int
	candidates []string
}

func (m *SectionStack) anchor() stackAnchor {
	id := m.TopID()

	anchor := stackAnchor{id: id}
	for _, span := range m.ranges {
		if span.id == id {
			anchor.row = m.Viewport.YOffset() - span.start
			break
		}

		anchor.candidates = append([]string{span.id}, anchor.candidates...)
	}

	var ancestors []string
	for parent := m.parents[id]; parent != ""; parent = m.parents[parent] {
		ancestors = append(ancestors, parent)
	}

	anchor.candidates = append(ancestors, anchor.candidates...)

	return anchor
}

func (m *SectionStack) restore(anchor stackAnchor) {
	for _, id := range append([]string{anchor.id}, anchor.candidates...) {
		for _, span := range m.ranges {
			if span.id == id {
				row := 0
				if id == anchor.id {
					row = min(anchor.row, span.end-span.start-1)
				}

				m.Viewport.SetYOffset(span.start + row)

				return
			}
		}
	}

	m.Viewport.GotoTop()
}

func (m *SectionStack) Init() tea.Cmd { return nil }
func (m *SectionStack) Close() error  { return nil }
