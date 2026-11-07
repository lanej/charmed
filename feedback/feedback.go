// Package feedback contains shared loading, progress, time, and help components.
package feedback

import (
	"time"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/paginator"
	"charm.land/bubbles/v2/progress"
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/stopwatch"
	"charm.land/bubbles/v2/timer"
	"github.com/charmbracelet/harmonica"
	"github.com/lanej/charmed/theme"
)

func NewSpinner(t theme.Theme) spinner.Model {
	m := spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(t.Heading()))
	if t.ASCII || t.ReducedMotion {
		m.Spinner = spinner.Line
	}
	return m
}
func NewProgress(width int, t theme.Theme) progress.Model {
	opts := []progress.Option{progress.WithWidth(max(1, width)), progress.WithoutPercentage()}
	if t.Accent != nil {
		opts = append(opts, progress.WithColors(t.Accent))
	}
	m := progress.New(opts...)
	if t.ASCII {
		m.Full = '#'
		m.Empty = '-'
	}
	return m
}
func NewPaginator(total, perPage int, t theme.Theme) paginator.Model {
	m := paginator.New()
	m.Type = paginator.Dots
	m.PerPage = max(1, perPage)
	m.SetTotalPages(max(0, total))
	m.ActiveDot = t.Heading().Render("●")
	m.InactiveDot = t.Subtle().Render("○")
	if t.ASCII {
		m.ActiveDot = "[x]"
		m.InactiveDot = "[ ]"
	}
	return m
}
func NewHelp(width int, t theme.Theme) help.Model {
	m := help.New()
	m.SetWidth(width)
	m.Styles = help.DefaultStyles(t.IsDark())
	return m
}
func NewTimer(duration time.Duration) timer.Model {
	return timer.New(duration, timer.WithInterval(time.Second))
}
func NewStopwatch() stopwatch.Model { return stopwatch.New(stopwatch.WithInterval(time.Second)) }

// Motion is an explicit spring state. Call Step from the owner's tick command.
type Motion struct {
	spring             harmonica.Spring
	Position, Velocity float64
	Reduced            bool
}

func NewMotion(t theme.Theme) Motion {
	return Motion{spring: harmonica.NewSpring(harmonica.FPS(60), 6, 0.8), Reduced: t.ReducedMotion}
}
func (m *Motion) Step(target float64) float64 {
	if m.Reduced {
		m.Position = target
		m.Velocity = 0
	} else {
		m.Position, m.Velocity = m.spring.Update(m.Position, m.Velocity, target)
	}
	return m.Position
}
