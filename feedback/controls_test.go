package feedback_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/feedback"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestFollowAtomUsesPausedStatusAndStableGeometry(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		for _, emphasis := range []text.Emphasis{text.Normal, text.Bright, text.Dim} {
			following, paused := feedback.Follow(true, emphasis, palette), feedback.Follow(false, emphasis, palette)
			if paused != feedback.Paused.View(emphasis, palette) || !strings.Contains(ansi.Strip(following), "FOLLOW") {
				t.Fatal("follow diverged from shared status contract")
			}

			if ansi.StringWidth(following) != ansi.StringWidth(paused) {
				t.Fatal("follow state changed allocation")
			}
		}
	}
}

func TestScrollbarShowsViewportFractionAndClampsOffsets(t *testing.T) {
	palette := theme.New(theme.Mono)

	for _, example := range []struct {
		offset int
		want   string
	}{
		{0, "##||||||||"}, {40, "||||##||||"}, {80, "||||||||##"}, {-100, "##||||||||"}, {1000, "||||||||##"},
	} {
		bar := feedback.Scrollbar{Total: 100, Visible: 20, Offset: example.offset, Height: 10}

		got := strings.ReplaceAll(ansi.Strip(bar.View(palette)), "\n", "")
		if got != example.want {
			t.Fatalf("offset %d: got %q want %q", example.offset, got, example.want)
		}
	}

	for _, bar := range []feedback.Scrollbar{{Total: 10, Visible: 10, Height: 4}, {Total: 0, Visible: 10, Height: 4}, {Total: 10, Visible: 0, Height: 4}} {
		if got := ansi.Strip(bar.View(palette)); got != " \n \n \n " {
			t.Fatal("non-scrollable gutter reflowed", got)
		}
	}

	if got := (feedback.Scrollbar{Height: 0}).View(palette); got != "" {
		t.Fatal("zero height rendered", got)
	}
}

func TestStatusAndFollowSharePresentationChoices(t *testing.T) {
	palette := theme.New(theme.Dark)
	for _, mode := range []text.Presentation{text.GlyphAndLabel, text.GlyphOnly, text.LabelOnly} {
		paused := feedback.Follow(false, text.Normal, palette, mode)
		if paused != feedback.Paused.View(text.Normal, palette, mode) {
			t.Fatal("follow bypassed status presentation")
		}

		if label := strings.Contains(ansi.Strip(paused), "PAUSED"); label != (mode != text.GlyphOnly) {
			t.Fatal("status ignored label choice", mode, paused)
		}
	}
}
