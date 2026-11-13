package pattern

import (
	"net/http"
	"strings"
	"testing"

	"github.com/lanej/charmed/collection"

	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func TestRetainedInspectorPaneRevisions(t *testing.T) {
	m, err := NewLogInspector(LogInspectorOptions{Theme: theme.New(theme.Dark), NoHelp: true, NoHeadline: true})
	if err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 160, Height: 44})
	m.Request.SetContent("original native content")
	first := m.ViewRetained(InspectorRevisions{})
	m.Request.SetContent("changed native content")

	if m.ViewRetained(InspectorRevisions{}) != first {
		t.Fatal("unchanged revision rebuilt pane")
	}

	revised := m.ViewRetained(InspectorRevisions{Details: 1})
	if !strings.Contains(revised, "changed native content") {
		t.Fatal("details revision did not refresh native view")
	}

	copied := *m
	m.Request.SetContent("immediate native content")

	if !strings.Contains(m.View(), "immediate native content") {
		t.Fatal("immediate view returned retained content")
	}

	if !strings.Contains(copied.ViewRetained(InspectorRevisions{}), "changed native content") {
		t.Fatal("copied model shared mutable cache")
	}
}

func TestRetainedInspectorIndependentPaneWork(t *testing.T) {
	revisions := InspectorRevisions{}
	m := &LogInspector{retainedRevisions: &revisions}
	size := layout.Size{Width: 20, Height: 10}
	calls := map[string]int{}

	draw := func(id string) string {
		return m.paneView(id, size, func(layout.Size) string { calls[id]++; return id })
	}
	for range 2 {
		draw(requestPane)
		draw(httpPane)
		draw(logPane)
	}

	revisions.Details++

	draw(requestPane)
	draw(httpPane)
	draw(logPane)

	if calls[requestPane] != 1 || calls[httpPane] != 2 || calls[logPane] != 1 {
		t.Fatalf("unrelated panes rebuilt: %v", calls)
	}
}

func TestRetainedInspectorReleasesHiddenPanes(t *testing.T) {
	m, err := NewLogInspector(LogInspectorOptions{Theme: theme.New(theme.Dark), NoHelp: true, NoHeadline: true})
	if err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 160, Height: 44})
	m.ViewRetained(InspectorRevisions{})

	if len(m.retainedPanes) < 2 {
		t.Fatal("wide panes not retained")
	}

	m.Resize(layout.Size{Width: 40, Height: 12})
	m.ViewRetained(InspectorRevisions{})

	if len(m.retainedPanes) != 1 {
		t.Fatalf("hidden panes retained: %d", len(m.retainedPanes))
	}
}

func TestInspectorGeometryResizeInvalidatesPreparedHeight(t *testing.T) {
	m, err := NewLogInspector(LogInspectorOptions{Theme: theme.New(theme.Dark), NoHelp: true, NoHeadline: true})
	if err != nil {
		t.Fatal(err)
	}

	narrow := layout.Size{Width: 100, Height: 70}
	m.Resize(narrow)
	m.Resize(layout.Size{Width: 240, Height: 70})

	records := []collection.HTTPExchange{{ID: "new", RequestHeaders: http.Header{"X-Long": {strings.Repeat("large header ", 100)}}}}
	if err = m.SetEntries(records, nil); err != nil {
		t.Fatal(err)
	}

	m.ResizeViewport(narrow)
	got := m.detailRegions
	m.Resize(narrow)

	if got != m.detailRegions {
		t.Fatal("geometry path reused an old request height")
	}

	m.Request.SetContent("native prepared payload")
	m.ResizeViewport(layout.Size{Width: 100, Height: 65})

	if m.Request.GetContent() != "native prepared payload" {
		t.Fatal("height change regenerated native payload")
	}

	m.Resize(layout.Size{Width: 100, Height: 65})

	if m.Request.GetContent() == "native prepared payload" {
		t.Fatal("full resize lost refresh behavior")
	}
}
