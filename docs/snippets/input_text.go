package snippets

import (
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleInputText() string {
	palette := theme.New(theme.Dark)
	field := input.NewText(input.TextOptions{
		Placeholder: "Type a value", Value: "production",
		Width: 36, Theme: palette,
	})
	// In Init, return field.Focus(); in Update, retain field.Update(msg)'s model and command.
	return layout.Panel("Value", field.View(),
		layout.Size{Width: 40, Height: 3}, palette, false)
}
