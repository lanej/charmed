package collection

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func TestSectionAnchorRetainsFallbackPriorityAndRelativeRow(t *testing.T) {
	model := &SectionStack{
		Viewport: NewViewport(strings.Repeat("line\n", 6), 20, 1),
		ranges: []sectionRange{
			{id: "outer", start: 0, end: 1},
			{id: "inner", start: 1, end: 2},
			{id: "previous", start: 2, end: 3},
			{id: "active", start: 3, end: 6},
		},
		parents: map[string]string{"active": "inner", "inner": "outer"},
	}
	model.Viewport.SetYOffset(4)

	anchor := model.anchor()
	if anchor.id != "active" || anchor.row != 1 {
		t.Fatalf("lost active section or relative row: %+v", anchor)
	}

	// The nearest ancestor wins before preceding headings in reverse order.
	want := []string{"inner", "outer", "previous", "inner", "outer"}
	if !slices.Equal(anchor.candidates, want) {
		t.Fatalf("fallback order = %v, want %v", anchor.candidates, want)
	}
}

func TestSectionAnchorAllocationBoundNearDocumentEnd(t *testing.T) {
	const sections = 2048

	model := &SectionStack{
		Viewport: NewViewport(strings.Repeat("line\n", sections), 20, 1),
		ranges:   make([]sectionRange, sections),
	}
	for i := range model.ranges {
		model.ranges[i] = sectionRange{id: fmt.Sprint(i), start: i, end: i + 1}
	}

	model.Viewport.SetYOffset(sections - 1)

	var anchor stackAnchor

	allocations := testing.AllocsPerRun(3, func() { anchor = model.anchor() })

	if len(anchor.candidates) != sections-1 || anchor.candidates[0] != "2046" || anchor.candidates[sections-2] != "0" {
		t.Fatal("anchor omitted or reordered preceding sections")
	}

	t.Logf("anchor capture: %.0f allocations for %d sections", allocations, sections)

	// Leave headroom for slice growth without allowing one allocation per ID.
	if allocations > 32 {
		t.Fatalf("anchor capture allocated %.0f times for %d sections; want at most 32", allocations, sections)
	}
}
