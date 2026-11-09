package collection

import (
	"strings"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
)

const timelineMinimumWidth = 36

// TimelineView projects retained requests onto one shared time axis. now is an
// explicit observation time for pending requests; zero leaves their end unknown.
// Completed stops require a known nonnegative duration. Group order is first-seen;
// requests retain ingestion order within each group.
func (m *HTTPLog) TimelineView(now time.Time, width, height int) string {
	width, height = max(1, width), max(1, height)
	if len(m.records) == 0 {
		return layout.Fit(m.Theme.Subtle().Render("No HTTP requests"), layout.Size{Width: width, Height: height})
	}

	if width < timelineMinimumWidth {
		return layout.Fit("Timeline needs at least 36 columns", layout.Size{Width: width, Height: height})
	}

	start, stop := timelineRange(m.records, now)
	if start.IsZero() {
		return layout.Fit(m.Theme.Subtle().Render("Request start times unavailable"), layout.Size{Width: width, Height: height})
	}

	span := stop.Sub(start)
	if span <= 0 {
		span = time.Millisecond
	}

	labelWidth := min(38, max(10, width/3))
	metadata := width >= 64

	plotWidth := width - labelWidth - 2
	if metadata {
		plotWidth -= 15
	}

	rows := []string{layout.Fit(m.Theme.Subtle().Render(start.UTC().Format("15:04:05.000Z")), layout.Size{Width: labelWidth, Height: 1}) + "  " + m.timelineAxis(span, plotWidth)}

	available := height - 1
	if m.TimelineLegend {
		available--
	}

	entries := m.timelineRows()
	selected := 0

	for i, row := range entries {
		if row.record.ID == m.selectedID {
			selected = i
			break
		}
	}

	visible := timelineWindow(entries, selected, max(0, available))
	for _, row := range visible {
		if row.header {
			heading := "  " + text.Inline(row.group)

			rule := strings.Repeat("─", max(0, width-ansi.StringWidth(heading)-1))
			if m.Theme.ASCII {
				rule = strings.Repeat("-", max(0, width-ansi.StringWidth(heading)-1))
			}

			rows = append(rows, ansi.Truncate(m.Theme.Subtle().Bold(true).Render(heading+" "+rule), width, "…"))

			continue
		}

		record := row.record

		marker := "  "
		if record.ID == m.selectedID {
			marker = m.Theme.Heading().Render("> ")
		}

		label := unpadHTTPAtom(text.HTTPMethod(record.Method, m.Theme)) + " " + m.target(record.URL)
		label = marker + ansi.Truncate(label, labelWidth-2, "…")

		line := layout.Fit(label, layout.Size{Width: labelWidth, Height: 1}) + "  " + m.timelinePlot(record, now, start, span, plotWidth)
		if metadata {
			line += " " + m.timelineMetadata(record, now)
		}

		rows = append(rows, line)
	}

	if m.TimelineLegend {
		legend := "├ start  ┤ stop  ▷ pending  ? unknown"
		if m.Theme.ASCII {
			legend = "[ start  ] stop  > pending  ? unknown"
		}

		rows = append(rows, ansi.Truncate(m.Theme.Subtle().Render(legend), width, "…"))
	}

	return layout.Fit(strings.Join(rows, "\n"), layout.Size{Width: width, Height: height})
}

func timelineSteps(span time.Duration, width int) int {
	if width < 32 || width < 4*(len(text.ShortDuration(span))+1)+4 {
		return 1
	}

	return 4
}

func (m *HTTPLog) timelineAxis(span time.Duration, width int) string {
	cells := make([]rune, width)
	for i := range cells {
		cells[i] = ' '
	}

	steps := timelineSteps(span, width)
	for i := 0; i <= steps; i++ {
		value := (span / time.Duration(steps)) * time.Duration(i)
		if i == steps {
			value = span
		}

		label := "+" + text.ShortDuration(value)
		if i == 0 {
			label = "0"
		}

		at := i * (width - 1) / steps

		at = min(at, width-len([]rune(label)))
		if at >= 0 {
			copy(cells[at:], []rune(label))
		}
	}

	return m.Theme.Subtle().Render(string(cells))
}

