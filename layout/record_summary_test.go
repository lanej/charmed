package layout_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestRecordSummaryPreservesFactsAndSelectedGeometry(t *testing.T) {
	palette := theme.New(theme.Dark)
	summary := layout.RecordSummary{
		Title:    palette.LinkStyle().Hyperlink("https://example.test/operations/42").Render("Operation 42"),
		Subtitle: "Controller request", Status: text.Pill("Outcome unknown", text.Warning, palette),
		Next: &layout.Pair{Label: "Next", Value: "Inspect remote evidence before retrying"},
		Fields: []layout.Pair{
			{Label: "Empty", Value: ""},
			{Label: "Count", Value: "0"},
			{Label: "Enabled", Value: "false"},
			{Label: "Observer", State: text.ValueMissing},
			{Label: "Token", State: text.ValueRedacted},
		},
	}
	fields := append([]layout.Pair(nil), summary.Fields...)
	next := *summary.Next

	for _, width := range []int{1, 2, 20, 40, 80, 132} {
		view := summary.View(width, palette, false)
		if selected := summary.View(width, palette, true); ansi.Strip(view) != ansi.Strip(selected) {
			t.Fatalf("width %d changed content or geometry on selection", width)
		}

		assertRecordBounds(t, view, width)
		compact := compactRecordText(view)

		for _, expected := range []string{"Operation42", "Outcomeunknown", "Inspectremoteevidencebeforeretrying", `""`, "0", "false", "missing", "redacted"} {
			if !strings.Contains(compact, expected) {
				t.Fatalf("width %d lost %q: %q", width, expected, ansi.Strip(view))
			}
		}

		if !strings.Contains(view, "https://example.test/operations/42") {
			t.Fatalf("width %d discarded the native hyperlink", width)
		}
	}

	if !reflect.DeepEqual(fields, summary.Fields) || next != *summary.Next {
		t.Fatal("rendering changed supplied facts")
	}

	if view := ansi.Strip(summary.View(40, palette, false)); !strings.Contains(view, "Next: Inspect") {
		t.Fatalf("next-step label did not share the available line with its value: %q", view)
	}

	if summary.View(0, palette, false) != "" || (layout.RecordSummary{}).View(40, palette, false) != "" {
		t.Fatal("empty allocation or empty summary produced content")
	}
}

func TestRecordSummaryEscapesFactsAndMarksImpossibleGraphemes(t *testing.T) {
	palette := theme.New(theme.Dark)
	summary := layout.RecordSummary{
		Title: "界 é", Next: &layout.Pair{Label: "Next", Value: "literal\x1b[2J"},
		Fields: []layout.Pair{{Label: "Value", Value: "literal\x1b]8;;https://untrusted.test\a"}},
	}
	wide := summary.View(80, palette, false)

	if strings.Contains(wide, "\x1b[2J") || strings.Contains(wide, "\x1b]8;;https://untrusted.test") || !strings.Contains(wide, `\u001b`) {
		t.Fatalf("facts interpreted terminal control data: %q", wide)
	}

	view := summary.View(1, palette, false)
	assertRecordBounds(t, view, 1)

	if !strings.Contains(view, "!") || strings.Contains(view, "界") || !strings.Contains(view, "é") {
		t.Fatalf("single-cell fallback lost or overflowed a grapheme: %q", view)
	}
}

func assertRecordBounds(t *testing.T, view string, width int) {
	t.Helper()

	for _, line := range strings.Split(view, "\n") {
		if ansi.StringWidth(line) > width {
			t.Fatalf("width %d overflowed with %d cells: %q", width, ansi.StringWidth(line), ansi.Strip(line))
		}
	}
}

func compactRecordText(view string) string {
	return strings.Join(strings.Fields(ansi.Strip(view)), "")
}
