package collection

import (
	"fmt"
	"slices"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

// KeyedRow keeps a stable identity separate from native display columns.
type KeyedRow[K comparable] struct {
	Key K
	Row table.Row
}

// KeyedTable manages rows and cursor on an existing native table. Styles, key
// bindings and focus survive replacement; zero-valued keys are valid identities.
type KeyedTable[K comparable] struct {
	Model *table.Model
	keys  []K
}

func NewKeyedTable[K comparable](model *table.Model, rows []KeyedRow[K]) (*KeyedTable[K], error) {
	m := &KeyedTable[K]{Model: model}
	if err := m.SetRows(rows); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *KeyedTable[K]) SetRows(rows []KeyedRow[K]) error { return m.SetData(m.Model.Columns(), rows) }

// SetData validates the complete replacement before touching the native model.
func (m *KeyedTable[K]) SetData(columns []table.Column, rows []KeyedRow[K]) error {
	seen := make(map[K]bool, len(rows))
	native := make([]table.Row, len(rows))

	keys := make([]K, len(rows))
	for i, row := range rows {
		if seen[row.Key] {
			return fmt.Errorf("duplicate table key %v", row.Key)
		}

		if len(row.Row) != len(columns) {
			return fmt.Errorf("row %d has %d cells; expected %d", i, len(row.Row), len(columns))
		}

		seen[row.Key] = true
		keys[i], native[i] = row.Key, slices.Clone(row.Row)
	}

	selected, present := m.Selected()
	m.Model.SetRows(nil)
	m.Model.SetColumns(slices.Clone(columns))
	m.Model.SetRows(native)
	m.keys = keys
	m.Model.GotoTop()

	if present {
		m.Select(selected)
	}

	return nil
}

func (m *KeyedTable[K]) Select(key K) bool {
	for i, k := range m.keys {
		if k == key {
			m.Model.GotoTop()
			m.Model.MoveDown(i)

			return true
		}
	}

	return false
}

func (m *KeyedTable[K]) Selected() (K, bool) {
	i := m.Model.Cursor()
	if i >= 0 && i < len(m.keys) {
		return m.keys[i], true
	}

	var zero K

	return zero, false
}

func (m *KeyedTable[K]) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	*m.Model, cmd = m.Model.Update(msg)

	return cmd
}

func (m *KeyedTable[K]) View() string { return m.Model.View() }
