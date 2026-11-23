package layout

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// StatusItem keeps a named condition and its explanation together. Label and
// Detail are escaped data; Value is trusted styled display text so callers can
// use the existing health, outcome, readiness or other independent state atoms.
// Charmed does not infer one kind of state from another. An empty Value is a
// present empty string, rendered explicitly as "".
type StatusItem struct{ Label, Value, Detail string }

// StatusList aligns named conditions where space permits and stacks them at
// narrow widths. Explanations remain next to their condition. Height is unbounded.
func StatusList(items []StatusItem, width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	labelWidth := 0
	for _, item := range items {
		labelWidth = max(labelWidth, ansi.StringWidth(text.Inline(item.Label)))
	}

	labelWidth = min(labelWidth, max(1, width/2-theme.SpaceMedium))
	parts := make([]string, 0, len(items))

	for _, item := range items {
		label, value := text.Inline(item.Label), item.Value
		if value == "" {
			value = `""`
		}

		var row string

		if StackedAt(width) || label == "" {
			if label != "" {
				label += ":"
			}

			row = (TextRow{Gap: theme.SpaceSmall, Items: []TextRowItem{
				{Text: t.Subtle().Render(label)}, {Text: value},
			}}).View(width)
		} else {
			row = Row(recordWrap(t.Subtle().Render(label), labelWidth), strings.Repeat(" ", theme.SpaceMedium), recordWrap(value, width-labelWidth-theme.SpaceMedium))
		}

		if item.Detail != "" {
			row += "\n" + detailIndent(t.Subtle().Render(text.Escape(item.Detail)), width)
		}

		parts = append(parts, row)
	}

	return strings.Join(parts, "\n")
}

// TimelineEvent is one caller-ordered event. Every field is escaped data. When
// is an application-formatted observation label, not a clock input; no sorting,
// duration inference or implicit present-time observation occurs while rendering.
type TimelineEvent struct{ When, Actor, Title, Detail string }

// Timeline presents complete events with their supplied time and actor labels.
// It preserves declaration order and returns content-driven, width-bounded rows.
func Timeline(events []TimelineEvent, width int, t theme.Theme) string {
	if width <= 0 {
		return ""
	}

	parts := make([]string, 0, len(events))
	for _, event := range events {
		metadata := make([]TextRowItem, 0, 2)
		if event.When != "" {
			metadata = append(metadata, TextRowItem{Text: t.Subtle().Render(text.Inline(event.When))})
		}

		if event.Actor != "" {
			metadata = append(metadata, TextRowItem{Text: t.Subtle().Render(text.Inline(event.Actor))})
		}

		gutter := min(theme.SpaceMedium, max(0, width-1))
		bodyWidth := width - gutter
		body := make([]string, 0, 3)

		if event.Title != "" {
			body = append(body, recordWrap(t.Text().Bold(true).Render(text.Inline(event.Title)), bodyWidth))
		}

		if len(metadata) > 0 {
			body = append(body, (TextRow{Gap: theme.SpaceMedium, Items: metadata}).View(bodyWidth))
		}

		if event.Detail != "" {
			body = append(body, recordWrap(text.Escape(event.Detail), bodyWidth))
		}

		if len(body) == 0 {
			body = append(body, recordWrap(`""`, bodyWidth))
		}

		lines := strings.Split(strings.Join(body, "\n"), "\n")
		for i, line := range lines {
			prefix := strings.Repeat(" ", gutter)
			if i == 0 && gutter > 0 {
				prefix = t.Subtle().Render(t.Glyph(theme.GlyphClock)) + strings.Repeat(" ", gutter-1)
			}

			lines[i] = prefix + line
		}

		parts = append(parts, strings.Join(lines, "\n"))
	}

	return strings.Join(parts, "\n\n")
}

// DetailSection groups related facts, conditions, events and explanations.
// Title and Paragraphs are escaped data; Fields retains Pair's presence states.
// Body is an optional pure styled renderer appended after typed content. It is
// called with the section's actual width. FullWidth keeps this section out of a
// paired row. Sections never fold or suppress supplied content.
//
// ID identifies the section to a containing inspector. Standalone rendering does
// not validate IDs or own navigation; the inspector validates unique, safe IDs.
type DetailSection struct {
	ID, Title  string
	Fields     []Pair
	Statuses   []StatusItem
	Timeline   []TimelineEvent
	Paragraphs []string
	FullWidth  bool
	Body       func(width int) string
}

// DetailDocument composes an optional identity header and summary with ordered
// sections. Adjacent sections share a row only when each receives at least
// MinSectionWidth cells (40 by default). FullWidth sections break the pair.
// There are at most two columns, preserving a predictable reading order.
type DetailDocument struct {
	Header          IdentityHeader
	Summary         RecordSummary
	Sections        []DetailSection
	MinSectionWidth int
}

