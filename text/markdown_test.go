package text_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestMarkdownRendersRichContentAcrossThemes(t *testing.T) {
	content := "# Guide\n\nUse **bold**, *emphasis*, and `theme.New`.\n\n- first\n- second\n\n[Docs](https://example.com)\n\n```go\nvar name = \"charmed\"\n```\n"
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono, theme.Dark} {
		out, err := text.Markdown(content, 60, theme.New(mode))
		if err != nil {
			t.Fatal(err)
		}

		plain := ansi.Strip(out)
		for _, want := range []string{"Guide", "bold", "emphasis", "theme.New", "first", "second", "https://example.com", `var name = "charmed"`} {
			if !strings.Contains(plain, want) {
				t.Fatalf("%s missing %q: %q", mode, want, plain)
			}
		}

		if mode == theme.Mono && plain != out {
			t.Fatal("monochrome Markdown emitted escape sequences")
		}

		if mode != theme.Mono && !strings.Contains(out, "38;2;") {
			t.Fatal("rich Markdown lost truecolor styling")
		}
	}
}

func TestCodePreservesLiteralContent(t *testing.T) {
	source := "  // literal Markdown\n  ~~~\n  ```go\n  <tag> & **bold**\n  ```\n  ~~~\n"

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)

		inline := text.InlineCode("a`b **c**", palette)
		if !strings.Contains(ansi.Strip(inline), "a`b **c**") {
			t.Fatal("inline code interpreted source")
		}

		for _, language := range []string{"text", "unknown-language", "go", "bash"} {
			out, err := text.CodeBlock(source, language, 80, palette)
			if err != nil {
				t.Fatal(err)
			}

			plain := ansi.Strip(out)
			for _, line := range strings.Split(strings.TrimSuffix(source, "\n"), "\n") {
				if !strings.Contains(plain, line) {
					t.Fatalf("%s/%s lost source %q in %q", mode, language, line, plain)
				}
			}

			if mode == theme.Mono && out != plain {
				t.Fatal("monochrome code emitted escape sequences")
			}
		}
	}
}

func TestMarkdownWrapsProse(t *testing.T) {
	for _, width := range []int{20, 40} {
		out, err := text.Markdown(strings.Repeat("word ", 30), width, theme.New(theme.Dark))
		if err != nil {
			t.Fatal(err)
		}

		for _, line := range strings.Split(out, "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatalf("line exceeds %d cells: %q", width, line)
			}
		}
	}

	for _, width := range []int{0, -1} {
		if _, err := text.Markdown("hello", width, theme.New(theme.Mono)); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCodeThemeSwitchingAndUnknownLanguage(t *testing.T) {
	source := "var name = \"charmed\""
	render := func(language string, mode theme.Mode) string {
		t.Helper()

		out, err := text.CodeBlock(source, language, 60, theme.New(mode))
		if err != nil {
			t.Fatal(err)
		}

		return out
	}
	dark := render("go", theme.Dark)
	light := render("go", theme.Light)

	if !strings.Contains(dark, "38;2;103;212;192") || !strings.Contains(light, "38;2;7;89;133") {
		t.Fatal("Go keyword colors do not match the theme-derived syntax palettes")
	}

	if dark == light {
		t.Fatal("syntax palette did not change with theme")
	}

	if again := render("go", theme.Dark); again != dark {
		t.Fatal("light rendering contaminated the dark syntax palette")
	}

	if got, want := render("unknown-language", theme.Dark), render("text", theme.Dark); got != want {
		t.Fatal("unknown lexer did not fall back to plain text")
	}
}
