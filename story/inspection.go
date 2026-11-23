package story

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/examples/inspection/fixture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
)

func inspectionDefinitions() []Definition {
	return []Definition{{
		ID: "pattern/list-inspector", Title: "Organisms/Patterns/List inspector", Level: "Organisms", Family: "pattern",
		Description: "Browse stable records with unequal row heights and independently scrollable evidence. Wide layouts preview the selection; Enter opens inspection and Escape restores the list.",
		Source:      "pattern/list_inspector.go", Symbol: "ListInspector",
		Controls: []string{"state"}, States: []string{"ready", "inspecting", "empty", "loading", "error", "confirmed"},
		Composes: []string{"collection/viewport", "foundations/layout"},
		Defaults: func() Config { c := DefaultConfig(); c.Columns = 112; c.Rows = 28; return c }(),
		Factory:  inspectionStory,
	}}
}

func inspectionStory(e Environment) (runtime.Scene, error) {
	m, err := pattern.NewListInspector(fixture.Records(e.Config.State, e.Config.Palette()), e.Config.Size(), e.Config.Palette())
	if err != nil {
		return nil, err
	}

	m.ListTitle, m.DetailTitle = "Operations", "Evidence"
	if e.Config.State == "loading" {
		m.Empty = "Waiting for the first observation"
	} else {
		m.Empty = "No operations observed"
	}

	m.Refresh()

	if e.Config.State == "inspecting" || e.Config.VerboseDetails {
		m.Inspect()
	}

	var help string

	size := e.Config.Size()
	resize := func(next layout.Size) {
		size = next
		navigation := "select"
		toggleKey, toggleLabel := "enter", "inspect"

		if m.Inspecting() {
			navigation = "scroll"
			toggleKey, toggleLabel = "esc", "back"
		}

		help = (layout.TextRow{Gap: 2, Items: []layout.TextRowItem{
			{Text: text.KeyHint("j/k", navigation, m.Theme)},
			{Text: text.KeyHint(toggleKey, toggleLabel, m.Theme)},
			{Text: text.KeyHint("[/]", "scroll", m.Theme)},
		}}).View(max(1, size.Width))
		m.Resize(layout.Size{Width: size.Width, Height: max(0, size.Height-lipgloss.Height(help))})
	}
	resize(size)

	update := func(msg tea.Msg) tea.Cmd {
		wasInspecting := m.Inspecting()
		cmd := m.Update(msg)

		if m.Inspecting() != wasInspecting {
			resize(size)
		}

		return cmd
	}

	return &scene{update: update, resize: resize, view: func(layout.Size) string { return layout.Column(m.View(), help) }}, nil
}
