package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionDataGrid() (string, error) {
	rows := []collection.GridAxis{{ID: "api", Label: "API"}, {ID: "worker", Label: "Workers"}}
	regions := []collection.GridAxis{{ID: "east", Label: "US East"}, {ID: "west", Label: "US West"}}
	entries := []collection.GridEntry[int]{
		{RowID: "api", ColumnID: "east", Value: 24},
		{RowID: "api", ColumnID: "west", Value: 0},
		{RowID: "worker", ColumnID: "east", Value: 48},
	}

	grid, err := collection.NewDataGrid(rows, regions, entries, layout.Size{Width: 54, Height: 4}, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	grid.Corner = "Service / Region"
	grid.Focus()
	// The parent forwards Update(msg), returns its command and calls Resize(size).
	// Read Selected() for stable IDs and presence; zero is a present value.
	// Set ShowHelp to opt into instructions, or RenderCell for custom styling.
	return grid.View(), nil
}
