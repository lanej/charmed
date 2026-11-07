package feedback_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/attention"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestStatusAtomsKeepUnavailableAndDisabledDistinct(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		for _, emphasis := range []text.Emphasis{text.Bright, text.Normal, text.Dim} {
			for _, status := range []feedback.Status{feedback.OK, feedback.Warn, feedback.Error, feedback.Unknown, feedback.Off} {
				if !strings.Contains(ansi.Strip(status.View(emphasis, theme.New(mode))), string(status)) {
					t.Fatalf("%s status lost its label", status)
				}
			}
		}
		palette := theme.New(mode)
		if feedback.Attention(attention.Off, palette) != "" || !strings.Contains(ansi.Strip(feedback.Attention(attention.Unknown, palette)), "UNKNOWN") {
			t.Fatal("off and unknown attention states were conflated")
		}
	}
}
