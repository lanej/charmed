package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextHttpMethod() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		text.HTTPMethod("GET", palette), text.HTTPMethod("POST", palette),
		text.HTTPMethod("PATCH", palette), text.HTTPMethod("DELETE", palette),
	)
}