func (m *HTTPLog) timelinePlot(record HTTPExchange, now, start time.Time, span time.Duration, width int) string {
	if record.Time.IsZero() {
		return layout.Fit(m.Theme.Subtle().Render("? start unavailable"), layout.Size{Width: width, Height: 1})
	}

	grid := make([]rune, width)
	for i := range grid {
		grid[i] = ' '
	}

	steps := timelineSteps(span, width)
	for i := 0; i <= steps; i++ {
		grid[i*(width-1)/steps] = '┊'
		if m.Theme.ASCII {
			grid[i*(width-1)/steps] = '|'
		}
	}

	left := timelineCell(record.Time, start, span, width)
	bar := ansi.Cut(m.timelineBar(record, now, start, span, width), left, width)
	end := min(width, left+ansi.StringWidth(bar))

	return m.Theme.Subtle().Render(string(grid[:left])) + bar + m.Theme.Subtle().Render(string(grid[end:]))
}

func (m *HTTPLog) timelineMetadata(record HTTPExchange, now time.Time) string {
	code := unpadHTTPAtom(text.HTTPStatusCode(record.StatusCode, m.Theme))
	elapsed := "?"

	if !record.Completed {
		if record.StatusCode == 0 {
			code = text.Pending(m.Theme)
		}

		if !record.Time.IsZero() && !now.IsZero() && !now.Before(record.Time) {
			elapsed = text.ShortDuration(now.Sub(record.Time))
		}
	} else if record.DurationKnown && record.Duration >= 0 {
		elapsed = text.ShortDuration(record.Duration)
	}

	code = ansi.Truncate(code, 5, "…")
	duration := ansi.Truncate(elapsed, 8, "…")

	return strings.Repeat(" ", max(0, 5-ansi.StringWidth(code))) + code + " " + strings.Repeat(" ", max(0, 8-ansi.StringWidth(duration))) + duration
}

func timelineEnd(record HTTPExchange, now time.Time) (time.Time, bool) {
	if record.Time.IsZero() {
		return time.Time{}, false
	}

	if record.Completed {
		if !record.DurationKnown || record.Duration < 0 {
			return record.Time, false
		}

		return record.Time.Add(record.Duration), true
	}

	if now.IsZero() || now.Before(record.Time) {
		return record.Time, false
	}

	return now, true
}

func timelineRange(records []HTTPExchange, now time.Time) (time.Time, time.Time) {
	var start, stop time.Time

	for _, record := range records {
		if record.Time.IsZero() {
			continue
		}

		if start.IsZero() || record.Time.Before(start) {
			start = record.Time
		}

		end, _ := timelineEnd(record, now)
		if stop.IsZero() || end.After(stop) {
			stop = end
		}
	}

	return start, stop
}

func timelineCell(at, start time.Time, span time.Duration, width int) int {
	return min(width-1, max(0, int(float64(at.Sub(start))/float64(span)*float64(width-1))))
}

func (m *HTTPLog) timelineBar(record HTTPExchange, now, start time.Time, span time.Duration, width int) string {
	if record.Time.IsZero() {
		return m.Theme.Subtle().Render("? start unavailable")
	}

	end, known := timelineEnd(record, now)
	left := timelineCell(record.Time, start, span, width)
	right := timelineCell(end, start, span, width)

	begin, fill, finish := "├", "━", "┤"
	if m.Theme.ASCII {
		begin, fill, finish = "[", "=", "]"
	}

	style := logStyle(m.Theme, m.Theme.Unknown)

	if !record.Completed {
		finish = "▷"
		if m.Theme.ASCII {
			finish = ">"
		}

		style = logStyle(m.Theme, m.Theme.Pending)
	} else if record.StatusCode >= 500 && record.StatusCode < 600 {
		style = logStyle(m.Theme, m.Theme.Danger)
	} else if record.StatusCode >= 400 && record.StatusCode < 500 {
		style = logStyle(m.Theme, m.Theme.Warning)
	} else if (record.StatusCode >= 100 && record.StatusCode < 200) || (record.StatusCode >= 300 && record.StatusCode < 400) {
		style = logStyle(m.Theme, m.Theme.Info)
	} else if record.StatusCode >= 200 && record.StatusCode < 300 {
		style = logStyle(m.Theme, m.Theme.Success)
	}

	if !known {
		finish = "?"
	}

	bar := finish
	if right > left {
		bar = begin + strings.Repeat(fill, right-left-1) + finish
	}

	return strings.Repeat(" ", left) + style.Render(bar)
}
