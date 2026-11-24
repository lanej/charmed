package pattern_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func structuredRecord() pattern.InspectorRecord {
	return pattern.InspectorRecord{
		ID: "operation",
		Summary: &layout.RecordSummary{
			Title: "Gateway update", Status: "Outcome unknown",
			Fields: []layout.Pair{{Label: "Owner", Value: "network"}},
			Next:   &layout.Pair{Label: "Next", Value: "Inspect the saved operation"},
		},
		Document: &layout.DetailDocument{
			Header: layout.IdentityHeader{Kind: "Operation", ID: "op-42"},
			Summary: layout.RecordSummary{
				Status: "Outcome unknown",
				Fields: []layout.Pair{{Label: "Observed", Value: "10:24"}},
				Next:   &layout.Pair{Label: "Next", Value: "Reconcile with the source"},
			},
			Sections: []layout.DetailSection{
				{ID: "identity", Title: "Identity", Fields: []layout.Pair{{Label: "Target", Value: "edge-west"}}},
				{ID: "evidence", Title: "Evidence", Statuses: []layout.StatusItem{{Label: "Receipt", Value: "Unconfirmed", Detail: "Timeout after acceptance"}}},
				{
					ID: "history", Title: "History", FullWidth: true,
					Timeline:   []layout.TimelineEvent{{When: "10:23", Actor: "controller", Title: "Request accepted", Detail: "Acceptance does not confirm the effect"}},
					Paragraphs: []string{strings.Repeat("Retained source evidence. ", 60)},
				},
			},
		},
	}
}

func TestListInspectorCopiesStructuredObservations(t *testing.T) {
	t.Parallel()

	record := structuredRecord()
	record.Document.Sections[1].Stages = []layout.StatusItem{{Label: "Delivery", Value: "Recorded"}}
	m := newInspector(t, []pattern.InspectorRecord{record}, layout.Size{Width: 132, Height: 20})
	list, detail := m.List.GetContent(), m.Detail.GetContent()

	record.Summary.Title = "caller changed title"
	record.Summary.Fields[0].Value = "caller changed owner"
	record.Summary.Next.Value = "caller changed next"
	record.Document.Header.ID = "caller changed identity"
	record.Document.Summary.Fields[0].Value = "caller changed observation"
	record.Document.Summary.Next.Value = "caller changed recovery"
	record.Document.Sections[0].Fields[0].Value = "caller changed target"
	record.Document.Sections[1].Statuses[0].Value = "caller changed receipt"
	record.Document.Sections[1].Stages[0].Value = "caller changed delivery"
	record.Document.Sections[2].Timeline[0].Title = "caller changed event"
	record.Document.Sections[2].Paragraphs[0] = "caller changed evidence"

	m.Refresh()

	if m.List.GetContent() != list || m.Detail.GetContent() != detail {
		t.Fatal("caller mutation changed the retained typed observation")
	}

	if !m.Inspect() || !strings.Contains(ansi.Strip(m.Detail.GetContent()), "Request accepted") {
		t.Fatal("typed document did not supply complete inspection")
	}

	if err := m.SetRecords([]pattern.InspectorRecord{record}); err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(m.Detail.GetContent()), "caller changed target") {
		t.Fatal("explicit replacement did not update the typed observation")
	}
}

func TestListInspectorRejectsAmbiguousSectionIdentitiesAtomically(t *testing.T) {
	t.Parallel()

	m := newInspector(t, []pattern.InspectorRecord{structuredRecord()}, layout.Size{Width: 80, Height: 12})
	m.Inspect()
	m.Detail.SetYOffset(3)
	before, offset := m.View(), m.Detail.YOffset()

	for _, id := range []string{"", " \t", "unsafe\nidentity", "identity"} {
		record := structuredRecord()
		record.Document.Sections[1].ID = id

		if err := m.SetRecords([]pattern.InspectorRecord{record}); err == nil {
			t.Fatalf("invalid section ID %q was accepted", id)
		}

		if m.View() != before || m.Detail.YOffset() != offset || !m.Inspecting() {
			t.Fatal("rejected document changed the retained view")
		}
	}
}

