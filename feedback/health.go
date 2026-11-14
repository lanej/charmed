package feedback

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Health is the supplied health-check result, independent of reachability and traffic.
type Health string

const (
	HealthUnknown Health = ""
	Healthy       Health = "Healthy"
	Unhealthy     Health = "Unhealthy"
)

// View preserves the supplied dimension; zero and unsupported values display Unknown.
func (s Health) View(emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	switch s {
	case Healthy:
		return text.Indicator(t.Glyph(theme.GlyphHealthy), "Healthy", text.Success, t, emphasis, options...)
	case Unhealthy:
		return text.Indicator(t.Glyph(theme.GlyphError), "Unhealthy", text.Danger, t, emphasis, options...)
	case HealthUnknown:
	default:
	}

	return text.Indicator(t.Glyph(theme.GlyphUnknown), "Unknown", text.Neutral, statusTheme(t, t.Unknown, t.DimUnknown), emphasis, options...)
}
