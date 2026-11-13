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

func TestFitPreservesCellAndEscapeSemantics(t *testing.T) {
	for _, content := range []string{
		"", "abcdef", "a\r\nb\r\n", "a\rb\nc", "a\tb", "a\n\nlast\n",
		"界e\u0301👩‍💻🙂", "\x1b[31mcolored\x1b[0m\nnext",
		"\x1b]8;;https://example.com\x1b\\linked\x1b]8;;\x1b\\",
		"\x1b[44m   \x1b[0m", "visible\n" + strings.Repeat("hidden\r\n", 1000),
	} {
		for _, width := range []int{0, 1, 2, 5, 20} {
			for _, height := range []int{0, 1, 2, 5} {
				size := layout.Size{Width: width, Height: height}
				if got, want := layout.Fit(content, size), referenceFit(content, size); got != want {
					t.Fatalf("at %v for %q: got %q, want %q", size, content, got, want)
				}
			}
		}
	}
}

// Retain the original complete-line cell algorithm as a compatibility oracle.
func referenceFit(content string, size layout.Size) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}

	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")

	out := make([]string, size.Height)
	for i := range out {
		line := ""
		if i < len(lines) {
			line = ansi.Truncate(lines[i], size.Width, "")
		}

		out[i] = line + strings.Repeat(" ", max(0, size.Width-lipgloss.Width(line)))
	}

	return strings.Join(out, "\n")
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
