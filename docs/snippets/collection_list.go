package snippets

import (
	"charm.land/bubbles/v2/list"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

// Define the item type at package scope in your application.
type environmentItem struct{ name, description string }

func (i environmentItem) Title() string       { return i.name }
func (i environmentItem) Description() string { return i.description }
func (i environmentItem) FilterValue() string { return i.name }

func ExampleCollectionList() string {
	items := []list.Item{
		environmentItem{"Production", "Primary service environment"},
		environmentItem{"Staging", "Integration environment"},
		environmentItem{"Preview", "Disposable review environment"},
	}
	environments := collection.NewList(items, 48, 14, theme.New(theme.Dark))
	environments.Title = "Environments"

	return environments.View()
}
