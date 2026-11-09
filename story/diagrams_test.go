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
	config.Columns, config.Rows = 100, 30

	model, err := story.New("diagram/state-transitions", config, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := model.Close(); err != nil {
			t.Error(err)
		}
	}()

	original := model.View().Content

	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyRight})
	if model.View().Content == original {
		t.Fatal("Right did not change active state marker")
	}

	_, _ = model.Update(tea.KeyPressMsg{Code: tea.KeyLeft})
	if model.View().Content != original {
		t.Fatal("Left did not restore the active state marker")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 40, Height: 12})

	narrow := ansi.Strip(model.View().Content)
	if !strings.Contains(narrow, "needs more rows") || strings.Contains(narrow, "retry") {
		t.Fatal("narrow view displayed a partial graph instead of allocation guidance")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 40, Height: 30})

	vertical := ansi.Strip(model.View().Content)
	for _, label := range []string{"Idle", "Running", "Done", "start", "retry", "finish", "cancel"} {
		if !strings.Contains(vertical, label) {
			t.Fatalf("narrow layout lost %q", label)
		}
	}

	if strings.Contains(vertical, "Crossings") {
		t.Fatal("narrow graph needs a legend")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 100, Height: 8})
	if !strings.Contains(ansi.Strip(model.View().Content), "needs more rows") {
		t.Fatal("short view silently clipped graph topology")
	}

	_, _ = model.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
	if model.View().Content != original {
		t.Fatal("resize lost the caller graph or active state")
	}
}
