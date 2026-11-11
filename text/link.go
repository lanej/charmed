package text

import (
	"errors"
	"net/url"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lanej/charmed/theme"
)

// Link associates a plain-text label with an exact HTTP or HTTPS destination.
// Rendering performs no I/O, terminal detection or opening. Applications choose
// View for hyperlink-capable output and Plain for logs or unsupported terminals.
type Link struct {
	Label string
	URL   string
}

// View renders an underlined label with the named link glyph and an OSC 8 target.
// The target covers only this atom; labels are escaped rather than interpreted
// as ANSI. Invalid links return an error and no terminal output.
func (l Link) View(t theme.Theme) (string, error) {
	if err := l.validate(); err != nil {
		return "", err
	}

	return t.Text().Foreground(t.Accent).Underline(true).Hyperlink(l.URL).
		Render(t.Glyph(theme.GlyphLink) + " " + Inline(l.Label)), nil
}

// Plain includes the full destination without emitting any terminal controls.
// Use this instead of stripping View, which would discard its hidden target.
func (l Link) Plain() (string, error) {
	if err := l.validate(); err != nil {
		return "", err
	}

	return Inline(l.Label) + " (" + l.URL + ")", nil
}

func (l Link) validate() error {
	if strings.TrimSpace(l.Label) == "" {
		return errors.New("link label must not be empty")
	}

	if !utf8.ValidString(l.URL) || strings.ContainsFunc(l.URL, func(r rune) bool {
		return unicode.IsControl(r) || unicode.IsSpace(r) || unicode.In(r, unicode.Cf)
	}) {
		return errors.New("link URL must not contain controls, whitespace or invalid UTF-8")
	}

	target, err := url.Parse(l.URL)
	if err != nil || target == nil || (target.Scheme != "http" && target.Scheme != "https") || target.Hostname() == "" || target.Opaque != "" {
		return errors.New("link URL must be an absolute HTTP or HTTPS URL with a host")
	}

	return nil
}
