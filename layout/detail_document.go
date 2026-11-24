package layout

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// DetailSection groups related facts, conditions, events and explanations.
// Title and Paragraphs are escaped data; Fields retains Pair's presence states.
// Body is an optional pure styled renderer appended after typed content. It is
// called with the section's actual width. FullWidth keeps this section out of a
// paired row. Sections never fold or suppress supplied content.
//
// ID identifies the section to a containing inspector. Standalone rendering does
// not validate IDs or own navigation; the inspector validates unique, safe IDs.
type DetailSection struct {
	ID, Title string
	Fields    []Pair
	Statuses  []StatusItem
	// Stages presents independently supplied observations in declaration order.
	Stages     []StatusItem
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
	Header  IdentityHeader
	Summary RecordSummary
	// ActionDetail explains the supplied next action using plain, escaped text.
	ActionDetail    string
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

	if heading := d.heading(width, t); heading != "" {
		appendBlock(heading)
	}

	if action := d.action(width, t); action != "" {
		appendBlock(action)
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
			title += strings.Repeat(" ", theme.SpaceMedium) + t.BorderStyle(false).Render(strings.Repeat(t.FrameBorder().Top, remaining))
		}

		parts = append(parts, recordWrap(title, width))
	}

	groups := make([]string, 0, len(section.Paragraphs)+4)
	if len(section.Fields) > 0 {
		groups = append(groups, detailFacts(section.Fields, width, t))
	}

	if len(section.Statuses) > 0 {
		groups = append(groups, StatusList(section.Statuses, width, t))
	}

	if len(section.Stages) > 0 {
		groups = append(groups, stageSequence(section.Stages, width, t))
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
