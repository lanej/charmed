package layout

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// LinkGroup lays out related destinations in declaration order with a two-cell
// gap. Links retain their individual presentation and emphasis. View never
// mutates the supplied slice or opens a target.
type LinkGroup struct {
	Links []text.Link
}

// View wraps between complete items, then within a single oversized label.
// Each wrapped label retains its exact target. Empty groups render nothing;
// nonempty groups require a positive width that can fit each label's graphemes.
// Errors identify the item and return no partial output. Height is unbounded:
// callers must allocate the returned rows or place the group in a viewport.
func (g LinkGroup) View(width int, t theme.Theme) (string, error) {
	if len(g.Links) == 0 {
		return "", nil
	}

	if width <= 0 {
		return "", fmt.Errorf("link group requires a positive width, got %d", width)
	}

	var (
		rows     []string
		row      string
		rowWidth int
	)

	gap := strings.Repeat(" ", theme.SpaceMedium)

	for i, link := range g.Links {
		view, err := link.View(t)
		if err != nil {
			return "", fmt.Errorf("link group item %d: %w", i, err)
		}

		itemWidth := ansi.StringWidth(view)
		if itemWidth > width {
			if row != "" {
				rows = append(rows, row)
				row, rowWidth = "", 0
			}

			wrapped := strings.Split(lipgloss.Wrap(view, width, ""), "\n")
			for _, line := range wrapped {
				if ansi.StringWidth(line) > width {
					return "", fmt.Errorf("link group item %d has a grapheme wider than %d cells", i, width)
				}
			}

			rows = append(rows, wrapped...)

			continue
		}

		if row != "" && rowWidth+theme.SpaceMedium+itemWidth > width {
			rows = append(rows, row)
			row, rowWidth = "", 0
		}

		if row != "" {
			row += gap
			rowWidth += theme.SpaceMedium
		}

		row += view
		rowWidth += itemWidth
	}

	if row != "" {
		rows = append(rows, row)
	}

	return strings.Join(rows, "\n"), nil
}

// Plain lists every label and full URI on its own line without terminal
// controls. Visual presentation and emphasis do not hide labels or targets.
func (g LinkGroup) Plain() (string, error) {
	rows := make([]string, 0, len(g.Links))
	for i, link := range g.Links {
		value, err := link.Plain()
		if err != nil {
			return "", fmt.Errorf("link group item %d: %w", i, err)
		}

		rows = append(rows, value)
	}

	return strings.Join(rows, "\n"), nil
}
