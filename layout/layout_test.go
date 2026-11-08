package layout_test

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestFitUsesDisplayCells(t *testing.T) {
	styled := lipgloss.NewStyle().Foreground(lipgloss.Color("#abcdef")).Render("界界界 hello") + "\nsecond\nthird"
	for _, size := range []layout.Size{{Width: 1, Height: 1}, {Width: 5, Height: 2}, {Width: 40, Height: 12}} {
		out := layout.Fit(styled, size)
		lines := strings.Split(out, "\n")
		if len(lines) != size.Height {
			t.Fatalf("got %d rows, want %d", len(lines), size.Height)
		}
		for _, line := range lines {
			if width := lipgloss.Width(line); width != size.Width {
				t.Fatalf("got %d cells, want %d: %q", width, size.Width, line)
			}
		}
	}
}

func TestPanelReservesGuttersInsideItsAllocation(t *testing.T) {
	for _, width := range []int{4, 5, 13, 40} {
		for _, focused := range []bool{false, true} {
			size := layout.Size{Width: width, Height: 4}

			view := layout.Panel("", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", size, theme.New(theme.Dark), focused)
			if lipgloss.Width(view) != width || lipgloss.Height(view) != size.Height {
				t.Fatalf("panel escaped allocation at width %d", width)
			}

			row := strings.Split(ansi.Strip(view), "\n")[1]
			if width >= 5 && (!strings.HasPrefix(row, "│ A") || !strings.HasSuffix(row, " │")) {
				t.Fatalf("missing content gutter: %q", row)
			}
		}
	}
}

func TestPanelIncludesItsBorderInAllocation(t *testing.T) {
	out := layout.Panel("A very long title", "wide 界 content", layout.Size{Width: 13, Height: 5}, theme.New(theme.Dark), true)
	if lipgloss.Width(out) != 13 || lipgloss.Height(out) != 5 {
		t.Fatalf("panel escaped allocation: %q", out)
	}
}
