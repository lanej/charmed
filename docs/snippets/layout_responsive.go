package snippets

import (
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func ExampleLayoutResponsive() (string, error) {
	screen, err := tui.NewScreen("main", "details")
	if err != nil {
		return "", err
	}

	root := &layout.Responsive{Name: "workspace", Variants: []layout.Variant{
		{Name: "wide", Min: layout.Size{Width: 70, Height: 12}, Root: &layout.SplitPane{First: &layout.Region{ID: "main"}, Second: &layout.Region{ID: "details"}, Gap: 1}},
		{Name: "focused", Root: &layout.Region{ID: "main", Exceptions: []layout.Exception{{Rule: layout.SingleRegion, Reason: "Single focused pane; parent selection reveals details."}}}, Hidden: []layout.Omission{{ID: "details", Reason: "Parent focus selects this pane at narrow widths."}}},
	}}
	if arrangeErr := screen.Arrange(root, layout.Size{Width: 80, Height: 16}); arrangeErr != nil {
		return "", arrangeErr
	}

	views := map[string]string{}

	for _, region := range screen.Plan().Regions() {
		size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
		views[region.ID] = layout.Panel(region.ID, "Native view", size, theme.New(theme.Dark), false)
	}

	frame, err := screen.Render(views)

	return frame.Content, err
}
