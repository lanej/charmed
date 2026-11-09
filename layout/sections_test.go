package layout_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestIdentityHeaderRetainsIdentifierAcrossWidths(t *testing.T) {
	const id = "01234567-89ab-cdef-0123-456789abcdef"
	for _, width := range []int{80, 40, 36, 20, 1, 0} {
		view := ansi.Strip((layout.IdentityHeader{Kind: "Request", ID: id, Actions: "refresh"}).View(width, theme.New(theme.Dark)))
		if width == 0 {
			if view != "" {
				t.Fatal("zero allocation rendered")
			}

			continue
		}

		var joined string

		for _, line := range strings.Split(view, "\n") {
			if ansi.StringWidth(line) != width {
				t.Fatalf("width %d: %q", width, line)
			}

			joined += strings.TrimSpace(line)
		}

		if !strings.Contains(joined, id) {
			t.Fatalf("width %d lost identifier: %q", width, view)
		}

		if width >= 36 && !strings.Contains(view, id) {
			t.Fatalf("identifier wrapped despite fitting: %q", view)
		}
	}

	view := (layout.IdentityHeader{Kind: "Request", ID: "id\x1b[2J\nnext"}).View(80, theme.New(theme.Mono))
	if strings.Contains(view, "\x1b[2J") || !strings.Contains(view, `\u001b`) {
		t.Fatalf("identifier controls not escaped: %q", view)
	}
}

func TestSectionFoldingReturnsHeightWithoutShowingContent(t *testing.T) {
	section := layout.Section{Title: "Timeline", Content: "captured request", Collapsed: true}
	if section.Height(20) != 1 || section.Height(0) != 0 {
		t.Fatal("folded height")
	}

	view := ansi.Strip(section.View(layout.Size{Width: 40, Height: 20}, theme.New(theme.Mono), true))
	if strings.Contains(view, "captured") || !strings.Contains(view, "Timeline") || strings.Contains(view, "\n") {
		t.Fatalf("folded panel: %q", view)
	}

	section.NoFolds = true

	view = ansi.Strip(section.View(layout.Size{Width: 40, Height: 20}, theme.New(theme.Mono), true))
	if section.Height(20) != 20 || !strings.Contains(view, "captured request") {
		t.Fatal("disabled folds did not restore content")
	}
}
