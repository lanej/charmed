package snippets

import (
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func ExampleCollectionCommandPalette() (string, error) {
	commands, err := collection.NewCommandPalette([]collection.Command{
		{ID: "refresh", Label: "Refresh observations", Shortcut: "r"},
		{ID: "trace", Label: "Scope: trace", DisabledReason: "No trace selected"},
		{ID: "layout", Label: "Layout: columns", State: "columns"},
	}, 64, 12, theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	commands.Open() // Return this command when opening the palette in your application.
	// Forward commands.Update(msg); handle collection.CommandSelectedMsg.ID to run an action.
	return commands.View(), nil
}