// DetailSectionRange identifies zero-based rows in a complete document. End is
// exclusive. Sections paired in one row share its complete range in declaration
// order; a parent can anchor the row to its first section while reflowing.
type DetailSectionRange struct {
	ID         string
	Start, End int
}

// DetailRendering includes unbounded content and section locations. Callers own
// viewport dimensions, scrolling and any retained section identity.
type DetailRendering struct {
	Content  string
	Sections []DetailSectionRange
}

// View is pure except for caller-supplied Body renderers, which must also be pure.
// It preserves complete section content; callers allocate or scroll its height.
func (d DetailDocument) View(width int, t theme.Theme) string {
	return d.Render(width, t).Content
}

// Render returns the same content as View with the section rows needed to retain
// a reader's place across replacement and responsive reflow. It does not mutate
// the document, infer state, or retain caches of caller data.
func (d DetailDocument) Render(width int, t theme.Theme) DetailRendering {
	if width <= 0 {
		return DetailRendering{}
	}

	parts := make([]string, 0, len(d.Sections)+2)
	height := 0
	appendBlock := func(block string) int {
		if len(parts) > 0 {
			height++
		}

		start := height

		parts = append(parts, block)
		height += lipgloss.Height(block)

		return start
	}

	if d.Header != (IdentityHeader{}) {
		header := d.Header
		if width == 1 {
			header.Kind = textRowSingleCell(header.Kind)
			header.ID = textRowSingleCell(header.ID)
			header.Context = textRowSingleCell(header.Context)
			header.Actions = textRowSingleCell(header.Actions)
		}

		appendBlock(recordWrap(header.View(width, t), width))
	}

	if summary := d.Summary.View(width, t, true); summary != "" {
		appendBlock(summary)
	}

	minimum := d.MinSectionWidth
	if minimum <= 0 {
		minimum = 40
	}

	ranges := make([]DetailSectionRange, 0, len(d.Sections))
	for i := 0; i < len(d.Sections); i++ {
		section := d.Sections[i]
		paired := i+1 < len(d.Sections) && !section.FullWidth && !d.Sections[i+1].FullWidth && minimum <= (width-theme.SpaceLarge)/2

		if !paired {
			body := detailSectionView(section, width, t)
			start := appendBlock(body)
			ranges = append(ranges, DetailSectionRange{ID: section.ID, Start: start, End: height})

			continue
		}

		left := (width - theme.SpaceLarge) / 2
		right := width - theme.SpaceLarge - left
		body := Row(detailSectionView(section, left, t), strings.Repeat(" ", theme.SpaceLarge), detailSectionView(d.Sections[i+1], right, t))
		start := appendBlock(body)
		ranges = append(ranges,
			DetailSectionRange{ID: section.ID, Start: start, End: height},
			DetailSectionRange{ID: d.Sections[i+1].ID, Start: start, End: height},
		)
		i++
	}

	return DetailRendering{Content: strings.Join(parts, "\n\n"), Sections: ranges}
}

func detailSectionView(section DetailSection, width int, t theme.Theme) string {
	parts := make([]string, 0, len(section.Paragraphs)+5)
	if section.Title != "" {
		title := t.Text().Bold(true).Render(text.Inline(section.Title))
		if remaining := width - ansi.StringWidth(title) - theme.SpaceMedium; remaining > 0 {
			title += strings.Repeat(" ", theme.SpaceMedium) + t.Subtle().Render(strings.Repeat(t.FrameBorder().Top, remaining))
		}

		parts = append(parts, recordWrap(title, width))
	}

	groups := make([]string, 0, len(section.Paragraphs)+4)
	if len(section.Fields) > 0 {
		groups = append(groups, recordWrap(KeyValue(section.Fields, width, t), width))
	}

	if len(section.Statuses) > 0 {
		groups = append(groups, StatusList(section.Statuses, width, t))
	}

	if len(section.Timeline) > 0 {
		groups = append(groups, Timeline(section.Timeline, width, t))
	}

	for _, paragraph := range section.Paragraphs {
		if paragraph != "" {
			groups = append(groups, recordWrap(text.Escape(paragraph), width))
		}
	}

	if section.Body != nil {
		if body := section.Body(width); body != "" {
			groups = append(groups, recordWrap(body, width))
		}
	}

	if len(groups) > 0 {
		parts = append(parts, strings.Join(groups, "\n\n"))
	}

	if len(parts) == 0 {
		return strings.Repeat(" ", width)
	}

	return strings.Join(parts, "\n")
}

func detailIndent(content string, width int) string {
	gutter := min(theme.SpaceMedium, max(0, width-1))
	lines := strings.Split(recordWrap(content, width-gutter), "\n")

	for i, line := range lines {
		lines[i] = strings.Repeat(" ", gutter) + line
	}

	return strings.Join(lines, "\n")
}
