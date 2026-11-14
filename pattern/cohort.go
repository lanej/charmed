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
	"github.com/charmbracelet/x/ansi"
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

type ComparisonVariant string

const (
	Stacked ComparisonVariant = "stacked"
	Bands   ComparisonVariant = "bands"
)

type CohortComparisonOptions struct {
	Title, Scope string
	// ReferenceID optionally selects a baseline by identity. Empty disables
	// reference-relative comparisons; Data accepts one or more cohorts.
	Data        cohort.Data
	ReferenceID string
	// MinSupport limits ratio member eligibility by denominator. Scalars remain eligible.
	MinSupport uint64
	Theme      theme.Theme
	// MinRequests is a compatibility alias; the larger of both thresholds applies.
	MinRequests uint64
	Sigma       int
	Emit        func(any)
	MetricTones map[string]text.Tone
	Boundaries  map[string]cohort.ComparisonPolicy
	Variant     ComparisonVariant
	MetricSigns map[string]text.SignTones
}

// CohortComparison links per-metric distributions, aggregate histories,
// and a member inspector. It owns UI state; callers supply measurements.
// Sigma is descriptive member spread, never a significance or rollout decision.
type CohortComparison struct {
	geometryOnly                 bool
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
	summaries                    [][]cohort.Summary
	referenceID                  string
	comparisons                  []map[string]cohort.ComparisonAttention
	plots                        []chart.TimeSeries
	historyNotes                 [][]string
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
	variant                      ComparisonVariant
	metricSigns                  map[string]text.SignTones
}

