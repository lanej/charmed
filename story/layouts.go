package story

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/x/ansi"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/tui"
)

func layoutDefinitions() []Definition {
	catalog := []struct{ id, title, level, family string }{
		{"layout/flex", "Weighted flex layout", "Molecules", "layout"},
		{"layout/grid", "Grid with spanning cells", "Molecules", "layout"},
		{"layout/flow", "Wrapping flow layout", "Molecules", "layout"},
		{"layout/dock", "Docked layout", "Molecules", "layout"},
		{"layout/split-pane", "Constrained split panes", "Molecules", "layout"},
		{"layout/responsive", "Responsive layout recipes", "Molecules", "layout"},
		{"tui/screen", "Declared screen", "Templates", "tui"},
		{"tui/regions", "Region inspector", "Templates", "tui"},
		{"tui/validation", "Strict layout validation", "Templates", "tui"},
		{"tui/static", "Static declared render", "Templates", "tui"},
	}

	var definitions []Definition

	for _, entry := range catalog {
		id := entry.id
		config := DefaultConfig()
		config.Columns, config.Rows = 100, 20
		composes := []string{"foundations/layout"}

		switch id {
		case "layout/dock", "layout/split-pane":
			composes = []string{"layout/flex"}
		case "layout/responsive":
			composes = []string{"layout/grid", "layout/split-pane"}
		case "tui/screen":
			composes = []string{"layout/responsive", "layout/dock"}
		case "tui/regions":
			composes = []string{"tui/screen"}
		case "tui/validation":
			composes = []string{"layout/flex", "tui/screen"}
		case "tui/static":
			composes = []string{"layout/responsive", "tui/screen"}
		}

		definitions = append(definitions, Definition{Composes: composes, ID: id, Title: entry.level + "/" + entry.family + "/" + entry.title, Level: entry.level, Family: entry.family, Description: "Allocate named terminal regions with explicit sizing, responsive intent and strict design validation.", Source: "story/layouts.go", Symbol: "declaredLayoutStory", Defaults: config, Factory: func(e Environment) (runtime.Scene, error) { return declaredLayoutStory(id, e) }})
	}

	return definitions
}

func specimenRegion(id string) *layout.Region {
	return &layout.Region{ID: id, Min: layout.Size{Width: max(8, len(id)+4), Height: 2}}
}

func layoutSpecimen(kind string, size layout.Size) (layout.Node, []string) {
	a, b, c := specimenRegion("primary"), specimenRegion("secondary"), specimenRegion("tools")
	ids := []string{a.ID, b.ID, c.ID}

	switch kind {
	case "layout/flex":
		return &layout.Flex{Gap: 1, Children: []layout.Slot{{Node: a, Track: layout.Track{Min: 8, Weight: 2}}, {Node: b, Track: layout.Grow(8)}, {Node: c, Track: layout.Fixed(10)}}}, ids
	case "layout/grid":
		return &layout.Grid{Columns: []layout.Track{layout.Grow(8), layout.Grow(8)}, Rows: []layout.Track{layout.Grow(3), layout.Grow(3)}, Gap: layout.Size{Width: 1, Height: 1}, Cells: []layout.Cell{{Node: a}, {Node: b, Column: 1}, {Node: c, Row: 1, ColumnSpan: 2}}}, ids
	case "layout/flow":
		box := layout.Size{Width: max(14, size.Width/3), Height: 3}
		return &layout.Flow{Gap: layout.Size{Width: 1, Height: 1}, Children: []layout.Box{{Node: a, Size: box}, {Node: b, Size: box}, {Node: c, Size: box}}}, ids
	case "layout/dock":
		c.Role = layout.ChromeRegion
		b.Role = layout.SupportingRegion

		return &layout.Dock{Top: c, TopHeight: 3, Left: b, LeftWidth: 16, Body: a, Gap: 1}, ids
	case "layout/split-pane":
		return &layout.SplitPane{First: a, Second: b, Gap: 1, Ratio: .65}, ids[:2]
	default:
		a.Exceptions = []layout.Exception{{Rule: layout.SingleRegion, Reason: "Focused narrow layout; secondary panes remain explicitly registered and hidden."}}

		return &layout.Responsive{Name: "workspace", Variants: []layout.Variant{
			{Name: "three-column", Min: layout.Size{Width: 100, Height: 18}, Root: &layout.Grid{Columns: []layout.Track{layout.Grow(8), layout.Grow(8), layout.Grow(8)}, Rows: []layout.Track{layout.Grow(3)}, Gap: layout.Size{Width: 1}, Cells: []layout.Cell{{Node: a}, {Node: b, Column: 1}, {Node: c, Column: 2}}}},
			{Name: "sidebar", Min: layout.Size{Width: 60, Height: 12}, Root: &layout.SplitPane{First: a, Second: b, Gap: 1}, Hidden: []layout.Omission{{ID: c.ID, Reason: "Tools move into the parent application's commands."}}},
			{Name: "focused", Root: a, Hidden: []layout.Omission{{ID: b.ID, Reason: "Parent selection switches the focused pane."}, {ID: c.ID, Reason: "Tools remain available through commands."}}},
		}}, ids
	}
}

