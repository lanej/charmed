package collection

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// InspectorSizes allocates a selection pane and a detail pane. At very short
// heights only the focused pane receives space; the parent retains Tab routing.
func InspectorSizes(size layout.Size, detailFocused bool) (layout.Size, layout.Size) {
	size.Width, size.Height = max(0, size.Width), max(0, size.Height)
	if size.Height < 8 {
		if detailFocused {
			return layout.Size{}, size
		}

		return size, layout.Size{}
	}

	if size.Width >= 90 {
		left := min(size.Width-30, size.Width*3/5)
		return layout.Size{Width: left, Height: size.Height}, layout.Size{Width: size.Width - left, Height: size.Height}
	}

	return layout.Size{Width: size.Width, Height: size.Height / 2}, layout.Size{Width: size.Width, Height: size.Height - size.Height/2}
}

// InspectorView joins already sized panes using the same geometry as InspectorSizes.
func InspectorView(left, detail string, size layout.Size, detailFocused bool) string {
	a, b := InspectorSizes(size, detailFocused)
	if a.Width == 0 || a.Height == 0 {
		return layout.Fit(detail, b)
	}

	if b.Width == 0 || b.Height == 0 {
		return layout.Fit(left, a)
	}

	if size.Width >= 90 {
		return lipgloss.JoinHorizontal(lipgloss.Top, layout.Fit(left, a), layout.Fit(detail, b))
	}

	return lipgloss.JoinVertical(lipgloss.Left, layout.Fit(left, a), layout.Fit(detail, b))
}

// GridInspector coordinates a native grid and scrollable details. Direct mutations
// of Grid require Refresh. The renderer receives missing cells with present=false.
type GridInspector[T any] struct {
	Grid            *DataGrid[T]
	Details         *ScrollPane
	RenderDetail    func(GridEntry[T], bool, int) string
	size            layout.Size
	detailFocused   bool
	rowID, columnID string
}

func NewGridInspector[T any](grid *DataGrid[T], render func(GridEntry[T], bool, int) string, size layout.Size, t theme.Theme) *GridInspector[T] {
	native := NewViewport("", 1, 1)
	native.SoftWrap = true
	m := &GridInspector[T]{Grid: grid, Details: NewScrollPane(&native, size, t), RenderDetail: render}
	m.Resize(size)
	m.FocusDetail(false)

	return m
}

func (m *GridInspector[T]) Resize(size layout.Size) {
	m.size = size
	a, b := InspectorSizes(size, m.detailFocused)
	m.Grid.Resize(a)
	m.Details.Resize(b)
	m.Refresh()
}

func (m *GridInspector[T]) FocusDetail(detail bool) {
	m.detailFocused = detail
	if detail {
		m.Grid.Blur()
	} else {
		m.Grid.Focus()
	}

	m.Resize(m.size)
}

func (m *GridInspector[T]) DetailFocused() bool { return m.detailFocused }

func (m *GridInspector[T]) Refresh() {
	entry, present := m.Grid.Selected()
	changed := entry.RowID != m.rowID || entry.ColumnID != m.columnID
	m.rowID, m.columnID = entry.RowID, entry.ColumnID

	content := ""
	if m.RenderDetail != nil {
		content = m.RenderDetail(entry, present, m.Details.Model.Width())
	}

	m.Details.SetContent(content)

	if changed {
		m.Details.Model.GotoTop()
	}
}

func (m *GridInspector[T]) SetData(rows, columns []GridAxis, entries []GridEntry[T]) error {
	if err := m.Grid.SetData(rows, columns, entries); err != nil {
		return err
	}

	m.Refresh()

	return nil
}

func (m *GridInspector[T]) Select(rowID, columnID string) bool {
	ok := m.Grid.Select(rowID, columnID)
	m.Refresh()

	return ok
}

func (m *GridInspector[T]) Selected() (GridEntry[T], bool) { return m.Grid.Selected() }

func (m *GridInspector[T]) Update(msg tea.Msg) tea.Cmd {
	if m.detailFocused {
		return m.Details.Update(msg)
	}

	_, cmd := m.Grid.Update(msg)
	m.Refresh()

	return cmd
}

func (m *GridInspector[T]) View() string {
	return InspectorView(m.Grid.View(), m.Details.View(), m.size, m.detailFocused)
}