func NewCohortComparison(o CohortComparisonOptions) (*CohortComparison, error) {
	if o.Variant == "" {
		o.Variant = Stacked
	}

	if o.Variant != Stacked && o.Variant != Bands {
		return nil, fmt.Errorf("unknown comparison variant %q", o.Variant)
	}

	if o.Sigma == 0 {
		o.Sigma = 3
	}

	if o.Sigma < 1 || o.Sigma > 3 {
		return nil, fmt.Errorf("sigma must be 1, 2 or 3")
	}

	m := &CohortComparison{
		title: o.Title, scope: o.Scope, theme: o.Theme, minRequests: max(1, o.MinRequests, o.MinSupport), referenceID: o.ReferenceID, sigma: o.Sigma, emit: o.Emit,
		focus: interaction.NewFocus("hosts", "metrics", "inspector"), window: -1, metricTones: maps.Clone(o.MetricTones), variant: o.Variant, metricSigns: maps.Clone(o.MetricSigns),
	}
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
func (m *CohortComparison) Variant() ComparisonVariant { return m.variant }

// SetVariant changes presentation on the existing model. No fixture, focus,
// selected host, metric, time window or attention policy is replaced.
func (m *CohortComparison) SetVariant(variant ComparisonVariant) error {
	if variant != Stacked && variant != Bands {
		return fmt.Errorf("unknown comparison variant %q", variant)
	}

	if variant == m.variant {
		return nil
	}

	m.variant = variant
	m.Panels.GotoTop()
	m.updateDetails()
	m.event("variant")

	return nil
}

// SetData is atomic on validation failure. Appends follow the latest window only
// if it was already selected. An inspected historical window stays selected.
func (m *CohortComparison) SetData(data cohort.Data) error {
	copy, err := data.Clone()
	if err != nil {
		return err
	}

	if err := validateReference(copy, m.referenceID); err != nil {
		return err
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
		case "b":
			variant := Bands
			if m.variant == Bands {
				variant = Stacked
			}

			_ = m.SetVariant(variant)

			return nil
		case "f6":
			_ = m.SetVariant(Stacked)
			return nil
		case "f7":
			_ = m.SetVariant(Bands)
			return nil
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
		m.emit(map[string]any{
			"type": kind, "cohort": m.selected.CohortID, "member": m.selected.MemberID,
			"reference": m.referenceID, "cohorts": m.cohortIDs(),
			"window": m.window, "metric": m.MetricID(), "sigma": m.sigma, "variant": m.variant,
		})
	}
}

// ResizeViewport changes geometry without refreshing unchanged-width stacked
// panel contents. Resize retains its full presentation refresh behavior.
func (m *CohortComparison) ResizeViewport(size layout.Size) {
	m.geometryOnly = true
	defer func() { m.geometryOnly = false }()

	m.Resize(size)
}

func (m *CohortComparison) Resize(size layout.Size) {
	previousGridWidth := m.gridSize.Width
	previousSideWidth := m.sideSize.Width
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
	m.Table.SetWidth(max(1, layout.PanelContentSize(m.sideSize).Width))
	m.Table.SetHeight(max(1, m.tableHeight-3))
	m.tableColumns()
	m.Panels.SetWidth(max(1, m.gridSize.Width))
	m.Panels.SetHeight(max(1, m.gridSize.Height))
	m.Panels.SetYOffset(m.Panels.YOffset())
	m.Inspector.SetHeight(max(1, m.inspectorHeight-3-lipgloss.Height(m.inspectorContext())))

	if m.geometryOnly && m.variant == Stacked && previousGridWidth == m.gridSize.Width && previousSideWidth == m.sideSize.Width && len(m.plots) == len(m.data.Metrics) {
		return
	}

	m.updateInspector()

	columns := 1
	if m.variant == Stacked && m.gridSize.Width >= 82 {
		columns = 2
	}

	width := max(1, layout.PanelContentSize(layout.Size{Width: (m.gridSize.Width - (columns - 1)) / columns}).Width)
	if len(m.plots) != len(m.data.Metrics) {
		m.rebuildPlots()
	} else {
		for i := range m.plots {
			m.plots[i].Resize(layout.Size{Width: width, Height: m.historyHeight()})
		}
	}

	m.Panels.SetContent(m.grid())
}

func (m *CohortComparison) refresh() {
	m.summaries = make([][]cohort.Summary, len(m.data.Metrics))
	for i := range m.summaries {
		m.summaries[i] = make([]cohort.Summary, len(m.data.Cohorts))
	}

	m.visible = nil
	if m.window >= 0 && m.window < len(m.data.Windows) {
		for i, metric := range m.data.Metrics {
			for c, group := range m.data.Cohorts {
				m.summaries[i][c] = cohort.AnalyzeMetric(m.data.Windows[m.window].Samples, group.ID, metric, m.minRequests)
			}
		}

		for c := range m.data.Cohorts {
			summary := m.memberSummary(m.metric, c)
			m.visible = append(m.visible, summary.Points...)
		}
	}

	changes := m.assessAttention()
	m.updateDeltaExtent()
	baseline := m.reference(m.metric)
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
		if m.referenceID != "" && da != db {
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
			z = m.signedNumber(m.metric, distance, 1, "", true).View(m.theme)
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
	if m.variant == Stacked && m.gridSize.Width >= 82 {
		columns = 2
	}

	panelWidth := (m.gridSize.Width - (columns - 1)) / columns
	width := max(1, layout.PanelContentSize(layout.Size{Width: panelWidth}).Width)

	m.plots = make([]chart.TimeSeries, len(m.data.Metrics))

	m.historyNotes = make([][]string, len(m.data.Metrics))
	for i, metric := range m.data.Metrics {
		series, notes := m.historySeries(metric)
		m.historyNotes[i] = notes
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

		low, high, _ = chart.PaddedRange(low, high, 0.15)
		m.plots[i] = chart.NewTimeSeries(series, layout.Size{Width: width, Height: m.historyHeight()}, m.theme,
			timeserieslinechart.WithYRange(low, high),
			timeserieslinechart.WithYLabelFormatter(func(_ int, value float64) string { return text.NumberLabel(value, 2, metric.DisplayUnit(), true) }),
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

func (m *CohortComparison) distribution(index int) chart.DistributionComparison {
	summaries := m.summaries[index]
	metric := m.data.Metrics[index]
	low, high := m.distributionRange(index)

	distribution := chart.DistributionComparison{Sigma: m.sigma, Min: low, Max: high, Unit: metric.DisplayUnit(), SpreadUnit: metric.DeltaUnit(), Markers: m.boundaryMarkers(index)}
	for c, s := range summaries {
		series := chart.DistributionSeries{Label: m.data.Cohorts[c].Label, Mean: s.Mean * metric.Scale(), StdDev: s.StdDev * metric.Scale(), Members: s.Eligible, HasSpread: s.HasSpread}
		for _, point := range s.Points {
			if point.MemberKey == m.selected && point.HasRate {
				value := point.Value * metric.Scale()
				series.Selected = &value
			}
		}

		distribution.Series = append(distribution.Series, series)
	}

	return distribution
}

func (m *CohortComparison) comparisonTitle(index, width int) string {
	return layout.Fit(m.metricLabel(index, true)+" "+feedback.Attention(m.assessments[index].Level, m.theme), layout.Size{Width: width, Height: 1})
}

func (m *CohortComparison) comparisonRows(index, width int) []string {
	metric := m.data.Metrics[index]

	var rows []string
	if m.referenceID == "" {
		return rows
	}

	for c, group := range m.data.Cohorts {
		if group.ID == m.referenceID {
			continue
		}

		delta, available := m.cohortDelta(index, c)
		state := m.comparisons[index][group.ID]
		property := chart.DeltaProperty{Value: delta, Extent: m.metricDeltaExtent(index), Available: available, Unit: metric.DeltaUnit(), Unitless: metric.DeltaUnit() == "", Style: feedback.AttentionStyle(state.Fleet.Level, m.theme), Tones: m.metricSigns[metric.ID]}
		labelWidth := max(0, width-15)

		label := ansi.Truncate(group.Label+" vs "+m.referenceLabel(), labelWidth, "~")
		if label != "" {
			label += " "
		}

		rows = append(rows, label+property.View(max(1, width-lipgloss.Width(label)), m.theme))
	}

	return rows
}

type metricPanel struct {
	title, content string
	height         int
}

func (m *CohortComparison) panelContent(index, width int) metricPanel {
	rows := m.comparisonRows(index, layout.PanelContentSize(layout.Size{Width: width}).Width)

	rows = append(rows, m.historyRows(index, layout.PanelContentSize(layout.Size{Width: width}).Width)...)
	if policy := m.boundaries[m.data.Metrics[index].ID]; policy.Fleet.Configured() {
		rows = append(rows, policy.Fleet.Summary(m.data.Metrics[index].DeltaUnit()))
	}

	card := chart.ComparisonPanel{Title: m.comparisonTitle(index, layout.PanelContentSize(layout.Size{Width: width}).Width), Summary: rows, Distribution: m.distribution(index), History: m.plots[index]}

	contentWidth := max(1, layout.PanelContentSize(layout.Size{Width: width}).Width)

	lines := append([]string(nil), rows...)
	lines = append(lines, card.Distribution.View(contentWidth, m.theme), card.History.View())
	height := lipgloss.Height(strings.Join(lines, "\n")) + 2*theme.BorderCells

	if len(card.Distribution.Series) == 0 {
		lines = append(rows, lines[len(lines)-1])
	}

	return metricPanel{title: card.Title, content: strings.Join(lines, "\n"), height: height}
}

func (m *CohortComparison) grid() string {
	if len(m.plots) == 0 {
		return ""
	}

	if m.variant == Bands {
		return m.bandVariantView()
	}

	columns := 1
	if m.gridSize.Width >= 82 {
		columns = 2
	}

	width := max(1, (m.gridSize.Width-(columns-1))/columns)

	var rows []string

	for i := 0; i < len(m.data.Metrics); i += columns {
		left := m.panelContent(i, width)
		height := left.height

		var right metricPanel

		otherWidth := m.gridSize.Width - width - 1

		paired := columns == 2 && i+1 < len(m.data.Metrics)
		if paired {
			right = m.panelContent(i+1, otherWidth)
			height = max(height, right.height)
		}

		row := layout.Panel(left.title, left.content, layout.Size{Width: width, Height: height}, m.theme, m.metric == i)
		if paired {
			other := layout.Panel(right.title, right.content, layout.Size{Width: otherWidth, Height: height}, m.theme, m.metric == i+1)
			row = layout.Row(row, " ", other)
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
		members := len(m.memberSummary(m.metric, c).Points)

		role := ""
		if group.ID == m.referenceID {
			role = "Reference "
		}

		groups = append(groups, m.theme.Text().Foreground(m.theme.SeriesColor(c)).Render(fmt.Sprintf("%s%s / %d members", role, group.Label, members)))
	}

	window := "No windows"
	if m.window >= 0 {
		w := m.data.Windows[m.window]
		window = fmt.Sprintf("%s-%s UTC [%d/%d]", w.Start.UTC().Format("15:04"), w.End.UTC().Format("15:04"), m.window+1, len(m.data.Windows))
	}

	baseline, outliers, eligible, excluded := m.reference(m.metric), 0, 0, 0
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

	info := fmt.Sprintf("%s | +/- %ds member spread | min ratio support %d | %d eligible / %d excluded", window, m.sigma, m.minRequests, eligible, excluded)
	if m.referenceID != "" && len(m.data.Cohorts) > 1 {
		info += fmt.Sprintf(" | delta +/-%.3g%s", m.deltaExtent, m.data.Metrics[m.metric].DeltaUnit())
	}

	if m.referenceID == "" {
		count = "No reference"
	} else if !baseline.HasSpread || baseline.StdDev == 0 {
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

func (m *CohortComparison) historyHeight() int {
	height := 9
	if m.variant == Bands {
		height = max(6, m.gridSize.Height-23)
	}

	if m.theme.ASCII {
		height = max(height, len(m.data.Cohorts)+1)
	}

	return height
}
