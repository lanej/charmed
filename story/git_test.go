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

func TestGitBrowserDeclaredLayoutsAndHiddenState(t *testing.T) {
	fixture, err := gitBrowser(Environment{Config: DefaultConfig()})
	if err != nil {
		t.Fatal(err)
	}

	m, ok := fixture.(*gitBrowserScene)
	if !ok {
		t.Fatal("unexpected browser scene")
	}

	sizes := []layout.Size{{Width: 40, Height: 12}, {Width: 100, Height: 24}, {Width: 120, Height: 32}, {Width: 160, Height: 40}}
	for _, focus := range []bool{false, true} {
		if _, err := m.screen.CheckLayouts(gitBrowserLayout(focus, 0), sizes...); err != nil {
			t.Fatal(err)
		}
	}

	for _, tc := range []struct {
		size    layout.Size
		variant string
	}{{sizes[0], "focused"}, {sizes[1], "sidebar"}, {sizes[2], "three-column"}, {sizes[3], "three-column"}} {
		m.Resize(tc.size)

		if m.layoutErr != nil {
			t.Fatal(m.layoutErr)
		}

		if m.screen.Plan().Variant() != tc.variant {
			t.Fatalf("at %v: %s", tc.size, m.screen.Plan().Variant())
		}

		if view := m.View(); strings.Contains(view, "layout rule") || !strings.Contains(view, "Repository browser") {
			t.Fatalf("invalid frame at %v: %s", tc.size, view)
		}
	}

	m.Resize(sizes[0])
	oldWidth, oldHeight := m.status.List.Viewport.Width(), m.status.List.Viewport.Height()
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})

	if _, visible := m.screen.Region("main"); visible {
		t.Fatal("narrow details view retained main allocation")
	}

	if _, visible := m.screen.Region("details"); !visible {
		t.Fatal("narrow focus did not reveal details")
	}

	if m.status.List.Viewport.Width() != oldWidth || m.status.List.Viewport.Height() != oldHeight {
		t.Fatal("hiding a pane resized its native state")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Resize(sizes[2])
	m.Update(tea.KeyPressMsg{Code: 'l', Mod: tea.ModAlt})

	if m.screen.Plan().Variant() != "sidebar" {
		t.Fatal("layout choice did not select sidebar")
	}

	m.Update(tea.KeyPressMsg{Code: 'l', Mod: tea.ModCtrl})

	if !m.showRegions || !strings.Contains(m.View(), "main") {
		t.Fatal("region map did not show allocation identities")
	}
}

func TestGitRegionInspectionPreservesIndependentHistoryScroll(t *testing.T) {
	fixture, err := gitBrowser(Environment{Config: DefaultConfig()})
	if err != nil {
		t.Fatal(err)
	}

	m, ok := fixture.(*gitBrowserScene)
	if !ok {
		t.Fatal("unexpected browser scene")
	}

	m.Resize(layout.Size{Width: 40, Height: 12})
	m.Update(tea.KeyPressMsg{Code: tea.KeyRight, Mod: tea.ModCtrl})

	for range 4 {
		m.Update(tea.KeyPressMsg{Code: 'e', Mod: tea.ModCtrl})
	}

	before := m.history.Viewport.YOffset()
	if before == 0 {
		t.Fatal("history did not scroll")
	}

	for range 2 {
		m.Update(tea.KeyPressMsg{Code: 'l', Mod: tea.ModCtrl})
	}

	if after := m.history.Viewport.YOffset(); after != before {
		t.Fatalf("inspection moved scroll: %d to %d", before, after)
	}

	m.Resize(layout.Size{Width: 120, Height: 32})

	for range 4 {
		m.Update(tea.KeyPressMsg{Code: 'e', Mod: tea.ModCtrl})
	}

	before = m.history.Viewport.YOffset()
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})

	if m.history.Viewport.YOffset() != before {
		t.Fatal("unchanged wide geometry reset hidden input focus state")
	}
}

func TestGitStatusScrollSurvivesUnchangedGeometry(t *testing.T) {
	fixture, err := gitBrowser(Environment{Config: DefaultConfig()})
	if err != nil {
		t.Fatal(err)
	}

	m, ok := fixture.(*gitBrowserScene)
	if !ok {
		t.Fatal("unexpected browser scene")
	}

	m.Resize(layout.Size{Width: 120, Height: 16})

	if m.layoutErr != nil {
		t.Fatal(m.layoutErr)
	}

	for range 4 {
		m.Update(tea.KeyPressMsg{Code: 'e', Mod: tea.ModCtrl})
	}

	before := m.status.List.Viewport.YOffset()
	if before == 0 {
		t.Fatal("status did not scroll")
	}

	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m.Update(tea.KeyPressMsg{Code: tea.KeyRight, Mod: tea.ModCtrl})
	m.Update(tea.KeyPressMsg{Code: tea.KeyLeft, Mod: tea.ModCtrl})

	if m.status.List.Viewport.YOffset() != before {
		t.Fatal("unchanged geometry reset status scroll")
	}
}
