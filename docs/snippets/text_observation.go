package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextObservation() string {
	palette := theme.New(theme.Dark)

	return layout.Column(
		text.Log("Log entry", palette),
		text.Trace("trace-c7", palette),
		text.Span("carrier", palette),
		text.Upstream("carrier.example", palette),
		text.Downstream("API consumer", palette),
	)
}
