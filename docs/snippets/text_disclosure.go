package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextDisclosure() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		text.Disclosure("Request headers", true, palette),
		text.Disclosure("Response body", false, palette),
	)
}
