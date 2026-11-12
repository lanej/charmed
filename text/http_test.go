package text_test

import (
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func TestHTTPAtomsKeepExactObservations(t *testing.T) {
	for _, mode := range []theme.Mode{theme.Dark, theme.Light, theme.Contrast, theme.Mono} {
		palette := theme.New(mode)
		for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "CONNECT", "TRACE", "PROPFIND", "QUERY", "get"} {
			if ansi.Strip(text.HTTPMethod(method, palette)) != " "+method+" " {
				t.Fatal("method changed or lost padding", method)
			}
		}

		for _, code := range []int{100, 199, 200, 299, 300, 399, 400, 499, 500, 599, 0, 799} {
			got := ansi.Strip(text.HTTPStatusCode(code, palette))
			if code == 0 && got != " ? " {
				t.Fatal(got)
			}

			if code == 799 && got != " 799? " {
				t.Fatal(got)
			}
		}
	}
}

func TestContentSyntaxPreservesEscapedSourceAndContentTypeDispatch(t *testing.T) {
	for contentType, language := range map[string]string{"application/problem+json; charset=utf-8": "json", "application/json": "json", "application/atom+xml": "xml", "text/html": "html", "text/css": "css", "text/javascript": "javascript", "application/x-www-form-urlencoded": "form", "text/plain": "text", "image/png": "", "broken": ""} {
		if text.ContentLanguage(contentType) != language {
			t.Fatalf("%s did not select %s", contentType, language)
		}
	}

	for _, language := range []string{"json", "xml", "html", "css", "javascript", "text", "form", "unknown"} {
		for _, source := range []string{"{\"x\": 1}", "{\n  \"x\": true\n}\n", "<node attr=\"v\"/>\n", "\x1b]52;c;payload\a\r\x9b", "abc\xff", "a= x\n", "a=one\n\ntwo&b=3"} {
			got := text.Syntax(source, language, theme.New(theme.Dark))
			if ansi.Strip(got) != text.Escape(source) {
				t.Fatalf("%s changed source: %q != %q", language, ansi.Strip(got), text.Escape(source))
			}

			if strings.Contains(ansi.Strip(got), "\x1b") {
				t.Fatal("terminal control escaped styling boundary")
			}
		}
	}

	colored := text.Syntax("{\"x\": 1}", "json", theme.New(theme.Dark))
	if !strings.Contains(colored, "38;2;") {
		t.Fatal("JSON did not receive lexical colors")
	}
}

func TestUntrustedTextControlsAreVisibleData(t *testing.T) {
	if got := text.Inline("\x1b[31mA\n\t\r"); got != `\u001b[31mA\n\t\r` {
		t.Fatalf("unsafe terminal text: %q", got)
	}
}
