package text_test

import (
	"image/color"
	"strings"
	"sync"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/alecthomas/chroma/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestMarkdownCustomThemeAndNativeOverrides(t *testing.T) {
	palette := theme.New(theme.Dark)
	palette.Accent = lipgloss.Color("#123456")
	palette.Muted = lipgloss.Color("#456789")
	palette.Border = lipgloss.Color("#789abc")
	palette.Series = []color.Color{lipgloss.Color("#234567"), lipgloss.Color("#345678")}
	content := "# Title\n\n> Quote\n\n---\n\n```go\n// comment\nvar name = \"literal\"\nvar n = 42\n```\n"

	out, err := text.Markdown(content, 80, palette)
	if err != nil {
		t.Fatal(err)
	}

	for _, rgb := range []string{"18;52;86", "69;103;137", "120;154;188", "35;69;103", "52;86;120"} {
		if !strings.Contains(out, rgb) {
			t.Fatalf("custom theme color %s missing from %q", rgb, out)
		}
	}

	styles := text.MarkdownStyles(palette)
	styles.H1.Prefix = "CUSTOM "
	styles.Document.Margin = new(uint(0))

	syntax, err := chroma.NewStyle("caller", chroma.StyleEntries{chroma.Keyword: "#abcdef"}) //nolint:exhaustive // Test a caller overriding only the keyword role.
	if err != nil {
		t.Fatal(err)
	}

	customized, err := text.Markdown(content, 80, palette, text.MarkdownOptions{Styles: &styles, Syntax: syntax})
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(ansi.Strip(customized), "CUSTOM Title") || !strings.Contains(customized, "171;205;239") {
		t.Fatalf("native overrides ignored: %q", customized)
	}

	mono, err := text.Markdown(content, 80, theme.New(theme.Mono), text.MarkdownOptions{Styles: &styles, Syntax: syntax})
	if err != nil {
		t.Fatal(err)
	}

	if mono != ansi.Strip(mono) {
		t.Fatal("native options bypassed monochrome policy")
	}
}

func TestMarkdownStylesAreOwned(t *testing.T) {
	palette := theme.New(theme.Dark)
	first := text.MarkdownStyles(palette)
	*first.Strong.Bold = false
	*first.Document.Margin = 99

	second := text.MarkdownStyles(palette)
	if !*second.Strong.Bold || *second.Document.Margin == 99 {
		t.Fatal("style mutation leaked into another render")
	}
}

func TestMarkdownNestedLiteralCodeAndConcurrentPalettes(t *testing.T) {
	content := "> ```go\n> var x = `\\path\\*`\n> ```\n\n- item\n\n    ```go\n    var n = 42\n    ```\n\n    \\*literal\\*\n"

	var wait sync.WaitGroup
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		wait.Go(func() {
			palette := theme.New(mode)
			for range 10 {
				out, err := text.Markdown(content, 80, palette)
				if err != nil {
					t.Error(err)
					return
				}

				plain := ansi.Strip(out)
				for _, literal := range []string{"var x = `\\path\\*`", "var n = 42"} {
					if !strings.Contains(plain, literal) {
						t.Errorf("%s lost literal source %q in %q", mode, literal, plain)
					}
				}
			}
		})
	}

	wait.Wait()
}

func TestMarkdownLiteralCodeEntitiesAndHeadingAccent(t *testing.T) {
	palette := theme.New(theme.Dark)

	palette.Accent = lipgloss.Color("#123456")
	for _, mode := range []theme.Mode{theme.Dark, theme.Mono} {
		palette.Mode = mode

		out, err := text.Markdown("`&lt;tag&gt; &amp; &#65; \\*`", 80, palette)
		if err != nil {
			t.Fatal(err)
		}

		if !strings.Contains(ansi.Strip(out), "&lt;tag&gt; &amp; &#65; \\*") {
			t.Fatalf("code span lost literal entities: %q", ansi.Strip(out))
		}
	}

	palette.Mode = theme.Dark

	out, err := text.Markdown("###### Level six", 80, palette)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(out, "38;2;18;52;86") || strings.Contains(out, "38;5;35") {
		t.Fatalf("H6 ignored theme accent: %q", out)
	}
}

func TestMarkdownExtensionsUseTerminalRenderer(t *testing.T) {
	content := "| Field | Value |\n| --- | --- |\n| name | charmed |\n\n- [x] done\n\nTerm\n: Definition\n"

	out, err := text.Markdown(content, 60, theme.New(theme.Dark))
	if err != nil {
		t.Fatal(err)
	}

	plain := ansi.Strip(out)
	for _, want := range []string{"Field", "Value", "name", "charmed", "done", "Term", "Definition"} {
		if !strings.Contains(plain, want) {
			t.Fatalf("missing rendered extension content %q: %q", want, plain)
		}
	}

	for _, tag := range []string{"<table>", "<tr>", "<td>", "<input", "<dl>"} {
		if strings.Contains(plain, tag) {
			t.Fatalf("HTML renderer leaked into terminal output: %q", plain)
		}
	}
}
