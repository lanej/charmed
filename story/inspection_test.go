package story_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
)

func TestListInspectorStoryStatesRenderSuppliedEvidence(t *testing.T) {
	t.Parallel()

	definition, err := story.Find("pattern/list-inspector")
	if err != nil {
		t.Fatal(err)
	}

	wants := map[string]string{
		"ready": "Outcome unknown", "inspecting": "j/k scroll",
		"empty": "No operations observed", "loading": "Waiting for the first observation",
		"error": "Refresh failed", "confirmed": "Route observer",
	}

	for _, state := range definition.States {
		t.Run(state, func(t *testing.T) {
			t.Parallel()

			config := definition.Defaults
			config.State = state

			frame, err := story.Frame(definition.ID, config)
			if err != nil {
				t.Fatal(err)
			}

			want, present := wants[state]
			if !present || !strings.Contains(ansi.Strip(frame), want) {
				t.Fatalf("%s lost supplied evidence %q: %s", state, want, ansi.Strip(frame))
			}

			if lipgloss.Width(frame) != config.Columns || lipgloss.Height(frame) != config.Rows {
				t.Fatalf("%s escaped its terminal allocation", state)
			}
		})
	}

	config := definition.Defaults
	config.State = "unregistered"

	if _, err := story.Frame(definition.ID, config); err == nil {
		t.Fatal("unknown fixture state accepted")
	}
}

func TestListInspectorStoryHelpTracksInspection(t *testing.T) {
	t.Parallel()

	config := story.DefaultConfig()
	config.Columns, config.Rows = 40, 18

	m, err := story.New("pattern/list-inspector", config, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := m.Close(); err != nil {
			t.Error(err)
		}
	})

	if !strings.Contains(ansi.Strip(m.Render().Content), "j/k select") {
		t.Fatal("list help did not describe record navigation")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	if view := ansi.Strip(m.Render().Content); !strings.Contains(view, "j/k scroll") || strings.Contains(view, "j/k select") {
		t.Fatal("inspection retained selection help")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})

	if view := ansi.Strip(m.Render().Content); !strings.Contains(view, "j/k select") || strings.Contains(view, "j/k scroll") {
		t.Fatal("returning to the list retained scrolling help")
	}
}
