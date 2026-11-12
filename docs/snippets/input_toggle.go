package snippets

import (
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func ExampleInputToggle() string {
	toggle := input.NewToggle("live", "Live updates", true, theme.New(theme.Dark))
	toggle.Focus()
	// input.Toggled carries the setting ID and new On value.
	return toggle.View()
}
