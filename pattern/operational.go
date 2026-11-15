package pattern

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// DashboardSection is an unbordered, titled region. Body must be a pure renderer:
// it receives content columns and runs on SetSnapshot and Resize, never in View.
// Keep application actions, clocks and data fetching outside Body.
type DashboardSection struct {
	Title string
	Body  func(width int) string
}

// DashboardMeter shows an exact count against a supplied capacity. Label defines
// the quantity (for example "REST remaining"); fill means Value / Total, not health.
// Missing/redacted/invalid observations have no bar. Zero capacity and values
// above capacity retain their exact counts without implying a valid proportion.
type DashboardMeter struct {
	Label, Unit  string
	Value, Total uint64
	State        text.ValueState
}

// DashboardSnapshot orders a monitoring task: summary, primary evidence, then
// optional inspection. Status is supplied text, independent of health or freshness.
// Summary values use layout.Pair's explicit missing/invalid value states.
type DashboardSnapshot struct {
	Title, Context, Status string
	StatusTone             text.Tone
	Summary                []layout.Pair
	Meters                 []DashboardMeter
	Sections, Details      []DashboardSection
}

// DashboardKeys are the dashboard-owned actions. Native scrolling bindings live
// on Overview.KeyMap and Inspection.KeyMap; help reads those same bindings.
type DashboardKeys struct{ Details, Back, Help key.Binding }

// OperationalDashboard composes a single scrolling region with quiet section
// headings, persistent identity and one contextual help area. Details open in a
// separate view, preserving each view's scroll position. Only the focused model
// handles input; its parent owns application actions and terminal lifetime.
//
// Native viewport styles and bindings survive SetSnapshot and Resize. Those
// operations replace viewport content and dimensions, clamping offsets as needed.
// Call Resize after changing Keys or native key help to remeasure the footer.
type OperationalDashboard struct {
	Overview, Inspection viewport.Model
	Keys                 DashboardKeys
	// AppKeys adds help for bindings handled by the parent before forwarding
	// messages here (for example refresh or quit). The dashboard never runs them.
	AppKeys             []key.Binding
	Theme               theme.Theme
	snapshot            DashboardSnapshot
	size                layout.Size
	focused, inspecting bool
	fullHelp            bool
	header, footer      string
	bodyHeight          int
}

func NewOperationalDashboard(snapshot DashboardSnapshot, size layout.Size, t theme.Theme) *OperationalDashboard {
	m := &OperationalDashboard{
		Overview: collection.NewViewport("", 1, 1), Inspection: collection.NewViewport("", 1, 1),
		Theme: t, focused: true,
		Keys: DashboardKeys{
			Details: key.NewBinding(key.WithKeys("i"), key.WithHelp("i", "details")),
			Back:    key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "overview")),
			Help:    key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "keys")),
		},
	}
	m.size = size
	m.SetSnapshot(snapshot)

	return m
}

// SetSnapshot retains disclosure and scroll state, unless details disappear.
// Slices are copied; Body closures and their captured observations remain owned
// by the caller and must be updated on the same Bubble Tea loop as this model.
func (m *OperationalDashboard) SetSnapshot(snapshot DashboardSnapshot) {
	snapshot.Summary = append([]layout.Pair(nil), snapshot.Summary...)
	snapshot.Meters = append([]DashboardMeter(nil), snapshot.Meters...)
	snapshot.Sections = append([]DashboardSection(nil), snapshot.Sections...)
	snapshot.Details = append([]DashboardSection(nil), snapshot.Details...)

	m.snapshot = snapshot
	if len(snapshot.Details) == 0 {
		m.inspecting = false
	}

	m.reflow()
}

func (m *OperationalDashboard) Resize(size layout.Size) { m.size = size; m.reflow() }
func (m *OperationalDashboard) Focus()                  { m.focused = true; m.reflow() }
func (m *OperationalDashboard) Blur()                   { m.focused = false; m.reflow() }
func (m *OperationalDashboard) Inspecting() bool        { return m.inspecting }

func (m *OperationalDashboard) active() *viewport.Model {
	if m.inspecting {
		return &m.Inspection
	}

	return &m.Overview
}

