package collection

import (
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// ScrollPane binds an existing viewport without replacing its pointer, styles or key map.
// It owns dimensions and content; SetContent preserves and clamps the current offset.
type ScrollPane struct {
	Model         *viewport.Model
	Theme         theme.Theme
	ShowScrollbar bool
	size          layout.Size
}

func NewScrollPane(model *viewport.Model, size layout.Size, t theme.Theme) *ScrollPane {
	m := &ScrollPane{Model: model, Theme: t, ShowScrollbar: true}
	m.Resize(size)

	return m
}

func (m *ScrollPane) Resize(size layout.Size) {
	m.size = layout.Size{Width: max(0, size.Width), Height: max(0, size.Height)}
	m.Model.SetWidth(max(0, m.size.Width-m.scrollbarWidth()))
	m.Model.SetHeight(m.size.Height)
	m.Model.SetYOffset(m.Model.YOffset())
	m.Model.SetXOffset(m.Model.XOffset())
}

func (m *ScrollPane) scrollbarWidth() int {
	if m.ShowScrollbar && m.size.Width > 1 {
		return 1
	}

	return 0
}

func (m *ScrollPane) SetContent(content string) {
	offset := m.Model.YOffset()
	m.Model.SetContent(content)
	m.Model.SetYOffset(offset)
	m.Model.SetXOffset(m.Model.XOffset())
}

func (m *ScrollPane) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	*m.Model, cmd = m.Model.Update(msg)

	return cmd
}

func (m *ScrollPane) View() string {
	if m.size.Width == 0 || m.size.Height == 0 {
		return ""
	}

	out := m.Model.View()
	if m.scrollbarWidth() > 0 {
		bar := feedback.Scrollbar{Total: m.Model.TotalLineCount(), Visible: max(0, m.Model.Height()-m.Model.Style.GetVerticalFrameSize()), Offset: m.Model.YOffset(), Height: m.Model.Height()}
		out = lipgloss.JoinHorizontal(lipgloss.Top, out, bar.View(m.Theme))
	}

	return layout.Fit(out, m.size)
}
