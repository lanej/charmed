package feedback

import (
	"math"
	"strings"

	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Follow distinguishes live-follow mode from paused history browsing.
func Follow(following bool, emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	if !following {
		return Paused.View(emphasis, t, options...)
	}

	mode := text.DefaultPresentation
	if len(options) > 0 {
		mode = options[0]
	}

	return text.Pill(mode.Format(t.Glyph(theme.GlyphRunning), "FOLLOW"), text.Neutral, statusTheme(t, t.Running, t.DimRunning), emphasis)
}

// Scrollbar describes vertical line geometry, not selection or elapsed time.
// Offset is clamped; no overflow retains a blank column to prevent reflow.
type Scrollbar struct{ Total, Visible, Offset, Height int }

func (s Scrollbar) View(t theme.Theme) string {
	if s.Height <= 0 {
		return ""
	}

	rows := make([]string, s.Height)
	for i := range rows {
		rows[i] = " "
	}

	if s.Total <= max(0, s.Visible) || s.Total <= 0 || s.Visible <= 0 {
		return strings.Join(rows, "\n")
	}

	thumb := max(1, min(s.Height, int(math.Round(float64(s.Visible)/float64(s.Total)*float64(s.Height)))))
	offset := min(max(0, s.Offset), s.Total-s.Visible)

	start := int(math.Round(float64(offset) / float64(s.Total-s.Visible) * float64(s.Height-thumb)))
	for i := range rows {
		if i >= start && i < start+thumb {
			rows[i] = t.Heading().Render(t.Glyph(theme.GlyphScrollThumb))
		} else {
			rows[i] = t.Subtle().Render(t.Glyph(theme.GlyphScrollTrack))
		}
	}

	return strings.Join(rows, "\n")
}
