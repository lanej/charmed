package text

import "net/url"

type URLFormat string

const (
	DefaultURLFormat URLFormat = ""
	FullURL          URLFormat = "full"
	PathOnly         URLFormat = "path"
	PathAndQuery     URLFormat = "path-query"
	HostAndPath      URLFormat = "host-path"
)

// FormatURL changes only display. Escaped paths and query ordering are retained;
// malformed or opaque targets fall back to the captured value without guessing.
func FormatURL(value string, format URLFormat) string {
	switch format {
	case DefaultURLFormat:
		format = PathAndQuery
	case FullURL:
		return Inline(value)
	case PathOnly, PathAndQuery, HostAndPath:
	default:
		return Inline(value)
	}

	parsed, err := url.Parse(value)
	if err != nil || parsed.Opaque != "" {
		return Inline(value)
	}

	path := parsed.EscapedPath()
	if path == "" && parsed.Host != "" {
		path = "/"
	}

	if format != PathOnly && (parsed.RawQuery != "" || parsed.ForceQuery) {
		path += "?" + parsed.RawQuery
	}

	if format == HostAndPath {
		path = parsed.Host + path
	}

	return Inline(path)
}
