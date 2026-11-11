package text

import (
	"strings"
	"unicode/utf8"

	"charm.land/lipgloss/v2"
	"github.com/rivo/uniseg"
)

// MatchRange is a half-open range of rune offsets into the original, unescaped text.
// Overlaps are combined; out-of-range portions are ignored. Each original rune is
// escaped before rendering so highlighting cannot turn source text into terminal controls.
type MatchRange struct{ Start, End int }

// HighlightMatches renders caller-computed matches without owning search policy.
// Styles are native Lip Gloss values; selection backgrounds and monochrome remain caller-owned.
func HighlightMatches(value string, ranges []MatchRange, normal, matched lipgloss.Style) string {
	runes := []rune(value)

	selected := make([]bool, len(runes))
	for _, span := range ranges {
		for i := max(0, span.Start); i < min(len(runes), span.End); i++ {
			selected[i] = true
		}
	}

	parts := make([]string, 0, len(runes))
	clusters := uniseg.NewGraphemes(value)
	offset := 0

	for clusters.Next() {
		cluster := clusters.Str()
		end := offset + utf8.RuneCountInString(cluster)
		style := normal

		for i := offset; i < end; i++ {
			if selected[i] {
				style = matched
				break
			}
		}

		parts = append(parts, style.Render(Inline(cluster)))
		offset = end
	}

	return strings.Join(parts, "")
}
