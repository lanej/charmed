package layout_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestChromePriorityAndAlignment(t *testing.T) {
	palette := theme.New(theme.Mono)

	header := layout.Header{Title: "Title", Context: "Production", Actions: "Refresh"}
	if got := ansi.Strip(header.View(16, palette)); got != "Title    Refresh\nProduction      " {
		t.Fatalf("header: %q", got)
	}

	if got := ansi.Strip(header.View(8, palette)); got != "Title   \nProducti" {
		t.Fatalf("narrow header: %q", got)
	}

	footer := layout.Footer{Help: "q quit", Status: "Online"}
	if got := ansi.Strip(footer.View(16, palette)); got != "q quit    Online" {
		t.Fatalf("footer: %q", got)
	}

	if got := ansi.Strip(footer.View(8, palette)); got != "q quit  " {
		t.Fatalf("narrow footer: %q", got)
	}
}

func TestChromeCellGeometry(t *testing.T) {
	palette := theme.New(theme.Dark)
	for width := 0; width < 30; width++ {
		for _, view := range []string{
			(layout.Header{Title: "界界\nignored", Context: "Context", Actions: "Refresh"}).View(width, palette),
			(layout.Footer{Help: "界界\nignored", Status: "Online"}).View(width, palette),
		} {
			if width == 0 && view != "" {
				t.Fatalf("zero allocation: %q", view)
			}

			for _, line := range strings.Split(view, "\n") {
				if ansi.StringWidth(line) != width {
					t.Fatalf("width %d: %q", width, line)
				}
			}
		}
	}
}
