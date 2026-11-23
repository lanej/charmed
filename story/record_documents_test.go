package story_test

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
)

func TestDetailDocumentStoryKeepsNarrowEvidenceReachable(t *testing.T) {
	t.Parallel()

	definition, err := story.Find("layout/detail-document")
	if err != nil {
		t.Fatal(err)
	}

	config := definition.Defaults
	config.Columns, config.Rows = 40, 16

	m, err := story.New(definition.ID, config, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := m.Close(); err != nil {
			t.Error(err)
		}
	})

	if !strings.Contains(ansi.Strip(m.Render().Content), "Outcome unknown") {
		t.Fatal("narrow inspection lost the supplied operation state")
	}

	for range 10 {
		m.Update(tea.KeyPressMsg{Code: tea.KeyPgDown})
	}

	view := ansi.Strip(m.Render().Content)
	if !strings.Contains(view, "Inspection notes") || !strings.Contains(view, "attribution gap") {
		t.Fatalf("narrow document made its final section unreachable:\n%s", view)
	}
}
