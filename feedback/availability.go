package feedback

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Availability is supplied reachability, independent of readiness and health.
type Availability string

const (
	AvailabilityUnknown Availability = ""
	Up                  Availability = "Up"
	Down                Availability = "Down"
)

// View preserves the supplied dimension; zero and unsupported values display Unknown.
func (s Availability) View(emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	switch s {
	case Up:
		return text.Indicator(t.Glyph(theme.GlyphUp), "Up", text.Success, t, emphasis, options...)
	case Down:
		return text.Indicator(t.Glyph(theme.GlyphDown), "Down", text.Danger, t, emphasis, options...)
	case AvailabilityUnknown:
	default:
	}

	return text.Indicator(t.Glyph(theme.GlyphUnknown), "Unknown", text.Neutral, statusTheme(t, t.Unknown, t.DimUnknown), emphasis, options...)
}
