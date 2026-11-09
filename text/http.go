package text

import (
	"fmt"
	"image/color"
	"net/http"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lanej/charmed/theme"
)

// Escape preserves visible text while making terminal controls and invalid bytes inspectable.
// Newlines are retained for multiline content; tabs and other controls are visible escapes.
func Escape(value string) string {
	var out []string

	for len(value) > 0 {
		r, size := utf8.DecodeRuneInString(value)
		if r == '\t' {
			out = append(out, `\t`)
		} else if r == '\r' {
			out = append(out, `\r`)
		} else if r == utf8.RuneError && size == 1 {
			out = append(out, fmt.Sprintf("\\x%02x", value[0]))
		} else if (unicode.IsControl(r) && r != '\n') || unicode.In(r, unicode.Cf) {
			out = append(out, fmt.Sprintf("\\u%04x", r))
		} else {
			out = append(out, string(r))
		}

		value = value[size:]
	}

	return strings.Join(out, "")
}

func Inline(value string) string {
	return strings.NewReplacer("\n", "\\n", "\t", "\\t").Replace(Escape(value))
}

func HTTPMethod(method string, t theme.Theme) string {
	foreground := t.HTTP.Other

	switch method {
	case http.MethodGet:
		foreground = t.HTTP.Read
	case http.MethodPost:
		foreground = t.HTTP.Create
	case http.MethodPut, http.MethodPatch:
		foreground = t.HTTP.Change
	case http.MethodDelete:
		foreground = t.HTTP.Delete
	case http.MethodHead, http.MethodOptions:
		foreground = t.HTTP.Meta
	case http.MethodConnect, http.MethodTrace:
		foreground = t.HTTP.Other
	}

	if method == "" {
		method = "?"
	}

	return httpAtom(Inline(method), foreground, t)
}

// HTTPStatusCode preserves invalid observed codes rather than coercing them into a class.
// Zero denotes an unavailable code; request lifecycle is displayed by the owning molecule.
func HTTPStatusCode(code int, t theme.Theme) string {
	foreground, label := t.Unknown, fmt.Sprint(code)
	switch {
	case code == 0:
		label = "?"
	case code >= 100 && code < 200:
		foreground = t.Info
	case code >= 200 && code < 300:
		foreground = t.Success
	case code >= 300 && code < 400:
		foreground = t.Info
	case code >= 400 && code < 500:
		foreground = t.Warning
	case code >= 500 && code < 600:
		foreground = t.Danger
	default:
		label += "?"
	}

	return httpAtom(label, foreground, t)
}

func httpAtom(label string, foreground color.Color, t theme.Theme) string {
	style := t.Control().Bold(true)
	if foreground != nil && t.Mode != theme.Mono {
		style = style.Foreground(foreground)
	}

	return style.Render(label)
}
