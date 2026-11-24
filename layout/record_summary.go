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
// Wide rows place identity, condition and next action in comparable columns.
// Selection changes title emphasis without adding or removing content. Callers
// allocate the returned height or place it in a viewport. No caller data changes.
func (s RecordSummary) View(width int, t theme.Theme, selected bool) string {
	if width <= 0 {
		return ""
	}

	// Each inventory column needs enough room for a record identifier, a
	// condition, or an instruction. Narrow lists keep the same reading order.
	if width >= 96 && s.Title != "" && s.Status != "" && s.Next != nil {
		gap := strings.Repeat(" ", theme.SpaceLarge)
		available := width - 2*theme.SpaceLarge
		identityWidth, conditionWidth := available/3, available/4
		identity := s.identity(identityWidth, t, selected)
		condition := recordWrap(s.Status, conditionWidth)
		actionWidth := available - identityWidth - conditionWidth
		action := s.action(actionWidth, t)
		row := Row(identity, gap, condition, gap, action)

		if len(s.Fields) > 0 {
			return row + "\n" + recordWrap(KeyValue(s.Fields, width, t), width)
		}

		return row
	}

	parts := make([]string, 0, 4)
	if identity := s.identity(width, t, selected); identity != "" {
		parts = append(parts, identity)
	}

	if s.Status != "" {
		parts = append(parts, recordWrap(s.Status, width))
	}

	if s.Next != nil {
		parts = append(parts, s.action(width, t))
	}

	if len(s.Fields) > 0 {
		parts = append(parts, recordWrap(KeyValue(s.Fields, width, t), width))
	}

	return strings.Join(parts, "\n")
}

func (s RecordSummary) identity(width int, t theme.Theme, selected bool) string {
	parts := make([]string, 0, 2)

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

	return strings.Join(parts, "\n")
}

func (s RecordSummary) action(width int, t theme.Theme) string {
	label := text.Inline(s.Next.Label)
	if label != "" {
		label = t.Heading().Render(label+":") + " "
	}

	return recordWrap(label+s.Next.State.View(s.Next.Value, text.Normal, t), width)
}

// recordWrap shares TextRow's styled wrapping and explicit impossible-grapheme
// marker. In particular, a two-cell grapheme never escapes a one-cell allocation.
func recordWrap(content string, width int) string {
	return (TextRow{Items: []TextRowItem{{Text: content}}}).View(width)
}