func (m *OperationalDashboard) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	if press, ok := msg.(tea.KeyPressMsg); ok {
		switch {
		case key.Matches(press, m.Keys.Back) && m.inspecting:
			m.inspecting = false
			m.reflow()

			return nil
		case key.Matches(press, m.Keys.Details) && len(m.snapshot.Details) > 0:
			m.inspecting = !m.inspecting
			m.reflow()

			return nil
		case key.Matches(press, m.Keys.Help):
			m.fullHelp = !m.fullHelp
			m.reflow()

			return nil
		}
	}

	var cmd tea.Cmd

	view := m.active()
	*view, cmd = view.Update(msg)

	return cmd
}

func (m *OperationalDashboard) reflow() {
	width := max(1, m.size.Width-2*theme.InlinePadding)
	m.header = m.heading(width)
	m.footer = m.help(width)

	m.bodyHeight = max(0, m.size.Height-lipgloss.Height(m.header)-lipgloss.Height(m.footer)-2)
	for _, view := range []*viewport.Model{&m.Overview, &m.Inspection} {
		view.SetWidth(width)
		view.SetHeight(max(1, m.bodyHeight))
	}

	parts := []string{}
	if len(m.snapshot.Meters) > 0 {
		parts = append(parts, dashboardMeters(m.snapshot.Meters, width, m.Theme))
	}

	if len(m.snapshot.Summary) > 0 {
		parts = append(parts, layout.KeyValue(m.snapshot.Summary, width, m.Theme))
	}

	parts = append(parts, dashboardSections(m.snapshot.Sections, width, m.Theme)...)
	if len(parts) == 0 {
		parts = append(parts, m.Theme.Subtle().Render("No observations supplied"))
	}

	m.Overview.SetContent(strings.Join(parts, "\n\n"))
	m.Inspection.SetContent(strings.Join(dashboardSections(m.snapshot.Details, width, m.Theme), "\n\n"))
}

func (m *OperationalDashboard) heading(width int) string {
	label := text.Inline(m.snapshot.Title)
	if label == "" {
		label = "Operations"
	}

	if m.inspecting {
		label += " / Details"
	}

	style := m.Theme.Text()
	if m.focused {
		style = m.Theme.Heading()
	}

	title := style.Render(dashboardWrap(label, width))
	context := m.Theme.Subtle().Render(dashboardWrap(text.Inline(m.snapshot.Context), width))
	status := text.ToneStyle(m.snapshot.StatusTone, text.Normal, m.Theme).Render(dashboardWrap(text.Inline(m.snapshot.Status), width))

	left := title
	if m.snapshot.Context != "" {
		left += "  " + context
	}

	if !strings.Contains(left+status, "\n") && ansi.StringWidth(left)+ansi.StringWidth(status)+2 <= width {
		return dashboardHeaderRow(left, status, width)
	}

	if m.snapshot.Status != "" && !strings.Contains(title+status, "\n") && ansi.StringWidth(title)+ansi.StringWidth(status)+2 <= width {
		return dashboardHeaderRow(title, status, width) + dashboardContext(context, m.snapshot.Context)
	}

	parts := []string{title}
	if m.snapshot.Context != "" {
		parts = append(parts, context)
	}

	if m.snapshot.Status != "" {
		parts = append(parts, lipgloss.NewStyle().Width(width).Align(lipgloss.Right).Render(status))
	}

	return strings.Join(parts, "\n")
}

func dashboardHeaderRow(left, status string, width int) string {
	if status == "" {
		return left
	}

	return left + strings.Repeat(" ", max(2, width-ansi.StringWidth(left)-ansi.StringWidth(status))) + status
}

func dashboardContext(rendered, raw string) string {
	if raw == "" {
		return ""
	}

	return "\n" + rendered
}

