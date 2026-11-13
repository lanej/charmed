package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleTemplateNestedColumns() (string, error) {
	screen, err := tui.NewScreen("hosts", "summary", "details")
	if err != nil {
		return "", err
	}

	summary, details := &layout.Region{ID: "summary"}, &layout.Region{ID: "details"}
	inspection := &layout.Flex{
		Name: "inspection", Axis: layout.Vertical, Gap: 1,
		Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "The summary introduces the evidence below it."}},
		Children:   []layout.Slot{{Node: summary, Track: layout.Fixed(5)}, {Node: details, Track: layout.Grow(5)}},
	}

	root := &layout.Flex{Gap: 1, Children: []layout.Slot{
		{Node: &layout.Region{ID: "hosts"}, Track: layout.Fixed(24)},
		{Node: inspection, Track: layout.Grow(30)},
	}}
	if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
		return "", arrangeErr
	}

	palette := theme.New(theme.Dark)
	content := map[string]string{"hosts": "host-01\nhost-02\nhost-03", "summary": "host-02 · v2.8.1\nPublic pool", "details": "Probe: HTTP 200\nLatency: 12 ms\nRetries: 0"}
	views := map[string]string{}

	for _, region := range screen.Plan().Regions() {
		size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
		views[region.ID] = layout.Panel(region.ID, content[region.ID], size, palette, false)
	}
	// Resize native children from each leaf's PanelContentSize before rendering.
	// The parent owns focus and forwards commands to its active child.
	frame, err := screen.Render(views)

	return frame.Content, err
}
