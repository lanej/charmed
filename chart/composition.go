package chart

import (
	"fmt"
	"image"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/NimbleMarkets/ntcharts/v2/linechart/timeserieslinechart"
	"github.com/NimbleMarkets/ntcharts/v2/picture"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

func MetricTrend(label, value, unit string, values []float64, size layout.Size, t theme.Theme) string {
	trend := NewSparkline(values, layout.Size{Width: size.Width, Height: max(1, size.Height-3)}, t)
	return layout.Fit(t.Subtle().Render(label)+"\n"+t.Heading().Render(value)+" "+t.Subtle().Render(unit)+"\n\n"+trend.View(), size)
}

type Candle struct {
	Time                   time.Time
	Open, High, Low, Close float64
}
type Candles struct {
	Plot  TimeSeries
	data  []Candle
	theme theme.Theme
}

func NewCandles(data []Candle, size layout.Size, t theme.Theme) (Candles, error) {
	series := []TimeSeriesData{{ID: "open"}, {ID: "high"}, {ID: "low"}, {ID: "close"}}
	for _, p := range data {
		if !finite(p.Open) || !finite(p.High) || !finite(p.Low) || !finite(p.Close) || p.High < p.Low || p.High < max(p.Open, p.Close) || p.Low > min(p.Open, p.Close) {
			return Candles{}, fmt.Errorf("invalid OHLC candle at %s", p.Time.Format(time.RFC3339))
		}
		for i, v := range []float64{p.Open, p.High, p.Low, p.Close} {
			series[i].Points = append(series[i].Points, timeserieslinechart.TimePoint{Time: p.Time, Value: v})
		}
	}
	m := Candles{Plot: NewTimeSeries(series, size, t), data: append([]Candle(nil), data...), theme: t}
	m.draw()
	return m, nil
}
func (m *Candles) draw() {
	m.Plot.Bubble.Clear()
	m.Plot.Bubble.DrawCandle("open", "high", "low", "close", m.theme.Text().Foreground(m.theme.Success), m.theme.Text().Foreground(m.theme.Danger))
}
func (m *Candles) Resize(size layout.Size) { m.Plot.Resize(size); m.draw() }
func (m Candles) Update(msg tea.Msg) (Candles, tea.Cmd) {
	var cmd tea.Cmd
	m.Plot, cmd = m.Plot.Update(msg)
	m.draw()
	return m, cmd
}
func (m Candles) View() string { return m.Plot.View() }

// NewPicture starts in portable glyph mode. Call the native model's Toggle only
// after the application has established Kitty support; forward every returned Cmd.
func NewPicture(img image.Image, size layout.Size) (picture.Model, tea.Cmd) {
	m := picture.New()
	sizeCmd := m.SetSize(max(1, size.Width), max(1, size.Height))
	imageCmd := m.SetImage(img)
	return m, tea.Batch(sizeCmd, imageCmd)
}
