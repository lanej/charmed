// Package input applies shared styles to concrete Bubbles input models.
package input

import (
	"strings"

	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
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
	palette  theme.Theme
}

func NewText(o TextOptions) Text {
	m := Text{Model: textinput.New(), Disabled: o.Disabled}
	m.Prompt = defaultPrompt
	m.Placeholder, m.Validate = o.Placeholder, o.Validate
	m.SetValue(o.Value)
	m.Resize(o.Width)
	m.SetVirtualCursor(true)
	m.SetTheme(o.Theme)
	return m
}

func (m *Text) SetTheme(t theme.Theme) {
	m.palette = t
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

const defaultPrompt = "> "

// View applies disabled presentation to a copy, retaining native customization.
func (m Text) View() string {
	native := m.Model
	if m.Disabled {
		native.Blur()

		if native.Prompt == defaultPrompt {
			native.Prompt = m.palette.Glyph(theme.GlyphDisabled) + " "
		}

		styles := native.Styles()

		dim := m.palette.DimInactive
		if dim == nil {
			dim = m.palette.Muted
		}

		if dim != nil {
			styles.Blurred.Text = styles.Blurred.Text.Foreground(dim)
			styles.Blurred.Prompt = styles.Blurred.Prompt.Foreground(dim)
			styles.Blurred.Placeholder = styles.Blurred.Placeholder.Foreground(dim)
		}

		native.SetStyles(styles)
	}

	return native.View()
}

// Cursor suppresses the native real cursor while the wrapper is disabled.
func (m Text) Cursor() *tea.Cursor {
	if m.Disabled {
		return nil
	}

	return m.Model.Cursor()
}

func (m Text) Focused() bool { return !m.Disabled && m.Model.Focused() }

// Resize includes the native prompt and the extra cell reserved for its cursor.
func (m *Text) Resize(columns int) {
	m.SetWidth(max(1, columns-ansi.StringWidth(m.Prompt)-1))
	position := m.Position()
	// WHY: Native overflow handling does not shrink a viewport containing the cursor.
	m.CursorEnd()
	m.SetCursor(position)
}

func (m Text) Init() tea.Cmd { return nil }
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
	s.Focused.CursorLine = t.Text().Background(t.Surface)
	s.Blurred.CursorLine = t.Text()
	s.Focused.LineNumber, s.Blurred.LineNumber = t.Subtle(), t.Subtle()
	s.Focused.CursorLineNumber, s.Blurred.CursorLineNumber = t.Heading(), t.Subtle()
	s.Focused.EndOfBuffer, s.Blurred.EndOfBuffer = t.Subtle(), t.Subtle()
	s.Focused.Selection, s.Blurred.Selection = t.Selection().Padding(0), t.Selection().Padding(0)
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
	control := f.Input.View()
	width := max(1, lipgloss.Width(control))

	var lines []string

	if f.Label != "" {
		lines = append(lines, f.Theme.Heading().Render(ansi.Wrap(f.Label, width, "")))
	}

	lines = append(lines, control)

	if f.Input.Err != nil {
		message := f.Theme.Glyph(theme.GlyphError) + " " + f.Input.Err.Error()
		lines = append(lines, f.Theme.Text().Foreground(f.Theme.Danger).Render(ansi.Wrap(message, width, "")))
	} else if f.Help != "" {
		lines = append(lines, f.Theme.Subtle().Render(ansi.Wrap(f.Help, width, "")))
	}

	return strings.Join(lines, "\n")
}
