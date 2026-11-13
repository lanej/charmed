package tui_test

import (
	"strings"
	"testing"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/tui"
)

func TestRetainedRevisionResizeAndRecovery(t *testing.T) {
	s := screen(t, "main", "details")
	root := responsive()

	c, err := layout.Compile(root)
	if err != nil {
		t.Fatal(err)
	}

	calls := 0
	render := func(layout.Size) string { calls++; return "content" }
	views := map[string]tui.RegionView{"main": {Render: render}, "details": {Render: render}}

	size := layout.Size{Width: 120, Height: 30}
	if err = s.ArrangeCompiled(c, size); err != nil {
		t.Fatal(err)
	}

	first, err := s.RenderRetained(views)
	if err != nil {
		t.Fatal(err)
	}

	if err = s.ArrangeCompiled(c, size); err != nil {
		t.Fatal(err)
	}

	second, err := s.RenderRetained(views)
	if err != nil || first.Content != second.Content || calls != 2 {
		t.Fatalf("cache miss: %d %v", calls, err)
	}

	views["main"] = tui.RegionView{Revision: 1, Render: render}
	if _, err = s.RenderRetained(views); err != nil || calls != 3 {
		t.Fatalf("revision: %d %v", calls, err)
	}

	if err = s.ArrangeCompiled(c, layout.Size{Width: 40, Height: 12}); err != nil {
		t.Fatal(err)
	}

	delete(views, "details")

	if _, err = s.RenderRetained(views); err != nil || calls != 4 {
		t.Fatalf("resize: %d %v", calls, err)
	}

	if err = s.ArrangeCompiled(c, layout.Size{}); err == nil {
		t.Fatal("invalid size accepted")
	}

	if _, err = s.RenderRetained(views); err == nil {
		t.Fatal("stale frame returned")
	}
}

func TestCompiledOwnsDeclaration(t *testing.T) {
	root := responsive()

	c, err := layout.Compile(root)
	if err != nil {
		t.Fatal(err)
	}

	responsiveRoot, ok := root.(*layout.Responsive)
	if !ok {
		t.Fatal("expected responsive declaration")
	}

	responsiveRoot.Variants[0].Name = "mutated"

	p, err := c.Resolve(layout.Size{Width: 120, Height: 30})
	if err != nil || p.Variant() == "mutated" {
		t.Fatalf("mutable compiled declaration: %v", err)
	}
}

func TestRetainedDiagnosticsCannotBeMutated(t *testing.T) {
	s := screen(t, "main")

	root := &layout.Region{ID: "main", Exceptions: []layout.Exception{{Rule: layout.SingleRegion, Reason: "Single test surface"}, {Rule: layout.ContentClipped, Reason: "Bounded test preview"}}}
	if err := s.Arrange(root, layout.Size{Width: 10, Height: 2}); err != nil {
		t.Fatal(err)
	}

	views := map[string]tui.RegionView{"main": {Render: func(layout.Size) string { return strings.Repeat("x", 100) }}}

	first, err := s.RenderRetained(views)
	if err != nil {
		t.Fatal(err)
	}

	if len(first.Diagnostics) == 0 {
		t.Fatal("missing clipping diagnostic")
	}

	for i := range first.Diagnostics {
		first.Diagnostics[i].Message = "tampered"
		if len(first.Diagnostics[i].Regions) > 0 {
			first.Diagnostics[i].Regions[0] = "tampered"
		}
	}

	second, err := s.RenderRetained(views)
	if err != nil {
		t.Fatal(err)
	}

	for _, diagnostic := range second.Diagnostics {
		if diagnostic.Message == "tampered" || (len(diagnostic.Regions) > 0 && diagnostic.Regions[0] == "tampered") {
			t.Fatal("returned diagnostics alias cache")
		}
	}

	copied := *s

	third, err := copied.RenderRetained(views)
	if err != nil || third.Content != second.Content {
		t.Fatalf("copied screen changed: %v", err)
	}
}
