package snippets

import (
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/navigation"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTemplateApplication() string {
	palette := theme.New(theme.Dark)
	tabs := navigation.NewTabs([]navigation.Tab{
		{ID: "overview", Label: "Overview"}, {ID: "services", Label: "Services"},
	}, palette)
	tabs.SetWidth(64)

	header := layout.Header{Title: "Service workbench", Context: "Production / us-west-2"}
	footer := layout.Footer{Help: "arrows tabs", Status: feedback.Connected.View(text.Normal, palette)}
	body := layout.Panel("Overview", "12 services / all healthy",
		layout.Size{Width: 64, Height: 8}, palette, false)

	return layout.Column(header.View(64, palette), tabs.View(), body, footer.View(64, palette))
}
