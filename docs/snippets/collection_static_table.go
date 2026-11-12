package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionStaticTable() string {
	table := collection.StaticTable{
		Columns: []collection.TableColumn{{Title: "HOST", Fixed: true}, {Title: "VERSION", Min: 10}, {Title: "STATE"}},
		Rows: [][]collection.TableCell{
			{collection.TableText("web1west"), collection.TableText("9f2c4e1a7b3d"), collection.TableEvidence(feedback.EvidenceHealthy, "")},
			{collection.TableText("web2west"), collection.TableMissing(), collection.TableEvidence(feedback.EvidenceUnknown, "")}, // Missing is not empty.
		},
	}

	return table.View(48, theme.New(theme.Dark))
}
