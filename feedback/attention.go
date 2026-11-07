package feedback

import (
	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Attention is a labeled state atom, independent of metric-category semantics.
func Attention(level attention.Level, t theme.Theme) string {
	switch level {
	case attention.Error:
		return Error.View(text.Bright, t)
	case attention.Warn:
		return Warn.View(text.Bright, t)
	case attention.Unknown:
		return Unknown.View(text.Dim, t)
	case attention.Normal:
		return OK.View(text.Dim, t)
	}
	return ""
}
func AttentionStyle(level attention.Level, t theme.Theme) lipgloss.Style {
	switch level {
	case attention.Error:
		return t.Text().Foreground(t.Danger).Bold(true)
	case attention.Warn:
		return t.Text().Foreground(t.Warning).Bold(true)
	}
	return t.Subtle()
}
