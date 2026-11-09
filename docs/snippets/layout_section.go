package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleLayoutSection() string {
	palette := theme.New(theme.Dark)
	section := layout.Section{
		Title: "Request timeline", Content: "POST /v2/shipments/shp_fixture/buy   500   12.5ms",
		Collapsed: false, // Set true to return the section's space to the response.
	}
	height := section.Height(3)
	timeline := section.View(layout.Size{Width: 64, Height: height}, palette, true)
	response := layout.Panel("Response", "Internal server error\nCould not purchase shipment",
		layout.Size{Width: 64, Height: 8 - height}, palette, false)

	return layout.Column(timeline, response)
}
