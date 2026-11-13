package story

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

// scene is only a fixture adapter: production widgets keep their concrete APIs.
type scene struct {
	size   layout.Size
	init   func() tea.Cmd
	update func(tea.Msg) tea.Cmd
	resize func(layout.Size)
	view   func(layout.Size) string
	close  func() error
}

func (s *scene) Init() tea.Cmd {
	if s.init != nil {
		return s.init()
	}

	return nil
}

func (s *scene) Update(msg tea.Msg) tea.Cmd {
	if s.update != nil {
		return s.update(msg)
	}

	return nil
}

func (s *scene) Resize(size layout.Size) {
	s.size = size
	if s.resize != nil {
		s.resize(size)
	}
}
func (s *scene) View() string { return layout.Fit(s.view(s.size), s.size) }
func (s *scene) Close() error {
	if s.close != nil {
		return s.close()
	}

	return nil
}

// Static fixtures depend only on their fixed inputs and the current allocation.
func static(view func(layout.Size) string) runtime.Scene {
	var (
		previous layout.Size
		frame    string
	)

	valid := false

	return &scene{view: func(size layout.Size) string {
		if !valid || previous != size {
			frame = view(size)
			previous, valid = size, true
		}

		return frame
	}}
}

func emit(e Environment, value any) {
	if e.Emit != nil {
		e.Emit(value)
	}
}

func stateMessage(e Environment) string {
	switch e.Config.State {
	case "empty":
		return "No data yet. Change the fixture state to ready."
	case "loading":
		return "Loading fixture data..."
	case "error":
		return "! Fixture error: service unavailable. Press Reset to try again."
	}

	return ""
}

func title(value string, e Environment) string {
	return e.Config.Palette().Heading().Render(value) + "\n\n"
}

func footer(value string, e Environment) string {
	return "\n\n" + e.Config.Palette().Subtle().Render(value)
}
func labelCounter(label string, n int) string { return fmt.Sprintf("%s %d", label, n) }
