package text

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/theme"
	"github.com/rivo/uniseg"
)

// Identifier stores the complete identity. Label abbreviates its escaped display
// to cells, preserving both ends when possible; never use Label as an identity key.
type Identifier struct {
	Value    string
	Emphasis Emphasis
}

func (id Identifier) Label(width int) string {
	if width <= 0 {
		return ""
	}

	value := Inline(id.Value)
	if value == "" {
		value = "unavailable"
	}

	if ansi.StringWidth(value) <= width {
		return value
	}
	// ASCII dots keep abbreviation portable across terminal glyph policies.
	marker := "..."
	if width < len(marker) {
		return marker[:width]
	}

	remaining := width - len(marker)
	head := (remaining + 1) / 2
	tail := remaining - head

	return ansi.Truncate(value, head, "") + marker + identifierSuffix(value, tail)
}

func identifierSuffix(value string, width int) string {
	var clusters []string

	iterator := uniseg.NewGraphemes(value)
	for iterator.Next() {
		clusters = append(clusters, iterator.Str())
	}

	start := len(clusters)
	for start > 0 {
		cells := ansi.StringWidth(clusters[start-1])
		if cells > width {
			break
		}

		width -= cells
		start--
	}

	return strings.Join(clusters[start:], "")
}

func (id Identifier) View(width int, t theme.Theme) string {
	return ToneStyle(Neutral, id.Emphasis, t).Render(id.Label(width))
}
