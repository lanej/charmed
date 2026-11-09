package collection

import (
	"strings"
	"time"

	"charm.land/bubbles/v2/viewport"
	"charm.land/log/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

type LogEntry struct {
	Time    time.Time
	Level   log.Level
	Message string
	Fields  []any
}

// LogViewer is updated by its owning model. Send log entries as messages when
// producers run on other goroutines instead of mutating UI state from a logger.
type LogViewer struct {
	viewport.Model
	Theme theme.Theme
	Limit int
	lines []string
}

func NewLogViewer(width, height int, t theme.Theme) LogViewer {
	return LogViewer{Model: NewViewport("", width, height), Theme: t, Limit: 500}
}

func (m *LogViewer) Append(entry LogEntry) {
	var out strings.Builder

	logger := log.NewWithOptions(&out, log.Options{})
	logger.SetColorProfile(colorprofile.NoTTY)
	logger.Print(entry.Message, entry.Fields...)
	line := m.Theme.Subtle().Render(entry.Time.Format("15:04:05")) + " " + feedback.LogLevel(entry.Level, text.Normal, m.Theme) + " " + m.Theme.Text().Render(strings.TrimSuffix(out.String(), "\n"))
	follow := m.AtBottom()

	m.lines = append(m.lines, line)
	if limit := max(1, m.Limit); len(m.lines) > limit {
		m.lines = append([]string(nil), m.lines[len(m.lines)-limit:]...)
	}

	m.SetContent(strings.Join(m.lines, "\n"))

	if follow {
		m.GotoBottom()
	}
}
