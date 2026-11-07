package story

import (
	"fmt"
	"math"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/barchart"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/heatmap"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/streamlinechart"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/wavelinechart"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func chartStory(kind string, e Environment) (runtime.Scene, error) {
	t := e.Config.Palette()
	step := 0
	currentSize := e.Config.Size()
	var render func() string
	var update func(tea.Msg) tea.Cmd
	var rebuild func(layout.Size)
	rebuild = func(size layout.Size) {
		currentSize = size
		plotSize := layout.Size{Width: size.Width, Height: max(1, size.Height-6)}
		update = nil
		switch kind {
		case "sparkline":
			values := sampleValues(48, step, e.Config.Amount)
			plot := chart.NewSparkline(values, layout.Size{Width: size.Width, Height: max(1, min(5, plotSize.Height))}, t)
			render = func() string { return plot.View() + fmt.Sprintf("\n\nLatest %.1f req/min", values[len(values)-1]) }
		case "timeseries":
			data := sampleSeries(e.Config.Series, step, e.Config.Amount)
			plot := chart.NewTimeSeries(data, plotSize, t)
			plot.Focus()
			render = func() string { return plot.View() + "\n" + chart.Legend(data, t) }
			update = func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; plot, cmd = plot.Update(msg); return cmd }
		case "line", "scatter":
			plot := chart.NewLine(plotSize, 0, 31, 0, 120, t)
			plot.Focus()
			draw := func() {
				plot.Clear()
				plot.DrawXYAxisAndLabel()
				for s := range e.Config.Series {
					values := sampleValues(32, step+s*3, e.Config.Amount-float64(s*10))
					style := lipgloss.NewStyle().Foreground(t.SeriesColor(s))
					for i, v := range values {
						p := canvas.Float64Point{X: float64(i), Y: v}
						if kind == "scatter" {
							plot.DrawRuneWithStyle(p, '●', style)
						} else if i > 0 {
							plot.DrawBrailleLineWithStyle(canvas.Float64Point{X: float64(i - 1), Y: values[i-1]}, p, style)
						}
					}
				}
			}
			draw()
			render = plot.View
			update = func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; plot, cmd = plot.Update(msg); draw(); return cmd }
		case "bars":
			opts := []barchart.Option{}
			if e.Config.Horizontal {
				opts = append(opts, barchart.WithHorizontalBars())
			}
			plot := chart.NewBars(plotSize, t, opts...)
			for i, name := range []string{"API", "Events", "Jobs", "Storage", "Schedule"} {
				values := []barchart.BarValue{}
				for s := range e.Config.Series {
					values = append(values, barchart.BarValue{Name: fmt.Sprintf("Series %d", s+1), Value: 20 + float64((i+s+step)*13%50), Style: lipgloss.NewStyle().Foreground(t.SeriesColor(s))})
				}
				plot.Push(barchart.BarData{Label: name, Values: values})
			}
			plot.Draw()
			render = plot.View
		case "heatmap":
			plot := chart.NewHeatmap(plotSize, t, heatmap.WithValueRange(0, 100), heatmap.WithCellSize(1, 1))
			plot.SetXYRange(0, 12, 0, 8)
			plot.SetViewXYRange(0, 12, 0, 8)
			for x := range 12 {
				for y := range 8 {
					plot.Push(heatmap.NewHeatPoint(float64(x), float64(y), 50+45*math.Sin(float64(x+y+step)*.4)))
				}
			}
			plot.Draw()
			render = plot.View
		case "stream":
			plot := chart.NewStream(plotSize, t, streamlinechart.WithYRange(0, 120))
			for _, v := range sampleValues(max(32, size.Width), step, e.Config.Amount) {
				plot.Push(v)
			}
			plot.DrawAll()
			render = plot.View
		case "wave":
			plot := chart.NewWave(plotSize, t, wavelinechart.WithXYRange(0, 31, 0, 120))
			for i, v := range sampleValues(32, step, e.Config.Amount) {
				plot.PlotDataSet("observed", canvas.Float64Point{X: float64(i), Y: v})
			}
			plot.DrawAll()
			render = plot.View
		case "ohlc":
			data := sampleSeries(4, step, e.Config.Amount)
			names := []string{"open", "high", "low", "close"}
			for i := range data {
				data[i].ID = names[i]
				data[i].Label = names[i]
				for p := range data[i].Points {
					base := sampleValues(32, step, e.Config.Amount)[p]
					offset := []float64{0, 9, -7, 4}[i]
					data[i].Points[p].Value = base + offset
				}
			}
			plot := chart.NewTimeSeries(data, plotSize, t)
			plot.Bubble.Clear()
			plot.Bubble.DrawCandle("open", "high", "low", "close", t.Text().Foreground(t.Success), t.Text().Foreground(t.Danger))
			render = plot.Bubble.View
		case "canvas":
			plot := chart.NewCanvas(plotSize, t)
			for x := range max(0, size.Width) {
				for y := range max(0, plotSize.Height) {
					if (x+y+step)%7 == 0 {
						plot.SetRuneWithStyle(canvas.Point{X: x, Y: y}, '·', t.Subtle())
					}
				}
			}
			plot.SetStringWithStyle(canvas.Point{X: 2, Y: 2}, "CHARMED / ntcharts canvas", t.Heading())
			render = plot.View
		}
	}
	tick := func() tea.Cmd { return tea.Tick(time.Second, func(time.Time) tea.Msg { return chartPulse{} }) }
	s := &scene{init: func() tea.Cmd {
		if e.Config.Live && !t.ReducedMotion {
			return tick()
		}
		return nil
	}, resize: rebuild, update: func(msg tea.Msg) tea.Cmd {
		_, manual := msg.(runtime.StepMsg)
		_, pulse := msg.(chartPulse)
		if manual || pulse {
			step++
			rebuild(currentSize)
			emit(e, map[string]any{"type": "chart-step", "step": step})
			if pulse && e.Config.Live && !t.ReducedMotion {
				return tick()
			}
			return nil
		}
		if update != nil {
			return update(msg)
		}
		return nil
	}, view: func(size layout.Size) string {
		if stateMessage(e) != "" {
			return title("Chart / "+kind, e) + stateMessage(e)
		}
		if size.Width < 20 || size.Height < 12 {
			return title(kind, e) + "Increase terminal size to see the chart."
		}
		if t.ASCII {
			return title(kind+" / text fallback", e) + fmt.Sprintf("%d series / latest %.1f\n\nASCII mode presents values without chart glyphs.", e.Config.Series, e.Config.Amount)
		}
		return title("Chart / "+kind, e) + render() + footer("Ctrl+N advances fixture data / arrows pan supported plots", e)
	}}
	rebuild(e.Config.Size())
	return s, nil
}

type chartPulse struct{}

func sampleValues(count, step int, amount float64) []float64 {
	values := make([]float64, count)
	for i := range values {
		values[i] = max(1, amount+14*math.Sin(float64(i+step)*.3)+5*math.Cos(float64(i)*.7))
	}
	return values
}
func sampleSeries(count, step int, amount float64) []chart.TimeSeriesData {
	base := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	data := make([]chart.TimeSeriesData, count)
	for s := range count {
		data[s] = chart.TimeSeriesData{ID: fmt.Sprintf("series-%d", s), Label: fmt.Sprintf("Series %d", s+1), Unit: "req/min"}
		for i, v := range sampleValues(32, step+s*4, amount+float64(s*12)) {
			data[s].Points = append(data[s].Points, timeserieslinechart.TimePoint{Time: base.Add(time.Duration(i) * time.Minute), Value: v})
		}
	}
	return data
}
