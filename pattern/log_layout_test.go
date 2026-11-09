package pattern_test

import (
	"strings"
	"testing"
	"time"

	binding "charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestInspectorLayoutsPreserveCorrelationAndNativeState(t *testing.T) {
	_, logs := story.LogFixture()

	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: story.TimelineFixture(), Logs: logs, DetachedLogDetails: true, TimelineGroups: []collection.HTTPGroup{collection.GroupByService, collection.GroupByEndpoint}, TimelineNow: time.Date(2026, 11, 8, 12, 4, 4, 0, time.UTC)})
	if err != nil {
		t.Fatal(err)
	}

	m.SelectRequest("http-2")
	m.SetScope(pattern.TraceScope)
	m.Logs.Follow = false
	m.Logs.Select("log-3")
	m.HTTP.Raw = true

	m.HTTP.Table.KeyMap.LineDown = binding.NewBinding(binding.WithKeys("x"))
	for _, preset := range []pattern.InspectorLayout{pattern.InspectorSidebar, pattern.InspectorTimeline, pattern.InspectorColumns, pattern.InspectorFocused, pattern.InspectorSidebar} {
		if !m.SetLayout(preset) {
			t.Fatal("valid preset rejected", preset)
		}

		m.Resize(layout.Size{Width: 160, Height: 44})

		if m.Layout() != preset || m.ActiveLayout() != preset || m.HTTP.SelectedID() != "http-2" || m.Logs.Selected() != "log-3" || len(m.Logs.Entries()) != 4 || m.Scope() != pattern.TraceScope || !m.HTTP.Raw {
			t.Fatalf("layout %s lost correlated/native state", preset)
		}

		if m.HTTP.Timeline != (preset == pattern.InspectorTimeline) {
			t.Fatal("request presentation disagrees with layout")
		}

		view := ansi.Strip(m.View())
		if preset != pattern.InspectorFocused && (!strings.Contains(view, "HTTP details") || !strings.Contains(view, "Related logs")) {
			t.Fatalf("layout %s dropped a pane", preset)
		}
	}

	before := m.Layout()
	if m.SetLayout(pattern.InspectorLayout("invalid")) || m.Layout() != before {
		t.Fatal("invalid preset changed layout")
	}

	m.FocusPane("requests")
	m.Update(tea.KeyPressMsg{Code: 'x'})

	if m.HTTP.SelectedID() != "http-3" || !strings.Contains(m.CorrelationMessage(), "No matching logs") {
		t.Fatal("native binding or correlation stopped after layout switch")
	}
}

func TestInspectorLayoutsFitAndAdapt(t *testing.T) {
	requests, logs := story.LogFixture()
	for _, preset := range []pattern.InspectorLayout{pattern.InspectorAuto, pattern.InspectorSidebar, pattern.InspectorTimeline, pattern.InspectorColumns, pattern.InspectorFocused} {
		m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark), Requests: requests, Logs: logs, Layout: preset, DetachedLogDetails: true})
		if err != nil {
			t.Fatal(err)
		}

		m.SelectRequest("http-2")

		for _, size := range []layout.Size{{Width: 160, Height: 44}, {Width: 120, Height: 38}, {Width: 80, Height: 24}, {Width: 40, Height: 16}, {Width: 1, Height: 1}, {}} {
			m.Resize(size)

			for _, pane := range []string{"requests", "http", "logs"} {
				m.FocusPane(pane)

				view := ansi.Strip(m.View())
				if size.Width == 0 {
					if view != "" {
						t.Fatal("zero allocation rendered")
					}

					continue
				}

				lines := strings.Split(view, "\n")
				if len(lines) != size.Height {
					t.Fatalf("%s %v incorrect height", preset, size)
				}

				for _, line := range lines {
					if ansi.StringWidth(line) != size.Width {
						t.Fatalf("%s %v overflow: %q", preset, size, line)
					}
				}
			}

			if size.Width <= 80 && m.ActiveLayout() != pattern.InspectorFocused {
				t.Fatalf("%s did not use focused fallback at %v", preset, size)
			}
		}
	}

	if _, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Layout: pattern.InspectorLayout("invalid")}); err == nil {
		t.Fatal("invalid initial layout accepted")
	}
}

func TestInspectorLayoutKeyboardCycle(t *testing.T) {
	m, err := pattern.NewLogInspector(pattern.LogInspectorOptions{Theme: theme.New(theme.Dark)})
	if err != nil {
		t.Fatal(err)
	}

	for _, want := range []pattern.InspectorLayout{pattern.InspectorTimeline, pattern.InspectorColumns, pattern.InspectorFocused, pattern.InspectorSidebar} {
		m.Update(tea.KeyPressMsg{Code: 'L'})

		if m.Layout() != want {
			t.Fatalf("layout cycle got %s want %s", m.Layout(), want)
		}
	}
}
