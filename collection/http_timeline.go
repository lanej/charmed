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
// Completed stops require a known nonnegative duration. Input order is retained.
func (m *HTTPLog) TimelineView(now time.Time, width, height int) string {
	width, height = max(1, width), max(1, height)

	if len(m.records) == 0 {
		return m.Theme.Subtle().Render("No HTTP requests")
	}

	if width < timelineMinimumWidth {
		return ansi.Truncate("Timeline needs at least 36 columns", width, "…")
	}

	labelWidth := min(36, width/3)
	plotWidth := width - labelWidth - 2

	start, stop := timelineRange(m.records, now)
	if start.IsZero() {
		return m.Theme.Subtle().Render("Request start times unavailable")
	}

	span := stop.Sub(start)
	if span <= 0 {
		span = time.Millisecond
	}

	axis := start.UTC().Format("15:04:05.000Z")
	endLabel := "+" + text.ShortDuration(span)
	axis += strings.Repeat(" ", max(1, plotWidth-ansi.StringWidth(axis)-ansi.StringWidth(endLabel))) + endLabel
	axis = ansi.Truncate(axis, plotWidth, "…")
	rows := []string{strings.Repeat(" ", labelWidth+2) + axis}

	reserved := 1
	if m.TimelineLegend {
		reserved++
	}

	available := max(0, height-reserved)

	first := max(0, m.Table.Cursor()-available+1)
	for _, record := range m.records[first:min(len(m.records), first+available)] {
		marker := "  "
		if record.ID == m.selectedID {
			marker = "> "
		}

		label := text.HTTPMethod(record.Method, m.Theme) + " " + m.target(record.URL)
		label = ansi.Truncate(label, labelWidth-2, "…")
		label = marker + label + strings.Repeat(" ", max(0, labelWidth-2-ansi.StringWidth(label)))
		rows = append(rows, label+"  "+m.timelineBar(record, now, start, span, plotWidth))
	}

	legend := "├ start  ┤ stop  ▷ pending  ? unknown"
	if m.Theme.ASCII {
		legend = "[ start  ] stop  > pending  ? unknown"
	}

	if m.TimelineLegend {
		rows = append(rows, ansi.Truncate(legend, width, "…"))
	}

	return layout.Fit(strings.Join(rows, "\n"), layout.Size{Width: width, Height: height})
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
