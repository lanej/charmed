package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutLinkGroup() (string, error) {
	group := layout.LinkGroup{Links: []text.Link{
		{Label: "Logs", URL: "https://example.test/builds/42/logs"},
		{Label: "Artifact", URL: "https://example.test/builds/42/artifact"},
		{Label: "Worker session", URL: "terminal://session/build/pane/tests", Emphasis: text.Dim},
	}}

	return group.View(24, theme.New(theme.Dark))
}