func declaredLayoutStory(id string, e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		if id == "tui/validation" {
			return validationSpecimen(size, e)
		}

		if id == "tui/static" {
			return staticSpecimen(size, e)
		}

		root, ids := layoutSpecimen(id, size)

		screen, err := tui.NewScreen(ids...)
		if err != nil {
			return err.Error()
		}

		if arrangeErr := screen.Arrange(root, size); arrangeErr != nil {
			return e.Config.Palette().Text().Foreground(e.Config.Palette().Danger).Render(arrangeErr.Error())
		}

		if id == "tui/regions" {
			return screen.RegionMap(e.Config.Palette())
		}

		views := map[string]string{}

		for _, region := range screen.Plan().Regions() {
			allocation := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
			views[region.ID] = layout.Panel(region.ID, fmt.Sprintf("%d×%d", allocation.Width, allocation.Height), allocation, e.Config.Palette(), region.ID == "primary")
		}

		frame, err := screen.Render(views)
		if err != nil {
			return err.Error()
		}

		return frame.Content
	}), nil
}

func validationSpecimen(size layout.Size, e Environment) string {
	screen, err := tui.NewScreen("first", "second")
	if err != nil {
		return err.Error()
	}

	column := &layout.Flex{Name: "steps", Axis: layout.Vertical, Children: []layout.Slot{{Node: specimenRegion("first"), Track: layout.Grow(2)}, {Node: specimenRegion("second"), Track: layout.Grow(2)}}}

	blocked := screen.Arrange(column, size)
	if blocked == nil {
		return "The declared minima require stacking at this width."
	}

	column.Exceptions = []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Chronological steps are read top to bottom."}}
	if err := screen.Arrange(column, size); err != nil {
		return err.Error()
	}

	t := e.Config.Palette()

	return ansi.Wrap(t.Text().Foreground(t.Danger).Render("BLOCKED: "+blocked.Error())+"\n\n"+t.Heading().Render("Accepted local exception")+"\n"+column.Exceptions[0].Reason, size.Width, "")
}

// staticSpecimen prints a finite report: the canvas height is measured, not supplied.
func staticSpecimen(size layout.Size, e Environment) string {
	t := e.Config.Palette()
	ids := []string{"summary", "detail"}
	text := map[string]string{
		"summary": "Totals for the period, grouped by outcome. The summary keeps its place beside the detail when the page is wide enough.",
		"detail":  "Each finding names its evidence and its owner. When the page narrows the detail is stacked below the summary instead of being cut off.",
	}

	declare := func(_ layout.Size, heights tui.StaticHeights) layout.Node {
		summary := &layout.Region{ID: "summary", Min: layout.Size{Width: 24, Height: heights.Height("summary")}}
		detail := &layout.Region{ID: "detail", Min: layout.Size{Width: 24, Height: heights.Height("detail")}}

		return &layout.Responsive{Name: "report", Variants: []layout.Variant{
			{Name: "side-by-side", Min: layout.Size{Width: 70, Height: 1}, Root: &layout.SplitPane{First: summary, Second: detail, Gap: 2}},
			{Name: "stacked", Root: &layout.Flex{Name: "stack", Axis: layout.Vertical, Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Printed reports read top to bottom."}}, Children: []layout.Slot{
				{Node: summary, Track: layout.Fixed(heights.Height("summary"))},
				{Node: detail, Track: layout.Fixed(heights.Height("detail"))},
			}}},
		}}
	}

	render := func(id string, width int) string {
		return t.Heading().Render(strings.ToUpper(id[:1])+id[1:]) + "\n" + ansi.Wrap(text[id], width, "")
	}

	frame, err := tui.RenderStatic(ids, size.Width, declare, render)
	if err != nil {
		return t.Text().Foreground(t.Danger).Render(err.Error())
	}

	return frame.Content
}
