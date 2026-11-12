package layout_test

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestOverlayRetainsCustomForegroundWithoutDimTokens(t *testing.T) {
	palette := theme.Theme{Foreground: lipgloss.Color("#123456")}
	background := palette.Text().Render("background")

	size := layout.Size{Width: 20, Height: 4}
	explicit := palette
	explicit.DimForeground = palette.Foreground

	if layout.CenterOverlay(background, "", size, palette, true) != layout.CenterOverlay(background, "", size, explicit, true) {
		t.Fatal("missing dim tokens cleared custom foreground")
	}
}

func TestCenteredOverlayClipsAndPreservesCellGeometry(t *testing.T) {
	for _, size := range []layout.Size{{Width: 20, Height: 7}, {Width: 4, Height: 2}, {Width: 1, Height: 1}, {}} {
		for _, dim := range []bool{false, true} {
			view := layout.CenterOverlay(strings.Repeat("界 background\n", 7), "[command]", size, theme.New(theme.Dark), dim)
			if size.Width == 0 {
				if view != "" {
					t.Fatal("zero allocation rendered")
				}

				continue
			}

			lines := strings.Split(ansi.Strip(view), "\n")
			if len(lines) != size.Height {
				t.Fatal("overlay height overflow")
			}

			for _, line := range lines {
				if ansi.StringWidth(line) != size.Width {
					t.Fatal("overlay changed geometry", size, line)
				}
			}

			if size.Width == 20 && !strings.Contains(lines[3], "[command]") {
				t.Fatal("overlay not centered", view)
			}
		}
	}
}
