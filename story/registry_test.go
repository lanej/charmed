package story_test

import (
	"regexp"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/story"
	"github.com/lanej/charmed/theme"
)

func TestCatalogFramesAreDeterministicAndBounded(t *testing.T) {
	seen := map[string]bool{}
	brokenEscape := regexp.MustCompile(`\[(?:38|48);2;`)
	for _, d := range story.Definitions() {
		if seen[d.ID] {
			t.Fatalf("duplicate story %s", d.ID)
		}
		seen[d.ID] = true
		t.Run(d.ID, func(t *testing.T) {
			for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
				for _, size := range [][2]int{{40, 12}, {80, 24}, {120, 40}} {
					c := d.Defaults
					c.Columns, c.Rows = size[0], size[1]
					c.Theme = mode
					a, err := story.Frame(d.ID, c)
					if err != nil {
						t.Fatal(err)
					}
					b, err := story.Frame(d.ID, c)
					if err != nil {
						t.Fatal(err)
					}
					if a != b {
						t.Fatalf("nondeterministic %s frame at %v\nfirst: %q\nsecond: %q", mode, size, a, b)
					}
					if lipgloss.Width(a) != c.Columns || lipgloss.Height(a) != c.Rows {
						t.Fatalf("escaped %dx%d allocation", c.Columns, c.Rows)
					}
					if brokenEscape.MatchString(ansi.Strip(a)) {
						t.Fatalf("literal color escape leaked into content: %q", ansi.Strip(a))
					}
					if mode == theme.Mono && strings.Contains(a, ";2;") {
						t.Fatal("monochrome frame contains RGB color")
					}
					if mode == theme.Mono {
						for _, r := range ansi.Strip(a) {
							if r > 127 {
								t.Fatalf("ASCII policy leaked Unicode glyph %q", r)
							}
						}
					}
				}
			}
		})
	}
}

func TestStoriesDoNotShareState(t *testing.T) {
	c := story.DefaultConfig()
	c.Value = ""
	a, err := story.New("input/text", c, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer a.Close()
	a.Init()
	b, err := story.New("input/text", c, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()
	b.Init()
	a.Update(tea.KeyPressMsg{Code: 'x', Text: "x"})
	if a.View().Content == b.View().Content {
		t.Fatal("input did not change its own scene")
	}
	if strings.Contains(ansi.Strip(b.View().Content), "> x") {
		t.Fatal("input state leaked between scenes")
	}
}

func TestRejectsInvalidStoryConfiguration(t *testing.T) {
	c := story.DefaultConfig()
	c.Columns = 100000
	if _, err := story.New("input/text", c, nil); err == nil {
		t.Fatal("unbounded dimensions accepted")
	}

	if _, err := story.New("not/a/story", story.DefaultConfig(), nil); err == nil {
		t.Fatal("unregistered story accepted")
	}
}

func TestNarrowPillSpecimenKeepsEveryToneAtEachEmphasis(t *testing.T) {
	c := story.DefaultConfig()
	c.Columns, c.Rows = 40, 16

	frame, err := story.Frame("text/pill", c)
	if err != nil {
		t.Fatal(err)
	}

	plain := ansi.Strip(frame)
	for _, tone := range []string{"neutral", "success", "info", "warning", "danger"} {
		if strings.Count(plain, tone) != 3 {
			t.Fatalf("narrow specimen clipped %s: %s", tone, plain)
		}
	}
}

func TestNarrowFoundationCatalogKeepsAllGlyphsAndSelection(t *testing.T) {
	c := story.DefaultConfig()
	c.Columns, c.Rows = 40, 16

	for _, mode := range []theme.Mode{theme.Dark, theme.Mono} {
		c.Theme = mode

		glyphConfig := c
		glyphConfig.Rows = (len(theme.Glyphs())+1)/2 + 2

		glyphs, err := story.Frame("foundations/glyphs", glyphConfig)
		if err != nil {
			t.Fatal(err)
		}

		for _, glyph := range theme.Glyphs() {
			if !strings.Contains(ansi.Strip(glyphs), string(glyph)) {
				t.Fatalf("narrow %s catalog clipped %s", mode, glyph)
			}
		}

		tokens, err := story.Frame("foundations/tokens", c)
		if err != nil {
			t.Fatal(err)
		}

		for _, label := range []string{"none", "small", "medium", "large", "Normal", "Focused", "selected | Value |"} {
			if !strings.Contains(ansi.Strip(tokens), label) {
				t.Fatalf("narrow %s token specimen clipped %s", mode, label)
			}
		}
	}
}
