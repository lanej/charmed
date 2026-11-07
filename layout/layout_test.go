package layout_test

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
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
func TestPanelIncludesItsBorderInAllocation(t *testing.T) {
	out := layout.Panel("A very long title", "wide 界 content", layout.Size{Width: 13, Height: 5}, theme.New(theme.Dark), true)
	if lipgloss.Width(out) != 13 || lipgloss.Height(out) != 5 {
		t.Fatalf("panel escaped allocation: %q", out)
	}
}
