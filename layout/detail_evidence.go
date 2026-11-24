package layout

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Short facts still form a readable column inside a paired detail section.
func detailFacts(fields []Pair, width int, t theme.Theme) string {
	labelWidth, valueWidth := 0, 0
	for _, field := range fields {
		labelWidth = max(labelWidth, ansi.StringWidth(text.Inline(field.Label)))
		valueWidth = max(valueWidth, ansi.StringWidth(field.State.View(field.Value, text.Normal, t)))
	}

	stacked := StackedAt(width) && labelWidth+theme.SpaceMedium+valueWidth > width

	return recordWrap(keyValue(fields, width, stacked, t), width)
}

// StatusItem keeps a named condition and its explanation together. Label and
// Detail are escaped data; Value is trusted styled display text. It can include
// a native state atom. An empty Value is a present empty string, rendered as "".
type StatusItem struct{ Label, Value, Detail string }

func (item StatusItem) value() string {
	if item.Value == "" {
		return `""`
	}

	return item.Value
}

// StatusList aligns named conditions and keeps explanations in the value column.
// At small widths the value follows its label on a wrapped line.
func StatusList(items []StatusItem, width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	longest := 0
	for _, item := range items {
		longest = max(longest, ansi.StringWidth(text.Inline(item.Label)))
	}

	labelWidth := min(longest, width/2)
	gap := strings.Repeat(" ", theme.SpaceMedium)
	rows := make([]string, 0, len(items))

	for _, item := range items {
		label := t.Subtle().Render(text.Inline(item.Label))
		valueWidth := width

		if width >= 32 && labelWidth > 0 {
			valueWidth -= labelWidth + theme.SpaceMedium
			rows = append(rows, Row(recordWrap(label, labelWidth), gap, statusValue(item, valueWidth, t)))

			continue
		}

		row := TextRow{Gap: theme.SpaceMedium, Items: []TextRowItem{{Text: label}, {Text: item.value()}}}
		rows = append(rows, row.View(width))

		if item.Detail != "" {
			rows = append(rows, recordWrap(t.Subtle().Render(text.Escape(item.Detail)), width))
		}
	}

	return strings.Join(rows, "\n")
}

func statusValue(item StatusItem, width int, t theme.Theme) string {
	value := recordWrap(item.value(), width)
	if item.Detail != "" {
		value += "\n" + recordWrap(t.Subtle().Render(text.Escape(item.Detail)), width)
	}

	return value
}

// stageSequence puts comparable stage labels over their observed values. Below
// the width needed for independent columns, each label sits next to its value.
// Connectors indicate reading order, independent of each supplied state.
func stageSequence(items []StatusItem, width int, t theme.Theme) string {
	const (
		minimumColumn = 16
		connection    = 3
	)

	if len(items) > (width+connection)/(minimumColumn+connection) {
		return StatusList(items, width, t)
	}

	available := width - (len(items)-1)*connection
	columns := make([]string, 0, 2*len(items)-1)
	arrow := t.BorderStyle(false).Render(" " + t.Glyph(theme.GlyphRight) + " ")

	for i, item := range items {
		cellWidth := available / len(items)
		if i < available%len(items) {
			cellWidth++
		}

		if i > 0 {
			columns = append(columns, arrow)
		}

		label := t.Text().Bold(true).Render(text.Inline(item.Label))
		columns = append(columns, recordWrap(label, cellWidth)+"\n"+statusValue(item, cellWidth, t))
	}

	return Row(columns...)
}

// TimelineEvent is one caller-ordered event. Every field is escaped data. When
// is the caller's formatted observation time. Rendering neither sorts nor infers one.
type TimelineEvent struct{ When, Actor, Title, Detail string }

// Timeline lines up event times beside a continuous rail. At small widths the
// supplied time and actor move below the title, preserving every observation.
func Timeline(events []TimelineEvent, width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	timeWidth := 0
	for _, event := range events {
		timeWidth = max(timeWidth, ansi.StringWidth(text.Inline(event.When)))
	}

	if width-timeWidth < 32 || timeWidth > width/3 {
		timeWidth = 0
	}

	rows := make([]string, 0, len(events))
	for i, event := range events {
		rows = append(rows, timelineEvent(event, width, timeWidth, i < len(events)-1, t))
	}

	return strings.Join(rows, "\n")
}

func timelineEvent(event TimelineEvent, width, timeWidth int, continued bool, t theme.Theme) string {
	railWidth := theme.SpaceMedium
	if width < 24 {
		railWidth = 0
	}

	bodyWidth := width - railWidth
	time := event.When

	if time != "" {
		time = t.Subtle().Render(text.Inline(time))
	}

	if timeWidth > 0 {
		bodyWidth -= timeWidth + theme.SpaceMedium
	}

	head := TextRow{Gap: theme.SpaceMedium}
	if event.Title != "" {
		head.Items = append(head.Items, TextRowItem{Text: t.Text().Bold(true).Render(text.Inline(event.Title))})
	}

	if event.Actor != "" {
		head.Items = append(head.Items, TextRowItem{Text: t.Subtle().Render(text.Inline(event.Actor))})
	}

	body := head.View(bodyWidth)
	if timeWidth == 0 && time != "" {
		body = joinEvidenceLines(body, recordWrap(time, bodyWidth))
	}

	if event.Detail != "" {
		body = joinEvidenceLines(body, recordWrap(t.Subtle().Render(text.Escape(event.Detail)), bodyWidth))
	}

	if body == "" && time == "" {
		body = recordWrap(`""`, bodyWidth)
	}

	lines := strings.Split(body, "\n")
	for i, line := range lines {
		prefix := ""

		if timeWidth > 0 {
			field := ""
			if i == 0 {
				field = time
			}

			prefix = Fit(field, Size{Width: timeWidth + theme.SpaceMedium, Height: 1})
		}

		if railWidth > 0 {
			mark := " "

			switch {
			case i == 0:
				mark = t.Glyph(theme.GlyphUnselected)
			case continued:
				mark = t.FrameBorder().Left
			default:
			}

			prefix += t.BorderStyle(false).Render(mark) + " "
		}

		lines[i] = prefix + line
	}

	return strings.Join(lines, "\n")
}

func joinEvidenceLines(first, next string) string {
	if first == "" {
		return next
	}

	return first + "\n" + next
}
