// Command review runs a file review, or an arbitrary document with -document.
package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/examples/review/fixture"
	"github.com/lanej/charmed/gitui"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/pattern"
	"github.com/lanej/charmed/theme"
)

type model struct {
	document *pattern.OutlineDocument
	review   *gitui.DiffReview
	palette  theme.Theme
}

func newModel(document bool) (*model, error) {
	palette := theme.New(theme.Dark)
	size := layout.Size{Width: 120, Height: 32}

	m := &model{palette: palette}
	if document {
		view, err := pattern.NewOutlineDocument(fixture.Document(), size, palette)
		if err != nil {
			return nil, err
		}

		m.document = view
	} else {
		view, err := gitui.NewFileDiffReview(fixture.Files(), size, palette)
		if err != nil {
			return nil, err
		}

		m.review, m.document = view, view.OutlineDocument
	}

	return m, nil
}

func (m *model) Init() tea.Cmd { return m.document.Init() }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if pressed, ok := msg.(tea.KeyPressMsg); ok {
		switch pressed.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab", "shift+tab":
			m.document.SetFocus(!m.document.DocumentFocused())
			return m, nil
		}
	}

	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.document.Resize(layout.Size{Width: size.Width, Height: size.Height})
		return m, nil
	}

	if m.review != nil {
		return m, m.review.Update(msg)
	}

	return m, m.document.Update(msg)
}

func (m *model) View() tea.View {
	view := tea.NewView(m.document.View())
	view.AltScreen = true
	view.MouseMode = tea.MouseModeCellMotion
	view.BackgroundColor = m.palette.Background

	return view
}

func run() error {
	document := flag.Bool("document", false, "Show an outline with non-file document content")

	flag.Parse()

	m, err := newModel(*document)
	if err != nil {
		return err
	}

	_, err = tea.NewProgram(m).Run()

	return err
}

func main() {
	if err := run(); err != nil {
		if _, writeErr := fmt.Fprintln(os.Stderr, err); writeErr != nil {
			os.Exit(1)
		}

		os.Exit(1)
	}
}
