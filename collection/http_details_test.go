package collection_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestHTTPDetailsUseHorizontalSpaceAndStackWhenNarrow(t *testing.T) {
	m := collection.NewHTTPLog(120, 60, theme.New(theme.Dark))

	record := collection.HTTPExchange{ID: "request", URL: "https://api.example/labels?expand=rate&expand=tracking", RequestHeaders: http.Header{"Content-Type": {"application/json"}}, RequestBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"parcel":"p_42"}`)}, ResponseBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"ok":true}`)}}
	if err := m.SetEntries([]collection.HTTPExchange{record}); err != nil {
		t.Fatal(err)
	}

	for _, width := range []int{120, 69, 68, 40, 120} {
		m.ResizeDetails(width, 60)

		frame := ansi.Strip(m.DetailView())
		if strings.Contains(frame, "?expand=") || !strings.Contains(frame, `expand: "rate"`) || !strings.Contains(frame, `expand: "tracking"`) {
			t.Fatal("default target duplicates parameters or hides values", frame)
		}

		sideBySide := false

		for _, line := range strings.Split(frame, "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatal("details exceed width", width, line)
			}

			if strings.Contains(line, "Parameters") && strings.Contains(line, "Response") {
				sideBySide = true
			}
		}

		if sideBySide != (width >= 69) {
			t.Fatal("responsive sections did not follow available content width", width, frame)
		}

		if !strings.Contains(frame, `"parcel": "p_42"`) || !strings.Contains(frame, `"ok": true`) || m.SelectedID() != "request" {
			t.Fatal("resize lost captured body or selection", frame)
		}
	}

	if strings.Contains(ansi.Strip(m.ListView()), "?expand=") {
		t.Fatal("default request list repeats query")
	}

	m.SetURLFormat(text.FullURL)

	if !strings.Contains(ansi.Strip(m.DetailView()), record.URL) {
		t.Fatal("explicit full URL format was ignored")
	}

	selected, _ := m.Selected()
	if selected.URL != record.URL {
		t.Fatal("display formatting altered capture")
	}
}

func TestHTTPDetailsKeepUUIDHeadersAndResponseLinesIntact(t *testing.T) {
	const (
		id      = "01234567-89ab-cdef-0123-456789abcdef"
		message = "Could not purchase shipment: the carrier returned an internal server error"
	)

	m := collection.NewHTTPLog(105, 60, theme.New(theme.Dark))

	record := collection.HTTPExchange{ID: "request", RequestID: id, RequestHeaders: http.Header{"X-Ep-Request-Uuid": {id}}, ResponseHeaders: http.Header{"X-Ep-Request-Uuid": {id}}, RequestBody: collection.HTTPBody{State: collection.Captured, Bytes: []byte(`{"rate":"rate_fixture"}`), ContentType: "application/json"}, ResponseBody: collection.HTTPBody{State: collection.Captured, Bytes: []byte(`{"message":"` + message + `"}`), ContentType: "application/json"}}
	if err := m.SetEntries([]collection.HTTPExchange{record}); err != nil {
		t.Fatal(err)
	}

	for _, width := range []int{105, 180, 105} {
		m.ResizeDetails(width, 60)

		view := ansi.Strip(m.DetailView())
		if strings.Count(view, "X-Ep-Request-Uuid: "+id) != 2 || !strings.Contains(view, message) {
			t.Fatalf("width %d fragmented captured evidence:\n%s", width, view)
		}
	}

	m.RequestOpen = false
	m.ResizeDetails(105, 60)

	if view := ansi.Strip(m.DetailView()); !strings.Contains(view, message) || strings.Contains(view, "rate_fixture") {
		t.Fatalf("folded request still occupies response space:\n%s", view)
	}
}

func TestRepeatedHTTPLabelsKeepPerRecordFormatting(t *testing.T) {
	m := collection.NewHTTPLog(100, 30, theme.New(theme.Mono))
	calls := 0

	m.SetURLFormatter(func(string) string { calls++; return fmt.Sprintf("target-%d", calls) })

	records := []collection.HTTPExchange{{ID: "one", Method: "GET", URL: "https://example.test/", StatusCode: 200, Completed: true}, {ID: "two", Method: "GET", URL: "https://example.test/", StatusCode: 200, Completed: true}, {ID: "three", Method: "GET", URL: "https://example.test/", StatusCode: 200, Completed: true}}
	if err := m.SetEntries(records); err != nil {
		t.Fatal(err)
	}

	rows := m.Table.Rows()
	if len(rows) != 3 || rows[0][len(rows[0])-1] == rows[1][len(rows[1])-1] {
		t.Fatal("record formatters were memoized")
	}

	m.Theme = theme.New(theme.Light)
	m.ResizeList(100, 12)

	if len(m.Table.Rows()) != 3 || calls < 6 {
		t.Fatal("refresh lost formatter or row content")
	}
}
