package collection

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// TableColumn describes one column of a StaticTable.
type TableColumn struct {
	Title string
	// Min is the narrowest the column may wrap to. Zero means the width of the title.
	Min int
	// Fixed keeps the column at its natural width; it never wraps.
	Fixed bool
}

// TableCell is one value. Exactly one of Styled or Text supplies the content.
type TableCell struct {
	// Text is untrusted plain text; control sequences are escaped.
	Text string
	// Styled is content the caller has already escaped and styled. It wins over Text.
	Styled string
	// State other than ValuePresent replaces the content with its marker, so a missing, redacted or invalid
	// value is never shown as an empty string.
	State text.ValueState
	// Evidence prefixes the content with the evidence state, whose word and glyph survive monochrome and ASCII.
	Evidence *feedback.Evidence
	// Omission follows the content when part of it was left out.
	Omission *text.Omission
}

// TableText is a cell of plain text.
func TableText(value string) TableCell { return TableCell{Text: value} }

// TableMissing is a cell whose value was not reported.
func TableMissing() TableCell { return TableCell{State: text.ValueMissing} }

// TableEvidence is a cell that leads with an evidence state, followed by optional detail.
func TableEvidence(state feedback.Evidence, detail string) TableCell {
	return TableCell{Evidence: &state, Text: detail}
}

// StaticTable is a non-interactive comparison table. The first column names the row. Cells wrap inside their
// column rather than truncate. At StackWidth columns and narrower, a table that would not fit whole stacks into one
// block per row, and a table that cannot give every column its Min width stacks at any width. A stacked table
// never truncates a value. Rows shorter than Columns show missing cells; cells beyond Columns are not shown.
//
// Use the collection Table for selection and keyboard navigation.
type StaticTable struct {
	Columns []TableColumn
	Rows    [][]TableCell
	// Empty replaces a table with no rows. It defaults to "no rows".
	Empty string
}

// Stacked reports whether View stacks at width.
func (s StaticTable) Stacked(width int, t theme.Theme) bool {
	_, stacked := s.layout(width, t)

	return stacked
}

// View renders the table into at most width columns.
func (s StaticTable) View(width int, t theme.Theme) string {
	width = max(1, width)

	if len(s.Columns) == 0 || len(s.Rows) == 0 {
		empty := s.Empty
		if empty == "" {
			empty = "no rows"
		}

		return ansi.Hardwrap(t.Subtle().Render(text.Inline(empty)), width, true)
	}

	cells := s.render(t)

	widths, stacked := s.layoutOf(cells, width)
	if stacked {
		return s.stack(cells, width, t)
	}

	return s.columns(cells, widths, t)
}

func (s StaticTable) layout(width int, t theme.Theme) ([]int, bool) {
	return s.layoutOf(s.render(t), max(1, width))
}

func (s StaticTable) layoutOf(cells [][]string, width int) ([]int, bool) {
	n := len(s.Columns)
	natural := make([]int, n)

	for i, column := range s.Columns {
		natural[i] = ansi.StringWidth(text.Inline(column.Title))
	}

	for _, row := range cells {
		for i, cell := range row {
			for _, line := range strings.Split(cell, "\n") {
				natural[i] = max(natural[i], ansi.StringWidth(line))
			}
		}
	}

	widths, least, whole := make([]int, n), 0, 0

	for i, column := range s.Columns {
		minimum := column.Min
		if minimum <= 0 {
			minimum = ansi.StringWidth(text.Inline(column.Title))
		}

		if column.Fixed {
			minimum = natural[i]
		}

		widths[i] = min(minimum, natural[i])
		least += widths[i]
		whole += natural[i]
	}

	available := width - 2*(n-1)
	if available < n || least > available || (layout.StackedAt(width) && whole > available) {
		return nil, true
	}

	for used := least; used < available; used++ {
		next := -1

		for i := range widths {
			if widths[i] < natural[i] && (next < 0 || widths[i] < widths[next]) {
				next = i
			}
		}

		if next < 0 {
			break
		}

		widths[next]++
	}

	return widths, false
}

func (s StaticTable) render(t theme.Theme) [][]string {
	rows := make([][]string, len(s.Rows))

	for r, row := range s.Rows {
		rows[r] = make([]string, len(s.Columns))

		for c := range s.Columns {
			cell := TableMissing()
			if c < len(row) {
				cell = row[c]
			}

			emphasis := text.Normal
			if c == 0 {
				emphasis = text.Bright
			}

			rows[r][c] = cell.view(emphasis, t)
		}
	}

	return rows
}

func (c TableCell) view(emphasis text.Emphasis, t theme.Theme) string {
	var parts []string

	if c.Evidence != nil {
		parts = append(parts, c.Evidence.View(text.Normal, t))
	}

	switch {
	case c.State != text.ValuePresent:
		parts = append(parts, c.State.View(c.Text, emphasis, t))
	case c.Styled != "":
		parts = append(parts, c.Styled)
	case c.Text != "" || c.Evidence == nil:
		parts = append(parts, text.ValuePresent.View(c.Text, emphasis, t))
	}

	if c.Omission != nil {
		if omitted := c.Omission.View(t); ansi.StringWidth(omitted) > 0 {
			parts = append(parts, omitted)
		}
	}

	return strings.Join(parts, " ")
}

// fitLines wraps a styled cell to width. Lipgloss closes and restores styles at each line break, so a background
// that wraps never tints the neighbouring column.
func fitLines(cell string, width int) []string {
	wrapped := lipgloss.NewStyle().Width(width).Render(ansi.Wrap(cell, width, ""))

	return strings.Split(ansi.Hardwrap(wrapped, width, true), "\n")
}

func (s StaticTable) columns(cells [][]string, widths []int, t theme.Theme) string {
	strong := text.ToneStyle(text.Neutral, text.Bright, t)
	headers := make([]string, len(s.Columns))

	for i, column := range s.Columns {
		headers[i] = strong.Render(text.Inline(column.Title))
	}

	var out []string

	for _, row := range append([][]string{headers}, cells...) {
		blocks, height := make([][]string, len(row)), 1

		for i, cell := range row {
			blocks[i] = fitLines(cell, widths[i])
			height = max(height, len(blocks[i]))
		}

		for line := range height {
			var b string

			for i := range blocks {
				cell := ""
				if line < len(blocks[i]) {
					cell = blocks[i][line]
				}

				b += cell

				if i < len(blocks)-1 {
					b += strings.Repeat(" ", widths[i]-ansi.StringWidth(cell)+2)
				}
			}

			out = append(out, strings.TrimRight(b, " "))
		}
	}

	return strings.Join(out, "\n")
}

func (s StaticTable) stack(cells [][]string, width int, t theme.Theme) string {
	labelMax := max(1, (width-2)/2)
	blocks := make([]string, 0, len(cells))

	for _, row := range cells {
		lines := fitLines(row[0], width)
		for n, line := range lines {
			lines[n] = strings.TrimRight(line, " ")
		}

		for i := 1; i < len(row); i++ {
			title := ansi.Truncate(text.Inline(s.Columns[i].Title), labelMax, t.Glyph(theme.GlyphTruncated))
			head := t.Subtle().Render(title+":") + " "
			indent := strings.Repeat(" ", ansi.StringWidth(title)+2)

			for n, line := range fitLines(row[i], max(1, width-ansi.StringWidth(indent))) {
				if n == 0 {
					line = head + line
				} else {
					line = indent + line
				}

				lines = append(lines, strings.TrimRight(line, " "))
			}
		}

		blocks = append(blocks, strings.Join(lines, "\n"))
	}

	return strings.Join(blocks, "\n\n")
}
