package collection_test

import (
	"strings"
	"testing"
	"time"

	"charm.land/bubbles/v2/tree"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/log/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestDetachedLogDetailsSelectionAndHierarchy(t *testing.T) {
	m := collection.NewLogTree(56, 24, theme.New(theme.Dark))

	records := []collection.LogRecord{
		{ID: "one", TraceID: "trace", SpanID: "span", LogEntry: collection.LogEntry{Time: time.Date(2026, 11, 8, 12, 4, 1, 0, time.UTC), Level: log.InfoLevel, Message: "Create label with complete message", Fields: []any{"parcel", map[string]any{"service": "express", "id": "p_42"}, "attempt", 2}}},
		{ID: "two", TraceID: "trace", SpanID: "span", LogEntry: collection.LogEntry{Level: log.ErrorLevel, Message: "Carrier failed", Fields: []any{"reason", "timeout"}}},
	}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	m.Select("one")
	m.SetDetachedDetails(true)

	if m.Selected() != "one" {
		t.Fatal("mode switch lost selected record")
	}

	hierarchy := ansi.Strip(m.Model.View())
	if strings.Contains(hierarchy, "attempt") || strings.Contains(hierarchy, "service") || !strings.Contains(hierarchy, "trace") {
		t.Fatalf("hierarchy polluted: %q", hierarchy)
	}

	details := ansi.Strip(m.Details.View())
	for _, want := range []string{"Create label with complete message", "2026-11-08T12:04:01Z", "express", "p_42", "attempt: 2"} {
		if !strings.Contains(details, want) {
			t.Fatalf("missing %q in %q", want, details)
		}
	}

	m.Select("two")

	if !strings.Contains(ansi.Strip(m.Details.View()), "timeout") || strings.Contains(ansi.Strip(m.Details.View()), "express") {
		t.Fatal("details did not track selection")
	}

	m.SetDetachedDetails(false)

	if m.Selected() != "two" || m.Model.Height() != 24 {
		t.Fatal("inline mode did not restore selection/allocation")
	}

	m.Model.ToggleCurrentNode()

	if !strings.Contains(ansi.Strip(m.Model.View()), "reason") {
		t.Fatal("inline fields not restored")
	}
}

func TestDetachedLogDetailsScrollResizeAndLimits(t *testing.T) {
	m := collection.NewLogTree(40, 12, theme.New(theme.Mono))

	record := collection.LogRecord{ID: "one", LogEntry: collection.LogEntry{Message: strings.Repeat("long message ", 30), Fields: []any{"control\x1b[31m", "value\ntext"}}}
	if err := m.SetEntries([]collection.LogRecord{record}); err != nil {
		t.Fatal(err)
	}

	m.SetDetachedDetails(true)
	m.Update(tea.KeyPressMsg{Code: 'D'})
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})

	if !m.DetailsFocused || m.Details.YOffset() == 0 || m.Follow {
		t.Fatal("detail scroll did not own focus and pause follow")
	}

	offset := m.Details.YOffset()
	if err := m.Upsert(record); err != nil {
		t.Fatal(err)
	}

	if m.Details.YOffset() != offset {
		t.Fatal("same-record update reset scroll")
	}

	for _, size := range [][2]int{{40, 12}, {20, 8}, {1, 1}, {80, 30}} {
		m.Resize(size[0], size[1])

		lines := strings.Split(m.View(), "\n")
		if len(lines) != size[1] {
			t.Fatalf("height overflow %v", size)
		}

		for _, line := range lines {
			if ansi.StringWidth(line) != size[0] {
				t.Fatalf("width overflow %v: %q", size, line)
			}
		}
	}

	m.FieldNodes = 2
	m.FieldDepth = 2
	cycle := map[string]any{}
	cycle["self"] = cycle

	record.Fields = []any{"cycle", cycle}
	if err := m.Upsert(record); err != nil {
		t.Fatal(err)
	}

	if m.Details.TotalLineCount() > 100 {
		t.Fatal("field budget did not bound cyclic fields")
	}

	if err := m.SetEntries(nil); err != nil {
		t.Fatal(err)
	}

	if m.Details.YOffset() != 0 || !strings.Contains(ansi.Strip(m.Details.View()), "Select a log entry") {
		t.Fatal("empty data retained stale details")
	}
}