func TestListInspectorStructuredDetailsRetainSectionThroughReflow(t *testing.T) {
	t.Parallel()

	record := structuredRecord()
	m := newInspector(t, []pattern.InspectorRecord{record}, layout.Size{Width: 80, Height: 10})
	m.Inspect()
	m.Detail.SetYOffset(contentRow(t, m.Detail.GetContent(), "Evidence"))

	inserted := layout.DetailSection{
		ID: "inserted", Title: "New context", FullWidth: true,
		Paragraphs: []string{strings.Repeat("Additional observed context. ", 8)},
	}
	record.Document.Sections = append([]layout.DetailSection{inserted}, record.Document.Sections...)

	if err := m.SetRecords([]pattern.InspectorRecord{record}); err != nil {
		t.Fatal(err)
	}

	assertDetailStartsWith(t, m, "Evidence")
	m.Resize(layout.Size{Width: 42, Height: 10})
	assertDetailStartsWith(t, m, "Evidence")

	// Removing the visible section returns to its nearest surviving predecessor.
	record.Document.Sections = append(record.Document.Sections[:2], record.Document.Sections[3:]...)
	if err := m.SetRecords([]pattern.InspectorRecord{record}); err != nil {
		t.Fatal(err)
	}

	assertDetailStartsWith(t, m, "Identity")

	other := structuredRecord()

	other.ID = "other-operation"
	if err := m.SetRecords([]pattern.InspectorRecord{record, other}); err != nil {
		t.Fatal(err)
	}

	m.Select(other.ID)

	if m.Detail.YOffset() != 0 {
		t.Fatal("changing record reused the previous record's section anchor")
	}
}

func TestListInspectorCustomRenderersOverrideTypedPresentation(t *testing.T) {
	t.Parallel()

	record := structuredRecord()
	record.Row = func(int, bool) string { return "Custom row" }
	record.Detail = func(int) string { return "Custom detail" }
	m := newInspector(t, []pattern.InspectorRecord{record}, layout.Size{Width: 132, Height: 12})

	if !strings.Contains(m.List.GetContent(), "Custom row") || strings.Contains(m.List.GetContent(), "Gateway") || m.Detail.GetContent() != "Custom detail" {
		t.Fatal("custom renderers did not override typed presentation")
	}
}

func TestListInspectorRetainsSectionSeparatorOffset(t *testing.T) {
	t.Parallel()

	m := newInspector(t, []pattern.InspectorRecord{structuredRecord()}, layout.Size{Width: 80, Height: 10})
	m.Inspect()
	separator := contentRow(t, m.Detail.GetContent(), "History") - 1
	m.Detail.SetYOffset(separator)
	before := m.Detail.View()
	m.Resize(layout.Size{Width: 80, Height: 10})
	m.Refresh()

	if m.Detail.YOffset() != separator || m.Detail.View() != before {
		t.Fatal("same-content reflow displaced a blank section separator")
	}
}

func TestListInspectorStructuredBodyRunsOnlyWhenPrepared(t *testing.T) {
	t.Parallel()

	calls := 0
	record := structuredRecord()
	record.Document.Sections[2].Body = func(int) string {
		calls++
		return "Custom source appendix"
	}
	m := newInspector(t, []pattern.InspectorRecord{record}, layout.Size{Width: 132, Height: 12})
	m.Inspect()

	prepared := calls

	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	m.View()
	m.View()

	if prepared == 0 || calls != prepared {
		t.Fatal("view or scrolling reran a document body renderer")
	}

	m.Refresh()

	if calls <= prepared {
		t.Fatal("refresh failed to prepare the custom document body")
	}
}

func TestListInspectorStructuredContentFitsTinyNativeAllocations(t *testing.T) {
	t.Parallel()

	for _, width := range []int{1, 2, 4, 24, 40, 80, 132} {
		record := structuredRecord()
		record.Summary.Title = "Gateway 界 e\u0301"
		record.Document.Sections[0].Fields[0].Value = "target 界 e\u0301"
		m := newInspector(t, []pattern.InspectorRecord{record}, layout.Size{Width: width, Height: 12})
		m.Theme = theme.New(theme.Dark)
		m.Inspect()

		for _, line := range strings.Split(m.View(), "\n") {
			if ansi.StringWidth(line) != width {
				t.Fatalf("structured inspection exceeded %d cells: %q", width, line)
			}
		}
	}
}

func contentRow(t *testing.T, content, text string) int {
	t.Helper()

	for row, line := range strings.Split(ansi.Strip(content), "\n") {
		if strings.Contains(line, text) {
			return row
		}
	}

	t.Fatalf("content does not contain %q", text)

	return 0
}

func assertDetailStartsWith(t *testing.T, m *pattern.ListInspector, text string) {
	t.Helper()

	first := strings.Split(ansi.Strip(m.Detail.View()), "\n")[0]
	if !strings.Contains(first, text) {
		t.Fatalf("detail anchor moved away from %q: %q", text, first)
	}
}
