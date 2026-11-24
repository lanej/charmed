package layout

import (
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func (d DetailDocument) heading(width int, t theme.Theme) string {
	lines := []string{}
	header := d.Header

	if header != (IdentityHeader{}) {
		if width == 1 {
			header.Kind = textRowSingleCell(header.Kind)
			header.ID = textRowSingleCell(header.ID)
			header.Context = textRowSingleCell(header.Context)
			header.Actions = textRowSingleCell(header.Actions)
		}

		lines = append(lines, recordWrap(header.View(width, t), width))
	}

	summary := d.Summary
	headline := TextRow{Gap: theme.SpaceMedium}

	if summary.Title != "" {
		headline.Items = append(headline.Items, TextRowItem{Text: t.Text().Bold(true).Render(summary.Title)})
	}

	headline.Items = append(headline.Items, TextRowItem{Text: summary.Status, Flex: true})

	for _, field := range summary.Fields {
		label := text.Inline(field.Label)
		if label != "" {
			label = t.Subtle().Render(label + " ")
		}

		headline.Items = append(headline.Items, TextRowItem{Text: label + field.State.View(field.Value, text.Normal, t)})
	}

	if row := headline.View(width); row != "" {
		lines = append(lines, row)
	}

	if summary.Subtitle != "" {
		lines = append(lines, recordWrap(t.Subtle().Render(summary.Subtitle), width))
	}

	return strings.Join(lines, "\n")
}

func (d DetailDocument) action(width int, t theme.Theme) string {
	next := d.Summary.Next
	if next == nil {
		if d.ActionDetail == "" {
			return ""
		}

		// Keep supplied context readable even in a partial snapshot.
		return recordWrap(t.Subtle().Render(text.Escape(d.ActionDetail)), width)
	}

	gutter := theme.SpaceMedium
	if width < 24 {
		gutter = 0
	}

	bodyWidth := width - gutter
	lines := make([]string, 0, 3)

	if label := text.Inline(next.Label); label != "" {
		lines = append(lines, recordWrap(t.Heading().Render(strings.ToUpper(label)), bodyWidth))
	}

	value := next.State.View(next.Value, text.Bright, t)
	lines = append(lines, recordWrap(value, bodyWidth))

	if d.ActionDetail != "" {
		lines = append(lines, recordWrap(t.Subtle().Render(text.Escape(d.ActionDetail)), bodyWidth))
	}

	edge := ""
	if gutter > 0 {
		edge = t.Heading().Render(t.Glyph(theme.GlyphScrollThumb)) + lipgloss.NewStyle().Background(t.Surface).Render(" ")
	}

	surface := lipgloss.NewStyle().Background(t.Surface)
	rows := strings.Split(strings.Join(lines, "\n"), "\n")

	for i, row := range rows {
		content := strings.TrimRight(row, " ")
		padding := strings.Repeat(" ", bodyWidth-ansi.StringWidth(content))
		rows[i] = edge + surface.Render(content) + surface.Render(padding)
	}

	return strings.Join(rows, "\n")
}
