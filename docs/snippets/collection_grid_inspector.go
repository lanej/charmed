package snippets

import (
	"fmt"

	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionGridInspector() (string, error) {
	size := layout.Size{Width: 70, Height: 12}
	axes := []collection.GridAxis{{ID: "api", Label: "API"}, {ID: "worker", Label: "Workers"}}

	grid, err := collection.NewDataGrid(axes, axes, []collection.GridEntry[int]{{RowID: "api", ColumnID: "api", Value: 24}}, size, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	inspector := collection.NewGridInspector(grid, func(e collection.GridEntry[int], present bool, _ int) string {
		return fmt.Sprintf("%s / %s\nObserved: %t\nReplicas: %d", e.RowID, e.ColumnID, present, e.Value)
	}, size, theme.New(theme.Dark))
	// Parent routes Tab using FocusDetail. Refresh after direct grid mutations.
	return inspector.View(), nil
}
