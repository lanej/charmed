package main

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/examples/operations/fixture"
)

// Copy this acceptance path with the starter and replace the fixture states
// with the consumer's observation messages; an import alone proves no UI behavior.
func TestOperationsConsumerAcceptance(t *testing.T) {
	m := newModel()
	initial := m.View()

	if initial.ForegroundColor == nil || initial.BackgroundColor == nil {
		t.Fatal("starter must own both default colors so unstyled content remains readable in light terminals")
	}

	red, green, blue, _ := initial.ForegroundColor.RGBA()
	if red != 0xd7d7 || green != 0xdfdf || blue != 0xd9d9 {
		t.Fatalf("dark starter default foreground = %04x/%04x/%04x, want Sage Midnight #d7dfd9", red, green, blue)
	}

	for _, size := range []tea.WindowSizeMsg{{Width: 132, Height: 40}, {Width: 80, Height: 24}, {Width: 40, Height: 18}} {
		m.Update(size)

		for _, state := range []string{"ready", "empty", "loading", "error"} {
			m.dashboard.SetSnapshot(fixture.Snapshot(state, m.dashboard.Theme))
			view := ansi.Strip(m.View().Content)
			lines := strings.Split(view, "\n")

			if len(lines) != size.Height {
				t.Fatalf("%v/%s: escaped height", size, state)
			}

			for _, line := range lines {
				if ansi.StringWidth(line) != size.Width {
					t.Fatalf("%v/%s: escaped width: %q", size, state, line)
				}
			}

			switch state {
			case "error":
				if !strings.Contains(view, "Refresh failed") || !strings.Contains(view, "3500") {
					t.Fatalf("%v: failed refresh lost last observation or its warning", size)
				}
			case "empty", "loading":
				if strings.Contains(view, "3500") || !strings.Contains(view, "missing") {
					t.Fatalf("%v/%s: absent observation looks populated", size, state)
				}
			}
		}

		m.dashboard.SetSnapshot(fixture.Snapshot("ready", m.dashboard.Theme))
		m.Update(tea.KeyPressMsg{Code: 'i'})

		if !strings.Contains(ansi.Strip(m.View().Content), "Scheduler") {
			t.Fatalf("%v: inspection is unreachable", size)
		}

		m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
		m.Update(tea.KeyPressMsg{Code: '?'})

		if !strings.Contains(ansi.Strip(m.View().Content), "page down") {
			t.Fatalf("%v: full keyboard help is unreachable", size)
		}

		m.Update(tea.KeyPressMsg{Code: '?'})
	}

	_, quit := m.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})
	if quit == nil {
		t.Fatal("parent did not handle quit")
	}

	if _, ok := quit().(tea.QuitMsg); !ok {
		t.Fatal("parent did not return terminal shutdown command")
	}
}
