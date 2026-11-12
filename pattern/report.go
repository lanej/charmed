package pattern

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// ReportGap is the blank cells between panels sitting side by side.
const ReportGap = 1

// ReportPanel is one titled panel of a ReportPage.
type ReportPanel struct {
	Title string
	// Body lays the panel's text out for the given content width, which excludes the border and gutters. Lines
	// wider than the width are wrapped, never cut. A nil Body, or one returning no text, omits the panel.
	Body func(width int) string
	// Beside asks to share a row with the previous panel when the page is wide enough. The first panel of a page
	// and every panel at or below layout.StackWidth stack regardless.
	Beside bool
	// Weight shares a row's width among its panels in proportion; zero counts as one.
	Weight int
	// Min is the narrowest outer width the panel accepts beside another; zero uses DefaultReportPanelMin.
	Min int
}

// DefaultReportPanelMin is the narrowest side-by-side panel when ReportPanel.Min is zero.
const DefaultReportPanelMin = 30

// ReportPage is a finite, non-interactive report: header chrome carrying the overall evidence state, a summary
// line, a responsive grid of titled panels and the evidence footer.
//
// The overall state is the worst counted state (feedback.Worst), so a page never reads healthier than its weakest
// evidence and no evidence reads unknown. The footer always carries that state. Rendering is a pure function of the
// value, the width and the theme.
type ReportPage struct {
	Title, Context string
	Counts         []feedback.EvidenceCount
	Panels         []ReportPanel
	Footer         feedback.EvidenceFooter
}

// Overall is the page's evidence state.
func (r ReportPage) Overall() feedback.Evidence { return feedback.Worst(r.Counts) }

// View renders the page at width. No line exceeds width; its height follows its content.
func (r ReportPage) View(width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	return strings.Join(r.parts(width, t), "\n")
}

// ViewSize renders the page into size. When the panels do not fit, whole lines are dropped from the end of the
// panel area and an omission row says how many; the header, summary and footer always remain.
func (r ReportPage) ViewSize(size layout.Size, t theme.Theme) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}

	parts := r.parts(size.Width, t)
	head, body, foot := parts[:2], parts[2], parts[3]

	chrome := len(strings.Split(strings.Join(head, "\n"), "\n")) + len(strings.Split(foot, "\n"))
	lines := strings.Split(body, "\n")

	if body != "" && chrome+len(lines) > size.Height {
		keep := max(0, size.Height-chrome-1)
		omitted := text.Omission{Count: len(lines) - keep, KnownCount: true}
		lines = append(lines[:keep:keep], omitted.View(t))
	}

	if body == "" {
		lines = nil
	}

	out := append(append(head[:2:2], lines...), foot) //nolint:gocritic // head is capacity-limited

	return layout.Fit(strings.Join(out, "\n"), size)
}

func (r ReportPage) parts(width int, t theme.Theme) []string {
	overall := r.Overall()
	footer := r.Footer
	footer.State = overall

	head := layout.Header{Title: r.Title, Context: r.Context, Actions: overall.View(text.Bright, t)}.View(width, t)

	return []string{head, feedback.SummaryLine(r.Counts, width, t), r.grid(width, t), footer.View(width, t)}
}

// grid arranges the panels in rows. A row's panels sit side by side only when every minimum fits and the width is
// above layout.StackWidth; otherwise each takes the full width in order.
func (r ReportPage) grid(width int, t theme.Theme) string {
	var rows [][]ReportPanel

	for _, p := range r.Panels {
		if p.Body == nil {
			continue
		}

		if p.Beside && len(rows) > 0 && !layout.StackedAt(width) {
			rows[len(rows)-1] = append(rows[len(rows)-1], p)
			continue
		}

		rows = append(rows, []ReportPanel{p})
	}

	var blocks []string

	for _, row := range rows {
		for _, group := range fitRow(row, width) {
			if block := renderRow(group, width, t); block != "" {
				blocks = append(blocks, block)
			}
		}
	}

	return strings.Join(blocks, "\n")
}

// fitRow splits a row into groups whose minimums fit the width, so a panel that cannot sit beside the others
// stacks below them.
func fitRow(row []ReportPanel, width int) [][]ReportPanel {
	var groups [][]ReportPanel

	used := 0

	for _, p := range row {
		need := panelMin(p)
		if len(groups) > 0 && used+ReportGap+need <= width {
			groups[len(groups)-1] = append(groups[len(groups)-1], p)
			used += ReportGap + need

			continue
		}

		groups = append(groups, []ReportPanel{p})
		used = need
	}

	return groups
}

func panelMin(p ReportPanel) int {
	if p.Min > 0 {
		return p.Min
	}

	return DefaultReportPanelMin
}

// renderRow shares width among panels by weight, never below a panel's minimum, and gives every panel the height
// of the tallest.
func renderRow(row []ReportPanel, width int, t theme.Theme) string {
	widths := shareWidths(row, width)

	bodies := make([]string, len(row))
	height := 0

	for i, p := range row {
		content := layout.PanelContentSize(layout.Size{Width: widths[i], Height: 3}).Width
		bodies[i] = wrapLines(p.Body(content), content)

		if bodies[i] != "" {
			height = max(height, len(strings.Split(bodies[i], "\n")))
		}
	}

	if height == 0 {
		return ""
	}

	rendered := make([]string, 0, 2*len(row))

	for i, p := range row {
		if i > 0 {
			rendered = append(rendered, strings.Repeat(" ", ReportGap))
		}

		rendered = append(rendered, layout.Panel(
			p.Title, bodies[i], layout.Size{Width: widths[i], Height: height + 2*theme.BorderCells}, t, false))
	}

	if len(row) == 1 {
		return rendered[0]
	}

	return layout.Row(rendered...)
}

func shareWidths(row []ReportPanel, width int) []int {
	widths := make([]int, len(row))
	if len(row) == 1 {
		widths[0] = width

		return widths
	}

	total, floor := 0, 0
	for _, p := range row {
		total += max(1, p.Weight)
		floor += panelMin(p)
	}

	spare := width - ReportGap*(len(row)-1) - floor
	given := 0

	for i, p := range row {
		widths[i] = panelMin(p) + spare*max(1, p.Weight)/total
		given += widths[i]
	}

	// WHY: Integer division leaves a few cells; the last panel absorbs them so the row spans the width exactly.
	widths[len(widths)-1] += width - ReportGap*(len(row)-1) - given

	return widths
}

// wrapLines wraps styled text at width, hard-wrapping words longer than a line, so panel content is never cut.
func wrapLines(body string, width int) string {
	if width <= 0 {
		return ""
	}

	return ansi.Hardwrap(ansi.Wrap(strings.TrimRight(body, "\n"), width, ""), width, true)
}
