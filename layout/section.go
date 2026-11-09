package layout

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Section is a panel that folds to a single heading row. Callers own the fold
// state and input bindings, and reallocate siblings using Height after a toggle.
// Title and Content accept styled text, as with Panel.
type Section struct {
	Title, Content     string
	Collapsed, NoFolds bool
	Symbols            text.Presentation
}

func (s Section) Height(expanded int) int {
	if s.Collapsed && !s.NoFolds {
		return min(1, max(0, expanded))
	}

	return max(0, expanded)
}

func (s Section) View(size Size, t theme.Theme, focused bool) string {
	title := s.Title
	if !s.NoFolds {
		if marker := text.Disclosure("", !s.Collapsed, t, s.Symbols); marker != "" {
			title = marker + " " + title
		}
	}

	if s.Collapsed && !s.NoFolds {
		style := t.Subtle()
		if focused {
			style = t.Heading()
		}

		return Fit(style.Render(title), Size{Width: size.Width, Height: s.Height(size.Height)})
	}

	return Panel(title, s.Content, size, t, focused)
}
