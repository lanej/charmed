package snippets

import (
	"github.com/charmbracelet/x/ansi"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/tui"
)

func ExampleTuiStatic() (string, error) {
	declare := func(_ layout.Size, heights tui.StaticHeights) layout.Node {
		summary := &layout.Region{ID: "summary", Min: layout.Size{Width: 20, Height: heights.Height("summary")}}
		detail := &layout.Region{ID: "detail", Min: layout.Size{Width: 20, Height: heights.Height("detail")}}

		return &layout.Responsive{Name: "report", Variants: []layout.Variant{
			{Name: "side-by-side", Min: layout.Size{Width: 60, Height: 1}, Root: &layout.SplitPane{First: summary, Second: detail, Gap: 2}},
			{Name: "stacked", Root: &layout.Flex{Name: "stack", Axis: layout.Vertical, Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Printed reports read top to bottom."}}, Children: []layout.Slot{
				{Node: summary, Track: layout.Fixed(heights.Height("summary"))},
				{Node: detail, Track: layout.Fixed(heights.Height("detail"))},
			}}},
		}}
	}

	text := map[string]string{
		"summary": "Totals for the period, grouped by outcome.",
		"detail":  "Each finding names its evidence and its owner.",
	}

	frame, err := tui.RenderStatic([]string{"summary", "detail"}, 72, declare, func(id string, width int) string {
		return ansi.Wrap(text[id], width, "")
	})

	return frame.Content, err
}
