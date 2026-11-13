package snippets

import (
	"charm.land/bubbles/v2/table"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionKeyedTable() (string, error) {
	columns := []table.Column{{Title: "Service", Width: 24}, {Title: "Replicas", Width: 12}}
	native := collection.NewTable(columns, nil, 40, 5, theme.New(theme.Dark))
	rows := []collection.KeyedRow[string]{{Key: "api", Row: table.Row{"API gateway", "24"}}, {Key: "worker", Row: table.Row{"Workers", "12"}}}

	keyed, err := collection.NewKeyedTable(&native, rows)
	if err != nil {
		return "", err
	}

	keyed.Select("worker")
	// Reordering and filtering retain the selected key when it survives.
	if err := keyed.SetRows([]collection.KeyedRow[string]{rows[1]}); err != nil {
		return "", err
	}
	// Forward Update commands; configure styles and focus on keyed.Model.
	return keyed.View(), nil
}
