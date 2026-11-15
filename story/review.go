package story

import (
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/examples/review/fixture"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
)

func reviewDefinitions() []Definition {
	defaults := DefaultConfig()
	defaults.Columns, defaults.Rows = 120, 32

	return []Definition{
		{ID: "collection/section-stack", Title: "Organisms/Collection/Section stack", Level: "Organisms", Family: "collection", Description: "One scrolling document of ordered, independently foldable sections with stable identities and rendered-row anchors.", Source: "collection/section_stack.go", Symbol: "SectionStack", Composes: []string{"collection/viewport", "text/disclosure"}, Defaults: defaults, Factory: sectionStackStory},
		{ID: "pattern/outline-document", Title: "Organisms/Collection/Outline document", Level: "Organisms", Family: "collection", Description: "Link any ordered outline to stacked sections. Selection, scrolling and subtree folding stay synchronized without file assumptions.", Source: "pattern/outline_document.go", Symbol: "OutlineDocument", Composes: []string{"collection/tree", "collection/section-stack"}, Defaults: defaults, Controls: []string{"state"}, States: []string{"ready", "folded", "empty"}, Factory: outlineStory},
		{ID: "git/review", Title: "Organisms/Git/Diff review", Level: "Organisms", Family: "git", Description: "Review stacked caller-supplied diffs with optional file grouping, statistics and independent reviewed marks. Changed diff content clears its mark.", Source: "gitui/review.go", Symbol: "DiffReview", Composes: []string{"pattern/outline-document", "git/hunk", "git/diff-summary"}, Defaults: defaults, Controls: []string{"state"}, States: []string{"ready", "folded", "reviewed", "empty"}, Factory: reviewStory},
	}
}

func sectionStackStory(e Environment) (runtime.Scene, error) {
	return collection.NewSectionStack(fixture.Document(), e.Config.Size(), e.Config.Palette())
}

func outlineStory(e Environment) (runtime.Scene, error) {
	entries := fixture.Document()
	if e.Config.State == "empty" {
		entries = nil
	}

	m, err := pattern.NewOutlineDocument(entries, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	if e.Config.State == "folded" {
		m.SetCollapsed("proposal", true)
	}

	return outlineScene(m, m.Update), nil
}

func reviewStory(e Environment) (runtime.Scene, error) {
	files := fixture.Files()
	if e.Config.State == "empty" {
		files = nil
	}

	m, err := gitui.NewFileDiffReview(files, e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	switch e.Config.State {
	case "reviewed":
		m.SetReviewed("routing", true)
	case "folded":
		m.SetCollapsed(m.Document.Sections()[0].ID, true)
	}

	return outlineScene(m.OutlineDocument, m.Update), nil
}

func outlineScene(m *pattern.OutlineDocument, update func(tea.Msg) tea.Cmd) runtime.Scene {
	return &scene{resize: m.Resize, view: func(layout.Size) string { return m.View() }, update: func(msg tea.Msg) tea.Cmd {
		if pressed, ok := msg.(tea.KeyPressMsg); ok && (pressed.String() == "tab" || pressed.String() == "shift+tab") {
			m.SetFocus(!m.DocumentFocused())
			return nil
		}

		return update(msg)
	}}
}
