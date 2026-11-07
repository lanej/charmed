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
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/cohort"
	"github.com/lanej/charmed/collection"
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
}

// CohortComparison links per-metric distributions, request-weighted histories,
// and a member inspector. It owns UI state; callers supply interval counts.
// Sigma is descriptive member spread, never a significance or rollout decision.
type CohortComparison struct {
	Table                        table.Model
	Panels                       viewport.Model
	Inspector                    viewport.Model
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
	m.Inspector = collection.NewViewport("", 38, 16)
	m.size = layout.Size{Width: 132, Height: 40}
	if err := m.SetData(o.Data); err != nil {
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
func (m *CohortComparison) Init() tea.Cmd { return m.syncFocus() }
func (m *CohortComparison) syncFocus() tea.Cmd {
	m.Table.Blur()
	if m.focus.Current() == "hosts" {
		m.Table.Focus()
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
	height := max(1, size.Height-6)
	m.sidebar = size.Width >= 124 && size.Height >= 28
	if m.sidebar {
		m.sideSize = layout.Size{Width: 40, Height: height}
		m.gridSize = layout.Size{Width: max(1, size.Width-41), Height: height}
	} else {
		sideHeight := min(15, max(6, height/2))
		m.sideSize = layout.Size{Width: size.Width, Height: sideHeight}
		m.gridSize = layout.Size{Width: size.Width, Height: max(1, height-sideHeight)}
	}
	m.tableHeight = max(3, m.sideSize.Height/2)
	m.inspectorHeight = max(3, m.sideSize.Height-m.tableHeight)
	m.Table.SetWidth(max(1, m.sideSize.Width-2))
	m.Table.SetHeight(max(1, m.tableHeight-3))
	m.Table.SetColumns([]table.Column{{Title: "MEMBER", Width: max(1, m.sideSize.Width-18)}, {Title: "ROLE", Width: 4}, {Title: "z(ref)", Width: 6}})
	m.Panels.SetWidth(max(1, m.gridSize.Width))
	m.Panels.SetHeight(max(1, m.gridSize.Height))
	m.Inspector.SetWidth(max(1, m.sideSize.Width-2))
	m.Inspector.SetHeight(max(1, m.inspectorHeight-2))
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
	baseline := m.summaries[m.metric][0]
	sort.SliceStable(m.visible, func(i, j int) bool {
		a, b := m.visible[i], m.visible[j]
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
		rows = append(rows, table.Row{p.MemberID, m.short(p.CohortID), z})
	}
	if cursor < 0 && len(m.visible) > 0 {
		cursor = 0
		m.selected = m.visible[0].MemberKey
	}
	if len(m.visible) == 0 {
		m.selected = cohort.MemberKey{}
	}
	m.Table.SetRows(rows)
	if cursor >= 0 {
		m.Table.SetCursor(cursor)
	}
	m.updateDetails()
}
func (m *CohortComparison) short(id string) string {
	for _, c := range m.data.Cohorts {
		if c.ID == id {
			return c.Short
		}
	}
	return "?"
}
func (m *CohortComparison) updateDetails() {
	var lines []string
	if m.selected.MemberID == "" {
		lines = []string{"No member observations.", "Missing data is not zero."}
	} else {
		lines = []string{m.selected.MemberID + " / " + m.short(m.selected.CohortID)}
		lines = append(lines, ansi.Hardwrap("ID: "+m.selected.CohortID, max(1, m.sideSize.Width-2), true))
		for i := range m.data.Metrics {
			for _, summary := range m.summaries[i] {
				for _, p := range summary.Points {
					if p.MemberKey != m.selected {
						continue
					}
					distance := "z n/a"
					if z, ok := cohort.Distance(p, m.summaries[i][0]); ok {
						distance = fmt.Sprintf("z %+.2f", z)
					}
					if !p.Eligible {
						distance = "low volume"
					}
					lines = append(lines, fmt.Sprintf("%s %s  %s", m.metricLabel(i, true), rate(p.Rate, p.HasRate), distance))
					if i == 0 {
						lines = append(lines, fmt.Sprintf("%d requests in window", p.Requests))
					}
				}
			}
		}
		lines = append(lines, "", "X = this member on its cohort band", "3 = member history in every plot", "z(ref) = distance from reference hosts", "Descriptive; not a confidence test.")
	}
	m.Inspector.SetContent(strings.Join(lines, "\n"))
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
			series[c].Label = group.Short
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
		m.plots[i] = chart.NewTimeSeries(series, layout.Size{Width: width, Height: 7}, m.theme,
			timeserieslinechart.WithYRange(max(0, low-padding), min(100, high+padding)),
			timeserieslinechart.WithYLabelFormatter(func(_ int, value float64) string { return fmt.Sprintf("%.2f", value) }),
			timeserieslinechart.WithXLabelFormatter(func(_ int, value float64) string {
				return time.UnixMilli(int64(math.Round(value * 1000))).UTC().Format("15:04")
			}))
	}
}
func rate(value float64, ok bool) string {
	if !ok {
		return "n/a"
	}
	return fmt.Sprintf("%.3f%%", value*100)
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
	delta := "n/a"
	if summaries[0].HasRate && summaries[1].HasRate {
		delta = fmt.Sprintf("%+.3fpp", (summaries[1].Rate-summaries[0].Rate)*100)
	}
	lines := []string{fmt.Sprintf("Fleet %s %s  %s %s", m.data.Cohorts[0].Short, rate(summaries[0].Rate, summaries[0].HasRate), m.data.Cohorts[1].Short, rate(summaries[1].Rate, summaries[1].HasRate)), "Delta " + delta + " (request weighted)"}
	low, high := math.Inf(1), math.Inf(-1)
	for _, s := range summaries {
		if s.Eligible > 0 {
			low = min(low, s.Mean-3*s.StdDev)
			high = max(high, s.Mean+3*s.StdDev)
		}
		for _, p := range s.Points {
			if p.MemberKey == m.selected && p.HasRate {
				low = min(low, p.Rate)
				high = max(high, p.Rate)
			}
		}
	}
	if !math.IsInf(low, 0) {
		pad := max((high-low)*0.07, 0.00001)
		low, high = max(0, low-pad), min(1, high+pad)
	}
	for c, s := range summaries {
		label := m.data.Cohorts[c].Short
		style := m.theme.Text().Foreground(m.theme.SeriesColor(c))
		if s.Eligible == 0 {
			lines = append(lines, style.Render(label+" host mean n/a"), "Insufficient member samples")
			continue
		}
		spread := "n/a"
		if s.HasSpread {
			spread = fmt.Sprintf("%.3fpp", s.StdDev*100)
		}
		lines = append(lines, style.Render(fmt.Sprintf("%s mean %.3f%% s=%s n=%d", label, s.Mean*100, spread, s.Eligible)))
		options := chart.BandOptions{Min: low, Max: high, Center: s.Mean, Series: c}
		if s.HasSpread {
			for k := m.sigma; k >= 1; k-- {
				options.Intervals = append(options.Intervals, chart.Interval{Low: s.Mean - float64(k)*s.StdDev, High: s.Mean + float64(k)*s.StdDev})
			}
		}
		for _, p := range s.Points {
			if p.MemberKey == m.selected && p.HasRate {
				v := p.Rate
				options.Selected = &v
			}
		}
		lines = append(lines, chart.IntervalBand(options, max(1, width-2), m.theme))
	}
	if math.IsInf(low, 0) {
		lines = append(lines, "No eligible host spread")
	} else {
		axis := fmt.Sprintf("%.3f%%", low*100)
		end := fmt.Sprintf("%.3f%%", high*100)
		lines = append(lines, axis+strings.Repeat(" ", max(1, width-2-len(axis)-len(end)))+end)
	}
	lines = append(lines, "History %: 1="+m.data.Cohorts[0].Short+" 2="+m.data.Cohorts[1].Short+" 3=member", m.plots[index].View())
	return layout.Panel(m.metricLabel(index, false), strings.Join(lines, "\n"), layout.Size{Width: width, Height: 17}, m.theme, m.metric == index)
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
	header := m.theme.Heading().Render(m.title) + " / " + m.scope
	var groups []string
	for c, group := range m.data.Cohorts {
		s := m.summaries[m.metric][c]
		groups = append(groups, m.theme.Text().Foreground(m.theme.SeriesColor(c)).Render(fmt.Sprintf("%s %s / %.0f req / %d hosts", group.Short, group.Label, s.Requests, s.Members)))
	}
	window := "No windows"
	if m.window >= 0 {
		w := m.data.Windows[m.window]
		window = fmt.Sprintf("%s-%s UTC [%d/%d]", w.Start.UTC().Format("15:04"), w.End.UTC().Format("15:04"), m.window+1, len(m.data.Windows))
	}
	info := fmt.Sprintf("%s | +/- %ds host spread | min %d req/host | %s=reference %s=candidate", window, m.sigma, m.minRequests, m.data.Cohorts[0].Short, m.data.Cohorts[1].Short)
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
	if !baseline.HasSpread || baseline.StdDev == 0 {
		count = "baseline s unavailable"
	}
	tableTitle := fmt.Sprintf("%s / %s", m.metricLabel(m.metric, true), count)
	tablePanel := layout.Panel(tableTitle, m.Table.View(), layout.Size{Width: m.sideSize.Width, Height: m.tableHeight}, m.theme, m.focus.Current() == "hosts")
	inspector := layout.Panel("Member inspector", m.Inspector.View(), layout.Size{Width: m.sideSize.Width, Height: m.inspectorHeight}, m.theme, m.focus.Current() == "inspector")
	side := layout.Column(tablePanel, inspector)
	content := layout.Column(m.Panels.View(), side)
	if m.sidebar {
		content = layout.Row(m.Panels.View(), " ", side)
	}
	footer := fmt.Sprintf("1/2/3 bands | Left/Right window | [/] metric | Tab: %s | %d eligible / %d excluded", m.focus.Current(), eligible, excluded)
	note := "s = sample host SD; bands clipped 0-100%; z(ref) is descriptive. Tab + arrows scroll."
	return layout.Fit(strings.Join([]string{header, strings.Join(groups, " | "), info, content, footer, m.theme.Subtle().Render(note)}, "\n"), m.size)
}
