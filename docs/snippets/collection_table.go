package snippets

import (
	"charm.land/bubbles/v2/table"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionTable() string {
	columns := []table.Column{
		{Title: "Service", Width: 20}, {Title: "State", Width: 12}, {Title: "Req/min", Width: 8},
	}
	rows := []table.Row{
		{"API gateway", "Healthy", "1,240"},
		{"Event stream", "Healthy", "860"},
		{"Job workers", "Degraded", "420"},
	}
	tableView := collection.NewTable(columns, rows, 48, 6, theme.New(theme.Dark))

	return tableView.View()
}
