package pattern_test

import (
	"fmt"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestLogInspectorCorrelatesExplicitIDsAndPreservesBrowsing(t *testing.T) {
	requests, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs})
	if err != nil {
		t.Fatal(err)
	}

	if !m.SelectRequest("http-2") || len(m.Logs.Entries()) != 3 {
		t.Fatal("exact request correlation failed")
	}

	if !m.SetScope(pattern.TraceScope) || len(m.Logs.Entries()) != 4 {
		t.Fatal("trace context did not widen")
	}

	if !m.SetScope(pattern.RequestScope) || len(m.Logs.Entries()) != 3 {
		t.Fatal("request scope did not narrow")
	}

	if err := m.UpsertLogs(collection.LogRecord{ID: "same-time", RequestID: "different", TraceID: "different", LogEntry: collection.LogEntry{Time: logs[0].Time, Message: "same text"}}); err != nil {
		t.Fatal(err)
	}

	if len(m.Logs.Entries()) != 3 {
		t.Fatal("correlation guessed from timestamps")
	}

	if !m.SelectRequest("http-4") || !strings.Contains(m.CorrelationMessage(), "unavailable") {
		t.Fatal("missing ID conflated with zero matches")
	}

	if m.SetScope(pattern.TraceScope) {
		t.Fatal("trace action enabled without an ID")
	}

	m.SelectRequest("http-3")

	if !strings.Contains(m.CorrelationMessage(), "No matching logs") {
		t.Fatal("known empty correlation not visible")
	}

	if err := m.UpsertRequests(collection.HTTPExchange{ID: "http-5", RequestID: "req-85"}); err != nil {
		t.Fatal(err)
	}

	if m.HTTP.SelectedID() != "http-3" {
		t.Fatal("new data stole paused request selection")
	}

	before := m.HTTP.SelectedID()
	if err := m.SetEntries([]collection.HTTPExchange{{ID: "other"}}, []collection.LogRecord{{ID: ""}}); err == nil {
		t.Fatal("invalid dataset accepted")
	}

	if m.HTTP.SelectedID() != before {
		t.Fatal("combined validation was not atomic")
	}
}

func TestLogInspectorFitsEveryPaneAndKeepsStateAcrossResize(t *testing.T) {
	requests, logs := story.LogFixture()
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(mode), Requests: requests, Logs: logs})
		if err != nil {
			t.Fatal(err)
		}

		m.SelectRequest("http-2")

		for _, size := range []layout.Size{{Width: 120, Height: 38}, {Width: 80, Height: 24}, {Width: 40, Height: 16}, {Width: 20, Height: 8}, {Width: 1, Height: 1}, {Width: 0, Height: 0}} {
			m.Resize(size)

			for pane := 0; pane < 3; pane++ {
				frame := ansi.Strip(m.View())
				if size.Width <= 0 || size.Height <= 0 {
					if frame != "" {
						t.Fatal("zero allocation rendered")
					}

					continue
				}

				lines := strings.Split(frame, "\n")
				if len(lines) != size.Height {
					t.Fatalf("height %d != %d", len(lines), size.Height)
				}

				for _, line := range lines {
					if ansi.StringWidth(line) != size.Width {
						t.Fatalf("%s width %d != %d", mode, ansi.StringWidth(line), size.Width)
					}
				}

				m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
			}

			if m.HTTP.SelectedID() != "http-2" {
				t.Fatal("resize or pane switch lost selection")
			}
		}
	}
}

func TestInspectorBatchRetentionKeepsOnlyLatestObservations(t *testing.T) {
	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Limit: 2})
	if err != nil {
		t.Fatal(err)
	}

	batch := make([]collection.LogRecord, 100)
	for i := range batch {
		id := fmt.Sprint(i)
		batch[i] = collection.LogRecord{ID: id, RequestID: id}
	}

	if err := m.UpsertLogs(batch...); err != nil {
		t.Fatal(err)
	}

	if err := m.UpsertRequests(collection.HTTPExchange{ID: "old", RequestID: "0"}, collection.HTTPExchange{ID: "latest", RequestID: "99"}); err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("old")

	if len(m.Logs.Entries()) != 0 {
		t.Fatal("evicted log remained correlated")
	}

	m.SelectRequest("latest")

	if len(m.Logs.Entries()) != 1 || m.Logs.Entries()[0].ID != "99" {
		t.Fatal("latest log was not retained")
	}
}
