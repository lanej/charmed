package story

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func dataGridDefinition() Definition {
	return Definition{ID: "collection/data-grid", Title: "Organisms/Collections/Data grid", Level: "Organisms", Family: "collection", Description: "Typed values, stable axes, pinned headers and two-axis keyboard navigation. Plain values render directly; caller renderers own cell meaning and styling. Instructions are opt-in.", Source: "collection/data_grid.go", Symbol: "DataGrid", ExampleSource: "docs/snippets/collection_data_grid.go", Controls: []string{"state"}, States: []string{"ready", "empty", "loading", "error"}, Composes: []string{"foundations/layout", "text/primitives"}, Defaults: DefaultConfig(), Factory: dataGridStory}
}

func dataGridStory(e Environment) (runtime.Scene, error) {
	rows := []collection.GridAxis{{ID: "api", Label: "API"}, {ID: "worker", Label: "Workers"}, {ID: "cache", Label: "Cache"}}
	columns := []collection.GridAxis{{ID: "east", Label: "US East"}, {ID: "west", Label: "US West"}, {ID: "eu", Label: "Europe"}, {ID: "ap", Label: "Asia Pacific"}}

	entries := []collection.GridEntry[int]{
		{RowID: "api", ColumnID: "east", Value: 24},
		{RowID: "api", ColumnID: "west", Value: 16},
		{RowID: "api", ColumnID: "eu", Value: 12},
		{RowID: "api", ColumnID: "ap", Value: 8},
		{RowID: "worker", ColumnID: "east", Value: 48},
		{RowID: "worker", ColumnID: "west", Value: 32},
		{RowID: "worker", ColumnID: "eu", Value: 16},
		{RowID: "cache", ColumnID: "east", Value: 6},
		{RowID: "cache", ColumnID: "west", Value: 4},
		{RowID: "cache", ColumnID: "eu", Value: 0},
	}
	if e.Config.State == "empty" {
		rows, columns, entries = nil, nil, nil
	}

	grid, err := collection.NewDataGrid(rows, columns, entries, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	grid.Corner = "Service / Region"
	grid.HidePosition = e.Config.Static
	grid.Focus()

	return &scene{resize: func(size layout.Size) {
		grid.Resize(layout.Size{Width: size.Width, Height: min(len(rows)+2-boolInt(e.Config.Static), max(0, size.Height-4))})
	}, update: func(msg tea.Msg) tea.Cmd {
		_, cmd := grid.Update(msg)

		return cmd
	}, view: func(size layout.Size) string {
		if message := stateMessage(e); message != "" && e.Config.State != "empty" {
			return title("Replica counts by region", e) + message
		}

		entry, present := grid.Selected()

		detail := "No selection"
		if entry.RowID != "" {
			detail = entry.RowID + " / " + entry.ColumnID + ": absent"
			if present {
				detail = fmt.Sprintf("%s / %s: %d replicas", entry.RowID, entry.ColumnID, entry.Value)
			}
		}

		return title("Replica counts by region", e) + grid.View() + "\n" + layout.Fit(detail, layout.Size{Width: size.Width, Height: 1})
	}}, nil
}

func boolInt(b bool) int {
	if b {
		return 1
	}

	return 0
}
