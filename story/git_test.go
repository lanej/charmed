package story

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
)

func TestGitBrowserWorktreeContextAndFocus(t *testing.T) {
	config := DefaultConfig()
	config.Columns, config.Rows = 120, 32

	fixture, err := gitBrowser(Environment{Config: config})
	if err != nil {
		t.Fatal(err)
	}

	m, ok := fixture.(*gitBrowserScene)
	if !ok {
		t.Fatal("unexpected browser scene")
	}

	for range 4 {
		m.Update(tea.KeyPressMsg{Code: tea.KeyRight, Mod: tea.ModCtrl})
	}

	if m.page != 4 {
		t.Fatal("section navigation failed")
	}

	m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})

	cmd := m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("worktree selection missing")
	}

	m.Update(cmd())

	if m.context != "~/src/charmed-search" {
		t.Fatal("worktree selection did not switch fixture context")
	}

	summary := ansi.Strip(m.status.View())
	if !strings.Contains(summary, "HEAD → search") || !strings.Contains(summary, "staged 0") {
		t.Fatalf("wrong worktree status: %s", summary)
	}

	for _, line := range strings.Split(ansi.Strip(m.branches.View()), "\n") {
		if strings.Contains(line, "[HEAD]") && !strings.Contains(line, "search") {
			t.Fatalf("stale branch HEAD: %s", line)
		}
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	before := m.worktrees.SelectedID()
	m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})

	if before != m.worktrees.SelectedID() {
		t.Fatal("detail focus moved worktree selection")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Update(tea.KeyPressMsg{Code: 'j', Text: "j"})

	cmd = m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("detached worktree selection missing")
	}

	m.Update(cmd())

	if view := m.status.View(); !strings.Contains(view, "Status unavailable") || strings.Contains(view, "staged 0") {
		t.Fatal("unknown status became clean")
	}

	for _, size := range []layout.Size{{Width: 40, Height: 12}, {Width: 120, Height: 32}} {
		m.Resize(size)

		if !strings.Contains(m.View(), "Repository browser") {
			t.Fatal("resize lost browser content")
		}
	}
}

func TestGitBrowserForwardsSearchPaste(t *testing.T) {
	fixture, err := gitBrowser(Environment{Config: DefaultConfig()})
	if err != nil {
		t.Fatal(err)
	}

	m, ok := fixture.(*gitBrowserScene)
	if !ok {
		t.Fatal("unexpected browser scene")
	}

	m.Update(tea.KeyPressMsg{Code: '/', Text: "/"})
	m.Update(tea.PasteMsg{Content: "no matching file"})

	if m.status.List.SelectedID() != "" {
		t.Fatal("browser dropped pasted filter")
	}

	if !strings.Contains(m.detail.View(), "No item selected") {
		t.Fatal("browser retained hidden details")
	}
}
