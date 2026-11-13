package snippets

import (
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func ExampleInputRadioGroup() (string, error) {
	group, err := input.NewRadioGroup([]input.RadioOption{
		{ID: "last", Label: "Last transition"},
		{ID: "three", Label: "Last three transitions"},
		{ID: "all", Label: "Complete trace"},
		{ID: "live", Label: "Live execution", DisabledReason: "Replay is inspection only"},
	}, "last", theme.New(theme.Dark))
	if err != nil {
		return "", err
	}

	group.Focus()
	// Arrows move focus only; Enter/Space commits and emits RadioSelected.
	// Select changes selection programmatically without emitting a command.
	return group.View(), nil
}
