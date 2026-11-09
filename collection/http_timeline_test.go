package collection_test

import (
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestHTTPTimelineSharedAxis(t *testing.T) {
	at := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	m := collection.NewHTTPLog(100, 30, theme.New(theme.Dark))

	records := []collection.HTTPExchange{
		{ID: "a", Time: at, Method: "GET", URL: "/a", Completed: true, DurationKnown: true, Duration: 2 * time.Second},
		{ID: "b", Time: at.Add(time.Second), Method: "POST", URL: "/b", Completed: true, DurationKnown: true, Duration: time.Second},
		{ID: "c", Time: at.Add(time.Second), Method: "GET", URL: "/c"},
		{ID: "d", Time: at.Add(time.Second), Method: "GET", URL: "/d", Completed: true},
	}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	lines := strings.Split(ansi.Strip(m.TimelineView(at.Add(4*time.Second), 100, 10)), "\n")
	firstStart, secondStart := ansi.StringWidth(strings.Split(lines[1], "├")[0]), ansi.StringWidth(strings.Split(lines[2], "├")[0])

	firstStop, secondStop := ansi.StringWidth(strings.Split(lines[1], "┤")[0]), ansi.StringWidth(strings.Split(lines[2], "┤")[0])
	if firstStart >= secondStart || firstStop != secondStop {
		t.Fatalf("shared start/stop positions incorrect: %q", lines)
	}

	if !strings.Contains(lines[3], "▷") || !strings.Contains(lines[4], "?") {
		t.Fatalf("pending and unknown must differ: %q", lines)
	}

	if !strings.Contains(lines[0], "12:00:00.000Z") || !strings.Contains(lines[0], "+4s") {
		t.Fatalf("incorrect clock: %q", lines[0])
	}

	for _, line := range lines {
		if ansi.StringWidth(line) > 100 {
			t.Fatalf("overflow: %q", line)
		}
	}
}

func TestHTTPTimelineUnknownAndBounds(t *testing.T) {
	at := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)

	m := collection.NewHTTPLog(40, 20, theme.New(theme.Mono))
	if err := m.SetEntries([]collection.HTTPExchange{{ID: "a", Time: at, URL: "/a"}, {ID: "b", URL: "/missing"}, {ID: "c", Time: at, Completed: true, DurationKnown: true, Duration: -time.Second}}); err != nil {
		t.Fatal(err)
	}

	view := ansi.Strip(m.TimelineView(time.Time{}, 40, 10))
	if strings.Contains(view, "=") || !strings.Contains(view, "? start unavailable") {
		t.Fatalf("unobserved ends must not be implied: %q", view)
	}

	for _, width := range []int{1, 20, 36, 40, 120} {
		view = m.TimelineView(at.Add(time.Second), width, 10)
		for _, line := range strings.Split(view, "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatalf("width %d overflow: %q", width, line)
			}
		}
	}
}

func TestHTTPTimelineSelectionWindow(t *testing.T) {
	m := collection.NewHTTPLog(80, 18, theme.New(theme.Dark))
	at := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)

	var records []collection.HTTPExchange
	for _, id := range []string{"first", "second", "third", "fourth", "last"} {
		records = append(records, collection.HTTPExchange{ID: id, Time: at, Method: "GET", URL: "/" + id})
	}

	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	for _, id := range []string{"first", "last", "second"} {
		if !m.Select(id) {
			t.Fatalf("select %s failed", id)
		}

		view := ansi.Strip(m.TimelineView(at.Add(time.Second), 80, 3))
		if !strings.Contains(strings.Join(strings.Fields(view), " "), "> GET /"+id) {
			t.Fatalf("selection %s hidden: %q", id, view)
		}

		if len(strings.Split(view, "\n")) != 3 {
			t.Fatalf("incorrect height: %q", view)
		}
	}
}
