package collection

import (
	"bytes"
	"crypto/sha256"
	"maps"
	"slices"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// JSONViewer adds structural navigation to a native viewport. Model may point
// to a caller-owned viewport. Refresh owns content but retains native geometry,
// styles and key maps. Up/Down/Left/Right and paging use Model.KeyMap; Toggle is
// separate. Configure fields before SetSource or Refresh. Source bytes are copied.
type JSONViewer struct {
	Model                      *viewport.Model
	Theme                      theme.Theme
	Limits                     JSONLimits
	Toggle                     key.Binding
	Symbols                    text.Presentation
	Prefix                     string
	Raw, Folding, Focused      bool
	source                     []byte
	fingerprint                [sha256.Size]byte
	parsedLimits               JSONLimits
	nodes                      []jsonNode
	closed                     map[int]bool
	selected                   int
	visible                    []int
	rows                       map[int]int
	parseErr, renderErr        error
	wasRaw                     bool
	formattedOffset, rawOffset int
	lastWidth                  int
}

// JSONState is a source-bound snapshot of folds, cursor and scroll positions.
// It contains no captured bytes or native styles/key maps. RestoreState refuses
// a different document. Hosts own retention and must evict discarded records.
type JSONState struct {
	fingerprint                          [sha256.Size]byte
	closed                               map[int]bool
	selected, formattedOffset, rawOffset int
	valid                                bool
}

func NewJSONViewer(width, height int, t theme.Theme) *JSONViewer {
	model := NewViewport("", width, height)
	return &JSONViewer{Model: &model, Theme: t, Limits: DefaultJSONLimits(), Folding: true, Focused: true, Toggle: key.NewBinding(key.WithKeys("enter"), key.WithHelp("Enter", "fold JSON")), closed: map[int]bool{}}
}

// SetSource replaces the document and resets structural state when bytes or
// limits change. Invalid/over-budget input clears stale structure, displays a
// bounded source fallback and returns the diagnostic. Identical bytes retain
// state. Number lexemes, member order and duplicate keys are never normalized.
func (m *JSONViewer) SetSource(source []byte) error {
	limits := m.Limits.normalized()

	fingerprint := sha256.Sum256(source)
	if fingerprint != m.fingerprint || m.parsedLimits != limits || m.source == nil {
		m.source = bytes.Clone(source[:min(len(source), limits.Bytes)])
		m.fingerprint, m.parsedLimits = fingerprint, limits
		m.nodes, m.parseErr = parseJSON(source, limits)
		m.closed, m.selected = map[int]bool{}, 0
		m.formattedOffset, m.rawOffset, m.wasRaw = 0, 0, m.Raw
		m.Model.GotoTop()
	}

	m.Refresh()

	return m.Err()
}

// Clear forgets a document without touching the host viewport's content.
func (m *JSONViewer) Clear() {
	m.source, m.nodes, m.visible, m.rows = nil, nil, nil, nil
	m.closed = map[int]bool{}
	m.parseErr, m.renderErr = nil, nil
	m.fingerprint = [sha256.Size]byte{}
	m.selected, m.formattedOffset, m.rawOffset = 0, 0, 0
}

func (m *JSONViewer) Err() error {
	if m.parseErr != nil {
		return m.parseErr
	}

	return m.renderErr
}

func (m *JSONViewer) Interactive() bool {
	return len(m.nodes) > 0 && m.Err() == nil && m.Folding && !m.Raw
}

func (m *JSONViewer) CanFold() bool {
	return m.Interactive() && len(m.nodes[m.selected].children) > 0
}

// Selected returns the preorder node index, stable only for identical source.
func (m *JSONViewer) Selected() int           { return m.selected }
func (m *JSONViewer) Collapsed(node int) bool { return m.Folding && m.closed[node] }

func (m *JSONViewer) Select(node int) bool {
	if node < 0 || node >= len(m.nodes) {
		return false
	}

	for parent := m.nodes[node].parent; parent >= 0; parent = m.nodes[parent].parent {
		delete(m.closed, parent)
	}

	m.selected = node
	m.Refresh()
	m.reveal()

	return true
}

func (m *JSONViewer) SetCollapsed(node int, collapsed bool) bool {
	if !m.Folding || node < 0 || node >= len(m.nodes) || len(m.nodes[node].children) == 0 {
		return false
	}

	if collapsed {
		m.closed[node] = true
		for ancestor := m.selected; ancestor >= 0; ancestor = m.nodes[ancestor].parent {
			if ancestor == node {
				m.selected = node
				break
			}
		}
	} else {
		delete(m.closed, node)
	}

	m.Refresh()
	m.reveal()

	return true
}

func (m *JSONViewer) Resize(width, height int) {
	m.Model.SetWidth(max(1, width))
	m.Model.SetHeight(max(1, height))
	m.Refresh()
}

// ResizeViewport updates geometry without refreshing presentation on height-only
// changes. Use Refresh after changing exposed presentation fields.
func (m *JSONViewer) ResizeViewport(width, height int) {
	changed := m.Model.Width() != max(1, width)
	m.Model.SetWidth(max(1, width))
	m.Model.SetHeight(max(1, height))

	if changed {
		m.Refresh()
	} else {
		m.Model.SetYOffset(m.Model.YOffset())
	}
}

func (m *JSONViewer) View() string { return m.Model.View() }

func (m *JSONViewer) Snapshot() JSONState {
	state := JSONState{fingerprint: m.fingerprint, closed: maps.Clone(m.closed), selected: m.selected, formattedOffset: m.formattedOffset, rawOffset: m.rawOffset, valid: len(m.nodes) > 0}
	if m.Raw {
		state.rawOffset = m.Model.YOffset()
	} else {
		state.formattedOffset = m.Model.YOffset()
	}

	return state
}

func (m *JSONViewer) RestoreState(state JSONState) bool {
	if !state.valid || state.fingerprint != m.fingerprint || state.selected >= len(m.nodes) {
		return false
	}

	m.closed = maps.Clone(state.closed)
	m.selected = state.selected
	m.formattedOffset, m.rawOffset = state.formattedOffset, state.rawOffset
	m.Refresh()

	offset := m.formattedOffset
	if m.Raw {
		offset = m.rawOffset
	}

	m.Model.SetYOffset(offset)

	return true
}

func (m *JSONViewer) reveal() {
	row, ok := m.rows[m.selected]
	if !ok {
		return
	}

	if row < m.Model.YOffset() {
		m.Model.SetYOffset(row)
	} else if row >= m.Model.YOffset()+m.Model.Height() {
		m.Model.SetYOffset(row - m.Model.Height() + 1)
	}
}

func (m *JSONViewer) move(delta int) {
	position := slices.Index(m.visible, m.selected)

	position = min(max(0, position+delta), len(m.visible)-1)
	if position >= 0 {
		m.selected = m.visible[position]
		m.Refresh()
		m.reveal()
	}
}
