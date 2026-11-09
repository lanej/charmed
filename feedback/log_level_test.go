package feedback_test

import (
	"image/color"
	"strings"
	"testing"

	"charm.land/log/v2"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestLogLevelColorsDistinguishEverySeverity(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast} {
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			seen := map[color.Color]bool{}

			for _, level := range []log.Level{feedback.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel} {
				cells := uv.NewStyledString(feedback.LogLevel(level, emphasis, theme.New(mode))).Lines(ansi.GraphemeWidth)

				ink := cells[0][1].Style.Fg
				if emphasis == text.Bright {
					ink = cells[0][1].Style.Bg
				}

				if ink == nil || seen[color.NRGBAModel.Convert(ink)] {
					t.Fatalf("%s/%s conflated severity %d", mode, emphasis, level)
				}

				seen[color.NRGBAModel.Convert(ink)] = true
			}
		}
	}
}

func TestLogLevelAtomsRetainSeverityAndGeometry(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			for _, tc := range []struct {
				level log.Level
				label string
			}{{feedback.TraceLevel, "TRACE"}, {log.DebugLevel, "DEBUG"}, {log.InfoLevel, "INFO"}, {log.WarnLevel, "WARN"}, {log.ErrorLevel, "ERROR"}, {log.FatalLevel, "FATAL"}} {
				rendered := feedback.LogLevel(tc.level, emphasis, theme.New(mode))
				if !strings.Contains(ansi.Strip(rendered), tc.label) || ansi.StringWidth(rendered) != 7 {
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
