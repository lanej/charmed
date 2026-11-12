package collection

import (
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// MatrixAxis retains the status-matrix axis API; grids share the same identity model.
type MatrixAxis = GridAxis

// MatrixEntry is a present relationship. Omitted pairs mean no mapping;
// a present cell with UNKNOWN status means the relationship's health is unknown.
type MatrixEntry struct {
	RowID, ColumnID string
	Cell            feedback.StatusCell
}

// StatusMatrix specializes DataGrid with semantic health cells and a single
// unavailable glyph for absent mappings. DataGrid remains accessible for customization.
// Its theme, geometry, help, renderer and selection survive SetData and Resize.
type StatusMatrix struct {
	*DataGrid[feedback.StatusCell]
}

func NewStatusMatrix(rows, columns []MatrixAxis, entries []MatrixEntry, size layout.Size, t theme.Theme) (*StatusMatrix, error) {
	grid, err := NewDataGrid(rows, columns, statusGridEntries(entries), size, t)
	if err != nil {
		return nil, err
	}

	grid.Corner, grid.EmptyLabel, grid.RenderCell = "Mapping", "No mappings", statusGridCell

	return &StatusMatrix{DataGrid: grid}, nil
}

// SetData validates and copies replacements atomically through the grid.
func (m *StatusMatrix) SetData(rows, columns []MatrixAxis, entries []MatrixEntry) error {
	return m.DataGrid.SetData(rows, columns, statusGridEntries(entries))
}

// Selected preserves the status-specific API, including IDs for absent mappings.
func (m *StatusMatrix) Selected() (MatrixEntry, bool) {
	entry, present := m.DataGrid.Selected()

	return MatrixEntry{RowID: entry.RowID, ColumnID: entry.ColumnID, Cell: entry.Value}, present
}

func (m *StatusMatrix) Update(msg tea.Msg) (*StatusMatrix, tea.Cmd) {
	_, cmd := m.DataGrid.Update(msg)

	return m, cmd
}

func statusGridEntries(entries []MatrixEntry) []GridEntry[feedback.StatusCell] {
	result := make([]GridEntry[feedback.StatusCell], len(entries))
	for i, entry := range entries {
		result[i] = GridEntry[feedback.StatusCell]{RowID: entry.RowID, ColumnID: entry.ColumnID, Value: entry.Cell}
	}

	return result
}

func statusGridCell(cell GridCell[feedback.StatusCell]) string {
	if cell.Present {
		return cell.Entry.Value.View(cell.Width, cell.Theme, cell.Selected)
	}

	left, right := " ", " "
	if cell.Selected {
		left, right = cell.Theme.Glyph(theme.GlyphFocusStart), cell.Theme.Glyph(theme.GlyphFocusEnd)
	}

	return cell.Theme.Subtle().Render(left + layout.Fit(cell.Theme.Glyph(theme.GlyphUnavailable), layout.Size{Width: cell.Width - 2, Height: 1}) + right)
}
