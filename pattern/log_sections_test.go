package pattern_test

import (
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestSparseInspectorReturnsTimelineSpaceAndOptionalHelp(t *testing.T) {
	requests, logs := story.LogFixture()
	requests = requests[1:2]
	requests[0].TraceID = ""

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest(requests[0].ID)
	m.Resize(layout.Size{Width: 160, Height: 44})

	if m.ActiveLayout() != pattern.InspectorTimeline {
		t.Fatal("default did not adapt to wide terminal")
	}

	view := ansi.Strip(m.View())
	detailRow := -1

	for index, line := range strings.Split(view, "\n") {
		if strings.Contains(line, "Response") {
			detailRow = index
		}
	}

	if detailRow != 8 {
		t.Fatalf("one request should place payload panels on row 8; got %d: %s", detailRow, view)
	}

	before := m.Response.Height()
	m.Update(tea.KeyPressMsg{Code: '?'})

	if m.HelpVisible() || m.Response.Height() != before+2 || strings.Contains(ansi.Strip(m.View()), "panes ·") {
		t.Fatal("hidden instructions did not return their two rows")
	}

	m.Update(tea.KeyPressMsg{Code: '?'})

	if !m.HelpVisible() || m.Response.Height() != before {
		t.Fatal("instructions did not restore")
	}

	for i := range 20 {
		if err := m.UpsertRequests(collection.HTTPExchange{ID: string(rune('a' + i)), Time: time.Date(2026, 11, 8, 12, 5, i, 0, time.UTC)}); err != nil {
			t.Fatal(err)
		}
	}

	if m.Response.Height() >= before || m.HTTP.SelectedID() != requests[0].ID {
		t.Fatal("arriving requests did not grow the bounded timeline or stole selection")
	}
}

func TestInspectorFoldAndHelpCommandsRemainAvailableWithoutInstructions(t *testing.T) {
	requests, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs, NoHelp: true, NoHeadline: true})
	if err != nil {
		t.Fatal(err)
	}

	if m.HelpVisible() || strings.HasPrefix(ansi.Strip(m.View()), "Request req-") {
		t.Fatal("optional chrome was still shown")
	}

	m.Update(collection.CommandSelectedMsg{ID: "fold/logs"})

	if !m.PaneCollapsed("logs") {
		t.Fatal("component fold command did not collapse logs")
	}

	m.Update(collection.CommandSelectedMsg{ID: "fold/logs"})
	m.Update(collection.CommandSelectedMsg{ID: "help/toggle"})
	m.SetLayout(pattern.InspectorColumns)
	m.Update(collection.CommandSelectedMsg{ID: "layout/auto"})

	if m.PaneCollapsed("logs") || !m.HelpVisible() || m.Layout() != pattern.InspectorAuto {
		t.Fatal("restore commands failed")
	}

	m.SetFolding(false)
	m.Update(collection.CommandSelectedMsg{ID: "fold/logs"})

	if m.PaneCollapsed("logs") {
		t.Fatal("disabled component folding accepted a command")
	}
}

func TestInspectorPaneFoldingPreservesStateAndReclaimsSpace(t *testing.T) {
	requests, logs := story.LogFixture()
	requests[1].ResponseBody = collection.HTTPBody{State: collection.Captured, ContentType: "text/plain", Bytes: []byte(strings.Repeat("Retained response line\n", 100))}

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs, Layout: pattern.InspectorTimeline, NoHelp: true})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("http-2")
	m.FocusPane("http")
	m.Resize(layout.Size{Width: 160, Height: 44})
	m.HTTP.Raw = true
	m.Resize(layout.Size{Width: 160, Height: 44})

	beforeWidth := m.Response.Width()
	if !m.SetPaneCollapsed("logs", true) || m.Response.Width() <= beforeWidth {
		t.Fatal("folded logs did not return horizontal space")
	}

	beforeHeight := m.Response.Height()
	m.SetPaneCollapsed("requests", true)

	if m.Response.Height() <= beforeHeight {
		t.Fatal("folded timeline did not return vertical space")
	}

	m.Response.SetYOffset(10)
	m.Update(tea.KeyPressMsg{Code: 'z'})
	m.Update(tea.KeyPressMsg{Code: 'r'})

	if !m.PaneCollapsed("http") || !m.HTTP.Raw {
		t.Fatal("hidden pane handled a body command")
	}

	for _, preset := range []pattern.InspectorLayout{pattern.InspectorSidebar, pattern.InspectorColumns, pattern.InspectorFocused, pattern.InspectorTimeline} {
		m.SetLayout(preset)
		m.Resize(layout.Size{Width: 40, Height: 24})
		m.Resize(layout.Size{Width: 160, Height: 44})

		if !m.PaneCollapsed("logs") || !m.PaneCollapsed("http") || !m.PaneCollapsed("requests") || m.HTTP.SelectedID() != "http-2" || !m.HTTP.Raw {
			t.Fatal("layout changed retained UI state")
		}

		if m.Response.YOffset() != 10 {
			t.Fatal("resizing a folded component lost its scroll position")
		}
	}

	m.SetFolding(false)

	if m.SetPaneCollapsed("logs", true) || m.PaneCollapsed("http") || !m.HTTP.RequestOpen || !m.HTTP.ResponseOpen {
		t.Fatal("disabled folding left content hidden")
	}

	m.SetFolding(true)

	if m.PaneCollapsed("logs") || m.SetPaneCollapsed("missing", true) {
		t.Fatal("fold state or invalid target accepted")
	}
}

func TestFoldedInspectorFitsAllLayouts(t *testing.T) {
	requests, logs := story.LogFixture()
	for _, preset := range []pattern.InspectorLayout{pattern.InspectorAuto, pattern.InspectorSidebar, pattern.InspectorTimeline, pattern.InspectorColumns, pattern.InspectorFocused} {
		m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Mono), Requests: requests, Logs: logs, Layout: preset})
		if err != nil {
			t.Fatal(err)
		}

		for mask := range 16 {
			for i, pane := range []string{"requests", "request", "response", "logs"} {
				m.SetPaneCollapsed(pane, mask&(1<<i) != 0)
			}

			for _, size := range []layout.Size{{Width: 160, Height: 44}, {Width: 100, Height: 30}, {Width: 40, Height: 16}, {Width: 1, Height: 1}, {}} {
				m.Resize(size)

				view := ansi.Strip(m.View())
				if size.Width == 0 {
					if view != "" {
						t.Fatal("zero allocation rendered")
					}

					continue
				}

				lines := strings.Split(view, "\n")
				if len(lines) != size.Height {
					t.Fatalf("%s mask %d height", preset, mask)
				}

				for _, line := range lines {
					if ansi.StringWidth(line) != size.Width {
						t.Fatalf("%s mask %d width: %q", preset, mask, line)
					}
				}
			}
		}
	}
}
