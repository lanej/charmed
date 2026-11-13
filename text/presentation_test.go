package text_test

import (
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestSymbolPresentationKeepsIndependentGlyphAndLabelChoices(t *testing.T) {
	for _, example := range []struct {
		mode text.Presentation
		want string
	}{
		{text.DefaultPresentation, "! WARN"}, {text.GlyphAndLabel, "! WARN"}, {text.GlyphOnly, "!"}, {text.Compact, "!"}, {text.LabelOnly, "WARN"}, {text.Presentation("invalid"), "! WARN"},
	} {
		if got := example.mode.Format("!", "WARN"); got != example.want {
			t.Fatalf("%s: %q != %q", example.mode, got, example.want)
		}
	}

	if got := text.GlyphAndLabel.Format("", " spaced "); got != " spaced " {
		t.Fatal("label whitespace changed", got)
	}

	if got := text.LabelOnly.Format("!", "\x1b[31m"); got != `\u001b[31m` {
		t.Fatal("unescaped label", got)
	}

	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Mono} {
		palette := theme.New(mode)
		if got := ansi.Strip(text.Trace("Trace", palette, text.LabelOnly)); got != "Trace" {
			t.Fatal("trace label mode", got)
		}

		if got := ansi.Strip(text.Disclosure("Response", true, palette, text.GlyphOnly)); got != palette.Glyph(theme.GlyphExpanded) {
			t.Fatal("disclosure glyph mode", got)
		}

		if got := ansi.Strip(text.PendingWithPresentation(text.LabelOnly, palette, 2)); got != " PENDING " {
			t.Fatal("pending label mode", got)
		}
	}
}

func TestURLFormatsPreserveEscapedPathsAndQueryOrder(t *testing.T) {
	captured := "https://viewer@example.test:8443/a%2Fb?z=2&a=x%2By#detail"
	for _, example := range []struct {
		format text.URLFormat
		want   string
	}{
		{text.DefaultURLFormat, "/a%2Fb?z=2&a=x%2By"}, {text.PathOnly, "/a%2Fb"}, {text.PathAndQuery, "/a%2Fb?z=2&a=x%2By"}, {text.HostAndPath, "example.test:8443/a%2Fb?z=2&a=x%2By"}, {text.FullURL, captured}, {text.URLFormat("invalid"), captured},
	} {
		if got := text.FormatURL(captured, example.format); got != example.want {
			t.Fatalf("%s: %q != %q", example.format, got, example.want)
		}
	}

	for _, value := range []string{"http://%zz", "host:443", "%broken", "\x1b[31m"} {
		if got := text.FormatURL(value, text.PathOnly); got != text.Inline(value) {
			t.Fatal("malformed/opaque target changed", got)
		}
	}

	for value, want := range map[string]string{"/path?": "/path?", "https://example.test": "/", "/a%2Fb?q=1&q=2": "/a%2Fb?q=1&q=2", "?query=1": "?query=1"} {
		if got := text.FormatURL(value, text.PathAndQuery); got != want {
			t.Fatal("relative/root target changed", got, want)
		}
	}
}
