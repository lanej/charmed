package text

import (
	"errors"
	"net/url"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/lanej/charmed/theme"
)

// Link associates a plain-text label with an exact absolute URI destination.
// Rendering performs no I/O, terminal detection or opening. Applications choose
// View for hyperlink-capable output and Plain for logs or unsupported terminals.
// Scheme handling and permission to open a target belong to the application.
type Link struct {
	Label        string
	URL          string
	Presentation Presentation
	Emphasis     Emphasis
}

// View renders an underlined label with the named link glyph and an OSC 8 target.
// The target covers only this atom; labels are escaped rather than interpreted
// as ANSI. Presentation and emphasis change only the label, never its target.
// Invalid links return an error and no terminal output.
func (l Link) View(t theme.Theme) (string, error) {
	if _, err := l.validate(); err != nil {
		return "", err
	}

	style := t.LinkStyle().Bold(l.Emphasis == Bright)
	if l.Emphasis == Dim && t.DimLink != nil {
		style = style.Foreground(t.DimLink)
	}

	return style.Hyperlink(l.URL).Render(l.Presentation.Format(t.Glyph(theme.GlyphLink), l.Label)), nil
}

// Plain includes the full destination without emitting any terminal controls.
// Use this instead of stripping View, which would discard its hidden target.
func (l Link) Plain() (string, error) {
	if _, err := l.validate(); err != nil {
		return "", err
	}

	return Inline(l.Label) + " (" + l.URL + ")", nil
}

// NewWebLink applies web-specific validation, then returns the general Link value.
// Both helpers share View and Plain; the destination is never normalized.
func NewWebLink(label, target string) (Link, error) {
	link := Link{Label: label, URL: target}

	parsed, err := link.validate()
	if err != nil {
		return Link{}, err
	}

	if (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Hostname() == "" || parsed.Opaque != "" {
		return Link{}, errors.New("web link URL must be an absolute HTTP or HTTPS URL with a host")
	}

	return link, nil
}

func (l Link) validate() (*url.URL, error) {
	if strings.TrimSpace(l.Label) == "" {
		return nil, errors.New("link label must not be empty")
	}

	if !utf8.ValidString(l.URL) || strings.ContainsFunc(l.URL, func(r rune) bool {
		return unicode.IsControl(r) || unicode.IsSpace(r) || unicode.In(r, unicode.Cf)
	}) {
		return nil, errors.New("link URL must not contain controls, whitespace or invalid UTF-8")
	}

	// rationale: url.Parse leaves opaque and query escapes unchecked. Validate
	// escapes across the original URI without using or emitting its decoded form.
	if _, err := url.PathUnescape(l.URL); err != nil {
		return nil, errors.New("link URL contains an invalid percent escape")
	}

	target, err := url.Parse(l.URL)
	if err != nil || target == nil || !target.IsAbs() {
		return nil, errors.New("link URL must be an absolute URI with a scheme")
	}

	return target, nil
}
