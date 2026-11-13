package story

import "github.com/lanej/charmed/layout"

func gitBrowserLayout(focusDetail bool, mode int) layout.Node {
	main := &layout.Region{ID: "main", Min: layout.Size{Width: 20, Height: 9}}
	details := &layout.Region{ID: "details", Min: layout.Size{Width: 20, Height: 5}}
	context := &layout.Region{ID: "context", Min: layout.Size{Width: 24, Height: 5}, Role: layout.SupportingRegion}

	focused := layout.Variant{Name: "focused", Root: main, Hidden: []layout.Omission{{ID: "context", Reason: "Context rail is available in the wide layout."}, {ID: "details", Reason: "Tab switches to the focused detail pane."}}}
	if focusDetail {
		focused.Root = details
		focused.Hidden[1] = layout.Omission{ID: "main", Reason: "Tab returns to the focused primary pane."}
	}

	variants := []layout.Variant{focused}
	if mode != 2 {
		variants = append([]layout.Variant{{Name: "sidebar", Min: layout.Size{Width: 100, Height: 13}, Root: &layout.SplitPane{First: main, Second: details, Ratio: .6, Gap: 1}, Hidden: []layout.Omission{{ID: "context", Reason: "Context is shown in the header at this allocation."}}}}, variants...)
	}

	if mode == 0 {
		variants = append([]layout.Variant{{Name: "three-column", Min: layout.Size{Width: 120, Height: 19}, Root: &layout.Grid{
			Columns: []layout.Track{layout.Fixed(28), {Min: 54, Weight: 3}, {Min: 30, Weight: 2}}, Rows: []layout.Track{layout.Grow(5)}, Gap: layout.Size{Width: 1}, Cells: []layout.Cell{{Node: context}, {Node: main, Column: 1}, {Node: details, Column: 2}},
		}}}, variants...)
	}

	return &layout.Dock{
		Top: &layout.Flex{Axis: layout.Vertical, Children: []layout.Slot{
			{Node: &layout.Region{ID: "header", Role: layout.ChromeRegion}, Track: layout.Fixed(1)},
			{Node: &layout.Region{ID: "tabs", Role: layout.ChromeRegion}, Track: layout.Fixed(1)},
		}}, TopHeight: 2,
		Body:   &layout.Responsive{Name: "git-workspace", Variants: variants},
		Bottom: &layout.Region{ID: "footer", Role: layout.ChromeRegion}, BottomHeight: 1,
	}
}
