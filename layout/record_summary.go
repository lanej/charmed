package layout

import (
	"strings"

	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// RecordSummary gives a record a consistent title, state, next step and metadata.
// Title, Subtitle and Status accept trusted, optionally styled display text, as
// TextRow does. Escape application data before styling it. Fields and Next use
// Pair's escaped values and explicit presence states; a nil Next is omitted.
// Applications supply status meaning and actions. Rendering performs no effects.
type RecordSummary struct {
	Title, Subtitle, Status string
	Fields                  []Pair
	Next                    *Pair
}

// View wraps complete content to the supplied cell width with unbounded height.
// Selection changes title emphasis without adding or removing content. Callers
// allocate the returned height or place it in a viewport. No caller data changes.
func (s RecordSummary) View(width int, t theme.Theme, selected bool) string {
	if width <= 0 {
		return ""
	}

	parts := make([]string, 0, 5)

	if s.Title != "" {
		style := t.Text().Bold(true)
		if selected {
			style = t.Heading()
		}

		parts = append(parts, recordWrap(style.Render(s.Title), width))
	}

	if s.Subtitle != "" {
		parts = append(parts, recordWrap(t.Subtle().Render(s.Subtitle), width))
	}

	if s.Status != "" {
		parts = append(parts, recordWrap(s.Status, width))
	}

	if s.Next != nil {
		label := text.Inline(s.Next.Label)
		if label != "" {
			label = t.Heading().Render(label+":") + " "
		}

		parts = append(parts, recordWrap(label+s.Next.State.View(s.Next.Value, text.Normal, t), width))
	}

	if len(s.Fields) > 0 {
		parts = append(parts, recordWrap(KeyValue(s.Fields, width, t), width))
	}

	return strings.Join(parts, "\n")
}

// recordWrap shares TextRow's styled wrapping and explicit impossible-grapheme
// marker. In particular, a two-cell grapheme never escapes a one-cell allocation.
func recordWrap(content string, width int) string {
	return (TextRow{Items: []TextRowItem{{Text: content}}}).View(width)
}
