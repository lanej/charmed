package collection_test

import (
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

	for _, width := range []int{120, 63, 62, 40, 120} {
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

		if sideBySide != (width >= 63) {
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
