package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextValueState() string {
	palette := theme.New(theme.Dark)
	// An empty string is present; redaction never renders the supplied secret.
	return text.ValuePresent.View("", text.Normal, palette) + "  " +
		text.ValueMissing.View("", text.Normal, palette) + "  " +
		text.ValueRedacted.View("secret", text.Normal, palette)
}
