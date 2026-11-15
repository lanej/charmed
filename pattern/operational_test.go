package pattern_test

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/text"
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

func TestOperationalDashboardHeaderHierarchy(t *testing.T) {
	snapshot := operationalFixture()
	m := pattern.NewOperationalDashboard(snapshot, layout.Size{Width: 132, Height: 24}, theme.New(theme.Dark))
	first := strings.Split(ansi.Strip(m.View()), "\n")[0]

	for _, label := range []string{snapshot.Title, snapshot.Context, snapshot.Status} {
		if !strings.Contains(first, label) {
			t.Fatalf("wide header does not keep %q on the first row: %q", label, first)
		}
	}

	if !strings.HasSuffix(first, snapshot.Status+" ") {
		t.Fatalf("panel status is not aligned to the right gutter: %q", first)
	}

	m.Resize(layout.Size{Width: 40, Height: 24})
	narrow := ansi.Strip(m.View())

	for _, label := range []string{snapshot.Title, snapshot.Context, snapshot.Status} {
		if !strings.Contains(narrow, label) {
			t.Fatalf("narrow header dropped %q: %q", label, narrow)
		}
	}
}

func TestOperationalDashboardHeaderRetainsWrappedContext(t *testing.T) {
	snapshot := operationalFixture()
	snapshot.Context = "A long source description that must wrap without losing its final observation marker"
	snapshot.Status = ""
	m := pattern.NewOperationalDashboard(snapshot, layout.Size{Width: 40, Height: 24}, theme.New(theme.Dark))
	view := strings.Join(strings.Fields(ansi.Strip(m.View())), " ")

	if !strings.Contains(view, snapshot.Context) {
		t.Fatalf("wrapped context disappeared: %q", view)
	}
}

func TestOperationalDashboardMetersPreserveCountsAndPresence(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.ASCII = true
	snapshot := pattern.DashboardSnapshot{Meters: []pattern.DashboardMeter{
		{Label: "Remaining", Value: 3500, Total: 5000, Unit: "requests"},
		{Label: "Empty", Value: 0, Total: 5000},
		{Label: "Missing", Value: 123, Total: 5000, State: text.ValueMissing},
		{Label: "Unbounded", Value: 0, Total: 0},
		{Label: "Over", Value: 18446744073709551615, Total: 5},
	}}

	for _, width := range []int{132, 80, 40, 20} {
		m := pattern.NewOperationalDashboard(snapshot, layout.Size{Width: width, Height: 40}, palette)

		view := ansi.Strip(m.View())
		for _, label := range []string{"3500", "5000", "requests", "missing", "no capacity", "18446744073709551615", "invalid ratio"} {
			// At short widths a long exact count may wrap across terminal rows.
			compact := strings.Join(strings.Fields(view), "")
			if !strings.Contains(compact, strings.ReplaceAll(label, " ", "")) {
				t.Fatalf("width %d: lost %q: %q", width, label, view)
			}
		}

		if strings.Contains(view, "123") || !strings.Contains(view, "#") {
			t.Fatalf("width %d: missing count disclosed or valid meter absent: %q", width, view)
		}
	}

	m := pattern.NewOperationalDashboard(snapshot, layout.Size{Width: 132, Height: 24}, palette)

	rows := strings.Split(ansi.Strip(m.View()), "\n")
	for _, row := range rows {
		if strings.Contains(row, "Remaining") && strings.Count(row, "#") != 20 {
			t.Fatalf("3500 of 5000 must round to 20 of 28 cells: %q", row)
		}

		if (strings.Contains(row, "Missing") || strings.Contains(row, "Unbounded") || strings.Contains(row, "Over")) && strings.Contains(row, "#") {
			t.Fatalf("unavailable or invalid ratio received a meter: %q", row)
		}
	}
}
