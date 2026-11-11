package text

import (
	"fmt"

	"github.com/lanej/charmed/theme"
)

type SortDirection string

const (
	SortNone       SortDirection = ""
	SortAscending  SortDirection = "ascending"
	SortDescending SortDirection = "descending"
)

// SortIndicator's Priority is one-based precedence, or zero to omit it.
// A direction remains explicit in ASCII output; unsorted never implies ascending.
type SortIndicator struct {
	Direction SortDirection
	Priority  int
	Emphasis  Emphasis
}

func (s SortIndicator) View(t theme.Theme) string {
	label := "UNSORTED"

	switch s.Direction {
	case SortNone:
	case SortAscending:
		label = t.Glyph(theme.GlyphUp) + " ASC"
	case SortDescending:
		label = t.Glyph(theme.GlyphDown) + " DESC"
	default:
		label = "UNKNOWN SORT"
	}

	if s.Priority > 0 && (s.Direction == SortAscending || s.Direction == SortDescending) {
		label += fmt.Sprintf(" %d", s.Priority)
	}

	return ToneStyle(Neutral, s.Emphasis, t).Render(label)
}

type ChangeKind string

const (
	Unchanged ChangeKind = ""
	Added     ChangeKind = "added"
	Modified  ChangeKind = "modified"
	Removed   ChangeKind = "removed"
)

func (c ChangeKind) View(emphasis Emphasis, t theme.Theme) string {
	label, tone := "= unchanged", Neutral

	switch c {
	case Unchanged:
	case Added:
		label, tone = "+ added", Success
	case Modified:
		label, tone = "~ modified", Warning
	case Removed:
		label, tone = "- removed", Danger
	default:
		label = "? unknown change"
	}

	return ToneStyle(tone, emphasis, t).Render(label)
}

type OmissionKind string

const (
	DisplayHidden   OmissionKind = ""
	SourceTruncated OmissionKind = "source-truncated"
)

// Omission describes suppressed evidence. Count is the omitted item count when
// KnownCount is true; a known zero renders nothing. Unknown counts stay explicit.
type Omission struct {
	Kind       OmissionKind
	Count      int
	KnownCount bool
	Emphasis   Emphasis
}

func (o Omission) Label() string {
	if o.KnownCount && o.Count < 0 {
		return "invalid omission count"
	}

	if o.KnownCount && o.Count == 0 {
		return ""
	}

	switch o.Kind {
	case DisplayHidden:
		if o.KnownCount {
			return fmt.Sprintf("+%d more", o.Count)
		}

		return "more hidden"
	case SourceTruncated:
		if o.KnownCount {
			return fmt.Sprintf("source truncated (%d omitted)", o.Count)
		}

		return "source truncated (count unknown)"
	default:
		return "unknown omission kind"
	}
}

func (o Omission) View(t theme.Theme) string {
	return ToneStyle(Neutral, o.Emphasis, t).Render(o.Label())
}
