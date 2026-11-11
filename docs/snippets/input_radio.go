package snippets

import (
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func ExampleInputRadio() string {
	palette := theme.New(theme.Dark)
	production := input.NewRadio("production", "Production", true, palette)
	staging := input.NewRadio("staging", "Staging", false, palette)
	// On input.RadioSelected, the parent deselects every other radio in this group.
	return production.View() + "  " + staging.View()
}
