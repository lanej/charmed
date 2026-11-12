// Package pattern contains composed, application-independent terminal patterns.
package pattern

import (
	"fmt"
	"math"
	"strings"
	"time"

	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/input"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type Service struct {
	ID, Name, Status string
	Rate, Latency    float64
}
type DashboardOptions struct {
	Theme      theme.Theme
	Services   []Service
	Live       bool
	ShowLegend bool
	Emit       func(any)
}

// Dashboard owns coordinated UI state; applications provide its data and events.
type Dashboard struct {
	Filter                                                        input.Text
	Table                                                         table.Model
	Details                                                       viewport.Model
	Plot                                                          chart.TimeSeries
	SelectedID                                                    string
	Step                                                          int
	services, visible                                             []Service
	focus                                                         interaction.Focus
	size                                                          layout.Size
	theme                                                         theme.Theme
	live                                                          bool
	showLegend                                                    bool
	emit                                                          func(any)
	leftWidth, rightWidth, tableHeight, chartHeight, detailHeight int
	split                                                         bool
}

func NewDashboard(o DashboardOptions) *Dashboard {
	m := &Dashboard{
		theme: o.Theme, services: append([]Service(nil), o.Services...), live: o.Live, emit: o.Emit, showLegend: o.ShowLegend,
		focus: interaction.NewFocus("filter", "table", "chart", "details"),
	}
	m.Filter = input.NewText(input.TextOptions{Placeholder: "Filter services…", Width: 80, Theme: o.Theme})
	if o.Theme.ASCII {
		m.Filter.Placeholder = "Filter services..."
	}
	m.Table = collection.NewTable(nil, nil, 35, 12, o.Theme)
	m.Details = collection.NewViewport("", 30, 8)
	m.Plot = chart.NewTimeSeries(nil, layout.Size{Width: 40, Height: 10}, o.Theme)
	m.Resize(layout.Size{Width: 120, Height: 34})
	m.applyFilter()
	return m
}

func (m *Dashboard) Init() tea.Cmd {
	var pulse tea.Cmd
	if m.live && !m.theme.ReducedMotion {
		pulse = m.nextTick()
	}
	return tea.Batch(m.syncFocus(), pulse)
}

type dashboardTick struct{}

func (m *Dashboard) nextTick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg { return dashboardTick{} })
}

func (m *Dashboard) Resize(size layout.Size) {
	m.size = size
	m.Filter.Resize(max(1, layout.PanelContentSize(size).Width))
	mainHeight := max(1, size.Height-5)
	m.leftWidth, m.rightWidth, m.split = layout.Split(size.Width, 1, 30, 48, 0.36)

	m.split = m.split && size.Height >= 24
	if m.split {
		m.tableHeight = min(mainHeight, max(3, len(m.services)+2))
		m.chartHeight = mainHeight
		m.detailHeight = mainHeight - m.tableHeight
	} else {
		m.leftWidth, m.rightWidth = size.Width, size.Width
		m.tableHeight = min(7, max(3, mainHeight/3))
		m.chartHeight = max(1, mainHeight-m.tableHeight)
		m.detailHeight = 0
	}
	current := m.focus.Current()
	ids := []string{"filter", "table"}
	if m.chartHeight >= 8 {
		ids = append(ids, "chart")
	}

	if m.detailHeight >= 3 {
		ids = append(ids, "details")
	}

	m.focus = interaction.NewFocus(ids...)
	m.focus.Set(current)

	const columnCount = 2

	paddingWidth := collection.TablePaddingWidth(columnCount)
	w := max(paddingWidth, layout.PanelContentSize(layout.Size{Width: m.leftWidth}).Width)

	statusWidth := ansi.StringWidth("STATE")
	for _, service := range m.services {
		statusWidth = max(statusWidth, ansi.StringWidth(serviceStatus(service.Status).View(text.Bright, m.theme)))
	}

	m.Table.SetColumns([]table.Column{{Title: "SERVICE", Width: max(1, w-statusWidth-paddingWidth)}, {Title: "STATE", Width: statusWidth}})
	m.Table.SetWidth(w)
	m.Table.SetHeight(max(1, m.tableHeight-1))
	m.Details.SetWidth(w)
	m.Details.SetHeight(max(1, layout.PanelContentSize(layout.Size{Height: m.detailHeight}).Height))
	m.Plot.Resize(chart.PanelPlotSize(layout.Size{Width: m.rightWidth, Height: m.chartHeight}, chart.PanelOptions{ShowLegend: m.showLegend}))
}

