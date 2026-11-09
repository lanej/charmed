package collection_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestHTTPGroupingCriteria(t *testing.T) {
	for _, example := range []struct{ url, host, port, endpoint string }{
		{"https://API.example.test/path", "api.example.test", "443", "api.example.test:443"},
		{"http://api.example.test:080/path", "api.example.test", "80", "api.example.test:80"},
		{"https://user@[2001:db8::1]:8443/path", "2001:db8::1", "8443", "[2001:db8::1]:8443"},
		{"/relative/path", "", "", ""},
		{"https://api.example.test:99999/path", "api.example.test", "", "api.example.test"},
		{"%broken", "", "", ""},
	} {
		record := collection.HTTPExchange{Service: "labels", URL: example.url}
		if collection.GroupByService(record) != "labels" || collection.GroupByHost(record) != example.host || collection.GroupByPort(record) != example.port || collection.GroupByEndpoint(record) != example.endpoint {
			t.Fatalf("incorrect grouping for %s", example.url)
		}
	}
}

func TestTimelineGroupingPreservesSourceAndVisualNavigation(t *testing.T) {
	at := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)
	records := []collection.HTTPExchange{
		{ID: "a", Service: "api", Time: at, Method: "GET", URL: "https://api.example.test/first"},
		{ID: "b", Service: "carrier", Time: at, Method: "POST", URL: "https://carrier.example:8443/second"},
		{ID: "c", Service: "api", Time: at, Method: "PATCH", URL: "https://api.example.test/third"},
	}
	m := collection.NewHTTPLog(100, 30, theme.New(theme.Dark))

	m.Timeline = true
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	criteria := []collection.HTTPGroup{collection.GroupByService, collection.GroupByEndpoint}
	m.SetTimelineGrouping(criteria...)
	criteria[0] = nil

	m.Select("a")

	view := ansi.Strip(m.TimelineView(at.Add(time.Second), 100, 15))
	if strings.Index(view, "/third") > strings.Index(view, "carrier · carrier.example:8443") || strings.Count(view, "api · api.example.test:443") != 1 {
		t.Fatalf("group order/criteria lost: %q", view)
	}

	for _, id := range []string{"c", "b"} {
		m.UpdateRequests(tea.KeyPressMsg{Code: tea.KeyDown})

		if m.SelectedID() != id {
			t.Fatalf("expected next visual record %s, got %s", id, m.SelectedID())
		}
	}

	if !reflect.DeepEqual(m.Entries(), records) {
		t.Fatal("grouping mutated captured records")
	}

	for _, height := range []int{2, 3, 4, 5} {
		view = ansi.Strip(m.TimelineView(at.Add(time.Second), 100, height))
		if !strings.Contains(view, "/second") {
			t.Fatalf("selected request hidden at height%d: %q", height, view)
		}
	}

	m.SetTimelineGrouping()
	m.Select("a")
	m.UpdateRequests(tea.KeyPressMsg{Code: tea.KeyDown})

	if m.SelectedID() != "b" {
		t.Fatal("clearing grouping did not restore ingestion order")
	}
}

func TestTimelineCustomGroupingAndEscaping(t *testing.T) {
	m := collection.NewHTTPLog(80, 30, theme.New(theme.Dark))
	at := time.Date(2026, 11, 8, 12, 0, 0, 0, time.UTC)

	records := []collection.HTTPExchange{{ID: "a", Time: at, Service: "a", URL: "bc"}, {ID: "b", Time: at, Service: "ab", URL: "c"}, {ID: "c", Time: at, Service: "\x1b[31m", URL: "/"}}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	m.SetTimelineGrouping(collection.GroupByService, func(record collection.HTTPExchange) string { return record.URL })

	view := ansi.Strip(m.TimelineView(at.Add(time.Second), 80, 20))
	for _, want := range []string{"a · bc", "ab · c", `\u001b[31m`} {
		if !strings.Contains(view, want) {
			t.Fatalf("custom/escaped group missing %q in %q", want, view)
		}
	}
}
