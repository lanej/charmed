package feedback

import (
	"image/color"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Status is display state, independent of attention-rule evaluation.
type Status string

const (
	OK           Status = "OK"
	Warn         Status = "WARN"
	Error        Status = "ERROR"
	Unknown      Status = "UNKNOWN"
	Off          Status = "OFF"
	Inactive     Status = "INACTIVE"
	Pending      Status = "PENDING"
	Running      Status = "RUNNING"
	Paused       Status = "PAUSED"
	Disabled     Status = "DISABLED"
	Connected    Status = "CONNECTED"
	Disconnected Status = "DISCONNECTED"
)

// Statuses returns the display catalog in its learning order.
func Statuses() []Status {
	return []Status{OK, Warn, Error, Unknown, Off, Inactive, Pending, Running, Paused, Disabled, Connected, Disconnected}
}

// MaxStatusWidth includes the glyph, separating space and horizontal insets.
// All emphasis levels retain this geometry.
func MaxStatusWidth(t theme.Theme) int {
	width := 0
	for _, status := range Statuses() {
		width = max(width, ansi.StringWidth(status.View(text.Normal, t)))
	}

	return width
}

func (s Status) View(emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	tone := text.Neutral
	glyph := theme.GlyphUnknown

	switch s {
	case OK:
		tone = text.Success
		glyph = theme.GlyphOK
	case Warn:
		tone = text.Warning
		glyph = theme.GlyphWarning
	case Error:
		tone = text.Danger
		glyph = theme.GlyphError
	case Unknown:
		t = statusTheme(t, t.Unknown, t.DimUnknown)
	case Off:
		glyph = theme.GlyphOff
		t = statusTheme(t, t.Inactive, t.DimInactive)
	case Inactive:
		glyph = theme.GlyphInactive
		t = statusTheme(t, t.Inactive, t.DimInactive)
	case Pending:
		glyph = theme.GlyphPending
		t = statusTheme(t, t.Pending, t.DimPending)
	case Running:
		glyph = theme.GlyphRunning
		t = statusTheme(t, t.Running, t.DimRunning)
	case Paused:
		glyph = theme.GlyphPaused
		t = statusTheme(t, t.Paused, t.DimPaused)
	case Connected:
		tone = text.Success
		glyph = theme.GlyphConnected
	case Disconnected:
		glyph = theme.GlyphDisconnected

		t = statusTheme(t, t.Inactive, t.DimInactive)
	case Disabled:
		glyph = theme.GlyphDisabled
		t = statusTheme(t, t.Inactive, t.DimInactive)
	default:
		s = Unknown
		t = statusTheme(t, t.Unknown, t.DimUnknown)
	}

	mode := text.DefaultPresentation
	if len(options) > 0 {
		mode = options[0]
	}

	return text.Pill(mode.Format(t.Glyph(glyph), string(s)), tone, t, emphasis)
}

func statusTheme(t theme.Theme, normal, dim color.Color) theme.Theme {
	if normal != nil {
		t.Foreground = normal
	}

	if dim != nil {
		t.DimForeground = dim
	}

	return t
}
