package story_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
)

func TestStateDiagramStoryUpdatesAndResizes(t *testing.T) {
	config := story.DefaultConfig()
	config.Columns, config.Rows = 120, 40

	model, err := story.New("diagram/state-transitions", config, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := model.Close(); err != nil {
			t.Error(err)
		}
	}()

	original := model.Render().Content

	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyRight})
	if model.Render().Content != original {
		t.Fatal("Right wrapped beyond final occurrence")
	}

	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyLeft})
	if model.Render().Content == original {
		t.Fatal("Left did not inspect the previous occurrence")
	}

	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyEnd})
	if model.Render().Content != original {
		t.Fatal("End did not restore final occurrence")
	}

	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyEnter})

	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	if !strings.Contains(ansi.Strip(model.Render().Content), "Last three") {
		t.Fatal("trace-window group missing")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 40, Height: 12})

	narrow := ansi.Strip(model.Render().Content)
	if !strings.Contains(narrow, "rows; available") || strings.Contains(narrow, "retry") {
		t.Fatal("narrow view displayed a partial graph instead of allocation guidance")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 40, Height: 40})

	vertical := ansi.Strip(model.Render().Content)
	for _, label := range []string{"Idle", "Running", "Done", "start", "retry", "finish", "cancel", "Failed", "success", "error", "fail"} {
		if !strings.Contains(vertical, label) {
			t.Fatalf("narrow layout lost %q", label)
		}
	}

	if strings.Contains(vertical, "Crossings") {
		t.Fatal("narrow graph needs a legend")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 100, Height: 8})
	if !strings.Contains(ansi.Strip(model.Render().Content), "rows; available") {
		t.Fatal("short view silently clipped graph topology")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	if !strings.Contains(model.Render().Content, "Done") {
		t.Fatal("resize lost the caller graph")
	}
}
