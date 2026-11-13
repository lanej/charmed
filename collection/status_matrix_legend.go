package collection

import (
	"slices"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Legend explains the statuses that appear in the matrix, each as the glyph and
// word its cells use, in the status catalog's order. It wraps whole entries to
// width and never exceeds it; an entry wider than width is truncated. It is
// static text, so a printed matrix explains its cells without an interactive
// terminal. An empty matrix has no legend.
//
// A status the catalog does not recognize, including the zero value, is shown
// as UNKNOWN, exactly as its cells are, and never shares the healthy glyph or
// word. Absent mappings are not statuses and are not listed.
func (m *StatusMatrix) Legend(width int) string {
	if width <= 0 {
		return ""
	}

	present := m.presentStatuses()
	if len(present) == 0 {
		return ""
	}

	var lines []string

	line := ""

	for _, status := range present {
		entry := ansi.Truncate(status.View(text.Normal, m.Theme), width, m.Theme.Glyph(theme.GlyphTruncated))
		if line != "" && ansi.StringWidth(line)+2+ansi.StringWidth(entry) > width {
			lines, line = append(lines, line), ""
		}

		if line != "" {
			line += "  "
		}

		line += entry
	}

	return strings.Join(append(lines, line), "\n")
}

// presentStatuses returns the distinct catalog statuses used by cells, in
// catalog order, with unrecognized values counted as UNKNOWN.
func (m *StatusMatrix) presentStatuses() []feedback.Status {
	catalog := feedback.Statuses()
	seen := map[feedback.Status]bool{}

	for _, cell := range m.cells {
		status := cell.Status
		if !slices.Contains(catalog, status) {
			status = feedback.Unknown
		}

		seen[status] = true
	}

	var present []feedback.Status

	for _, status := range catalog {
		if seen[status] {
			present = append(present, status)
		}
	}

	return present
}

// View renders the grid and, when ShowLegend is set, the legend below it. The
// legend takes rows from the allocation; the grid keeps the rest.
func (m *StatusMatrix) View() string {
	if !m.ShowLegend {
		return m.DataGrid.View()
	}

	legend := m.Legend(m.size.Width)
	if legend == "" {
		return m.DataGrid.View()
	}

	size := m.size
	rows := strings.Count(legend, "\n") + 1

	m.size.Height = max(0, size.Height-rows)
	grid := m.DataGrid.View()
	m.size = size

	return layout.Fit(grid+"\n"+legend, size)
}
