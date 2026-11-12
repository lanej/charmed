package text_test

import (
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestJSONSyntaxPreservesLexemesAndSemanticColors(t *testing.T) {
	palette := theme.New(theme.Dark)
	source := "{\n  \"same\": 9007199254740993,\n  \"same\": -1.2300e+04,\n  \"literal\": \"\\u0061\\\"true\\\\\",\n  \"ok\": true\n}\n"
	view := text.Syntax(source, "json", palette)

	if ansi.Strip(view) != source {
		t.Fatalf("JSON highlighting changed source: %q", ansi.Strip(view))
	}

	for _, number := range []string{"9007199254740993", "-1.2300e+04"} {
		if !strings.Contains(view, palette.Text().Foreground(palette.Syntax.Number).Render(number)) {
			t.Fatalf("number did not retain one lexical color: %s", number)
		}
	}

	if strings.Count(view, palette.Text().Foreground(palette.Syntax.Name).Render(`"same"`)) != 2 {
		t.Fatal("duplicate object keys lost their independent lexical color")
	}

	if !strings.Contains(view, palette.Text().Foreground(palette.Syntax.String).Render(`"\u0061\"true\\"`)) || !strings.Contains(view, palette.Text().Foreground(palette.Syntax.Keyword).Render("true")) {
		t.Fatal("escaped string content was confused with a JSON keyword")
	}
}

func TestJSONSyntaxColorsRemainStableAfterIdle(t *testing.T) {
	palette := theme.New(theme.Dark)
	source := "[" + strings.Repeat(`{"id":9007199254740993,"rate":1.2300e+04,"label":"escaped\u0061"},`, 31) + "null]"
	expected := text.Syntax(source, "json", palette)
	// Exercise the upstream regex clock too: other languages still use Chroma.
	text.Syntax("const ready = true;", "javascript", palette)
	// WHY: The upstream timeout clock stops after idle and previously expired
	// in-flight JSON matches immediately on restart, changing individual colors.
	time.Sleep(3 * time.Second)

	for range 128 {
		if got := text.Syntax(source, "json", palette); got != expected {
			t.Fatal("JSON token colors changed after the lexer was idle")
		}
	}
}
