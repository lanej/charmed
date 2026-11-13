package tui_test

import (
	"image"
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
	"github.com/lanej/charmed/tui"
)

func responsive() layout.Node {
	return &layout.Responsive{Name: "workspace", Variants: []layout.Variant{
		{Name: "wide", Min: layout.Size{Width: 80, Height: 16}, Root: &layout.SplitPane{First: &layout.Region{ID: "main", Min: layout.Size{Width: 30, Height: 2}}, Second: &layout.Region{ID: "details", Min: layout.Size{Width: 30, Height: 2}}, Gap: 1}},
		{Name: "focused", Root: &layout.Region{ID: "main", Min: layout.Size{Width: 10, Height: 2}, Exceptions: []layout.Exception{{Rule: layout.SingleRegion, Reason: "Narrow focus view; details are available through the parent focus switch."}}}, Hidden: []layout.Omission{{ID: "details", Reason: "Focused view selected by terminal size."}}},
	}}
}

func screen(t *testing.T, ids ...string) *tui.Screen {
	t.Helper()

	s, err := tui.NewScreen(ids...)
	if err != nil {
		t.Fatal(err)
	}

	return s
}

func TestScreenCoverageResizeAndValidation(t *testing.T) {
	s := screen(t, "main", "details")

	cases, err := s.CheckLayouts(responsive(), layout.Size{Width: 120, Height: 30}, layout.Size{Width: 40, Height: 12})
	if err != nil || len(cases) != 2 {
		t.Fatalf("validation: %v", err)
	}

	if _, err := s.CheckLayouts(responsive(), layout.Size{Width: 120, Height: 30}); err == nil {
		t.Fatal("unexercised responsive variant accepted")
	}

	if err := s.Arrange(responsive(), layout.Size{Width: 120, Height: 30}); err != nil {
		t.Fatal(err)
	}

	if bounds, ok := s.Region("details"); !ok || bounds != image.Rect(60, 0, 120, 30) {
		t.Fatalf("unexpected bounds: %v", bounds)
	}

	if err := s.Arrange(responsive(), layout.Size{Width: 40, Height: 12}); err != nil {
		t.Fatal(err)
	}

	if _, ok := s.Region("details"); ok {
		t.Fatal("hidden region has an allocation")
	}

	if !strings.Contains(s.Inspect(), "Focused view selected") {
		t.Fatal("hidden reason is not inspectable")
	}

	if err := s.Arrange(&layout.Region{ID: "main"}, layout.Size{Width: 40, Height: 12}); err == nil {
		t.Fatal("missing region accepted")
	}

	if _, err := s.Render(map[string]string{"main": "old"}); err == nil {
		t.Fatal("failed resize kept rendering stale arrangement")
	}

	if _, ok := s.Region("main"); ok {
		t.Fatal("failed resize exposed stale bounds")
	}
}

func TestFrameRejectsMissingUnknownAndClippedViews(t *testing.T) {
	s := screen(t, "main", "details")
	if err := s.Arrange(responsive(), layout.Size{Width: 80, Height: 20}); err != nil {
		t.Fatal(err)
	}

	for _, views := range []map[string]string{{"main": "x"}, {"main": "x", "unknown": "y"}, {"main": strings.Repeat("x", 100), "details": "ok"}} {
		frame, err := s.Render(views)
		if err == nil || frame.Content != "" {
			t.Fatal("invalid render returned a frame")
		}
	}

	left := lipgloss.NewStyle().Bold(true).Render("界e\u0301")

	frame, err := s.Render(map[string]string{"main": left, "details": "inspect"})
	if err != nil {
		t.Fatal(err)
	}

	if lipgloss.Width(frame.Content) != 80 || lipgloss.Height(frame.Content) != 20 {
		t.Fatal("frame escaped its allocation")
	}

	if !strings.Contains(ansi.Strip(frame.Content), "界e\u0301") {
		t.Fatal("grapheme data was lost")
	}

	if !strings.Contains(s.RegionMap(theme.New(theme.Dark)), "details") {
		t.Fatal("region map missing identity")
	}
}

func TestClippingExceptionsRemainScopedAndInspectable(t *testing.T) {
	s := screen(t, "label", "value")

	root := &layout.Flex{Children: []layout.Slot{
		{Node: &layout.Region{ID: "label", Exceptions: []layout.Exception{{Rule: layout.ContentClipped, Reason: "The full identifier is available in the inspector."}}}, Track: layout.Fixed(5)},
		{Node: &layout.Region{ID: "value"}, Track: layout.Grow(1)},
	}}
	if err := s.Arrange(root, layout.Size{Width: 20, Height: 2}); err != nil {
		t.Fatal(err)
	}

	frame, err := s.Render(map[string]string{"label": "long identifier", "value": "ok"})
	if err != nil {
		t.Fatal(err)
	}

	if len(frame.Diagnostics) != 1 || frame.Diagnostics[0].Exception == "" {
		t.Fatal("clipping exception not reported")
	}

	if _, err := s.Render(map[string]string{"label": "ok", "value": strings.Repeat("x", 30)}); err == nil {
		t.Fatal("exception leaked into another region")
	}
}
