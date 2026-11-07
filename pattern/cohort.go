package pattern

import (
	"fmt"
	"maps"
	"math"
	"sort"
	"strings"
	"time"

	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/interaction"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type CohortComparisonOptions struct {
	Title, Scope string
	// Data orders the reference cohort first and candidate second. IDs retain
	// their application identity; display aliases and roles do not replace them.
	Data        cohort.Data
	Theme       theme.Theme
	MinRequests uint64
	Sigma       int
	Emit        func(any)
	MetricTones map[string]text.Tone
	Boundaries  map[string]cohort.ComparisonPolicy
}

// CohortComparison links per-metric distributions, request-weighted histories,
// and a member inspector. It owns UI state; callers supply interval counts.
// Sigma is descriptive member spread, never a significance or rollout decision.
type CohortComparison struct {
	Table                        table.Model
	Panels                       viewport.Model
	Inspector                    table.Model
	data                         cohort.Data
	title, scope                 string
	theme                        theme.Theme
	minRequests                  uint64
	sigma, window, metric        int
	selected                     cohort.MemberKey
	visible                      []cohort.MemberRate
	summaries                    [][2]cohort.Summary
	plots                        []chart.TimeSeries
	focus                        interaction.Focus
	size                         layout.Size
	gridSize, sideSize           layout.Size
	tableHeight, inspectorHeight int
	sidebar                      bool
	emit                         func(any)
	metricTones                  map[string]text.Tone
	boundaries                   map[string]cohort.ComparisonPolicy
	assessments                  []cohort.ComparisonAttention
	attentionLevels              map[string]attention.Level
	initialized                  bool
	deltaExtent                  float64
}

func NewCohortComparison(o CohortComparisonOptions) (*CohortComparison, error) {
	if o.Sigma == 0 {
		o.Sigma = 3
	}
	if o.Sigma < 1 || o.Sigma > 3 {
		return nil, fmt.Errorf("sigma must be 1, 2 or 3")
	}
	m := &CohortComparison{title: o.Title, scope: o.Scope, theme: o.Theme, minRequests: max(1, o.MinRequests), sigma: o.Sigma, emit: o.Emit,
		focus: interaction.NewFocus("hosts", "metrics", "inspector"), window: -1, metricTones: maps.Clone(o.MetricTones)}
	m.Table = collection.NewTable([]table.Column{{Title: "MEMBER", Width: 22}, {Title: "ROLE", Width: 4}, {Title: "z(ref)", Width: 6}}, nil, 38, 14, o.Theme)
	m.Panels = collection.NewViewport("", 100, 36)
	m.Inspector = collection.NewDenseTable(nil, nil, 38, 4, o.Theme)
	m.size = layout.Size{Width: 132, Height: 40}
	if err := m.SetData(o.Data); err != nil {
		return nil, err
	}
	if err := m.SetBoundaries(o.Boundaries); err != nil {
		return nil, err
	}
	m.Resize(m.size)
	return m, nil
}

func (m *CohortComparison) Selected() cohort.MemberKey { return m.selected }
func (m *CohortComparison) WindowIndex() int           { return m.window }
func (m *CohortComparison) Sigma() int                 { return m.sigma }
func (m *CohortComparison) MetricID() string           { return m.data.Metrics[m.metric].ID }

