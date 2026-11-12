package snippets

import (
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func ExampleInputCheckbox() string {
	palette := theme.New(theme.Dark)
	checkbox := input.NewCheckbox("services", "All services", input.Mixed, palette)
	checkbox.Focus()
	// Retain checkbox.Update(msg)'s model and command; handle input.CheckChanged in the parent.
	return checkbox.View()
}
