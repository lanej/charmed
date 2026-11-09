package feedback

import (
	"fmt"

	"charm.land/log/v2"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// LogLevel renders severity independently of message content and emphasis.
// Unknown native levels retain their numeric value rather than implying INFO.
func LogLevel(level log.Level, emphasis text.Emphasis, t theme.Theme) string {
	var label string

	tone := text.Neutral

	switch level {
	case log.DebugLevel:
		label = "DEBUG"
	case log.InfoLevel:
		label, tone = "INFO", text.Info
	case log.WarnLevel:
		label, tone = "WARN", text.Warning
	case log.ErrorLevel:
		label, tone = "ERROR", text.Danger
	case log.FatalLevel:
		label, tone = "FATAL", text.Danger
	default:
		label = fmt.Sprintf("LEVEL %d", level)
	}

	return text.Pill(fmt.Sprintf("%-5s", label), tone, t, emphasis)
}