// SetData is atomic on validation failure. Appends follow the latest window only
// if it was already selected. An inspected historical window stays selected.
func (m *CohortComparison) SetData(data cohort.Data) error {
	copy, err := data.Clone()
	if err != nil {
		return err
	}
	if len(copy.Cohorts) != 2 {
		return fmt.Errorf("comparison needs exactly two cohorts")
	}
	if _, err := copyPolicies(m.boundaries, copy.Metrics); err != nil {
		return err
	}
	index := len(copy.Windows) - 1
	if m.window >= 0 && m.window < len(m.data.Windows)-1 {
		start := m.data.Windows[m.window].Start
		for i, w := range copy.Windows {
			if w.Start.Equal(start) {
				index = i
				break
			}
		}
	}
	metricID := ""
	if len(m.data.Metrics) > 0 {
		metricID = m.data.Metrics[m.metric].ID
	}
	m.data, m.window, m.metric = copy, index, len(copy.Metrics)-1
	for i, metric := range copy.Metrics {
		if metric.ID == metricID {
			m.metric = i
		}
	}
	m.refresh()
	return nil
}
func (m *CohortComparison) SetWindow(index int) error {
	if index < 0 || index >= len(m.data.Windows) {
		return fmt.Errorf("window index out of range")
	}
	m.window = index
	m.refresh()
	m.event("window")
	return nil
}
func (m *CohortComparison) SetSigma(sigma int) error {
	if sigma < 1 || sigma > 3 {
		return fmt.Errorf("sigma must be 1, 2 or 3")
	}
	m.sigma = sigma
	m.refresh()
	m.event("sigma")
	return nil
}
func (m *CohortComparison) SetMetric(id string) error {
	for i, metric := range m.data.Metrics {
		if metric.ID == id {
			m.metric = i
			m.refresh()
			m.event("metric")
			return nil
		}
	}
	return fmt.Errorf("unknown metric %q", id)
}
func (m *CohortComparison) Init() tea.Cmd { m.initialized = true; return m.syncFocus() }
func (m *CohortComparison) syncFocus() tea.Cmd {
	m.Table.Blur()
	m.Inspector.Blur()
	if m.focus.Current() == "hosts" {
		m.Table.Focus()
	}
	if m.focus.Current() == "inspector" {
		m.Inspector.Focus()
	}
	return nil
}
func (m *CohortComparison) Update(msg tea.Msg) tea.Cmd {
	if m.focus.Handle(msg) {
		return m.syncFocus()
	}
	if _, ok := msg.(runtime.StepMsg); ok {
		if len(m.data.Windows) > 0 {
			_ = m.SetWindow((m.window + 1) % len(m.data.Windows))
		}
		return nil
	}
	if key, ok := msg.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "1", "2", "3":
			_ = m.SetSigma(int(key.String()[0] - '0'))
			return nil
		case "left":
			if m.window > 0 {
				_ = m.SetWindow(m.window - 1)
			}
			return nil
		case "right":
			if m.window < len(m.data.Windows)-1 {
				_ = m.SetWindow(m.window + 1)
			}
			return nil
		case "[":
			m.metric = (m.metric + len(m.data.Metrics) - 1) % len(m.data.Metrics)
			m.refresh()
			m.event("metric")
			return nil
		case "]":
			m.metric = (m.metric + 1) % len(m.data.Metrics)
			m.refresh()
			m.event("metric")
			return nil
		}
	}
	var cmd tea.Cmd
	switch m.focus.Current() {
	case "hosts":
		m.Table, cmd = m.Table.Update(msg)
		if i := m.Table.Cursor(); i >= 0 && i < len(m.visible) && m.selected != m.visible[i].MemberKey {
			m.selected = m.visible[i].MemberKey
			m.updateDetails()
			m.event("selected")
		}
	case "metrics":
		m.Panels, cmd = m.Panels.Update(msg)
	case "inspector":
		m.Inspector, cmd = m.Inspector.Update(msg)
		if i := m.Inspector.Cursor(); i >= 0 && i < len(m.Inspector.Rows()) && i != m.metric {
			m.metric = i
			m.refresh()
			m.event("metric")
		}
	}
	return cmd
}
func (m *CohortComparison) event(kind string) {
	if m.emit != nil {
		m.emit(map[string]any{"type": kind, "cohort": m.selected.CohortID, "member": m.selected.MemberID,
			"reference": m.data.Cohorts[0].ID, "candidate": m.data.Cohorts[1].ID,
			"window": m.window, "metric": m.MetricID(), "sigma": m.sigma})
	}
}
func (m *CohortComparison) Resize(size layout.Size) {
	m.size = size
	height := max(1, size.Height-3)
	m.sidebar = size.Width >= 124 && size.Height >= 28
	if m.sidebar {
		m.sideSize = layout.Size{Width: 40, Height: height}
		m.gridSize = layout.Size{Width: max(1, size.Width-41), Height: height}
	} else {
		sideHeight := min(15, max(6, height/2))
		m.sideSize = layout.Size{Width: size.Width, Height: sideHeight}
		m.gridSize = layout.Size{Width: size.Width, Height: max(1, height-sideHeight)}
	}
	m.inspectorHeight = min(10, max(3, m.sideSize.Height-3))
	m.tableHeight = max(3, m.sideSize.Height-m.inspectorHeight)
	m.Table.SetWidth(max(1, m.sideSize.Width-2))
	m.Table.SetHeight(max(1, m.tableHeight-3))
	m.tableColumns()
	m.Panels.SetWidth(max(1, m.gridSize.Width))
	m.Panels.SetHeight(max(1, m.gridSize.Height))
	m.updateDetails()
}
func (m *CohortComparison) refresh() {
	m.summaries = make([][2]cohort.Summary, len(m.data.Metrics))
	m.visible = nil
	if m.window >= 0 && m.window < len(m.data.Windows) {
		for i, metric := range m.data.Metrics {
			for c, group := range m.data.Cohorts {
				m.summaries[i][c] = cohort.Analyze(m.data.Windows[m.window].Samples, group.ID, metric.ID, m.minRequests)
			}
		}
		for _, summary := range m.summaries[m.metric] {
			m.visible = append(m.visible, summary.Points...)
		}
	}
	changes := m.assessAttention()
	m.updateDeltaExtent()
	baseline := m.summaries[m.metric][0]
	sort.SliceStable(m.visible, func(i, j int) bool {
		a, b := m.visible[i], m.visible[j]
		pa, pb := attention.Priority(m.assessments[m.metric].Members[a.MemberKey].Level), attention.Priority(m.assessments[m.metric].Members[b.MemberKey].Level)
		if pa != pb {
			return pa > pb
		}
		if a.Eligible != b.Eligible {
			return a.Eligible
		}
		da, db := math.Abs(a.Rate-baseline.Mean), math.Abs(b.Rate-baseline.Mean)
		if da != db {
			return da > db
		}
		if a.CohortID != b.CohortID {
			return a.CohortID < b.CohortID
		}
		return a.MemberID < b.MemberID
	})
	rows, cursor := []table.Row{}, -1
	for i, p := range m.visible {
		if p.MemberKey == m.selected {
			cursor = i
		}
		z := "n/a"
		if distance, ok := cohort.Distance(p, baseline); ok {
			z = fmt.Sprintf("%+.1f", distance)
		}
		row := table.Row{p.MemberID, p.CohortID, z}
		if m.boundaries[m.MetricID()].Hosts.Configured() {
			row = append(row, feedback.Attention(m.assessments[m.metric].Members[p.MemberKey].Level, m.theme))
		}
		rows = append(rows, row)
	}
	if cursor < 0 && len(m.visible) > 0 {
		cursor = 0
		m.selected = m.visible[0].MemberKey
	}
	if len(m.visible) == 0 {
		m.selected = cohort.MemberKey{}
	}
	m.Table.SetRows(nil)
	m.tableColumns()
	m.Table.SetRows(rows)
	if cursor >= 0 {
		m.Table.SetCursor(cursor)
	}
	m.updateDetails()
	for _, change := range changes {
		m.emit(change)
	}
}
func (m *CohortComparison) updateDetails() {
	m.updateInspector()
	m.rebuildPlots()
	m.Panels.SetContent(m.grid())
}
func (m *CohortComparison) rebuildPlots() {
	columns := 1
	if m.gridSize.Width >= 82 {
		columns = 2
	}
	width := max(20, (m.gridSize.Width-(columns-1))/columns-2)
	m.plots = make([]chart.TimeSeries, len(m.data.Metrics))
	for i, metric := range m.data.Metrics {
		series := []chart.TimeSeriesData{{ID: "cohort:0"}, {ID: "cohort:1"}, {ID: "selected"}}
		for _, w := range m.data.Windows {
			for c, group := range m.data.Cohorts {
				s := cohort.Analyze(w.Samples, group.ID, metric.ID, m.minRequests)
				if s.HasRate {
					series[c].Points = append(series[c].Points, timeserieslinechart.TimePoint{Time: w.End, Value: s.Rate * 100})
				}
			}
			for _, s := range w.Samples {
				if s.MemberKey == m.selected && s.Requests > 0 {
					series[2].Points = append(series[2].Points, timeserieslinechart.TimePoint{Time: w.End, Value: float64(s.Counts[metric.ID]) / float64(s.Requests) * 100})
				}
			}
		}
		for c, group := range m.data.Cohorts {
			series[c].Label = group.ID
			series[c].Unit = "%"
		}
		series[2].Label, series[2].Unit = "member", "%"
		low, high := math.Inf(1), math.Inf(-1)
		for _, s := range series {
			for _, p := range s.Points {
				low = min(low, p.Value)
				high = max(high, p.Value)
			}
		}
		if math.IsInf(low, 0) {
			low, high = 0, 1
		}
		padding := max((high-low)*0.15, 0.001)
		m.plots[i] = chart.NewTimeSeries(series, layout.Size{Width: width, Height: 9}, m.theme,
			timeserieslinechart.WithYRange(max(0, low-padding), min(100, high+padding)),
			timeserieslinechart.WithYLabelFormatter(func(_ int, value float64) string { return fmt.Sprintf("%.2f%%", value) }),
			timeserieslinechart.WithXLabelFormatter(func(_ int, value float64) string {
				return time.UnixMilli(int64(math.Round(value * 1000))).UTC().Format("15:04")
			}))
	}
}
func (m *CohortComparison) metricLabel(index int, compact bool) string {
	metric := m.data.Metrics[index]
	value := metric.Label
	if compact {
		value = metric.ID
	}
	if tone, ok := m.metricTones[metric.ID]; ok {
		return text.Pill(value, tone, m.theme)
	}
	return value
}

