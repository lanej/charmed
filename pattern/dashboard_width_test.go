package pattern_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

func TestDashboardStatusAllocationPreservesServiceNames(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Mono} {
		m := pattern.NewDashboard(pattern.DashboardOptions{Theme: theme.New(mode), Services: pattern.SampleServices()})
		m.Init()

		for _, width := range []int{90, 100} {
			m.Resize(layout.Size{Width: width, Height: 34})
			view := ansi.Strip(m.Table.View())

			for _, name := range []string{"API gateway", "Event stream", "Job workers", "Storage", "Scheduler"} {
				if !strings.Contains(view, name) {
					t.Fatalf("%s at %d columns clipped %q: %s", mode, width, name, view)
				}
			}
		}

		columns := m.Table.Columns()
		for _, character := range "api" {
			m.Update(tea.KeyPressMsg{Code: character, Text: string(character)})
		}

		if m.Filter.Value() != "api" || len(m.Table.Rows()) != 1 {
			t.Fatal("input did not filter the table to one service")
		}

		m.Resize(layout.Size{Width: 100, Height: 34})

		if got := m.Table.Columns(); got[1].Width != columns[1].Width {
			t.Fatal("filter changed reserved status width")
		}
	}
}

func TestEmptyDashboardKeepsNativeStatusHeading(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Mono} {
		m := pattern.NewDashboard(pattern.DashboardOptions{Theme: theme.New(mode)})
		for _, width := range []int{40, 90, 100} {
			m.Resize(layout.Size{Width: width, Height: 34})

			if view := ansi.Strip(m.Table.View()); !strings.Contains(view, "STATE") {
				t.Fatalf("%s at %d columns clipped empty status heading: %s", mode, width, view)
			}
		}
	}
}
