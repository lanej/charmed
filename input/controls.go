package input

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Control holds presentation and activation bindings shared by single controls.
// Parents route messages and own traversal. A nonempty DisabledReason also disables activation.
type Control struct {
	ID, Label      string
	Disabled       bool
	DisabledReason string
	Keys           key.Binding
	Theme          theme.Theme
	focused        bool
}

func newControl(id, label string, t theme.Theme) Control {
	return Control{ID: id, Label: label, Theme: t, Keys: key.NewBinding(key.WithKeys("space", "enter"), key.WithHelp("space/enter", "activate"))}
}

func (c *Control) Focus()            { c.focused = true }
func (c *Control) Blur()             { c.focused = false }
func (c *Control) Focused() bool     { return c.focused && !c.unavailable() }
func (c *Control) unavailable() bool { return c.Disabled || c.DisabledReason != "" }
func (c *Control) activates(msg tea.Msg) bool {
	press, ok := msg.(tea.KeyPressMsg)
	return ok && c.Focused() && key.Matches(press, c.Keys)
}

func (c *Control) view(mark string, busy bool) string {
	label := text.Inline(c.Label)
	if mark != "" {
		label = mark + " " + label
	}

	if busy {
		label += " [BUSY]"
	}

	if c.unavailable() {
		label += " [DISABLED]"
		if c.DisabledReason != "" {
			label += " " + text.Inline(c.DisabledReason)
		}

		return text.ToneStyle(text.Neutral, text.Dim, c.Theme).Padding(0, theme.InlinePadding).Render(label)
	}

	if c.focused {
		return c.Theme.Focus().Render(label)
	}

	return c.Theme.Control().Render(label)
}

type CheckState uint8

const (
	Unchecked CheckState = iota
	Checked
	Mixed
)

// CheckChanged is emitted after a checkbox updates its local state.
type (
	CheckChanged struct {
		ID    string
		State CheckState
	}
	RadioSelected struct{ ID string }
	Toggled       struct {
		ID string
		On bool
	}
	Activated struct{ ID string }
)

type Checkbox struct {
	Control
	State CheckState
}

func NewCheckbox(id, label string, state CheckState, t theme.Theme) Checkbox {
	return Checkbox{Control: newControl(id, label, t), State: state}
}

// Mixed activates to Checked. Invalid states remain visible and cannot activate.
func (m Checkbox) Update(msg tea.Msg) (Checkbox, tea.Cmd) {
	if !m.activates(msg) {
		return m, nil
	}

	switch m.State {
	case Unchecked, Mixed:
		m.State = Checked
	case Checked:
		m.State = Unchecked
	default:
		return m, nil
	}

	event := CheckChanged{ID: m.ID, State: m.State}

	return m, func() tea.Msg { return event }
}

func (m Checkbox) View() string {
	mark := "[?]"

	switch m.State {
	case Unchecked:
		mark = "[ ]"
	case Checked:
		mark = "[x]"
	case Mixed:
		mark = "[-]"
	default:
	}

	return m.view(mark, false)
}

// Radio owns only its selection. The parent deselects other radios in the group.
type Radio struct {
	Control
	Selected bool
}

func NewRadio(id, label string, selected bool, t theme.Theme) Radio {
	return Radio{Control: newControl(id, label, t), Selected: selected}
}

func (m Radio) Update(msg tea.Msg) (Radio, tea.Cmd) {
	if !m.activates(msg) || m.Selected {
		return m, nil
	}

	m.Selected = true

	return m, func() tea.Msg { return RadioSelected{ID: m.ID} }
}

func (m Radio) View() string {
	mark := "( )"
	if m.Selected {
		mark = "(*)"
	}

	return m.view(mark, false)
}

type Toggle struct {
	Control
	On bool
}

func NewToggle(id, label string, on bool, t theme.Theme) Toggle {
	return Toggle{Control: newControl(id, label, t), On: on}
}

func (m Toggle) Update(msg tea.Msg) (Toggle, tea.Cmd) {
	if !m.activates(msg) {
		return m, nil
	}

	m.On = !m.On

	return m, func() tea.Msg { return Toggled{ID: m.ID, On: m.On} }
}

func (m Toggle) View() string {
	mark := "[OFF]"
	if m.On {
		mark = "[ON ]"
	}

	return m.view(mark, false)
}

// Button emits intent without running the action or changing Busy automatically.
// The parent sets Busy while an operation is outstanding and supplies its result.
type Button struct {
	Control
	Busy bool
}

func NewButton(id, label string, t theme.Theme) Button {
	return Button{Control: newControl(id, label, t)}
}

func (m Button) Update(msg tea.Msg) (Button, tea.Cmd) {
	if m.Busy || !m.activates(msg) {
		return m, nil
	}

	return m, func() tea.Msg { return Activated{ID: m.ID} }
}

func (m Button) View() string {
	c := m.Control
	c.Label = "[ " + c.Label + " ]"

	return c.view("", m.Busy)
}
