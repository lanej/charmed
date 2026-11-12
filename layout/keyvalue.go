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
	class, labelWidth := WidthClassOf(width), 0

	for _, p := range pairs {
		labelWidth = max(labelWidth, ansi.StringWidth(text.Inline(p.Label)))
	}

	rows := make([]string, 0, len(pairs))

	for _, p := range pairs {
		label, value := text.Inline(p.Label), p.State.View(p.Value, text.Normal, t)

		if class.Stacked() {
			head := label + ": "
			rows = append(rows, t.Subtle().Render(head)+ansi.Wrap(value, max(width-ansi.StringWidth(head), 1), ""))

			continue
		}

		gap := strings.Repeat(" ", labelWidth-ansi.StringWidth(label)+2)
		rows = append(rows, Row(t.Subtle().Render(label+gap), ansi.Wrap(value, max(width-labelWidth-2, 1), "")))
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