func (m *CohortComparison) panel(index, width int) string {
	summaries := m.summaries[index]
	state := m.assessments[index]
	distribution := chart.DistributionComparison{Sigma: m.sigma, Min: 0, Max: 100, Unit: "%", SpreadUnit: "pp", Markers: m.boundaryMarkers(index)}
	for c, s := range summaries {
		series := chart.DistributionSeries{Label: m.data.Cohorts[c].ID, Mean: s.Mean * 100, StdDev: s.StdDev * 100, Members: s.Eligible, HasSpread: s.HasSpread}
		for _, point := range s.Points {
			if point.MemberKey == m.selected && point.HasRate {
				value := point.Rate * 100
				series.Selected = &value
			}
		}
		distribution.Series = append(distribution.Series, series)
	}
	delta, available := m.delta(index)
	style := feedback.AttentionStyle(state.Fleet.Level, m.theme)
	titleLabel := m.metricLabel(index, true)
	propertyWidth := max(0, width-4-lipgloss.Width(titleLabel)-9)
	property := chart.DeltaProperty{Value: delta, Extent: m.deltaExtent, Available: available, Style: style}
	title := titleLabel + " " + property.View(propertyWidth, m.theme) + " " + layout.Fit(feedback.Attention(state.Level, m.theme), layout.Size{Width: 7, Height: 1})
	rows := []string{""}
	if policy := m.boundaries[m.data.Metrics[index].ID]; policy.Fleet.Configured() {
		rows[0] = policy.Fleet.Summary("pp")
	}
	card := chart.ComparisonPanel{Title: title, Summary: rows, Distribution: distribution, History: m.plots[index]}
	return card.View(layout.Size{Width: width, Height: 17}, m.theme, m.metric == index)
}
func (m *CohortComparison) grid() string {
	if len(m.plots) == 0 {
		return ""
	}
	columns := 1
	if m.gridSize.Width >= 82 {
		columns = 2
	}
	width := max(1, (m.gridSize.Width-(columns-1))/columns)
	var rows []string
	for i := 0; i < len(m.data.Metrics); i += columns {
		row := m.panel(i, width)
		if columns == 2 && i+1 < len(m.data.Metrics) {
			row = layout.Row(row, " ", m.panel(i+1, m.gridSize.Width-width-1))
		}
		rows = append(rows, row)
	}
	return layout.Column(rows...)
}
func (m *CohortComparison) View() string {
	overall := attention.Off
	for _, a := range m.assessments {
		overall = attention.Worst(overall, a.Level)
	}
	header := m.theme.Heading().Render(m.title) + " " + feedback.Attention(overall, m.theme) + " / " + m.scope
	var groups []string
	for c, group := range m.data.Cohorts {
		s := m.summaries[m.metric][c]
		role := "Reference"
		if c == 1 {
			role = "Candidate"
		}
		groups = append(groups, m.theme.Text().Foreground(m.theme.SeriesColor(c)).Render(fmt.Sprintf("%s %s / %.0f req / %d hosts", role, group.ID, s.Requests, s.Members)))
	}
	window := "No windows"
	if m.window >= 0 {
		w := m.data.Windows[m.window]
		window = fmt.Sprintf("%s-%s UTC [%d/%d]", w.Start.UTC().Format("15:04"), w.End.UTC().Format("15:04"), m.window+1, len(m.data.Windows))
	}
	baseline, outliers, eligible, excluded := m.summaries[m.metric][0], 0, 0, 0
	for _, s := range m.summaries[m.metric] {
		eligible += s.Eligible
		excluded += s.Excluded
		for _, p := range s.Points {
			if z, ok := cohort.Distance(p, baseline); ok && math.Abs(z) > float64(m.sigma) {
				outliers++
			}
		}
	}
	count := fmt.Sprintf("%d beyond %ds", outliers, m.sigma)
	info := fmt.Sprintf("%s | +/- %ds host spread | min %d req/host | %d eligible / %d excluded | delta +/-%.3gpp", window, m.sigma, m.minRequests, eligible, excluded, m.deltaExtent)
	if !baseline.HasSpread || baseline.StdDev == 0 {
		count = "baseline s unavailable"
	}
	tableTitle := fmt.Sprintf("%s / %s %s", m.metricLabel(m.metric, true), count, feedback.Attention(m.assessments[m.metric].Level, m.theme))
	tablePanel := layout.Panel(tableTitle, m.Table.View(), layout.Size{Width: m.sideSize.Width, Height: m.tableHeight}, m.theme, m.focus.Current() == "hosts")
	inspector := m.inspectorView()
	side := layout.Column(tablePanel, inspector)
	content := layout.Column(m.Panels.View(), side)
	if m.sidebar {
		content = layout.Row(m.Panels.View(), " ", side)
	}
	return layout.Fit(strings.Join([]string{header, strings.Join(groups, " | "), info, content}, "\n"), m.size)
}
