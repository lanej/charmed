package pattern_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestRequestAndResponseAreIndependentColumns(t *testing.T) {
	requests, logs := story.LogFixture()
	requests = requests[1:2]
	requests[0].RequestBody = collection.HTTPBody{State: collection.Captured, ContentType: "text/plain", Bytes: []byte(strings.Repeat("Request payload\n", 100))}
	requests[0].ResponseBody = collection.HTTPBody{State: collection.Captured, ContentType: "text/plain", Bytes: []byte(strings.Repeat("Response payload\n", 100))}

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: requests, Logs: logs, Theme: theme.New(theme.Dark), NoHelp: true})
	if err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 160, Height: 44})

	if m.Response.Width() <= m.Request.Width() || m.Response.Height() != m.Request.Height() {
		t.Fatal("request and response did not get full-height columns with response priority")
	}

	if !strings.Contains(ansi.Strip(m.Request.View()), "Request payload") || strings.Contains(ansi.Strip(m.Response.View()), "Request payload") {
		t.Fatal("directional payloads were mixed")
	}

	m.Request.KeyMap.Down = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "down"))
	m.FocusPane("request")
	m.Update(tea.KeyPressMsg{Code: 'x'})

	if m.Request.YOffset() != 1 || m.Response.YOffset() != 0 {
		t.Fatal("scrolling request moved response")
	}

	m.FocusPane("response")
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})

	if m.Response.YOffset() != 1 || m.Request.YOffset() != 1 {
		t.Fatal("response did not scroll independently")
	}

	before := m.Response.Width()
	m.SetPaneCollapsed("request", true)

	if m.Response.Width() <= before || m.Request.YOffset() != 1 {
		t.Fatal("request folding failed to return width or preserve offset")
	}

	m.SetPaneCollapsed("request", false)

	for _, width := range []int{60, 160} {
		m.Resize(layout.Size{Width: width, Height: 44})

		if m.Request.YOffset() != 1 || m.Response.YOffset() != 1 {
			t.Fatal("responsive reflow lost independent scroll state")
		}
	}

	m.FocusPane("request")
	m.Update(tea.KeyPressMsg{Code: 'x'})

	if m.Request.YOffset() != 2 {
		t.Fatal("resize replaced native request key map")
	}
}

func TestHiddenInstructionsKeepPayloadsAndControls(t *testing.T) {
	requests, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: requests, Logs: logs, Theme: theme.New(theme.Mono), NoHelp: true})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("http-2")
	m.Resize(layout.Size{Width: 160, Height: 44})

	for _, pane := range []string{"requests", "request", "response", "logs"} {
		m.FocusPane(pane)

		view := ansi.Strip(m.View())
		for _, hint := range []string{"Tab panes", "views", "commands", "scroll ·", "sections ·", "fold pane", "? help"} {
			if strings.Contains(view, hint) {
				t.Fatalf("hidden footer still contains %q", hint)
			}
		}
	}

	m.Update(tea.KeyPressMsg{Code: '?'})

	if !m.HelpVisible() || !strings.Contains(ansi.Strip(m.View()), "Tab panes") {
		t.Fatal("hidden instructions disabled the restore binding")
	}

	m.SetHelpVisible(false)
	m.OpenCommands()

	if m.Palette.HelpVisible() || strings.Contains(ansi.Strip(m.Palette.View()), "cancel") {
		t.Fatal("palette retained keyboard instructions")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if m.Palette.IsOpen() {
		t.Fatal("hiding instructions disabled palette cancellation")
	}
}

func TestPayloadFoldKeysRestoreFocusedComponent(t *testing.T) {
	requests, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: requests, Logs: logs, Theme: theme.New(theme.Mono)})
	if err != nil {
		t.Fatal(err)
	}

	for _, pane := range []struct {
		id   string
		code rune
	}{{"request", '1'}, {"response", '2'}} {
		m.FocusPane(pane.id)
		m.Update(tea.KeyPressMsg{Code: pane.code})

		if !m.PaneCollapsed(pane.id) {
			t.Fatalf("%s was not folded", pane.id)
		}

		m.Update(tea.KeyPressMsg{Code: pane.code})

		if m.PaneCollapsed(pane.id) {
			t.Fatalf("%s could not be reopened with its fold binding", pane.id)
		}
	}
}

func TestInspectorSymbolsRefreshSharedSummary(t *testing.T) {
	requests, _ := story.LogFixture()
	requests = requests[:1]
	requests[0].Completed, requests[0].StatusCode = false, 0

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: requests, Theme: theme.New(theme.Mono)})
	if err != nil {
		t.Fatal(err)
	}

	m.Resize(layout.Size{Width: 80, Height: 40})
	m.FocusPane("response")
	m.SetSymbols(text.LabelOnly)

	if !strings.Contains(ansi.Strip(m.View()), "PENDING") {
		t.Fatal("symbol change did not refresh the shared summary")
	}

	m.SetSymbols(text.GlyphOnly)

	if strings.Contains(ansi.Strip(m.View()), "PENDING") {
		t.Fatal("summary retained stale pending label")
	}
}

func TestRenderingCollapsedPayloadRetainsNativeViewport(t *testing.T) {
	for _, id := range []string{"request", "response"} {
		t.Run(id, func(t *testing.T) {
			requests, logs := story.LogFixture()
			requests = requests[1:2]
			requests[0].RequestBody = collection.HTTPBody{State: collection.Captured, ContentType: "text/plain", Bytes: []byte(strings.Repeat("Request payload\n", 100))}
			requests[0].ResponseBody = collection.HTTPBody{State: collection.Captured, ContentType: "text/plain", Bytes: []byte(strings.Repeat("Response payload\n", 100))}

			m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Requests: requests, Logs: logs, Theme: theme.New(theme.Dark), NoHelp: true})
			if err != nil {
				t.Fatal(err)
			}

			m.Resize(layout.Size{Width: 160, Height: 44})

			native := &m.Request
			if id == "response" {
				native = &m.Response
			}

			native.SetYOffset(6)
			width, height, offset := native.Width(), native.Height(), native.YOffset()

			m.SetPaneCollapsed(id, true)

			for _, columns := range []int{60, 160, 40, 160} {
				m.Resize(layout.Size{Width: columns, Height: 44})
				_ = m.View()

				if native.Width() != width || native.Height() != height || native.YOffset() != offset {
					t.Fatalf("rendering folded %s at %d columns changed viewport from %dx%d offset %d to %dx%d offset %d", id, columns, width, height, offset, native.Width(), native.Height(), native.YOffset())
				}
			}

			m.SetPaneCollapsed(id, false)

			if native.YOffset() != offset {
				t.Fatal("reopening payload lost its retained offset")
			}
		})
	}
}
