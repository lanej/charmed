package collection_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/theme"
)

func TestHTTPParametersDecodeAndFoldIndependently(t *testing.T) {
	m := collection.NewHTTPLog(80, 60, theme.New(theme.Dark))

	record := collection.HTTPExchange{ID: "request", URL: "/labels?expand=rate&expand=tracking&name=Josh+Lane&empty=&quotes=%22%22&escaped=%26%3D#ignored?fake=1", RequestBody: collection.HTTPBody{State: collection.Captured, ContentType: "application/x-www-form-urlencoded", Bytes: []byte("parcel=p_42&service=express")}}
	if err := m.SetEntries([]collection.HTTPExchange{record}); err != nil {
		t.Fatal(err)
	}

	frame := ansi.Strip(m.DetailView())
	for _, want := range []string{"Parameters", "Query", `expand: "rate"`, `expand: "tracking"`, `name: "Josh Lane"`, `empty: ""`, `quotes: "\"\""`, `escaped: "&="`, "Form", `parcel: "p_42"`, `service: "express"`} {
		if !strings.Contains(frame, want) {
			t.Fatal("missing decoded parameter", want, frame)
		}
	}

	m.UpdateDetails(tea.KeyPressMsg{Code: '3'})

	if m.ParametersOpen || !m.RequestOpen || !m.ResponseOpen || strings.Contains(ansi.Strip(m.DetailView()), `expand: "rate"`) {
		t.Fatal("parameter fold affected another section or retained expanded values")
	}

	m.SetFolding(false)
	m.UpdateDetails(tea.KeyPressMsg{Code: '3'})

	if !m.ParametersOpen || !strings.Contains(ansi.Strip(m.DetailView()), `expand: "tracking"`) {
		t.Fatal("foldless parameters hidden")
	}

	selected, _ := m.Selected()
	if selected.URL != record.URL || string(selected.RequestBody.Bytes) != string(record.RequestBody.Bytes) {
		t.Fatal("parameter display mutated observations")
	}
}

func TestHTTPParametersPreserveMalformedSourceAndCaptureBoundaries(t *testing.T) {
	for _, example := range []struct {
		name, query string
		body        collection.HTTPBody
		want        string
	}{
		{name: "invalid escape", query: "ok=1&bad=%ZZ", want: "ok=1&bad=%ZZ"},
		{name: "semicolon", query: "ok=1;bad=2", want: "ok=1;bad=2"},
		{name: "truncated form", body: collection.HTTPBody{State: collection.Truncated, ContentType: "application/x-www-form-urlencoded", Bytes: []byte("secret=partial")}},
		{name: "redacted form", body: collection.HTTPBody{State: collection.Redacted, ContentType: "application/x-www-form-urlencoded", Bytes: []byte("secret=private")}},
		{name: "oversized form", body: collection.HTTPBody{State: collection.Captured, ContentType: "application/x-www-form-urlencoded", Bytes: []byte("secret=" + strings.Repeat("a", collection.DefaultBodyPreviewBytes))}},
		{name: "unsupported charset", body: collection.HTTPBody{State: collection.Captured, ContentType: "application/x-www-form-urlencoded; charset=latin1", Bytes: []byte("secret=unknown")}},
		{name: "encoded form", body: collection.HTTPBody{State: collection.Captured, ContentType: "application/x-www-form-urlencoded", Encoding: "gzip", Bytes: []byte("secret=compressed")}},
	} {
		t.Run(example.name, func(t *testing.T) {
			m := collection.NewHTTPLog(80, 40, theme.New(theme.Dark))
			if err := m.SetEntries([]collection.HTTPExchange{{ID: "request", URL: "/?" + example.query, RequestBody: example.body}}); err != nil {
				t.Fatal(err)
			}

			frame := ansi.Strip(m.DetailView())
			if example.want != "" {
				if !strings.Contains(frame, "Invalid parameter") || !strings.Contains(frame, example.want) {
					t.Fatal("malformed source lost", frame)
				}
			} else if strings.Contains(frame, "Parameters") {
				t.Fatal("incomplete or unavailable capture decoded as parameters", frame)
			}
		})
	}
}
