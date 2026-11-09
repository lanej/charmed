package pattern_test

import (
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestInspectorVimPresetAndNativeOverridesSurviveUpdates(t *testing.T) {
	requests, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs, KeyLayout: interaction.VimKeys, BackdropMode: layout.BackdropNone})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("http-2")
	m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})

	if m.HTTP.SelectedID() != "http-3" || m.Palette.KeyLayout() != interaction.VimKeys || m.Palette.BackdropMode() != layout.BackdropNone {
		t.Fatal("composed preset not applied")
	}

	for _, binding := range []key.Binding{m.HTTP.Details.KeyMap.Down, m.Logs.Model.KeyMap.Down, m.Logs.Details.KeyMap.Down} {
		if !key.Matches(tea.KeyPressMsg{Code: 'j'}, binding) {
			t.Fatal("detail/tree navigation missing")
		}
	}

	m.HTTP.Table.KeyMap.LineDown = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "down"))
	m.Resize(layout.Size{Width: 120, Height: 38})

	if upsertErr := m.UpsertRequests(requests[2]); upsertErr != nil {
		t.Fatal(upsertErr)
	}

	m.Update(tea.KeyPressMsg{Code: 'x'})

	if m.HTTP.SelectedID() != "http-4" {
		t.Fatal("update replaced native movement override")
	}

	m.OpenCommands()
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if m.Palette.Mode() != collection.PaletteNavigate {
		t.Fatal("inspector did not route Esc to modal palette")
	}

	if _, layoutErr := m.SetKeyLayout("unknown"); layoutErr == nil || m.Palette.Mode() != collection.PaletteNavigate {
		t.Fatal("invalid preset mutated mode")
	}

	cmd, err := m.SetKeyLayout(interaction.StandardKeys)
	if err != nil || cmd == nil || m.Palette.Mode() != collection.PaletteSearch {
		t.Fatal("runtime preset did not restore search focus")
	}
}

func TestInspectorPaletteIsolatesInputAndRestoresFocus(t *testing.T) {
	requests, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("http-2")
	m.FocusPane("logs")
	selected := m.Logs.Selected()
	m.Palette.Keys.Open = key.NewBinding(key.WithKeys("f4"), key.WithHelp("F4", "commands"))
	m.Update(tea.KeyPressMsg{Code: tea.KeyF4})

	if !m.Palette.IsOpen() {
		t.Fatal("custom opener ignored")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Update(tea.KeyPressMsg{Code: 'F'})
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})

	if m.FocusedPane() != "logs" || m.Logs.Selected() != selected {
		t.Fatal("palette leaked input")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if m.Palette.IsOpen() || m.FocusedPane() != "logs" || m.HTTP.SelectedID() != "http-2" {
		t.Fatal("dismissal lost state")
	}

	m.OpenCommands()
	m.Palette.SetQuery("Layout: columns")

	cmd := m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("no action returned")
	}

	m.Update(cmd())

	if m.Palette.IsOpen() || m.Layout() != pattern.InspectorColumns || m.FocusedPane() != "logs" || m.HTTP.SelectedID() != "http-2" {
		t.Fatal("layout command lost state")
	}

	m.OpenCommands()
	m.Palette.SetQuery("Focus: HTTP details")

	cmd = m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("no pane action")
	}

	m.Update(cmd())

	if m.FocusedPane() != "http" {
		t.Fatal("intentional focus command ignored")
	}
}

func TestInspectorPaletteKeepsObservationsLiveAndRechecksAvailability(t *testing.T) {
	requests, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("http-2")
	m.OpenCommands()

	if m.Init() == nil {
		t.Fatal("initially open inspector palette lost focus command")
	}

	m.Palette.SetQuery("Show raw body")

	requests[1].RequestBody = collection.HTTPBody{}

	requests[1].ResponseBody = collection.HTTPBody{}
	if err := m.UpsertRequests(requests[1]); err != nil {
		t.Fatal(err)
	}

	if m.Update(tea.KeyPressMsg{Code: tea.KeyEnter}) != nil || !m.Palette.IsOpen() || m.HTTP.Raw {
		t.Fatal("stale enabled command executed")
	}

	for _, size := range []layout.Size{{Width: 120, Height: 40}, {Width: 40, Height: 16}, {Width: 1, Height: 1}, {}} {
		m.Resize(size)

		if !m.Palette.IsOpen() || m.Palette.Search.Value() != "Show raw body" {
			t.Fatal("resize reset palette")
		}

		_ = m.View()
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	m.Update(collection.CommandSelectedMsg{ID: "scope/request"})

	if m.Scope() != pattern.RequestScope {
		t.Fatal("request scope command invalid")
	}

	m.Update(collection.CommandSelectedMsg{ID: "scope/trace"})

	if m.Scope() != pattern.TraceScope {
		t.Fatal("trace scope command invalid")
	}

	m.SelectRequest("http-4")

	for _, command := range m.Commands() {
		if command.ID == "section/parameters" && command.DisabledReason == "" {
			t.Fatal("parameter command enabled without parameters")
		}
	}

	m.Update(collection.CommandSelectedMsg{ID: "url/full"})

	if m.HTTP.URLFormat != "full" {
		t.Fatal("URL command ignored")
	}
}
