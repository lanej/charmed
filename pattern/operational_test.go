package pattern_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func operationalFixture() pattern.DashboardSnapshot {
	return pattern.DashboardSnapshot{
		Title: "API usage", Context: "Synthetic / UTC", Status: "Last observed 00:16",
		Summary:  []layout.Pair{{Label: "REST remaining", Value: "3500 / 5000"}},
		Sections: []pattern.DashboardSection{{Title: "Traffic", Body: func(int) string { return strings.Repeat("request evidence\n", 60) + "overview end" }}},
		Details:  []pattern.DashboardSection{{Title: "Scheduler", Body: func(int) string { return strings.Repeat("dispatch evidence\n", 60) + "details end" }}},
	}
}

func TestOperationalDashboardInspectionRetainsIndependentOffsets(t *testing.T) {
	m := pattern.NewOperationalDashboard(operationalFixture(), layout.Size{Width: 80, Height: 24}, theme.New(theme.Dark))
	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})

	offset := m.Overview.YOffset()
	if offset == 0 {
		t.Fatal("overview did not scroll")
	}

	m.Update(tea.KeyPressMsg{Code: 'i'})

	if !m.Inspecting() || !strings.Contains(ansi.Strip(m.View()), "Scheduler") {
		t.Fatal("inspection did not open")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	detailOffset := m.Inspection.YOffset()
	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if m.Inspecting() || m.Overview.YOffset() != offset {
		t.Fatal("closing inspection lost overview position")
	}

	m.Update(tea.KeyPressMsg{Code: 'i'})

	if m.Inspection.YOffset() != detailOffset {
		t.Fatal("reopening inspection lost its position")
	}

	m.Resize(layout.Size{Width: 40, Height: 18})
	m.SetSnapshot(operationalFixture())

	if !m.Inspecting() || m.Inspection.YOffset() != detailOffset || m.Overview.YOffset() != offset {
		t.Fatal("resize or data replacement lost inspection or scroll state")
	}

	m.SetSnapshot(pattern.DashboardSnapshot{Title: "No details"})

	if m.Inspecting() || strings.Contains(ansi.Strip(m.View()), "i details") {
		t.Fatal("removed inspection is still available")
	}
}

func TestOperationalDashboardNativeBindingsAndFocus(t *testing.T) {
	m := pattern.NewOperationalDashboard(operationalFixture(), layout.Size{Width: 80, Height: 24}, theme.New(theme.Dark))
	m.Overview.KeyMap.Down = key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "next line"))
	m.Overview.Style = m.Overview.Style.Bold(true)
	m.SetSnapshot(operationalFixture())
	m.Resize(layout.Size{Width: 40, Height: 18})

	if !strings.Contains(ansi.Strip(m.View()), "n next line") || !m.Overview.Style.GetBold() {
		t.Fatal("resize or refresh lost native customization")
	}

	m.Blur()
	m.Update(tea.KeyPressMsg{Code: 'i'})
	m.Update(tea.KeyPressMsg{Code: 'n'})

	if m.Inspecting() || m.Overview.YOffset() != 0 {
		t.Fatal("inactive dashboard handled input")
	}

	m.Focus()
	m.Update(tea.KeyPressMsg{Code: 'n'})

	if m.Overview.YOffset() != 1 {
		t.Fatal("custom native key was not routed")
	}

	m.Update(tea.KeyPressMsg{Code: 'd'})

	if m.Inspecting() || m.Overview.YOffset() <= 1 {
		t.Fatal("native half-page scrolling conflicts with inspection")
	}
}

func TestOperationalDashboardBoundsAndReachability(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		for _, size := range []layout.Size{{Width: 1, Height: 1}, {Width: 20, Height: 8}, {Width: 40, Height: 18}, {Width: 80, Height: 24}, {Width: 132, Height: 50}} {
			palette := theme.New(mode)
			m := pattern.NewOperationalDashboard(operationalFixture(), size, palette)

			for _, action := range []rune{0, '?', 'i'} {
				if action != 0 {
					m.Update(tea.KeyPressMsg{Code: action})
				}

				view := m.View()

				lines := strings.Split(view, "\n")
				if len(lines) != size.Height {
					t.Fatalf("%v height %d", size, len(lines))
				}

				for _, line := range lines {
					if ansi.StringWidth(line) != size.Width {
						t.Fatalf("%v width escaped: %q", size, line)
					}
				}
			}
		}
	}

	m := pattern.NewOperationalDashboard(operationalFixture(), layout.Size{Width: 40, Height: 18}, theme.New(theme.Dark))
	for range 100 {
		m.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	}

	if !strings.Contains(ansi.Strip(m.View()), "overview end") {
		t.Fatal("bottom evidence cannot be reached by keyboard")
	}
}

func TestOperationalDashboardViewDoesNotPrepareContent(t *testing.T) {
	calls := 0
	snapshot := operationalFixture()
	snapshot.Sections[0].Body = func(int) string { calls++; return "observation" }
	m := pattern.NewOperationalDashboard(snapshot, layout.Size{Width: 80, Height: 24}, theme.New(theme.Dark))
	prepared := calls

	m.View()
	m.View()

	if prepared == 0 || calls != prepared {
		t.Fatal("View ran content preparation")
	}
}
