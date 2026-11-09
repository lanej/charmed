package snippets

import (
	"errors"

	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/theme"
)

func ExampleInputField() string {
	palette := theme.New(theme.Dark)
	field := input.Field{
		Label: "Environment", Help: "Use at least three characters.", Theme: palette,
		Input: input.NewText(input.TextOptions{
			Value: "production", Width: 40, Theme: palette,
			Validate: func(value string) error {
				if len(value) < 3 {
					return errors.New("use at least three characters")
				}

				return nil
			},
		}),
	}
	// In Init, return field.Input.Focus(); in Update, forward messages to field.Input.
	return field.View()
}
