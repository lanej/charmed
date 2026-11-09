package story_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/story"
)

func TestApplicationShellSelectionSurvivesResize(t *testing.T) {
	config := story.DefaultConfig()

	m, err := story.New("template/application", config, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := m.Close(); err != nil {
			t.Error(err)
		}
	})

	_, cmd := m.Update(tea.KeyPressMsg{Code: tea.KeyRight})
	if cmd == nil {
		t.Fatal("section change did not return a selection command")
	}

	selected, ok := cmd().(interaction.Selected)
	if !ok || selected.ID != "services" {
		t.Fatalf("section event: %+v", selected)
	}

	m.Update(selected)

	for _, width := range []int{40, 100} {
		m.Update(tea.WindowSizeMsg{Width: width, Height: 12})
		view := ansi.Strip(m.View().Content)

		for _, want := range []string{" Services ", "api          running", "Ctrl+C quit", "Service workbench"} {
			if !strings.Contains(view, want) {
				t.Fatalf("width %d lost %q: %s", width, want, view)
			}
		}

		if strings.Contains(view, "12 services / all healthy") {
			t.Fatal("old tab body remained visible")
		}

		lines := strings.Split(view, "\n")
		if len(lines) != 12 {
			t.Fatalf("height: %d", len(lines))
		}

		for _, line := range lines {
			if ansi.StringWidth(line) != width {
				t.Fatalf("width %d: %q", width, line)
			}
		}
	}
}
