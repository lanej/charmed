package main

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

func TestInspectionConsumerGeometryAndInspection(t *testing.T) {
	t.Parallel()

	for _, size := range []tea.WindowSizeMsg{{Width: 132, Height: 40}, {Width: 80, Height: 24}, {Width: 40, Height: 18}} {
		m, err := newModel()
		if err != nil {
			t.Fatal(err)
		}

		m.Update(size)

		if !strings.Contains(ansi.Strip(m.View().Content), "Outcome uncertain") {
			t.Fatalf("%v: list did not preserve the supplied unknown condition", size)
		}

		if !strings.Contains(ansi.Strip(m.footer), "j/k select") {
			t.Fatalf("%v: list help did not describe record navigation", size)
		}

		m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

		if !m.inspector.Inspecting() {
			t.Fatalf("%v: Enter did not open evidence", size)
		}

		if footer := ansi.Strip(m.footer); !strings.Contains(footer, "j/k scroll") || strings.Contains(footer, "j/k select") {
			t.Fatalf("%v: inspection help did not describe native scrolling", size)
		}

		view := m.View()
		if view.ForegroundColor == nil || view.BackgroundColor == nil {
			t.Fatal("consumer did not own native default colors")
		}

		lines := strings.Split(ansi.Strip(view.Content), "\n")
		if len(lines) != size.Height {
			t.Fatalf("%v: frame escaped its height", size)
		}

		for _, line := range lines {
			if ansi.StringWidth(line) != size.Width {
				t.Fatalf("%v: frame escaped its width: %q", size, line)
			}
		}

		m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

		if m.inspector.Inspecting() {
			t.Fatalf("%v: Escape did not return to the list", size)
		}

		if footer := ansi.Strip(m.footer); !strings.Contains(footer, "j/k select") || strings.Contains(footer, "j/k scroll") {
			t.Fatalf("%v: list help retained inspection navigation", size)
		}

		m.Update(tea.KeyPressMsg{Code: 'j'})

		if id, present := m.inspector.SelectedID(); !present || id != "example/service#43" {
			t.Fatalf("%v: application did not forward record navigation", size)
		}
	}
}

func TestInspectionConsumerOwnsRefreshAndQuit(t *testing.T) {
	t.Parallel()

	m, err := newModel()
	if err != nil {
		t.Fatal(err)
	}

	m.Update(tea.WindowSizeMsg{Width: 80, Height: 24})
	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	m.Update(tea.KeyPressMsg{Code: ']'})

	if m.inspector.Detail.YOffset() == 0 {
		t.Fatal("native bracket paging did not scroll the evidence")
	}

	offset := m.inspector.Detail.YOffset()
	m.Update(tea.KeyPressMsg{Code: 'r'})

	if !strings.Contains(m.status, "Refresh failed") || m.inspector.Detail.YOffset() != offset {
		t.Fatal("failed application refresh lost the retained observation position")
	}

	m.inspector.Detail.GotoTop()

	if !strings.Contains(ansi.Strip(m.View().Content), "Outcome uncertain") {
		t.Fatal("failed refresh strengthened the operation's evidence")
	}

	m.Update(tea.KeyPressMsg{Code: 'r'})

	if id, present := m.inspector.SelectedID(); !present || id != "example/service#42" || !m.inspector.Inspecting() {
		t.Fatal("replacement evidence lost the inspected record identity")
	}

	if !strings.Contains(ansi.Strip(m.View().Content), "Merge confirmed") {
		t.Fatal("later supplied evidence did not reach the inspector")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	m.Update(tea.KeyPressMsg{Code: 'j'})

	if id, present := m.inspector.SelectedID(); !present || id != "example/service#43" {
		t.Fatal("returning from inspection changed the record order")
	}

	_, quit := m.Update(tea.KeyPressMsg{Code: 'q'})
	if quit == nil {
		t.Fatal("parent did not handle quit")
	}

	if _, ok := quit().(tea.QuitMsg); !ok {
		t.Fatal("parent did not return the terminal shutdown command")
	}
}
