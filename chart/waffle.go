package chart

import (
	"fmt"
	"math"
	"math/bits"
	"sort"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// WafflePart is an exact count with a caller-selected semantic tone.
type WafflePart struct {
	Label string
	Count int
	Tone  text.Tone
}

// Waffle renders ordered categories as equal-area cells. Patterned square markers
// match Legend; ASCII uses letters A-Z. At most 26 parts are supported.
// Counts must be nonnegative and their sum must fit an int. Zero totals are empty.
type Waffle struct {
	Parts []WafflePart
}

// View fills a row-major grid with a space between cells. A width of 19 holds
// ten columns. Largest-remainder rounding favors input order for ties. Each cell represents
// total/cells observations, so small categories
// may receive no cells. Keep Legend or Header visible to retain exact counts.
// Allocations above 10,000 cells return an unpadded fallback.
func (w Waffle) View(size layout.Size, t theme.Theme) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}

	if size.Width > 10000/size.Height {
		return "Waffle too large"
	}

	total, valid := w.total()
	if !valid {
		return layout.Fit("Invalid waffle", size)
	}

	if total == 0 {
		return layout.Fit("No observations", size)
	}

	columns := (size.Width + 1) / 2
	counts := w.cells(columns*size.Height, total)
	cells := make([]string, 0, columns*size.Height)

	for i, count := range counts {
		style := text.ToneStyle(w.Parts[i].Tone, text.Normal, t)
		for range count {
			cells = append(cells, style.Render(waffleMarker(i, t)))
		}
	}

	rows := make([]string, size.Height)
	for y := range rows {
		rows[y] = strings.Join(cells[y*columns:(y+1)*columns], " ")
	}

	return layout.Fit(strings.Join(rows, "\n"), size)
}

// Legend retains exact counts, including categories rounded out of the grid.
func (w Waffle) Legend(t theme.Theme) string {
	total, valid := w.total()
	if !valid {
		return "Invalid waffle"
	}

	if total == 0 {
		return "No observations"
	}

	parts := make([]string, len(w.Parts))
	for i, part := range w.Parts {
		label := fmt.Sprintf("%s %s %d", waffleMarker(i, t), waffleLabel(part.Label), part.Count)
		parts[i] = text.ToneStyle(part.Tone, text.Normal, t).Render(label)
	}

	return strings.Join(parts, " / ")
}

// Header composes a title, up to 20 waffle cells and an exact legend on one row.
// When space is tight it drops the grid first, then clips the textual summary.
func (w Waffle) Header(title string, width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	summary := w.Legend(t)
	if title != "" {
		summary = t.Heading().Render(waffleLabel(title)) + "  " + summary
	}

	remaining := width - lipgloss.Width(summary) - 3

	total, valid := w.total()
	if valid && total > 0 && remaining >= 5 {
		grid := w.View(layout.Size{Width: min(39, remaining), Height: 1}, t)
		summary += " [" + grid + "]"
	}

	return layout.Fit(summary, layout.Size{Width: width, Height: 1})
}

// StatusHeader places a 20-cell micro-grid beside a title and exact counts in
// two rows. If either row cannot fit, it drops the grid before clipping text.
// A one-row allocation uses Header instead.
func (w Waffle) StatusHeader(title string, size layout.Size, t theme.Theme) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}

	if size.Height == 1 {
		return w.Header(title, size.Width, t)
	}

	total, valid := w.total()
	if !valid || total == 0 {
		return layout.Fit(w.Header(title, size.Width, t), size)
	}

	heading := t.Heading().Render(waffleLabel(title)) + t.Subtle().Render(fmt.Sprintf("  /  %d total", total))
	summary := w.Legend(t)
	content := heading + "\n" + summary

	const gridWidth = 19

	const gap = "   "

	if size.Width >= gridWidth+len(gap)+max(lipgloss.Width(heading), lipgloss.Width(summary)) {
		grid := w.View(layout.Size{Width: gridWidth, Height: 2}, t)
		content = layout.Row(grid, gap+heading+"\n"+gap+summary)
	}

	return layout.Fit(content, size)
}

func waffleMarker(index int, t theme.Theme) string {
	if index < 0 || index >= 26 {
		return "?"
	}

	if !t.ASCII {
		patterns := []string{"■", "▧", "▨", "▦"}
		if index < len(patterns) {
			return patterns[index]
		}
	}

	return string("ABCDEFGHIJKLMNOPQRSTUVWXYZ"[index])
}

func (w Waffle) total() (int, bool) {
	if len(w.Parts) > 26 {
		return 0, false
	}

	total := 0
	for _, part := range w.Parts {
		if part.Count < 0 || part.Count > math.MaxInt-total {
			return 0, false
		}

		total += part.Count
	}

	return total, true
}

func (w Waffle) cells(capacity, total int) []int {
	counts := make([]int, len(w.Parts))
	remainders := make([]uint64, len(w.Parts))
	order := make([]int, len(w.Parts))
	used := 0

	for i, part := range w.Parts {
		// WHY: A widened integer product keeps rounding exact even near MaxInt.
		high, low := bits.Mul64(uint64(part.Count), uint64(capacity)) //nolint:gosec // WHY: View validates nonnegative counts and a positive capacity <= 10,000.
		quotient, remainder := bits.Div64(high, low, uint64(total))   //nolint:gosec // WHY: View validates a positive total before allocating cells.
		counts[i] = int(quotient)                                     //nolint:gosec // WHY: Each count is at most total, so the quotient is bounded by capacity <= 10,000.
		remainders[i] = remainder
		order[i] = i
		used += counts[i]
	}

	sort.SliceStable(order, func(i, j int) bool { return remainders[order[i]] > remainders[order[j]] })

	for i := 0; i < capacity-used; i++ {
		counts[order[i]]++
	}

	return counts
}

func waffleLabel(label string) string { return strings.Join(strings.Fields(label), " ") }
