package story

import (
	"fmt"
	"image"
	"image/color"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/log/v2"
	"github.com/lanej/charmed/chart"
	"github.com/lanej/charmed/collection"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/runtime"
)

func metricStory(e Environment) (runtime.Scene, error) {
	return static(func(size layout.Size) string {
		return chart.MetricTrend("Throughput", fmt.Sprintf("%.0f", e.Config.Amount*20), "req/min", sampleValues(48, 0, e.Config.Amount), layout.Size{Width: size.Width, Height: min(size.Height, 8)}, e.Config.Palette())
	}), nil
}
func panelStory(e Environment) (runtime.Scene, error) {
	data := sampleSeries(e.Config.Series, 0, e.Config.Amount)
	plot := chart.NewTimeSeries(data, e.Config.Size(), e.Config.Palette())
	plot.Focus()
	return &scene{resize: func(size layout.Size) { plot.Resize(chart.PanelPlotSize(size)) }, update: func(msg tea.Msg) tea.Cmd { var cmd tea.Cmd; plot, cmd = plot.Update(msg); return cmd }, view: func(size layout.Size) string {
		return chart.Panel("Request throughput", plot, data, size, e.Config.Palette(), true)
	}}, nil
}
func pictureStory(e Environment) (runtime.Scene, error) {
	img := image.NewRGBA(image.Rect(0, 0, 120, 60))
	for y := range 60 {
		for x := range 120 {
			img.SetRGBA(x, y, color.RGBA{R: uint8(x * 2), G: uint8(70 + y*2), B: uint8(180 - x), A: 255})
		}
	}
	plot, init := chart.NewPicture(img, layout.Size{Width: e.Config.Columns, Height: e.Config.Rows - 5})
	return &scene{init: func() tea.Cmd { return init }, resize: func(s layout.Size) { plot.SetSize(s.Width, max(1, s.Height-5)) }, update: plot.Update, view: func(layout.Size) string {
		if e.Config.Palette().ASCII {
			return title("Picture / text fallback", e) + "Embedded gradient fixture: 120 x 60 pixels"
		}
		return title("Picture / portable glyph rendering", e) + plot.View().Content + footer("This preview uses half-block glyphs. Kitty graphics are capability-dependent.", e)
	}}, nil
}
func logStory(e Environment) (runtime.Scene, error) {
	m := collection.NewLogViewer(e.Config.Columns, e.Config.Rows-4, e.Config.Palette())
	step := 0
	appendLine := func() {
		level := log.InfoLevel
		if step%5 == 4 {
			level = log.WarnLevel
		}
		m.Append(collection.LogEntry{Time: time.Date(2026, 1, 1, 12, 0, step, 0, time.UTC), Level: level, Message: "Fixture request completed", Fields: []any{"service", "api", "sequence", step}})
		step++
	}
	for range 15 {
		appendLine()
	}
	return &scene{resize: func(s layout.Size) { m.SetWidth(s.Width); m.SetHeight(max(1, s.Height-4)) }, update: func(msg tea.Msg) tea.Cmd {
		if _, ok := msg.(runtime.StepMsg); ok {
			appendLine()
			return nil
		}
		var cmd tea.Cmd
		m.Model, cmd = m.Model.Update(msg)
		return cmd
	}, view: func(layout.Size) string {
		return title("Structured logs", e) + m.View() + footer("Arrows scroll / Ctrl+N appends a deterministic log entry", e)
	}}, nil
}
