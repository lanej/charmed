package feedback

import (
	"fmt"

	"charm.land/log/v2"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// TraceLevel extends Charm Log's native levels below DEBUG for display.
// Charm Log's String and ParseLevel methods do not recognize this value.
const TraceLevel log.Level = -8

// LogLevel renders severity independently of message content and emphasis.
// Unknown native levels retain their numeric value rather than implying INFO.
func LogLevel(level log.Level, emphasis text.Emphasis, t theme.Theme) string {
	var label string

	tone := text.Neutral

	switch level {
	case TraceLevel:
		label = "TRACE"
		t = statusTheme(t, t.Trace, t.DimTrace)
	case log.DebugLevel:
		label = "DEBUG"
		t = statusTheme(t, t.Debug, t.DimDebug)
	case log.InfoLevel:
		label, tone = "INFO", text.Info
	case log.WarnLevel:
		label, tone = "WARN", text.Warning
	case log.ErrorLevel:
		label, tone = "ERROR", text.Danger
	case log.FatalLevel:
		label = "FATAL"
		t = statusTheme(t, t.Fatal, t.DimFatal)
	default:
		label = fmt.Sprintf("LEVEL %d", level)
	}

	return text.Pill(fmt.Sprintf("%-5s", label), tone, t, emphasis)
}
