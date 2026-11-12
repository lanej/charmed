package pattern_test

import (
	"net/http"
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func jsonExchange(id string) collection.HTTPExchange {
	return collection.HTTPExchange{
		ID: id, RequestID: id, Method: "POST", URL: "/buy", Completed: true, StatusCode: 200,
		RequestHeaders: http.Header{"Content-Type": {"application/json"}}, ResponseHeaders: http.Header{"Content-Type": {"application/problem+json"}},
		RequestBody:  collection.HTTPBody{State: collection.Captured, Bytes: []byte(`{"shipment":{"id":"shp_42"},"exact":9007199254740993}`)},
		ResponseBody: collection.HTTPBody{State: collection.Captured, Bytes: []byte(`{"rates":[{"cost":8.500},{"cost":12.50}],"ok":true}`)},
	}
}

func TestInspectorJSONFoldsAreIndependentAndSurviveSelection(t *testing.T) {
	a, b := jsonExchange("a"), jsonExchange("b")

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: []collection.HTTPExchange{a, b}, Theme: theme.New(theme.Mono), NoHelp: true})
	if err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 160, Height: 40})
	m.SelectRequest("a")
	m.FocusPane("response")
	m.Response.KeyMap.Down = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "next"))
	m.Update(tea.KeyPressMsg{Code: 'x'})
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if !m.ResponseJSON.Collapsed(1) || m.RequestJSON.Collapsed(1) || strings.Contains(ansi.Strip(m.Response.View()), "8.500") {
		t.Fatal("response JSON did not fold independently")
	}

	m.Update(tea.KeyPressMsg{Code: 'z'})
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	m.Update(tea.KeyPressMsg{Code: 'z'})

	if !m.ResponseJSON.Collapsed(1) {
		t.Fatal("hidden component handled JSON input")
	}

	m.Update(tea.KeyPressMsg{Code: 'r'})

	if !strings.Contains(ansi.Strip(m.Response.View()), "8.500") || m.ResponseJSON.Interactive() {
		t.Fatal("raw mode lost captured body or retained JSON routing")
	}

	m.Update(tea.KeyPressMsg{Code: 'r'})
	m.Resize(layout.Size{Width: 60, Height: 32})
	m.Resize(layout.Size{Width: 160, Height: 40})
	m.SelectRequest("b")

	if m.ResponseJSON.Collapsed(1) {
		t.Fatal("another exchange inherited folds for identical bytes")
	}

	m.SelectRequest("a")

	if !m.ResponseJSON.Collapsed(1) || m.ResponseJSON.Selected() != 1 {
		t.Fatal("returning to a retained exchange lost its JSON state")
	}

	m.Update(collection.CommandSelectedMsg{ID: "json/toggle"})

	if m.ResponseJSON.Collapsed(1) {
		t.Fatal("palette JSON command failed")
	}

	m.Update(tea.KeyPressMsg{Code: 'x'})

	if m.ResponseJSON.Selected() != 2 {
		t.Fatal("native navigation customization was overwritten")
	}

	if strings.Contains(ansi.Strip(m.View()), "fold JSON") {
		t.Fatal("hidden instructions leaked JSON key hints")
	}
}

func TestInspectorJSONBodyReplacementAndRetention(t *testing.T) {
	a, b := jsonExchange("a"), jsonExchange("b")

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: []collection.HTTPExchange{a, b}, Limit: 2, Theme: theme.New(theme.Mono)})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("a")
	m.RequestJSON.SetCollapsed(1, true)
	m.ResponseJSON.SetCollapsed(1, true)

	a.ResponseBody.Bytes = []byte(`{"replacement":{"value":1}}`)
	if err := m.UpsertRequests(a); err != nil {
		t.Fatal(err)
	}

	if m.ResponseJSON.Collapsed(1) || !m.RequestJSON.Collapsed(1) || m.ResponseJSON.Selected() != 0 {
		t.Fatal("replacing one body failed to reset only its structure")
	}

	m.SelectRequest("b")

	if err := m.UpsertRequests(jsonExchange("c")); err != nil {
		t.Fatal(err)
	}

	if err := m.UpsertRequests(a); err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("a")

	if m.RequestJSON.Collapsed(1) {
		t.Fatal("an evicted exchange retained JSON state")
	}

	m.ResponseJSON.SetCollapsed(1, true)
	m.SetFolding(false)
	m.SetFolding(true)

	if m.ResponseJSON.Collapsed(1) || !strings.Contains(ansi.Strip(m.Response.View()), "value") {
		t.Fatal("disabled folding retained hidden JSON data")
	}
}

func TestInspectorJSONRespectsCapturePolicy(t *testing.T) {
	for _, test := range []struct {
		name, notice string
		body         collection.HTTPBody
	}{
		{name: "truncated", body: collection.HTTPBody{State: collection.Truncated, ContentType: "application/json", Bytes: []byte(`{"secret":1}`)}, notice: "capture truncated"},
		{name: "redacted", body: collection.HTTPBody{State: collection.Redacted, ContentType: "application/json", Bytes: []byte(`{"secret":1}`)}, notice: "redacted"},
		{name: "invalid", body: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Bytes: []byte(`{"a":`)}, notice: "invalid JSON"},
		{name: "encoding", body: collection.HTTPBody{State: collection.Captured, ContentType: "application/json", Encoding: "gzip", Bytes: []byte(`{"secret":1}`)}, notice: "encoding gzip"},
		{name: "charset", body: collection.HTTPBody{State: collection.Captured, ContentType: "application/json; charset=iso-8859-1", Bytes: []byte(`{"secret":1}`)}, notice: "unsupported charset"},
	} {
		t.Run(test.name, func(t *testing.T) {
			a := jsonExchange("a")
			a.ResponseBody = test.body

			m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: []collection.HTTPExchange{a}, Theme: theme.New(theme.Mono)})
			if err != nil {
				t.Fatal(err)
			}

			m.Resize(layout.Size{Width: 160, Height: 40})

			if m.ResponseJSON.Interactive() || !strings.Contains(ansi.Strip(m.Response.View()), test.notice) {
				t.Fatalf("capture policy lost: %s", ansi.Strip(m.Response.View()))
			}

			if test.name == "redacted" && strings.Contains(ansi.Strip(m.View()), "secret") {
				t.Fatal("redacted payload exposed captured bytes")
			}
		})
	}
}
