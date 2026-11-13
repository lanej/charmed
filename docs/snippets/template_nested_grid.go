package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleTemplateNestedGrid() (string, error) {
	screen, err := tui.NewScreen("hosts", "cpu", "memory", "details")
	if err != nil {
		return "", err
	}

	metrics := &layout.Grid{Columns: []layout.Track{layout.Grow(14), layout.Grow(14)}, Rows: []layout.Track{layout.Grow(5)}, Gap: layout.Size{Width: 1}, Cells: []layout.Cell{
		{Node: &layout.Region{ID: "cpu"}}, {Node: &layout.Region{ID: "memory"}, Column: 1},
	}}
	inspection := &layout.Flex{
		Name: "inspection", Axis: layout.Vertical, Gap: 1,
		Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Metrics precede their supporting evidence in this column."}},
		Children:   []layout.Slot{{Node: metrics, Track: layout.Fixed(5)}, {Node: &layout.Region{ID: "details"}, Track: layout.Grow(5)}},
	}
	columns := &layout.Flex{Gap: 1, Children: []layout.Slot{{Node: &layout.Region{ID: "hosts"}, Track: layout.Fixed(24)}, {Node: inspection, Track: layout.Grow(30)}}}

	root := &layout.Responsive{Name: "workspace", Variants: []layout.Variant{
		{Name: "columns", Min: layout.Size{Width: 80}, Root: columns},
		{Name: "inspection", Root: inspection, Hidden: []layout.Omission{{ID: "hosts", Reason: "Keep the selected host's metrics visible in narrow terminals."}}},
	}}
	if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
		return "", arrangeErr
	}

	content := map[string]string{"hosts": "host-01\nhost-02\nhost-03", "cpu": "43%\n4 cores", "memory": "61%\n16 GiB", "details": "host-02 · v2.8.1\nProbe: HTTP 200\nLatency: 12 ms"}
	views := map[string]string{}

	for _, region := range screen.Plan().Regions() {
		size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
		views[region.ID] = layout.Panel(region.ID, content[region.ID], size, theme.New(theme.Dark), false)
	}
	// Every leaf is registered once; the compact recipe explicitly hides hosts.
	// Call Arrange again on resize; retained native models stay with the parent.
	frame, err := screen.Render(views)

	return frame.Content, err
}
