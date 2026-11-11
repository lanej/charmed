package story_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
)

func TestControlStoriesHandleNativeActivationThroughTypedEvents(t *testing.T) {
	for _, id := range []string{"input/checkbox", "input/radio", "input/toggle", "input/button"} {
		config := story.DefaultConfig()
		config.Columns, config.Rows = 70, 18

		scene, err := story.New(id, config, nil)
		if err != nil {
			t.Fatal(err)
		}

		_, cmd := scene.Update(tea.KeyPressMsg{Code: tea.KeySpace})
		if cmd == nil {
			t.Fatalf("%s did not activate", id)
		}

		scene.Update(cmd())

		view := ansi.Strip(scene.View().Content)
		if !strings.Contains(view, "Changed:") && !strings.Contains(view, "Selected:") && !strings.Contains(view, "Toggled:") && !strings.Contains(view, "Requested:") {
			t.Fatalf("%s did not receive its event: %q", id, view)
		}

		if err := scene.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
