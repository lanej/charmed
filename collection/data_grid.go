package collection

import (
	"fmt"
	"slices"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// GridAxis gives a row or column stable identity independent of its label.
type GridAxis struct {
	ID, Label string
}

// GridEntry is a present value at a row/column pair. Presence is independent of
// T's zero value. The grid assigns no meaning to omitted pairs.
type GridEntry[T any] struct {
	RowID, ColumnID string
	Value           T
}

// GridCell supplies a renderer with the full typed value and its allocation.
// Absent cells still carry their axis IDs. Renderers own content and focus
// styling; the grid clips and pads the result to Width cells on one line.
type GridCell[T any] struct {
	Entry             GridEntry[T]
	Present, Selected bool
	Width             int
	Theme             theme.Theme
}

type gridPair struct{ row, column string }

// DataGrid owns two-axis navigation, pinned labels and bounded cell geometry.
// SetData copies axes and entries by value; references within T remain caller-owned.
// Selection follows IDs on replacement; a removed coordinate resets to the first
// axis item. Resize and SetData preserve theme, renderer and geometry settings.
type DataGrid[T any] struct {
	Theme          theme.Theme
	Corner         string
	CellWidth      int
	RowHeaderWidth int
	EmptyLabel     string
	// RenderCell defaults to plain text values and blank absent cells.
	RenderCell func(GridCell[T]) string
	// ShowHelp adds keyboard instructions below the range footer; off by default.
	ShowHelp bool
	// HidePosition omits the "Rows x-y/n" line when every row and column is
	// visible, as in a static render that shows all rows. When anything is
	// scrolled out of view the line remains so the omission stays visible.
	// The default keeps the line.
	HidePosition  bool
	rows, columns []GridAxis
	cells         map[gridPair]T
	row, column   int
	size          layout.Size
	focused       bool
}

func NewDataGrid[T any](rows, columns []GridAxis, entries []GridEntry[T], size layout.Size, t theme.Theme) (*DataGrid[T], error) {
	m := &DataGrid[T]{Theme: t, Corner: "Grid", EmptyLabel: "No data", CellWidth: 16, RowHeaderWidth: 18, size: size}
	if err := m.SetData(rows, columns, entries); err != nil {
		return nil, err
	}

	return m, nil
}

// SetData rejects empty/duplicate axis IDs, unknown references and duplicate
// pairs. An error leaves the existing data and selection unchanged.
func (m *DataGrid[T]) SetData(rows, columns []GridAxis, entries []GridEntry[T]) error {
	if err := validateGridAxes(rows); err != nil {
		return fmt.Errorf("grid rows: %w", err)
	}

	if err := validateGridAxes(columns); err != nil {
		return fmt.Errorf("grid columns: %w", err)
	}

	cells := make(map[gridPair]T, len(entries))
	for _, entry := range entries {
		if axisIndex(rows, entry.RowID) < 0 || axisIndex(columns, entry.ColumnID) < 0 {
			return fmt.Errorf("grid pair %q/%q references an unknown axis", entry.RowID, entry.ColumnID)
		}

		pair := gridPair{entry.RowID, entry.ColumnID}
		if _, exists := cells[pair]; exists {
			return fmt.Errorf("duplicate grid pair %q/%q", entry.RowID, entry.ColumnID)
		}

		cells[pair] = entry.Value
	}

	previous, _ := m.Selected()
	m.rows, m.columns, m.cells = slices.Clone(rows), slices.Clone(columns), cells
	m.row = max(0, axisIndex(rows, previous.RowID))
	m.column = max(0, axisIndex(columns, previous.ColumnID))

	return nil
}

func validateGridAxes(axes []GridAxis) error {
	seen := make(map[string]bool, len(axes))
	for _, axis := range axes {
		if axis.ID == "" || seen[axis.ID] {
			return fmt.Errorf("empty or duplicate axis ID %q", axis.ID)
		}

		seen[axis.ID] = true
	}

	return nil
}

func axisIndex(axes []GridAxis, id string) int {
	return slices.IndexFunc(axes, func(axis GridAxis) bool { return axis.ID == id })
}

func (m *DataGrid[T]) Resize(size layout.Size) { m.size = size }
func (m *DataGrid[T]) Focus()                  { m.focused = true }
func (m *DataGrid[T]) Blur()                   { m.focused = false }
func (m *DataGrid[T]) Focused() bool           { return m.focused }

// Select addresses a cell, including an absent value, by stable axis IDs.
// Unknown IDs leave selection unchanged.
func (m *DataGrid[T]) Select(rowID, columnID string) bool {
	r, c := axisIndex(m.rows, rowID), axisIndex(m.columns, columnID)
	if r < 0 || c < 0 {
		return false
	}

	m.row, m.column = r, c

	return true
}

// Selected returns axis IDs even for an absent pair (present=false). Empty axes
// return a zero entry and false. Values are shallow copies, as in SetData.
func (m *DataGrid[T]) Selected() (entry GridEntry[T], present bool) {
	if len(m.rows) == 0 || len(m.columns) == 0 {
		return GridEntry[T]{}, false
	}

	entry.RowID, entry.ColumnID = m.rows[m.row].ID, m.columns[m.column].ID
	entry.Value, present = m.cells[gridPair{entry.RowID, entry.ColumnID}]

	return entry, present
}

// Update handles arrows and Vim h/j/k/l while focused. Tab remains parent-owned.
func (m *DataGrid[T]) Update(msg tea.Msg) (*DataGrid[T], tea.Cmd) {
	k, ok := msg.(tea.KeyPressMsg)
	if !ok || !m.focused || len(m.rows) == 0 || len(m.columns) == 0 {
		return m, nil
	}

	switch k.String() {
	case "left", "h":
		m.column = max(0, m.column-1)
	case "right", "l":
		m.column = min(len(m.columns)-1, m.column+1)
	case "up", "k":
		m.row = max(0, m.row-1)
	case "down", "j":
		m.row = min(len(m.rows)-1, m.row+1)
	}

	return m, nil
}

func (m *DataGrid[T]) View() string {
	if len(m.rows) == 0 || len(m.columns) == 0 {
		return layout.Fit(m.Theme.Subtle().Render(text.CellLabel(m.EmptyLabel)), m.size)
	}

	headerWidth, cellWidth := max(3, m.RowHeaderWidth), max(5, m.CellWidth)
	columnCount := min(len(m.columns), (m.size.Width-headerWidth)/(cellWidth+1))

	helpHeight := 0
	if m.ShowHelp {
		helpHeight = 1
	}

	positionHeight := 1
	if m.HidePosition && len(m.columns) <= columnCount && len(m.rows) <= m.size.Height-1-helpHeight {
		positionHeight = 0
	}

	rowCount := min(len(m.rows), m.size.Height-1-positionHeight-helpHeight)
	if columnCount < 1 || rowCount < 1 {
		return layout.Fit("More space needed", m.size)
	}

	firstRow := min(max(0, m.row-rowCount+1), len(m.rows)-rowCount)
	firstColumn := min(max(0, m.column-columnCount+1), len(m.columns)-columnCount)

	lines := []string{gridHeading(m.Corner, headerWidth, m.Theme.Subtle(), m.Theme)}
	for _, column := range m.columns[firstColumn : firstColumn+columnCount] {
		lines[0] += " " + gridHeading(gridAxisLabel(column), cellWidth, m.Theme.Heading().Underline(true), m.Theme)
	}

	for r := firstRow; r < firstRow+rowCount; r++ {
		line := gridHeading(gridAxisLabel(m.rows[r]), headerWidth, m.Theme.Text(), m.Theme)
		for c := firstColumn; c < firstColumn+columnCount; c++ {
			cell, present := m.cells[gridPair{m.rows[r].ID, m.columns[c].ID}]
			selected := m.focused && r == m.row && c == m.column
			context := GridCell[T]{Entry: GridEntry[T]{RowID: m.rows[r].ID, ColumnID: m.columns[c].ID, Value: cell}, Present: present, Selected: selected, Width: cellWidth, Theme: m.Theme}

			renderer := m.RenderCell
			if renderer == nil {
				renderer = defaultGridCell[T]
			}

			value := layout.Fit(renderer(context), layout.Size{Width: cellWidth, Height: 1})

			line += " " + value
		}

		lines = append(lines, line)
	}

	if positionHeight > 0 {
		lines = append(lines, m.Theme.Subtle().Render(fmt.Sprintf("Rows %d-%d/%d  Columns %d-%d/%d", firstRow+1, firstRow+rowCount, len(m.rows), firstColumn+1, firstColumn+columnCount, len(m.columns))))
	}

	if m.ShowHelp {
		keys := m.Theme.Glyph(theme.GlyphUp) + " " + m.Theme.Glyph(theme.GlyphDown) + " " + m.Theme.Glyph(theme.GlyphLeft) + " " + m.Theme.Glyph(theme.GlyphRight)
		lines = append(lines, m.Theme.Subtle().Render(keys+" / h j k l: select and scroll"))
	}

	return layout.Fit(strings.Join(lines, "\n"), m.size)
}

func gridHeading(label string, width int, style lipgloss.Style, t theme.Theme) string {
	value := ansi.Truncate(text.CellLabel(label), width-2, t.Glyph(theme.GlyphTruncated))

	return style.Render(layout.Fit(" "+value, layout.Size{Width: width, Height: 1}))
}

func gridAxisLabel(axis GridAxis) string {
	if axis.Label == "" {
		return axis.ID
	}

	return axis.Label
}

func defaultGridCell[T any](cell GridCell[T]) string {
	label := ""
	if cell.Present {
		label = text.CellLabel(fmt.Sprint(cell.Entry.Value))
	}

	left, right := " ", " "
	if cell.Selected {
		left, right = cell.Theme.Glyph(theme.GlyphFocusStart), cell.Theme.Glyph(theme.GlyphFocusEnd)
	}

	return cell.Theme.Text().Render(left + layout.Fit(label, layout.Size{Width: cell.Width - 2, Height: 1}) + right)
}
