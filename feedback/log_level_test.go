package feedback_test

import (
	"strings"
	"testing"

	"charm.land/log/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestLogLevelAtomsRetainSeverityAndGeometry(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			for _, level := range []log.Level{log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel} {
				rendered := feedback.LogLevel(level, emphasis, theme.New(mode))
				if !strings.Contains(ansi.Strip(rendered), strings.ToUpper(level.String())) || ansi.StringWidth(rendered) != 7 {
					t.Fatalf("severity or width lost: %q", rendered)
				}

				if mode == theme.Mono && !strings.HasPrefix(ansi.Strip(rendered), "[") {
					t.Fatal("monochrome lost its ASCII pill")
				}
			}

			if !strings.Contains(ansi.Strip(feedback.LogLevel(log.Level(19), emphasis, theme.New(mode))), "LEVEL 19") {
				t.Fatal("unknown level was lost")
			}
		}
	}
}
