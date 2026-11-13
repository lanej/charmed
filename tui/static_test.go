package tui_test

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/tui"
)

var staticIDs = []string{"summary", "detail"}

// staticDeclare is a Responsive recipe: side by side from 80 columns, stacked
// from 40 and a single summary below that.
func staticDeclare(canvas layout.Size, heights tui.StaticHeights) layout.Node {
	summary := func(exceptions ...layout.Exception) *layout.Region {
		return &layout.Region{ID: "summary", Min: layout.Size{Width: 10, Height: heights.Height("summary")}, Exceptions: exceptions}
	}
	detail := &layout.Region{ID: "detail", Min: layout.Size{Width: 20, Height: heights.Height("detail")}}

	return &layout.Responsive{Name: "report", Variants: []layout.Variant{
		{Name: "side-by-side", Min: layout.Size{Width: 80, Height: 1}, Root: &layout.SplitPane{First: summary(), Second: detail, Gap: 2}},
		{Name: "stacked", Min: layout.Size{Width: 30, Height: 1}, Root: &layout.Flex{Name: "stack", Axis: layout.Vertical, Exceptions: []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Narrow printed pages read top to bottom."}}, Children: []layout.Slot{
			{Node: summary(), Track: layout.Fixed(heights.Height("summary"))},
			{Node: detail, Track: layout.Fixed(heights.Height("detail"))},
		}}},
		{Name: "summary-only", Root: summary(layout.Exception{Rule: layout.SingleRegion, Reason: "Very narrow output keeps the summary."}), Hidden: []layout.Omission{{ID: "detail", Reason: "Too narrow to print beside the summary."}}},
	}}
}

func staticRender(id string, width int) string {
	text := "plain words that wrap at the allocated width to produce several rows of text"
	if id == "detail" {
		text = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ff8800")).Render("styled detail text that is longer than the summary and wraps as well, adding rows")
	}

	return ansi.Wrap(text, width, "")
}

func TestRenderStaticAcrossWidths(t *testing.T) {
	variants := map[string]bool{}

	for width := 20; width <= 200; width++ {
		frame, err := tui.RenderStatic(staticIDs, width, staticDeclare, staticRender)
		if err != nil {
			t.Fatalf("width %d: %v", width, err)
		}

		variants[frame.Plan.Variant()] = true
		lines := strings.Split(frame.Content, "\n")

		for i, line := range lines {
			if got := lipgloss.Width(line); got > width {
				t.Fatalf("width %d line %d is %d cells", width, i, got)
			}

			if got := ansi.Strip(line); got != strings.TrimRight(got, " ") {
				t.Fatalf("width %d line %d has trailing padding %q", width, i, got)
			}
		}

		if ansi.Strip(lines[len(lines)-1]) == "" {
			t.Fatalf("width %d ends with a blank row", width)
		}

		if frame.Plan.Size().Width != width {
			t.Fatalf("plan width %d", frame.Plan.Size().Width)
		}
	}

	for _, name := range []string{"side-by-side", "stacked", "summary-only"} {
		if !variants[name] {
			t.Errorf("recipe %s was never selected", name)
		}
	}
}

func TestRenderStaticHiddenRegionsKeepReasons(t *testing.T) {
	frame, err := tui.RenderStatic(staticIDs, 25, staticDeclare, func(id string, width int) string {
		if id == "detail" {
			t.Fatal("hidden region was rendered")
		}

		return staticRender(id, width)
	})
	if err != nil {
		t.Fatal(err)
	}

	hidden := frame.Plan.Hidden()
	if len(hidden) != 1 || hidden[0].ID != "detail" || hidden[0].Reason == "" {
		t.Fatalf("hidden regions: %#v", hidden)
	}

	if strings.Contains(ansi.Strip(frame.Content), "styled") {
		t.Fatal("hidden content rendered")
	}
}

func TestRenderStaticBoundaryWidthsCheckLayouts(t *testing.T) {
	screen := screen(t, staticIDs...)
	sizes := []layout.Size{{Width: 29, Height: 20}, {Width: 30, Height: 20}, {Width: 79, Height: 20}, {Width: 80, Height: 20}}

	if _, err := screen.CheckLayouts(staticDeclare(sizes[0], tui.StaticHeights{}), sizes...); err != nil {
		t.Fatal(err)
	}

	for _, want := range []struct {
		width   int
		variant string
	}{{29, "summary-only"}, {30, "stacked"}, {79, "stacked"}, {80, "side-by-side"}} {
		frame, err := tui.RenderStatic(staticIDs, want.width, staticDeclare, staticRender)
		if err != nil || frame.Plan.Variant() != want.variant {
			t.Fatalf("width %d: variant %q err %v", want.width, frame.Plan.Variant(), err)
		}
	}
}

func TestRenderStaticUsesSmallestCanvas(t *testing.T) {
	frame, err := tui.RenderStatic(staticIDs, 100, staticDeclare, staticRender)
	if err != nil {
		t.Fatal(err)
	}

	rows := len(strings.Split(frame.Content, "\n"))
	if frame.Plan.Size().Height != rows {
		t.Fatalf("canvas %d rows for %d rendered rows", frame.Plan.Size().Height, rows)
	}
}

func TestRenderStaticIsDeterministic(t *testing.T) {
	first, err := tui.RenderStatic(staticIDs, 90, staticDeclare, staticRender)
	if err != nil {
		t.Fatal(err)
	}

	second, err := tui.RenderStatic(staticIDs, 90, staticDeclare, staticRender)
	if err != nil || first.Content != second.Content {
		t.Fatalf("renders differ: %v", err)
	}
}

func TestRenderStaticErrors(t *testing.T) {
	wide := func(id string, width int) string { return strings.Repeat("x", width+1) }
	missing := func(layout.Size, tui.StaticHeights) layout.Node {
		return &layout.Region{ID: "summary", Min: layout.Size{Width: 5, Height: 1}, Exceptions: []layout.Exception{{Rule: layout.SingleRegion, Reason: "Only one."}}}
	}
	tooWide := func(layout.Size, tui.StaticHeights) layout.Node {
		return &layout.SplitPane{First: &layout.Region{ID: "summary", Min: layout.Size{Width: 30, Height: 1}}, Second: &layout.Region{ID: "detail", Min: layout.Size{Width: 30, Height: 1}}}
	}

	for name, call := range map[string]func() error{
		"zero width":          func() error { _, err := tui.RenderStatic(staticIDs, 0, staticDeclare, staticRender); return err },
		"nil render":          func() error { _, err := tui.RenderStatic(staticIDs, 80, staticDeclare, nil); return err },
		"no regions":          func() error { _, err := tui.RenderStatic(nil, 80, staticDeclare, staticRender); return err },
		"content too wide":    func() error { _, err := tui.RenderStatic(staticIDs, 80, staticDeclare, wide); return err },
		"unregistered":        func() error { _, err := tui.RenderStatic(staticIDs, 80, missing, staticRender); return err },
		"minima exceed width": func() error { _, err := tui.RenderStatic(staticIDs, 40, tooWide, staticRender); return err },
	} {
		if err := call(); err == nil {
			t.Errorf("%s: expected an error", name)
		} else {
			t.Logf("%s: %v", name, err)
		}
	}
}
