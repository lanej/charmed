package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func ExampleFoundationsLayout() string {
	palette := theme.New(theme.Dark)
	left, right, sideBySide := layout.Split(80, 1, 24, 32, .4)
	summary := layout.Panel("Summary", "12 services / all healthy",
		layout.Size{Width: left, Height: 6}, palette, true)

	details := layout.Panel("Details", "Region: us-west-2",
		layout.Size{Width: right, Height: 6}, palette, false)
	if sideBySide {
		return layout.Row(summary, " ", details)
	}

	return layout.Column(summary, details)
}
