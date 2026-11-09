package collection

import (
	"mime"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/lanej/charmed/text"
)

// Parameters are decoded from supplied observations; route parameters require
// application routing knowledge and are not inferred from path segments.
func (m *HTTPLog) parameterSources(record HTTPExchange) (string, string) {
	target, _, _ := strings.Cut(record.URL, "#")
	_, query, _ := strings.Cut(target, "?")
	form := ""
	body := bodyWithContentType(record.RequestBody, record.RequestHeaders)
	media, params, err := mime.ParseMediaType(body.ContentType)

	charset := strings.ToLower(params["charset"])
	if err == nil && media == "application/x-www-form-urlencoded" && body.State == Captured && (body.Encoding == "" || strings.EqualFold(body.Encoding, "identity")) && (charset == "" || charset == "utf-8" || charset == "us-ascii") && len(body.Bytes) <= max(1, m.BodyPreviewBytes) {
		form = string(body.Bytes)
	}

	return query, form
}

func (m *HTTPLog) HasParameters() bool {
	record, ok := m.Selected()
	if !ok {
		return false
	}

	query, form := m.parameterSources(record)

	return query != "" || form != ""
}

func (m *HTTPLog) parameterView(record HTTPExchange, width int) string {
	query, form := m.parameterSources(record)

	if query == "" && form == "" {
		return ""
	}

	heading := m.Theme.Heading().Render("Parameters")
	if m.Folding {
		if marker := text.Disclosure("", m.ParametersOpen, m.Theme, m.Symbols); marker != "" {
			heading = marker + " " + heading
		}
	}

	if m.Folding && !m.ParametersOpen {
		return heading
	}

	rows := []string{heading}

	for _, source := range []struct{ name, raw string }{{"Query", query}, {"Form", form}} {
		if source.raw == "" {
			continue
		}

		rows = append(rows, m.Theme.Subtle().Render(source.name))

		values, parseErr := url.ParseQuery(source.raw)
		if parseErr != nil {
			rows = append(rows, m.Theme.Subtle().Render("Invalid parameter encoding · captured source"), text.Inline(source.raw))
			continue
		}

		keys := make([]string, 0, len(values))
		for key := range values {
			keys = append(keys, key)
		}

		slices.Sort(keys)

		for _, key := range keys {
			for _, value := range values[key] {
				rows = append(rows, hangingWrap(logStyle(m.Theme, m.Theme.Syntax.Name).Render(text.Inline(key))+": ", logStyle(m.Theme, m.Theme.Syntax.String).Render(text.Inline(strconv.Quote(value))), width))
			}
		}
	}

	return strings.Join(rows, "\n")
}
