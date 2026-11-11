package text_test

import (
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestLinkKeepsExactTargetAndVisibleAffordance(t *testing.T) {
	link := text.Link{Label: "PR 界é", URL: "https://example.test/a%2Fb?b=2&a=1&a=3#output"}

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)

		view, err := link.View(palette)
		if err != nil {
			t.Fatal(err)
		}

		want := "↗ PR 界é"
		if mode == theme.Mono {
			want = "@ PR 界é"
		}

		if got := ansi.Strip(view); got != want {
			t.Fatalf("%s: visible label = %q, want %q", mode, got, want)
		}

		for _, profile := range []string{"truecolor", "ansi256", "ansi16", "none"} {
			config := story.DefaultConfig()
			config.Profile = profile
			rendered := story.FormatFrame(view, config)
			cells := uv.NewStyledString(rendered + "!").Lines(ansi.GraphemeWidth)[0]

			for _, cell := range cells {
				if cell.Width == 0 {
					continue
				}

				if cell.Content == "!" {
					if cell.Link.URL != "" {
						t.Fatal("link leaked onto adjacent content")
					}

					continue
				}

				if cell.Link.URL != link.URL {
					t.Fatalf("%s/%s: cell %q target = %q", mode, profile, cell.Content, cell.Link.URL)
				}

				if profile != "none" && cell.Style.Underline == uv.UnderlineNone {
					t.Fatalf("%s/%s: label lost its underline", mode, profile)
				}
			}
		}
	}

	plain, err := link.Plain()
	if err != nil {
		t.Fatal(err)
	}

	if want := "PR 界é (" + link.URL + ")"; plain != want || ansi.Strip(plain) != plain {
		t.Fatalf("plain destination lost: %q", plain)
	}
}

func TestLinkEscapesLabelsAndRejectsInvalidTargets(t *testing.T) {
	palette := theme.New(theme.Dark)
	link := text.Link{Label: "Logs\n\x1b]8;;https://wrong.test\a", URL: "http://localhost:8080/logs"}

	view, err := link.View(palette)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(view), `Logs\n\u001b]8;;https://wrong.test\u0007`) {
		t.Fatalf("label controls were not shown as data: %q", view)
	}

	for _, line := range uv.NewStyledString(view + "!").Lines(ansi.GraphemeWidth) {
		for _, cell := range line {
			if cell.Content != "!" && cell.Width > 0 && cell.Link.URL != link.URL {
				t.Fatal("caller label replaced the destination")
			}
		}
	}

	for _, target := range []string{"", "/relative", "https:opaque", "https:///no-host", "https://:443/path", "https://example.test/%zz", "file:///tmp/test", "javascript:alert(1)", "terminal://session/worker", "https://example.test/with space", "https://example.test/\n", "https://example.test/\x1b\\", "https://example.test/\a", "https://example.test/\u009c", "https://example.test/\u202e", "https://example.test/\xff"} {
		invalid := text.Link{Label: "Logs", URL: target}
		if got, err := invalid.View(palette); err == nil || got != "" {
			t.Fatalf("invalid target emitted output: %q => %q, %v", target, got, err)
		}

		if got, err := invalid.Plain(); err == nil || got != "" {
			t.Fatalf("plain accepted invalid target: %q", target)
		}
	}

	for _, label := range []string{"", " \t\n"} {
		invalid := text.Link{Label: label, URL: link.URL}
		if got, err := invalid.View(palette); err == nil || got != "" {
			t.Fatal("empty label was accepted")
		}
	}
}

func TestLinkCompositionSurvivesCellClipping(t *testing.T) {
	first := text.Link{Label: "Build 界é", URL: "https://example.test/first"}
	second := text.Link{Label: "Logs", URL: "https://example.test/second"}
	palette := theme.New(theme.Dark)

	a, err := first.View(palette)
	if err != nil {
		t.Fatal(err)
	}

	b, err := second.View(palette)
	if err != nil {
		t.Fatal(err)
	}

	for _, width := range []int{1, 4, 9, 14, 30} {
		frame := layout.Fit(a+" | "+b, layout.Size{Width: width, Height: 1}) + "!"
		if ansi.StringWidth(frame) != width+1 {
			t.Fatalf("link escaped its %d-cell allocation", width)
		}

		line := uv.NewStyledString(frame).Lines(ansi.GraphemeWidth)[0]
		for _, cell := range line {
			if cell.Content == "!" || cell.Content == "|" {
				if cell.Link.URL != "" {
					t.Fatalf("clipped link leaked at width %d onto %q", width, cell.Content)
				}
			}
		}
	}
}
