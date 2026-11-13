package layout_test

import (
	"errors"
	"image"
	"testing"

	"github.com/lanej/charmed/layout"
)

func area(id string, w, h int) *layout.Region {
	return &layout.Region{ID: id, Min: layout.Size{Width: w, Height: h}}
}

func assertRect(t *testing.T, plan layout.Plan, id string, want image.Rectangle) {
	t.Helper()

	got, ok := plan.Bounds(id)
	if !ok || got != want {
		t.Fatalf("%s: got %v (%v), want %v", id, got, ok, want)
	}
}

func requirePlan(t *testing.T, node layout.Node, size layout.Size) layout.Plan {
	t.Helper()

	plan, err := layout.Resolve(node, size)
	if err != nil {
		t.Fatal(err)
	}

	return plan
}

func TestFlexAllocationAndAlignment(t *testing.T) {
	root := &layout.Flex{Gap: 2, Children: []layout.Slot{
		{Node: area("rail", 20, 1), Track: layout.Fixed(20)},
		{Node: area("main", 10, 1), Track: layout.Track{Min: 10, Weight: 1}},
		{Node: area("details", 10, 1), Track: layout.Track{Min: 10, Weight: 2}},
	}}
	plan := requirePlan(t, root, layout.Size{Width: 100, Height: 20})
	assertRect(t, plan, "rail", image.Rect(0, 0, 20, 20))
	assertRect(t, plan, "main", image.Rect(22, 0, 51, 20))
	assertRect(t, plan, "details", image.Rect(53, 0, 100, 20))

	root.Children[1].Track.Max = 15
	plan = requirePlan(t, root, layout.Size{Width: 100, Height: 20})
	assertRect(t, plan, "main", image.Rect(22, 0, 37, 20))
	assertRect(t, plan, "details", image.Rect(39, 0, 100, 20))

	root = &layout.Flex{Gap: 1, Align: layout.Center, Justify: layout.End, Children: []layout.Slot{{Node: area("a", 1, 1), Track: layout.Fixed(4), Cross: 2}, {Node: area("b", 1, 1), Track: layout.Fixed(5)}}}
	plan = requirePlan(t, root, layout.Size{Width: 20, Height: 10})
	assertRect(t, plan, "a", image.Rect(10, 4, 14, 6))
	assertRect(t, plan, "b", image.Rect(15, 0, 20, 10))
}

func TestGridSpansAndFlowWrapping(t *testing.T) {
	grid := &layout.Grid{Columns: []layout.Track{layout.Grow(20), layout.Grow(20)}, Rows: []layout.Track{layout.Fixed(2), layout.Grow(5)}, Gap: layout.Size{Width: 1, Height: 1}, Cells: []layout.Cell{
		{Node: area("a", 1, 1)},
		{Node: area("b", 1, 1), Column: 1},
		{Node: area("bottom", 1, 1), Row: 1, ColumnSpan: 2},
	}}
	plan := requirePlan(t, grid, layout.Size{Width: 81, Height: 21})
	assertRect(t, plan, "a", image.Rect(0, 0, 40, 2))
	assertRect(t, plan, "b", image.Rect(41, 0, 81, 2))
	assertRect(t, plan, "bottom", image.Rect(0, 3, 81, 21))

	flow := &layout.Flow{Gap: layout.Size{Width: 1, Height: 1}, Children: []layout.Box{{Node: area("a", 1, 1), Size: layout.Size{Width: 12, Height: 2}}, {Node: area("b", 1, 1), Size: layout.Size{Width: 12, Height: 3}}, {Node: area("c", 1, 1), Size: layout.Size{Width: 12, Height: 2}}}}
	plan = requirePlan(t, flow, layout.Size{Width: 30, Height: 8})
	assertRect(t, plan, "a", image.Rect(0, 0, 12, 2))
	assertRect(t, plan, "b", image.Rect(13, 0, 25, 3))
	assertRect(t, plan, "c", image.Rect(0, 4, 12, 6))

	if _, err := layout.Resolve(flow, layout.Size{Width: 30, Height: 5}); err == nil {
		t.Fatal("wrapped overflow accepted")
	}
}

func TestDockAndSplitBounds(t *testing.T) {
	top, bottom := area("header", 1, 1), area("footer", 1, 1)
	top.Role, bottom.Role = layout.ChromeRegion, layout.ChromeRegion
	dock := &layout.Dock{Top: top, TopHeight: 2, Bottom: bottom, BottomHeight: 1, Left: area("rail", 20, 1), LeftWidth: 20, Body: area("main", 10, 2), Gap: 1}
	plan := requirePlan(t, dock, layout.Size{Width: 100, Height: 30})
	assertRect(t, plan, "header", image.Rect(0, 0, 100, 2))
	assertRect(t, plan, "main", image.Rect(21, 3, 100, 28))
	assertRect(t, plan, "rail", image.Rect(0, 3, 20, 28))
	assertRect(t, plan, "footer", image.Rect(0, 29, 100, 30))
	plan = requirePlan(t, &layout.SplitPane{First: area("first", 20, 1), Second: area("second", 20, 1), Ratio: .6, Gap: 1}, layout.Size{Width: 101, Height: 10})
	assertRect(t, plan, "first", image.Rect(0, 0, 60, 10))
	assertRect(t, plan, "second", image.Rect(61, 0, 101, 10))
}

