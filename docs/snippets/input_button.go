package snippets

import (
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func ExampleInputButton() string {
	button := input.NewButton("retry", "Retry", theme.New(theme.Dark))
	button.Focus()
	// Handle input.Activated in the application; set Busy while its operation is outstanding.
	button.Busy = true

	return button.View()
}
