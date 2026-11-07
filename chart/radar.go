package chart

import (
	"fmt"
	"math"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/NimbleMarkets/ntcharts/v2/canvas"
	"github.com/NimbleMarkets/ntcharts/v2/canvas/graph"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// RadarValue is one observed axis value and optional descriptive intervals.
// Valid=false leaves a gap in the profile, including its adjoining segments.
type RadarValue struct {
	Value     float64
	Valid     bool
	Intervals []Interval
}

type RadarSeries struct {
	Label  string
	Marker rune
	Values []RadarValue
}

// RadarOptions uses one shared, caller-supplied numeric domain. Reference and
// Rings use the same units as Values. Smooth interpolates radius with a cosine
// ease between axes: C1 continuous, exact at observations, with no overshoot.
// Interpolation only connects categories; it does not estimate observations.
type RadarOptions struct {
	Axes         []string
	Series       []RadarSeries
	Min, Max     float64
	Reference    float64
	Rings        []float64
	Smooth       bool
	SelectedAxis int
}

func radarValid(o RadarOptions) bool {
	if len(o.Axes) < 3 || !finite(o.Min) || !finite(o.Max) || o.Max <= o.Min || !finite(o.Max-o.Min) ||
		!finite(o.Reference) || o.Reference < o.Min || o.Reference > o.Max {
		return false
	}
	for _, ring := range o.Rings {
		if !finite(ring) || ring < o.Min || ring > o.Max {
			return false
		}
	}
	for _, s := range o.Series {
		if len(s.Values) != len(o.Axes) {
			return false
		}
		for _, p := range s.Values {
			if !p.Valid {
				continue
			}
			if !finite(p.Value) || p.Value < o.Min || p.Value > o.Max {
				return false
			}
			for _, interval := range p.Intervals {
				if !finite(interval.Low) || !finite(interval.High) || interval.Low > interval.High || interval.Low < o.Min || interval.High > o.Max {
					return false
				}
			}
		}
	}
	return true
}

// radarSegment samples either a straight Cartesian segment or a bounded smooth
// polar segment. Missing endpoints intentionally produce no connecting segment.
func radarSegment(a, b RadarValue, angle, step, min, max float64, smooth bool, samples int) []canvas.Float64Point {
	if !a.Valid || !b.Valid {
		return nil
	}
	ra, rb := (a.Value-min)/(max-min), (b.Value-min)/(max-min)
	points := make([]canvas.Float64Point, samples+1)
	for i := range points {
		t := float64(i) / float64(samples)
		if smooth {
			r := ra + (rb-ra)*(1-math.Cos(math.Pi*t))/2
			points[i] = canvas.Float64Point{X: r * math.Cos(angle+step*t), Y: r * math.Sin(angle+step*t)}
		} else {
			points[i] = canvas.Float64Point{
				X: (1-t)*ra*math.Cos(angle) + t*rb*math.Cos(angle+step),
				Y: (1-t)*ra*math.Sin(angle) + t*rb*math.Sin(angle+step),
			}
		}
	}
	return points
}

// Radar draws real ntcharts Braille geometry, axis markers and nested radial
// whiskers. ASCII and small allocations keep the same values in a text fallback.
func Radar(o RadarOptions, size layout.Size, t theme.Theme) string {
	if size.Width <= 0 || size.Height <= 0 {
		return ""
	}
	if !radarValid(o) {
		return layout.Fit("Invalid radar geometry", size)
	}
	if t.ASCII || size.Width < 30 || size.Height < 10 {
		lines := []string{fmt.Sprintf("Radar range %+.2f..%+.2f; ref %+.2f", o.Min, o.Max, o.Reference)}
		for i, axis := range o.Axes {
			parts := []string{axis}
			for _, s := range o.Series {
				v := "n/a"
				if p := s.Values[i]; p.Valid {
					v = fmt.Sprintf("%+.2f", p.Value)
					if len(p.Intervals) > 0 {
						p := p.Intervals[0]
						v += fmt.Sprintf(" [%+.2f,%+.2f]", p.Low, p.High)
					}
				}
				parts = append(parts, fmt.Sprintf("%c %s", radarMarker(s), v))
			}
			lines = append(lines, strings.Join(parts, " / "))
		}
		return layout.Fit(strings.Join(lines, "\n"), size)
	}
	plot := NewCanvas(size, t)
	cx, cy := float64(size.Width-1)/2, float64(size.Height-1)/2
	ry := math.Min(float64(size.Height-4)/2, float64(size.Width-12)/4)
	rx := ry * 2 // Terminal cells are approximately twice as tall as wide.
	step := -2 * math.Pi / float64(len(o.Axes))
	project := func(p canvas.Float64Point) canvas.Float64Point {
		return canvas.Float64Point{X: cx + p.X*rx, Y: cy + p.Y*ry}
	}
	cell := func(p canvas.Float64Point) canvas.Point {
		p = project(p)
		return canvas.Point{X: int(math.Round(p.X)), Y: size.Height - 1 - int(math.Round(p.Y))}
	}
	valuePoint := func(i int, value float64) canvas.Float64Point {
		r := (value - o.Min) / (o.Max - o.Min)
		a := math.Pi/2 + step*float64(i)
		return canvas.Float64Point{X: r * math.Cos(a), Y: r * math.Sin(a)}
	}
	draw := func(paths [][]canvas.Float64Point, style lipgloss.Style, dashed bool) {
		grid := graph.NewBrailleGrid(size.Width, size.Height, 0, float64(size.Width-1), 0, float64(size.Height-1))
		for _, path := range paths {
			for j := 1; j < len(path); j++ {
				if dashed && j%8 >= 5 {
					continue
				}
				a, b := grid.GridPoint(project(path[j-1])), grid.GridPoint(project(path[j]))
				for _, p := range graph.GetLinePoints(a, b) {
					grid.Set(p)
				}
			}
		}
		graph.DrawBraillePatterns(&plot, canvas.Point{}, grid.BraillePatterns(), style)
	}
	circle := func(value float64) []canvas.Float64Point {
		points := make([]canvas.Float64Point, 241)
		r := (value - o.Min) / (o.Max - o.Min)
		for i := range points {
			a := 2 * math.Pi * float64(i) / 240
			points[i] = canvas.Float64Point{X: r * math.Cos(a), Y: r * math.Sin(a)}
		}
		return points
	}
	var structure [][]canvas.Float64Point
	for _, ring := range o.Rings {
		structure = append(structure, circle(ring))
	}
	for i := range o.Axes {
		structure = append(structure, []canvas.Float64Point{{}, valuePoint(i, o.Max)})
	}
	draw(structure, lipgloss.NewStyle().Foreground(t.Border), false)
	draw([][]canvas.Float64Point{circle(o.Reference)}, t.Subtle(), true)
	for c, s := range o.Series {
		style := t.Text().Foreground(t.SeriesColor(c))
		var paths [][]canvas.Float64Point
		for i, p := range s.Values {
			paths = append(paths, radarSegment(p, s.Values[(i+1)%len(s.Values)], math.Pi/2+step*float64(i), step, o.Min, o.Max, o.Smooth, 100))
		}
		draw(paths, style, c == 0)
		for i, p := range s.Values {
			if !p.Valid {
				continue
			}
			// Whiskers sit on either side of the axis, leaving its mean marker
			// unobscured. Cap glyphs distinguish nested 3/2/1 intervals.
			for k, interval := range p.Intervals {
				lo, hi := cell(valuePoint(i, interval.Low)), cell(valuePoint(i, interval.High))
				shift := c*2 - 1
				if math.Abs(math.Sin(math.Pi/2+step*float64(i))) > .5 {
					lo.X, hi.X = lo.X+shift, hi.X+shift
				} else {
					lo.Y, hi.Y = lo.Y+shift, hi.Y+shift
				}
				for _, q := range graph.GetLinePoints(lo, hi) {
					plot.SetRuneWithStyle(q, '·', style)
				}
				cap := []rune{'┊', '┼', '╋'}[min(k+max(0, 3-len(p.Intervals)), 2)]
				plot.SetRuneWithStyle(lo, cap, style)
				plot.SetRuneWithStyle(hi, cap, style)
			}
		}
	}
	// Draw all exact observation markers last so curves and whiskers cannot
	// erase a selected host or version's value at an axis.
	for c, s := range o.Series {
		for i, p := range s.Values {
			if p.Valid {
				plot.SetRuneWithStyle(cell(valuePoint(i, p.Value)), radarMarker(s), t.Text().Foreground(t.SeriesColor(c)).Bold(true))
			}
		}
	}
	// Label the reference first, then any grid labels that fit without sharing
	// a terminal row. A crowded small plot omits labels rather than merging them.
	usedRows := map[int]bool{}
	for _, value := range append([]float64{o.Reference}, o.Rings...) {
		r := (value - o.Min) / (o.Max - o.Min)
		p := cell(canvas.Float64Point{X: r * math.Cos(math.Pi/4), Y: r * math.Sin(math.Pi/4)})
		p.X++
		label := fmt.Sprintf("%+.0f", value)
		if value == 0 {
			label = "0"
		}
		if !usedRows[p.Y] && p.X+len(label) < size.Width {
			plot.SetStringWithStyle(p, label, t.Subtle())
			usedRows[p.Y] = true
		}
	}
	for i, axis := range o.Axes {
		p := valuePoint(i, o.Max)
		p.X *= 1.18
		p.Y *= 1.18
		pos := cell(p)
		label := axis
		style := t.Text()
		if i == o.SelectedAxis {
			label, style = "["+axis+"]", t.Focus()
		}
		pos.X = max(0, min(size.Width-lipgloss.Width(label), pos.X-lipgloss.Width(label)/2))
		pos.Y = max(0, min(size.Height-1, pos.Y))
		plot.SetStringWithStyle(pos, label, style)
	}
	return layout.Fit(plot.View(), size)
}

func radarMarker(s RadarSeries) rune {
	if s.Marker == 0 {
		return 'o'
	}
	return s.Marker
}