func dashboardMeters(meters []DashboardMeter, width int, t theme.Theme) string {
	if width == 1 {
		return "!"
	}

	labels, values := make([]string, len(meters)), make([]string, len(meters))
	labelWidth, valueWidth := 0, 0

	for i, meter := range meters {
		labels[i] = text.Inline(meter.Label)

		values[i] = fmt.Sprintf("%d / %d", meter.Value, meter.Total)
		if meter.Unit != "" {
			values[i] += " " + text.Inline(meter.Unit)
		}

		if meter.State != text.ValuePresent {
			values[i] = meter.State.View("", text.Normal, t)
		} else if meter.Total == 0 {
			values[i] += " (no capacity)"
		} else if meter.Value > meter.Total {
			values[i] += " (invalid ratio)"
		}

		labelWidth = max(labelWidth, ansi.StringWidth(labels[i]))
		valueWidth = max(valueWidth, ansi.StringWidth(values[i]))
	}

	barWidth := min(28, width-labelWidth-valueWidth-4)
	rows := make([]string, 0, len(meters))

	for i, meter := range meters {
		label := labels[i] + strings.Repeat(" ", labelWidth-ansi.StringWidth(labels[i]))

		valid := meter.State == text.ValuePresent && meter.Total > 0 && meter.Value <= meter.Total
		if barWidth >= 8 && valid {
			bar := feedback.NewProgress(barWidth, t)
			bar.EmptyColor = t.Border
			rows = append(rows, label+"  "+bar.ViewAs(float64(meter.Value)/float64(meter.Total))+"  "+values[i])

			continue
		}

		rows = append(rows, dashboardWrap(labels[i]+"  "+values[i], width))
		if valid && width >= 8 {
			bar := feedback.NewProgress(min(28, width), t)
			bar.EmptyColor = t.Border
			rows = append(rows, bar.ViewAs(float64(meter.Value)/float64(meter.Total)))
		}
	}

	return strings.Join(rows, "\n")
}

func dashboardSections(sections []DashboardSection, width int, t theme.Theme) []string {
	parts := []string{}

	for _, section := range sections {
		if section.Body == nil {
			continue
		}

		body := dashboardWrap(section.Body(width), width)
		if section.Title != "" {
			body = t.Text().Bold(true).Render(dashboardWrap(text.Inline(section.Title), width)) + "\n" + body
		}

		parts = append(parts, body)
	}

	return parts
}

func dashboardWrap(content string, width int) string {
	return ansi.Hardwrap(ansi.Wrap(content, width, ""), width, true)
}

func (m *OperationalDashboard) help(width int) string {
	if !m.focused {
		return m.Theme.Subtle().Render("Inactive")
	}

	view := m.active()

	bindings := []key.Binding{view.KeyMap.Down, view.KeyMap.Up}
	if m.fullHelp {
		bindings = append(bindings, view.KeyMap.PageDown, view.KeyMap.PageUp, view.KeyMap.HalfPageDown, view.KeyMap.HalfPageUp)
	}

	if m.inspecting {
		bindings = append(bindings, m.Keys.Back)
	} else if len(m.snapshot.Details) > 0 {
		bindings = append(bindings, m.Keys.Details)
	}

	bindings = append(bindings, m.AppKeys...)
	bindings = append(bindings, m.Keys.Help)
	labels := []string{}

	for _, binding := range bindings {
		if binding.Enabled() {
			hint := binding.Help()
			labels = append(labels, text.KeyHint(hint.Key, hint.Desc, m.Theme))
		}
	}

	content := strings.Join(labels, "  ")
	if m.Theme.ASCII {
		content = layout.ASCII(strings.ReplaceAll(content, "½", "half"))
	}

	return ansi.Hardwrap(ansi.Wrap(content, width, ""), width, true)
}

func (m *OperationalDashboard) View() string {
	if m.size.Width <= 0 || m.size.Height <= 0 {
		return ""
	}

	if m.bodyHeight == 0 {
		return layout.Fit("Enlarge terminal to inspect observations", m.size)
	}

	view := m.active()
	position := fmt.Sprintf("%d-%d of %d", view.YOffset()+1, min(view.TotalLineCount(), view.YOffset()+m.bodyHeight), view.TotalLineCount())
	footer := m.Theme.Subtle().Render(position) + "\n" + m.footer
	content := m.header + "\n\n" + view.View() + "\n" + footer

	return layout.Fit(lipgloss.NewStyle().Padding(0, theme.InlinePadding).Render(content), m.size)
}
