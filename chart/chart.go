// Package chart gives ntcharts semantic styles, bounded sizes, and chart compositions.
package chart

import (
	"fmt"
	"math"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/NimbleMarkets/ntcharts/v2/sparkline"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type Sparkline struct {
	Bubble sparkline.Model
	values []float64
	size   layout.Size
	theme  theme.Theme
}

func NewSparkline(values []float64, size layout.Size, t theme.Theme) Sparkline {
	m := Sparkline{values: append([]float64(nil), values...), size: size, theme: t}
	m.rebuild()
	return m
}

func (m *Sparkline) rebuild() {
	m.Bubble = sparkline.New(max(1, m.size.Width), max(1, m.size.Height), sparkline.WithStyle(m.theme.SeriesStyle(0)))
	for _, v := range m.values {
		if !math.IsNaN(v) && !math.IsInf(v, 0) {
			m.Bubble.Push(v)
		}
	}
	m.Bubble.Draw()
}
func (m *Sparkline) Resize(size layout.Size) { m.size = size; m.rebuild() }
func (m *Sparkline) SetData(values []float64) {
	m.values = append([]float64(nil), values...)
	m.rebuild()
}

func (m Sparkline) View() string {
	if len(m.values) == 0 {
		return layout.Fit("No data", m.size)
	}
	if m.theme.ASCII {
		return layout.Fit(fmt.Sprintf("latest %.1f", m.values[len(m.values)-1]), m.size)
	}
	return layout.Fit(m.Bubble.View(), m.size)
}

type TimeSeriesData struct {
	ID, Label, Unit string
	Points          []timeserieslinechart.TimePoint
}

type TimeSeries struct {
	Bubble         timeserieslinechart.Model
	series         []TimeSeriesData
	size           layout.Size
	theme          theme.Theme
	focused        bool
	rangeAvailable bool
	options        []timeserieslinechart.Option
}

func NewTimeSeries(series []TimeSeriesData, size layout.Size, t theme.Theme, options ...timeserieslinechart.Option) TimeSeries {
	m := TimeSeries{size: size, theme: t, options: append([]timeserieslinechart.Option(nil), options...)}
	m.SetData(series)
	return m
}

func (m *TimeSeries) SetData(series []TimeSeriesData) {
	m.series = make([]TimeSeriesData, len(series))
	for i, s := range series {
		m.series[i] = s
		m.series[i].Points = append([]timeserieslinechart.TimePoint(nil), s.Points...)
	}
	m.rebuild()
}

func (m *TimeSeries) Resize(size layout.Size) {
	if m.size == size {
		return
	}
	minX, maxX, minY, maxY := m.Bubble.ViewMinX(), m.Bubble.ViewMaxX(), m.Bubble.ViewMinY(), m.Bubble.ViewMaxY()
	m.size = size
	m.rebuild()
	m.Bubble.SetViewXYRange(minX, maxX, minY, maxY)
	m.Bubble.DrawBrailleAll()
}
func (m *TimeSeries) Focus() { m.focused = true; m.Bubble.Focus() }
func (m *TimeSeries) Blur()  { m.focused = false; m.Bubble.Blur() }
func (m *TimeSeries) rebuild() {
	var first, last time.Time
	low, high := math.Inf(1), math.Inf(-1)
	for _, s := range m.series {
		for _, p := range s.Points {
			if !finite(p.Value) {
				continue
			}
			if first.IsZero() || p.Time.Before(first) {
				first = p.Time
			}
			if last.IsZero() || p.Time.After(last) {
				last = p.Time
			}
			low = math.Min(low, p.Value)
			high = math.Max(high, p.Value)
		}
	}
	if first.IsZero() {
		first = time.Unix(0, 0).UTC()
		last = first.Add(time.Hour)
		low, high = 0, 1
	}
	if !last.After(first) {
		last = first.Add(time.Hour)
	}

	low, high, m.rangeAvailable = PaddedRange(low, high, 0.15)
	if !m.rangeAvailable {
		low, high = 0, 1
	}
	defaults := []timeserieslinechart.Option{
		timeserieslinechart.WithTimeRange(first, last), timeserieslinechart.WithYRange(low, high),
		timeserieslinechart.WithXLabelFormatter(timeserieslinechart.HourTimeLabelFormatter()),
		timeserieslinechart.WithYLabelFormatter(func(_ int, value float64) string { return text.NumberLabel(value, 0, "", true) }),
		timeserieslinechart.WithXYSteps(3, 2), timeserieslinechart.WithAxesStyles(m.theme.Subtle(), m.theme.Subtle()),
	}

	height := m.size.Height
	if !m.rangeAvailable {
		height--
	}

	m.Bubble = timeserieslinechart.New(max(20, m.size.Width), max(6, height), append(defaults, m.options...)...)
	for i, s := range m.series {
		m.Bubble.SetDataSetStyle(s.ID, lipgloss.NewStyle().Foreground(m.theme.SeriesColor(i)))
		for _, p := range s.Points {
			if finite(p.Value) {
				m.Bubble.PushDataSet(s.ID, p)
			}
		}
	}
	if m.focused {
		m.Bubble.Focus()
	}
	m.Bubble.DrawBrailleAll()
}

