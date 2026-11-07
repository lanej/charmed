// Package input applies shared styles to concrete Bubbles input models.
package input

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/theme"
)

type TextOptions struct {
	Placeholder, Value string
	Width              int
	Disabled           bool
	Validate           textinput.ValidateFunc
	Theme              theme.Theme
}

// Text embeds its upstream model so applications retain the full native API.
type Text struct {
	textinput.Model
	Disabled bool
}

func NewText(o TextOptions) Text {
	m := Text{Model: textinput.New(), Disabled: o.Disabled}
	m.Prompt = "> "
	m.Placeholder, m.Validate = o.Placeholder, o.Validate
	m.SetValue(o.Value)
	m.SetWidth(max(1, o.Width-2))
	m.SetVirtualCursor(true)
	m.SetTheme(o.Theme)
	return m
}
func (m *Text) SetTheme(t theme.Theme) {
	s := textinput.DefaultStyles(t.IsDark())
	s.Focused.Text, s.Blurred.Text = t.Text(), t.Text()
	s.Focused.Prompt, s.Blurred.Prompt = t.Heading(), t.Subtle()
	s.Focused.Placeholder, s.Blurred.Placeholder = t.Subtle(), t.Subtle()
	s.Focused.Suggestion, s.Blurred.Suggestion = t.Subtle(), t.Subtle()
	s.Cursor.Color, s.Cursor.Blink = t.Accent, !t.ReducedMotion
	m.SetStyles(s)
}
func (m *Text) Focus() tea.Cmd {
	if m.Disabled {
		return nil
	}
	return m.Model.Focus()
}
func (m *Text) Resize(columns int) { m.SetWidth(max(1, columns-2)) }
func (m Text) Init() tea.Cmd       { return nil }
func (m Text) Update(msg tea.Msg) (Text, tea.Cmd) {
	if m.Disabled {
		return m, nil
	}
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

type Area struct{ textarea.Model }

func NewArea(width, height int, t theme.Theme) Area {
	m := Area{textarea.New()}
	m.SetWidth(max(1, width))
	m.SetHeight(max(1, height))
	m.SetVirtualCursor(true)
	s := textarea.DefaultStyles(t.IsDark())
	s.Focused.Text, s.Blurred.Text = t.Text(), t.Text()
	s.Focused.Prompt, s.Blurred.Prompt = t.Heading(), t.Subtle()
	s.Focused.Placeholder, s.Blurred.Placeholder = t.Subtle(), t.Subtle()
	s.Cursor.Color, s.Cursor.Blink = t.Accent, !t.ReducedMotion
	m.SetStyles(s)
	return m
}
func (m Area) Update(msg tea.Msg) (Area, tea.Cmd) {
	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

// Field is a molecule: one input, its label, help, and validation feedback.
type Field struct {
	Label, Help string
	Input       Text
	Theme       theme.Theme
}

func (f Field) View() string {
	s := f.Theme.Heading().Render(f.Label) + "\n" + f.Input.View()
	if f.Input.Err != nil {
		return s + "\n" + f.Theme.Text().Foreground(f.Theme.Danger).Render("! "+f.Input.Err.Error())
	}
	if f.Help != "" {
		s += "\n" + f.Theme.Subtle().Render(f.Help)
	}
	return s
}
