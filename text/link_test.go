package text_test

import (
	"image/color"
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

	for _, target := range []string{"", "/relative", "//example.test/path", "1bad:target", "https://[::", "https://example.test/%zz", "https://example.test/?q=%zz", "mailto:dev@example.test?subject=%zz", "custom:%zz", "custom:incomplete%", "https://example.test/with space", "https://example.test/\n", "https://example.test/\x1b\\", "https://example.test/\a", "https://example.test/\u009c", "https://example.test/\u202e", "https://example.test/\xff"} {
		invalid := text.Link{Label: "Logs", URL: target}
		if got, err := invalid.View(palette); err == nil || got != "" {
			t.Fatalf("invalid target emitted output: %q => %q, %v", target, got, err)
		}

		if got, err := invalid.Plain(); err == nil || got != "" {
			t.Fatalf("plain accepted invalid target: %q", target)
		}

		if got, err := text.NewWebLink("Logs", target); err == nil || got != (text.Link{}) {
			t.Fatalf("web helper accepted invalid target: %q", target)
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

func TestLinkSupportsHierarchicalAndOpaqueURIs(t *testing.T) {
	palette := theme.New(theme.Dark)

	for _, target := range []string{
		"file:///tmp/build%20output.log",
		"ssh://dev@example.test:2222/work",
		"mailto:dev@example.test?subject=Build%2042",
		"urn:example:build:42",
		"terminal://session/build/pane/tests",
		"custom:opaque-destination",
	} {
		link := text.Link{Label: "Destination", URL: target}

		view, err := link.View(palette)
		if err != nil {
			t.Fatalf("%s: %v", target, err)
		}

		for _, cell := range uv.NewStyledString(view).Lines(ansi.GraphemeWidth)[0] {
			if cell.Width > 0 && cell.Link.URL != target {
				t.Fatalf("general link changed target %q to %q", target, cell.Link.URL)
			}
		}

		plain, err := link.Plain()
		if err != nil || plain != "Destination ("+target+")" {
			t.Fatalf("general link lost plain destination: %q, %v", plain, err)
		}
	}
}

func TestWebLinkLayersValidationOverGeneralLink(t *testing.T) {
	for _, target := range []string{"https://example.test/a%2Fb?b=2&a=1#output", "http://localhost:8080/logs", "HTTPS://EXAMPLE.test/logs"} {
		link, err := text.NewWebLink("Logs", target)
		if err != nil {
			t.Fatal(err)
		}

		if link != (text.Link{Label: "Logs", URL: target}) {
			t.Fatalf("web helper changed the general link value: %#v", link)
		}
	}

	for _, target := range []string{"https:opaque", "https:///no-host", "https://:443/path", "file:///tmp/test", "mailto:dev@example.test", "urn:example:42", "terminal://session/worker", "javascript:alert(1)"} {
		if got, err := text.NewWebLink("Logs", target); err == nil || got != (text.Link{}) {
			t.Fatalf("web helper accepted %q: %#v, %v", target, got, err)
		}
	}

	if got, err := text.NewWebLink("", "https://example.test"); err == nil || got != (text.Link{}) {
		t.Fatal("web helper skipped shared label validation")
	}
}

func TestLinkPresentationAndEmphasisKeepTargetAndGeometry(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		for _, presentation := range []text.Presentation{text.DefaultPresentation, text.GlyphAndLabel, text.GlyphOnly, text.LabelOnly, text.Compact} {
			expected := palette.Glyph(theme.GlyphLink) + " Logs"

			switch presentation {
			case text.DefaultPresentation, text.GlyphAndLabel:
			case text.GlyphOnly, text.Compact:
				expected = palette.Glyph(theme.GlyphLink)
			case text.LabelOnly:
				expected = "Logs"
			default:
				t.Fatal("unknown test presentation")
			}

			for _, emphasis := range []text.Emphasis{text.Normal, text.Dim, text.Bright} {
				link := text.Link{Label: "Logs", URL: "https://example.test/logs", Presentation: presentation, Emphasis: emphasis}

				view, err := link.View(palette)
				if err != nil {
					t.Fatal(err)
				}

				if ansi.Strip(view) != expected || ansi.StringWidth(view) != ansi.StringWidth(expected) {
					t.Fatalf("presentation/emphasis changed label geometry: %q", view)
				}

				foreground := palette.Link
				if emphasis == text.Dim {
					foreground = palette.DimLink
				}

				for _, cell := range uv.NewStyledString(view).Lines(ansi.GraphemeWidth)[0] {
					if cell.Width == 0 {
						continue
					}

					if cell.Link.URL != link.URL || cell.Style.Underline == uv.UnderlineNone {
						t.Fatal("presentation or emphasis removed the target or underline")
					}

					if cell.Content != " " && (cell.Style.Attrs&uv.AttrBold != 0) != (emphasis == text.Bright) {
						t.Fatalf("%s/%s/%s cell %q attributes %v lost distinct weight", mode, presentation, emphasis, cell.Content, cell.Style.Attrs)
					}

					if color.NRGBAModel.Convert(cell.Style.Fg) != color.NRGBAModel.Convert(foreground) {
						t.Fatalf("%s/%s ignored the semantic link color", mode, emphasis)
					}
				}

				plain, err := link.Plain()
				if err != nil || plain != "Logs (https://example.test/logs)" {
					t.Fatal("visual options hid the plain label or destination")
				}
			}
		}
	}
}
