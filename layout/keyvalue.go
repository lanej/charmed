package layout

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Pair is one label and value. State separates a missing, redacted or invalid value from an empty string: a
// present empty value renders as "" and never as a blank.
type Pair struct {
	Label, Value string
	State        text.ValueState
}

// KeyValue aligns labels in a column with their values beside them. At 55 columns and narrower each pair stacks as
// "label: value". Values wrap within the width that is left and are never truncated.
func KeyValue(pairs []Pair, width int, t theme.Theme) string {
	stacked, labelWidth := StackedAt(width), 0

	// The label column takes at most half the allocation, so the value always keeps room.
	labelMax := max(1, (width-2)/2)
	labels := make([]string, len(pairs))

	for i, p := range pairs {
		labels[i] = fitLabel(text.Inline(p.Label), labelMax, t)
		labelWidth = max(labelWidth, ansi.StringWidth(labels[i]))
	}

	rows := make([]string, 0, len(pairs))

	for i, p := range pairs {
		label, value := labels[i], p.State.View(p.Value, text.Normal, t)

		if stacked {
			head := label + ": "
			rows = append(rows, t.Subtle().Render(head)+ansi.Hardwrap(value, max(width-ansi.StringWidth(head), 1), true))

			continue
		}

		gap := strings.Repeat(" ", labelWidth-ansi.StringWidth(label)+2)
		rows = append(rows, Row(t.Subtle().Render(label+gap), ansi.Hardwrap(value, max(width-labelWidth-2, 1), true)))
	}

	return trimColumn(rows)
}

// trimColumn stacks blocks and drops the padding that joining adds after short lines.
func trimColumn(blocks []string) string {
	lines := strings.Split(Column(blocks...), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " ")
	}

	return strings.Join(lines, "\n")
}

// fitLabel shortens a label to width cells with an ellipsis, so a cut label is never mistaken for a full one.
func fitLabel(label string, width int, t theme.Theme) string {
	if ansi.StringWidth(label) <= width {
		return label
	}

	return ansi.Truncate(label, width, t.Glyph(theme.GlyphTruncated))
}
