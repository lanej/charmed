package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextHttpStatus() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		text.HTTPStatusCode(200, palette)+" OK",
		text.HTTPStatusCode(404, palette)+" Not Found",
		text.HTTPStatusCode(502, palette)+" Bad Gateway",
	)
}