func (m *Dashboard) syncFocus() tea.Cmd {
	m.Filter.Blur()
	m.Table.Blur()
	m.Plot.Blur()
	switch m.focus.Current() {
	case "filter":
		return m.Filter.Focus()
	case "table":
		m.Table.Focus()
	case "chart":
		m.Plot.Focus()
	}
	return nil
}

func (m *Dashboard) Update(msg tea.Msg) tea.Cmd {
	if m.focus.Handle(msg) {
		return m.syncFocus()
	}
	switch msg.(type) {
	case dashboardTick:
		m.advance()
		if m.live && !m.theme.ReducedMotion {
			return m.nextTick()
		}
		return nil
	case runtime.StepMsg:
		m.advance()
		return nil
	}
	if _, key := msg.(tea.KeyPressMsg); key {
		switch m.focus.Current() {
		case "filter":
			before := m.Filter.Value()
			var cmd tea.Cmd
			m.Filter, cmd = m.Filter.Update(msg)
			if before != m.Filter.Value() {
				m.applyFilter()
			}
			return cmd
		case "table":
			var cmd tea.Cmd
			m.Table, cmd = m.Table.Update(msg)
			m.selectRow()
			return cmd
		case "chart":
			var cmd tea.Cmd
			m.Plot, cmd = m.Plot.Update(msg)
			return cmd
		case "details":
			var cmd tea.Cmd
			m.Details, cmd = m.Details.Update(msg)
			return cmd
		}
		return nil
	}
	// Cursor and asynchronous child messages still reach their owner when unfocused.
	var cmd tea.Cmd
	m.Filter, cmd = m.Filter.Update(msg)
	return cmd
}

func (m *Dashboard) advance() {
	m.Step++
	m.updateSelection()
	if m.emit != nil {
		m.emit(map[string]any{"type": "tick", "step": m.Step, "selected": m.SelectedID})
	}
}

func (m *Dashboard) applyFilter() {
	needle := strings.ToLower(m.Filter.Value())
	m.visible = nil
	var rows []table.Row
	selected := 0
	for _, s := range m.services {
		if strings.Contains(strings.ToLower(s.Name+" "+s.Status), needle) {
			if s.ID == m.SelectedID {
				selected = len(rows)
			}
			m.visible = append(m.visible, s)
			rows = append(rows, table.Row{s.Name, serviceStatus(s.Status).View(text.Bright, m.theme)})
		}
	}
	m.Table.SetRows(rows)
	if len(rows) > 0 {
		m.Table.SetCursor(selected)
	}
	m.selectRow()
}

func (m *Dashboard) selectRow() {
	old := m.SelectedID
	m.SelectedID = ""
	if i := m.Table.Cursor(); i >= 0 && i < len(m.visible) {
		m.SelectedID = m.visible[i].ID
	}
	m.updateSelection()
	if old != m.SelectedID && m.emit != nil {
		m.emit(interaction.Selected{Component: "services", ID: m.SelectedID})
	}
}

func (m *Dashboard) selected() (Service, bool) {
	for _, s := range m.visible {
		if s.ID == m.SelectedID {
			return s, true
		}
	}
	return Service{}, false
}

func (m *Dashboard) updateSelection() {
	s, ok := m.selected()
	if !ok {
		m.Plot.SetData(nil)
		m.Details.SetContent("No services match your filter.")
		return
	}
	m.Plot.SetData(m.series(s))
	m.Details.SetContent(fmt.Sprintf("%s  %s\n\nThroughput   %.0f req/min\nLatency p95  %.0f ms", s.Name, serviceStatus(s.Status).View(text.Bright, m.theme), s.Rate, s.Latency))
}

