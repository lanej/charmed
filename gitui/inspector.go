package gitui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// ChangesInspector links a status selection to caller-supplied diffs.
// Its owner routes focus with SetFocus; Tab remains parent-owned.
type ChangesInspector struct {
	Status    *Status
	Diff      *DiffViewer
	diffs     map[changeIdentity]FileDiff
	size      layout.Size
	focusDiff bool
	selection changeIdentity
}

func NewChangesInspector(repo Repository, diffs []FileDiff, width, height int, t theme.Theme) (*ChangesInspector, error) {
	status, err := NewStatus(repo, width, height, t)
	if err != nil {
		return nil, err
	}

	m := &ChangesInspector{Status: status, Diff: NewDiffViewer(FileDiff{State: UnavailableDiff, Reason: "select a change"}, width, height, t)}
	if err := m.SetDiffs(diffs); err != nil {
		return nil, err
	}

	m.Resize(layout.Size{Width: width, Height: height})

	return m, nil
}

// SetDiffs copies all nested slices. Duplicate file/section identities are rejected.
func (m *ChangesInspector) SetDiffs(diffs []FileDiff) error {
	indexed := make(map[changeIdentity]FileDiff, len(diffs))
	for _, diff := range diffs {
		key := changeIdentity{id: diff.FileID, section: diff.Section}
		if key.id == "" {
			return fmt.Errorf("diff has no file ID")
		}

		if _, exists := indexed[key]; exists {
			return fmt.Errorf("duplicate diff for %q in %q", key.id, key.section)
		}

		indexed[key] = cloneDiff(diff)
	}

	m.diffs = indexed
	m.syncDiff(true)

	return nil
}

func cloneDiff(diff FileDiff) FileDiff {
	hunks := make([]Hunk, len(diff.Hunks))
	for i, hunk := range diff.Hunks {
		hunks[i] = Hunk{Header: hunk.Header, Lines: append([]DiffLine(nil), hunk.Lines...)}
	}

	diff.Hunks = hunks

	return diff
}

func (m *ChangesInspector) syncDiff(force bool) {
	id, section := m.Status.SelectedChange()

	key := changeIdentity{id: id, section: section}
	if !force && key == m.selection {
		return
	}

	m.selection = key

	diff, exists := m.diffs[key]
	if !exists {
		diff = FileDiff{Path: id, State: UnavailableDiff, Reason: "content was not supplied"}
	}

	m.Diff.SetDiff(diff)
}

func (m *ChangesInspector) SetFocus(diff bool) { m.focusDiff = diff; m.Status.List.SetFocused(!diff) }
func (m *ChangesInspector) Init() tea.Cmd      { return nil }
func (m *ChangesInspector) Close() error       { return nil }
func (m *ChangesInspector) Update(msg tea.Msg) tea.Cmd {
	if m.focusDiff {
		return m.Diff.Update(msg)
	}

	cmd := m.Status.Update(msg)
	m.syncDiff(false)

	return cmd
}

func (m *ChangesInspector) Resize(size layout.Size) {
	m.size = size
	left, right := inspectorSizes(size)
	m.Status.Resize(layout.PanelContentSize(left))
	m.Diff.Resize(layout.PanelContentSize(right))
}

func inspectorSizes(size layout.Size) (layout.Size, layout.Size) {
	if size.Width < 90 {
		return layout.Size{Width: size.Width, Height: size.Height / 2}, layout.Size{Width: size.Width, Height: size.Height - size.Height/2}
	}

	return layout.Size{Width: size.Width / 2, Height: size.Height}, layout.Size{Width: size.Width - size.Width/2, Height: size.Height}
}

func (m *ChangesInspector) View() string {
	m.syncDiff(false)
	left, right := inspectorSizes(m.size)
	a := layout.Panel("Status", m.Status.View(), left, m.Status.List.Theme, !m.focusDiff)

	b := layout.Panel("Diff", m.Diff.View(), right, m.Diff.Theme, m.focusDiff)
	if m.size.Width < 90 {
		return strings.Join([]string{a, b}, "\n")
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, a, b)
}
