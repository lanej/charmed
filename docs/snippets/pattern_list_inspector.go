package snippets

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExamplePatternListInspector() (string, error) {
	palette := theme.New(theme.Dark)
	row := func(title, condition, next string) func(int, bool) string {
		return func(width int, selected bool) string {
			style := palette.Text()
			if selected {
				style = palette.Heading()
			}

			return ansi.Hardwrap(style.Render(title)+"\n"+condition+"\nNext: "+next, width, true)
		}
	}
	records := []pattern.InspectorRecord{
		{ID: "route-0042", Heading: "Network", Row: row("Gateway route update", "Outcome unknown", "inspect remote evidence"), Detail: func(width int) string {
			return layout.KeyValue([]layout.Pair{
				{Label: "Operation", Value: "route-0042"},
				{Label: "Condition", Value: "Outcome unknown"},
				{Label: "Evidence", Value: "Submission recorded; response lost. Remote route has not been observed."},
				{Label: "Next", Value: "Read the saved request's remote state before changing traffic."},
			}, width, palette)
		}},
		{ID: "pool-0007", Heading: "Compute", Row: row("Worker capacity", "Outcome confirmed", "check queue latency"), Detail: func(int) string {
			return "An independent observer found twelve workers.\nCapacity does not establish healthy workers or acceptable latency."
		}},
		{ID: "index-0031", Row: row("Indexer restart", "Success reported", "read process observer"), Detail: func(int) string {
			return "The operator reported success.\nNo independent process sample confirms recovery."
		}},
	}

	inspector, err := pattern.NewListInspector(records, layout.Size{Width: 104, Height: 19}, palette)
	if err != nil {
		return "", err
	}

	inspector.ListTitle, inspector.DetailTitle = "Operations", "Evidence"
	inspector.Refresh()

	help := text.KeyHint("j/k", "select", palette) + "  " + text.KeyHint("enter", "inspect", palette) + "  " + text.KeyHint("esc", "back", palette) + "  " + text.KeyHint("[/]", "scroll", palette)

	return layout.Column(inspector.View(), help), nil
}