func TestFocusedLogLeafStaysInHierarchyPosition(t *testing.T) {
	m := collection.NewLogTree(72, 30, theme.New(theme.Dark))

	records := []collection.LogRecord{
		{ID: "first", TraceID: "trace", SpanID: "span", LogEntry: collection.LogEntry{Message: "first message", Fields: []any{"value", 1}}},
		{ID: "second", TraceID: "trace", SpanID: "span", LogEntry: collection.LogEntry{Message: "second message"}},
		{ID: "third", TraceID: "trace", SpanID: "span", LogEntry: collection.LogEntry{Message: "third message"}},
	}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	m.Follow = false
	m.Select("first")
	m.SetDetachedDetails(true)
	view := ansi.Strip(m.View())

	title := strings.Index(view, "Log · first")
	if title < 0 || strings.Index(view, "Span span") > title || strings.Index(view, "second message") < title || strings.Index(view, "third message") < strings.Index(view, "second message") {
		t.Fatalf("focused leaf moved from its position: %q", view)
	}

	if strings.Count(view, "first message") != 1 {
		t.Fatalf("leaf duplicated outside focused container: %q", view)
	}

	var boundary string

	for _, line := range strings.Split(view, "\n") {
		if strings.Contains(line, "Log · first") {
			boundary = line
		}
	}

	border := m.Theme.FrameBorder()
	if !strings.HasPrefix(boundary, border.TopLeft) || !strings.HasSuffix(strings.TrimRight(boundary, " "), border.TopRight) {
		t.Fatalf("boundary must span full container width: %q", boundary)
	}

	m.Resize(72, 60)

	styles := tree.DefaultStyles(true)
	styles.TreeStyle = lipgloss.NewStyle().PaddingTop(1)
	m.Model.SetStyles(styles)
	m.Model.SetCursorCharacter("@")
	m.Select("first")

	view = ansi.Strip(m.View())
	if strings.Count(view, "first message") != 1 || !strings.Contains(view, "Log · first") {
		t.Fatalf("native frame shifted replacement: %q", view)
	}

	styles.ParentNodeStyle = lipgloss.NewStyle().Padding(1, 0)
	styles.SelectedNodeStyle = lipgloss.NewStyle().Padding(1, 0)
	m.Model.SetStyles(styles)
	m.Select("first")

	view = ansi.Strip(m.View())
	if strings.Contains(view, "Log · first") || strings.Count(view, "first message") != 1 {
		t.Fatalf("unsupported native geometry must retain native view: %q", view)
	}

	m.Update(tea.KeyPressMsg{Code: 'D'})

	if m.DetailsFocused {
		t.Fatal("fallback focused hidden details")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})

	if m.Selected() == "first" {
		t.Fatal("fallback navigation did not move visible tree")
	}

	m.Select("first")

	m.FocusedLeafGeometry = func(tree.Model) (int, int, bool) { return 16, 23, true }

	view = ansi.Strip(m.View())
	if !strings.Contains(view, "Log · first") || strings.Count(view, "first message") != 1 || strings.Count(view, "second message") != 1 {
		t.Fatalf("explicit native geometry failed: %q", view)
	}

	m.FocusedLeafGeometry = nil
	m.Model.SetStyles(tree.DefaultStyles(true))

	m.Model.SetYOffset(0)
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if strings.Contains(ansi.Strip(m.View()), "Log · first") || m.DetailsFocused {
		t.Fatal("collapsed parent retained stale focused leaf")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	m.Select("second")

	if !strings.Contains(ansi.Strip(m.View()), "Log · second") {
		t.Fatal("reopened hierarchy did not restore focused leaf")
	}
}

func TestFocusedLogFieldsUseAvailableWidth(t *testing.T) {
	for _, ascii := range []bool{false, true} {
		palette := theme.New(theme.Dark)
		palette.ASCII = ascii
		m := collection.NewLogTree(36, 24, palette)

		record := collection.LogRecord{ID: "one", LogEntry: collection.LogEntry{Message: "Carrier response delayed", Fields: []any{"upstream", map[string]any{"host": "carrier.example", "timeout_ms": 800}}}}
		if err := m.SetEntries([]collection.LogRecord{record}); err != nil {
			t.Fatal(err)
		}

		m.SetDetachedDetails(true)

		if !strings.Contains(ansi.Strip(m.Details.View()), `host: "carrier.example"`) {
			t.Fatalf("host wrapped despite available cells: %q", ansi.Strip(m.Details.View()))
		}

		for _, line := range strings.Split(m.View(), "\n") {
			if ansi.StringWidth(line) != 36 {
				t.Fatalf("field overflow: %q", line)
			}
		}
	}
}