func TestRulesBlockUnlessLocallyExempted(t *testing.T) {
	column := &layout.Flex{Name: "comparison", Axis: layout.Vertical, Children: []layout.Slot{{Node: area("a", 30, 2), Track: layout.Grow(2)}, {Node: area("b", 30, 2), Track: layout.Grow(2)}}}
	if _, err := layout.Resolve(column, layout.Size{Width: 100, Height: 20}); err == nil {
		t.Fatal("unexplained wide stacking accepted")
	} else {
		var rule *layout.RuleError
		if !errors.As(err, &rule) || rule.Diagnostics[0].Rule != layout.StackOpportunity {
			t.Fatalf("wrong error: %v", err)
		}
	}

	column.Exceptions = []layout.Exception{{Rule: layout.StackOpportunity, Reason: "Chronological steps are read top to bottom."}}

	plan := requirePlan(t, column, layout.Size{Width: 100, Height: 20})
	if len(plan.Diagnostics()) != 1 || plan.Diagnostics()[0].Exception == "" {
		t.Fatal("exception was not retained")
	}

	column.Exceptions[0].Reason = " "
	if _, err := layout.Resolve(column, layout.Size{Width: 100, Height: 20}); err == nil {
		t.Fatal("blank exception accepted")
	}

	column.Exceptions = []layout.Exception{{Rule: layout.ContentClipped, Reason: "Wrong scope"}}
	if _, err := layout.Resolve(column, layout.Size{Width: 100, Height: 20}); err == nil {
		t.Fatal("exception for another scope accepted")
	}

	single := area("document", 10, 2)
	if _, err := layout.Resolve(single, layout.Size{Width: 40, Height: 10}); err == nil {
		t.Fatal("unexplained single surface accepted")
	}

	single.Exceptions = []layout.Exception{{Rule: layout.SingleRegion, Reason: "This is a single document editor."}}
	plan = requirePlan(t, single, layout.Size{Width: 40, Height: 10})
	plan.Regions()[0].ID = "mutated"

	plan.Diagnostics()[0].Regions[0] = "mutated"
	if plan.Regions()[0].ID != "document" || plan.Diagnostics()[0].Regions[0] != "document" {
		t.Fatal("plan leaked mutable slices")
	}

	single.Exceptions = append(single.Exceptions, layout.Exception{Rule: layout.BelowPreferred, Reason: "Narrow document window."})
	single.Preferred = layout.Size{Width: 80}
	requirePlan(t, single, layout.Size{Width: 40, Height: 10})

	if _, err := layout.Resolve(single, layout.Size{Width: 5, Height: 10}); err == nil {
		t.Fatal("exception bypassed geometry minimum")
	}
}

func TestMalformedLayoutsNeverReturnPartialPlans(t *testing.T) {
	cycle := &layout.Flex{}
	cycle.Children = []layout.Slot{{Node: cycle, Track: layout.Grow(1)}}

	var nilRegion *layout.Region
	for _, root := range []layout.Node{
		nil, nilRegion, cycle,
		&layout.Flex{Children: []layout.Slot{{Node: area("same", 1, 1), Track: layout.Grow(1)}, {Node: area("same", 1, 1), Track: layout.Grow(1)}}},
		&layout.Flex{Children: []layout.Slot{{Node: area("a", 10, 1), Track: layout.Fixed(1)}, {Node: area("b", 1, 1), Track: layout.Grow(1)}}},
		&layout.Grid{Columns: []layout.Track{layout.Grow(1)}, Rows: []layout.Track{layout.Grow(1)}, Cells: []layout.Cell{{Node: area("a", 1, 1)}, {Node: area("b", 1, 1)}}},
		&layout.Dock{Body: area("a", 1, 1), TopHeight: 2},
		&layout.Responsive{Name: "test", Variants: []layout.Variant{{Name: "empty", Root: area("a", 1, 1), Hidden: []layout.Omission{{ID: "a", Reason: "duplicate visible ID"}}}}},
	} {
		plan, err := layout.Resolve(root, layout.Size{Width: 80, Height: 20})
		if err == nil || len(plan.Regions()) != 0 {
			t.Fatalf("invalid layout accepted: %T", root)
		}
	}
}