func (m TimeSeries) Update(msg tea.Msg) (TimeSeries, tea.Cmd) {
	var cmd tea.Cmd
	m.Bubble, cmd = m.Bubble.Update(msg)
	m.Bubble.DrawBrailleAll()
	return m, cmd
}

func (m TimeSeries) View() string {
	count := 0
	for _, s := range m.series {
		count += len(s.Points)
	}

	if count == 0 {
		return layout.Fit("No data for this selection", m.size)
	}
	if m.size.Width < 20 || m.size.Height < 6 {
		return layout.Fit("Chart needs 20 columns / 6 rows", m.size)
	}
	if m.theme.ASCII {
		var lines []string
		for _, s := range m.series {
			if len(s.Points) > 0 {
				lines = append(lines, fmt.Sprintf("%s: %s %s", s.Label, text.NumberLabel(s.Points[len(s.Points)-1].Value, 1, "", true), s.Unit))
			}
		}
		return layout.Fit(strings.Join(lines, "\n"), m.size)
	}

	content := m.Bubble.View()
	if !m.rangeAvailable {
		content = "History range unavailable (full span)\n" + content
	}

	return layout.Fit(content, m.size)
}

// Legend renders a single-row key using each series color.
func Legend(series []TimeSeriesData, t theme.Theme) string {
	entries := seriesLegendEntries(series, t)
	width := 0

	for _, entry := range entries {
		width += lipgloss.Width(entry.glyph+" "+entry.label) + 2
	}

	return renderLegend(entries, width, t)
}

// SeriesLegend wraps series keys to the supplied width without dropping labels.
func SeriesLegend(series []TimeSeriesData, width int, t theme.Theme) string {
	return renderLegend(seriesLegendEntries(series, t), width, t)
}

func seriesLegendEntries(series []TimeSeriesData, t theme.Theme) []legendEntry {
	entries := make([]legendEntry, len(series))
	for i, s := range series {
		label := s.Label
		if label == "" {
			label = s.ID
		}

		entries[i] = legendEntry{glyph: fmt.Sprint(i + 1), label: label, style: lipgloss.NewStyle().Foreground(t.SeriesColor(i))}
	}

	return entries
}

// PanelOptions controls optional annotations. The zero value omits the legend.
type PanelOptions struct{ ShowLegend bool }

// PlotSize reserves the actual wrapped legend height for the supplied series.
// Use this allocation when constructing or resizing the native child chart.
func (o PanelOptions) PlotSize(size layout.Size, series []TimeSeriesData, t theme.Theme) layout.Size {
	content := layout.PanelContentSize(size)
	if o.ShowLegend {
		if legend := SeriesLegend(series, content.Width, t); legend != "" {
			content.Height = max(0, content.Height-lipgloss.Height(legend))
		}
	}

	return content
}

// Panel composes a chart allocated to options.PlotSize(size, series, theme).
// Rendering does not rebuild or mutate the native model or its custom styles.
func Panel(title string, plot TimeSeries, series []TimeSeriesData, size layout.Size, t theme.Theme, focused bool, options ...PanelOptions) string {
	var o PanelOptions
	if len(options) > 0 {
		o = options[0]
	}

	legend := ""
	if o.ShowLegend {
		legend = SeriesLegend(series, layout.PanelContentSize(size).Width, t)
	}

	if legend != "" && lipgloss.Height(legend) > layout.PanelContentSize(size).Height {
		return layout.Panel(title, t.Subtle().Render(fmt.Sprintf("Legend needs %d rows", lipgloss.Height(legend))), size, t, focused)
	}

	content := layout.Fit(plot.View(), o.PlotSize(size, series, t))
	if legend != "" {
		if content != "" {
			content += "\n"
		}

		content += legend
	}

	return layout.Panel(title, content, size, t, focused)
}

// PanelPlotSize retains the single-row legend allocation for existing callers.
// Use PanelOptions.PlotSize to measure legends that can occupy multiple rows.
func PanelPlotSize(size layout.Size, options ...PanelOptions) layout.Size {
	legendRows := 0
	if len(options) > 0 && options[0].ShowLegend {
		legendRows = 1
	}

	content := layout.PanelContentSize(size)
	content.Height = max(0, content.Height-legendRows)

	return content
}

func finite(v float64) bool { return !math.IsNaN(v) && !math.IsInf(v, 0) }