func (m *Dashboard) series(s Service) []chart.TimeSeriesData {
	base := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	data := []chart.TimeSeriesData{{ID: "observed", Label: "Observed", Unit: "req/min"}, {ID: "baseline", Label: "Baseline", Unit: "req/min"}}
	for i := range 32 {
		t := base.Add(time.Duration(i) * time.Minute)
		phase := float64(i+m.Step)*0.27 + float64(len(s.ID))
		data[0].Points = append(data[0].Points, timeserieslinechart.TimePoint{Time: t, Value: s.Rate * (0.9 + 0.15*math.Sin(phase))})
		data[1].Points = append(data[1].Points, timeserieslinechart.TimePoint{Time: t, Value: s.Rate * 0.94})
	}
	return data
}

func (m *Dashboard) View() string {
	header := m.theme.Heading().Render("Service workbench") + "  " + m.theme.Subtle().Render(fmt.Sprintf("%d services · %d samples", len(m.visible), m.Step))
	filter := layout.Panel("", m.Filter.View(), layout.Size{Width: m.size.Width, Height: 3}, m.theme, m.focus.Current() == "filter")
	// WHY: Names and status pills are self-describing; retain native table navigation without its column-heading row.
	_, rows, _ := strings.Cut(m.Table.View(), "\n")
	left := layout.Panel("", rows, layout.Size{Width: m.leftWidth, Height: m.tableHeight}, m.theme, m.focus.Current() == "table")
	title := "Throughput / req per minute"
	data := []chart.TimeSeriesData{}
	if s, ok := m.selected(); ok {
		title = s.Name + " / throughput"
		data = m.series(s)
	}
	plot := chart.Panel(title, m.Plot, data, layout.Size{Width: m.rightWidth, Height: m.chartHeight}, m.theme, m.focus.Current() == "chart", chart.PanelOptions{ShowLegend: m.showLegend})
	main := ""
	if m.split {
		details := layout.Panel("", m.Details.View(), layout.Size{Width: m.leftWidth, Height: m.detailHeight}, m.theme, m.focus.Current() == "details")
		main = layout.Row(layout.Column(left, details), strings.Repeat(" \n", max(0, m.chartHeight-1))+" ", plot)
	} else {
		main = layout.Column(left, plot)
	}
	footer := m.theme.Subtle().Render("Tab focus   arrows navigate   Ctrl+N step   Ctrl+C quit")
	if m.theme.ASCII {
		header = strings.ReplaceAll(header, "·", "/")
	}
	return layout.Fit(layout.Column(header, filter, main, footer), m.size)
}

func SampleServices() []Service {
	return []Service{
		{ID: "api", Name: "API gateway", Status: "Healthy", Rate: 1240, Latency: 78},
		{ID: "events", Name: "Event stream", Status: "Healthy", Rate: 860, Latency: 42},
		{ID: "workers", Name: "Job workers", Status: "Degraded", Rate: 420, Latency: 198},
		{ID: "storage", Name: "Storage", Status: "Healthy", Rate: 640, Latency: 24},
		{ID: "scheduler", Name: "Scheduler", Status: "Healthy", Rate: 180, Latency: 16},
	}
}

// serviceStatus maps the example's service health vocabulary to shared display atoms.
func serviceStatus(value string) feedback.Status {
	switch strings.ToLower(value) {
	case "healthy", "ok":
		return feedback.OK
	case "degraded", "warn":
		return feedback.Warn
	case "failed", "error":
		return feedback.Error
	case "off":
		return feedback.Off
	case "disabled":
		return feedback.Disabled
	case "pending":
		return feedback.Pending
	case "running":
		return feedback.Running
	case "paused":
		return feedback.Paused
	default:
		return feedback.Unknown
	}
}