func TestStrictStackingCannotBeHiddenByDefaultSizesOrWrappers(t *testing.T) {
	for _, wrapped := range []bool{false, true} {
		var (
			a layout.Node = &layout.Region{ID: "a"}
			b layout.Node = &layout.Region{ID: "b"}
		)

		if wrapped {
			a = &layout.Flex{Children: []layout.Slot{{Node: a, Track: layout.Grow(0)}}}
			b = &layout.Responsive{Name: "wrapped", Variants: []layout.Variant{{Name: "only", Root: b}}}
		}

		root := &layout.Flex{Axis: layout.Vertical, Children: []layout.Slot{{Node: a, Track: layout.Grow(0)}, {Node: b, Track: layout.Grow(0)}}}
		if _, err := layout.Resolve(root, layout.Size{Width: 100, Height: 20}); err == nil {
			t.Fatalf("stacking bypass: wrapped=%v", wrapped)
		}
	}
}

func TestGridDerivesSharedSpanAndSingleCellMinima(t *testing.T) {
	grid := &layout.Grid{Columns: []layout.Track{layout.Grow(0), layout.Grow(0)}, Rows: []layout.Track{layout.Grow(0)}, Cells: []layout.Cell{{Node: area("wide", 20, 1)}, {Node: area("small", 1, 1), Column: 1}}}
	plan := requirePlan(t, grid, layout.Size{Width: 30, Height: 10})
	wide, _ := plan.Bounds("wide")

	small, _ := plan.Bounds("small")
	if wide.Dx() < 20 || small.Dx() < 1 || wide.Max.X != small.Min.X || small.Max.X != 30 {
		t.Fatal("grid failed to reserve child minima")
	}

	grid = &layout.Grid{Columns: []layout.Track{layout.Grow(0), layout.Grow(0), layout.Grow(0)}, Rows: []layout.Track{layout.Fixed(1), layout.Fixed(1)}, Cells: []layout.Cell{{Node: area("a", 10, 1), ColumnSpan: 2}, {Node: area("b", 10, 1), Row: 1, Column: 1, ColumnSpan: 2}}}
	plan = requirePlan(t, grid, layout.Size{Width: 12, Height: 2})
	a, _ := plan.Bounds("a")

	b, _ := plan.Bounds("b")
	if a.Dx() < 10 || b.Dx() < 10 {
		t.Fatal("spanning constraints were not shared")
	}
}

func TestVerticalFlexMeasuresWrappedFlowBeforeSharingHeight(t *testing.T) {
	flow := &layout.Flow{Children: []layout.Box{{Node: area("a", 1, 1), Size: layout.Size{Width: 10, Height: 5}}, {Node: area("b", 1, 1), Size: layout.Size{Width: 10, Height: 5}}}}
	root := &layout.Flex{Axis: layout.Vertical, Children: []layout.Slot{{Node: flow, Track: layout.Grow(0)}, {Node: area("last", 1, 5), Track: layout.Grow(0)}}}
	plan := requirePlan(t, root, layout.Size{Width: 10, Height: 15})
	assertRect(t, plan, "a", image.Rect(0, 0, 10, 5))
	assertRect(t, plan, "b", image.Rect(0, 5, 10, 10))
	assertRect(t, plan, "last", image.Rect(0, 10, 10, 15))
}

func TestResponsiveContentGroupsUseSelectedSubtreeMinima(t *testing.T) {
	group := func(name string) layout.Node {
		pair := &layout.Flex{Children: []layout.Slot{{Node: area(name+"-a", 30, 1), Track: layout.Grow(0)}, {Node: area(name+"-b", 30, 1), Track: layout.Grow(0)}}}
		return &layout.Responsive{Name: name, Variants: []layout.Variant{{Name: "pair", Root: pair}}}
	}
	root := &layout.Flex{Axis: layout.Vertical, Children: []layout.Slot{{Node: group("first"), Track: layout.Grow(0)}, {Node: group("second"), Track: layout.Grow(0)}}}
	requirePlan(t, root, layout.Size{Width: 100, Height: 10})

	if _, err := layout.Resolve(root, layout.Size{Width: 120, Height: 10}); err == nil {
		t.Fatal("fitting selected groups did not trigger the stacking rule")
	}
}

func TestWeightedRoundingUsesLargestRemainder(t *testing.T) {
	root := &layout.Flex{Children: []layout.Slot{{Node: area("small", 1, 1), Track: layout.Track{Min: 1, Weight: 1}}, {Node: area("large", 1, 1), Track: layout.Track{Min: 1, Weight: 1000}}}}
	plan := requirePlan(t, root, layout.Size{Width: 3, Height: 1})
	assertRect(t, plan, "small", image.Rect(0, 0, 1, 1))
	assertRect(t, plan, "large", image.Rect(1, 0, 3, 1))
}
