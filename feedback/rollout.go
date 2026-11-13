package feedback

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Rollout describes production traffic exposure, independent of health and log-follow mode.
type Rollout string

const (
	RolloutUnknown Rollout = ""
	Canary         Rollout = "Canary"
	Live           Rollout = "Live"
)

// View preserves the supplied dimension; zero and unsupported values display Unknown.
func (s Rollout) View(emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	switch s {
	case Canary:
		return text.Indicator(t.Glyph(theme.GlyphCanary), "Canary", text.Info, t, emphasis, options...)
	case Live:
		return text.Indicator(t.Glyph(theme.GlyphLive), "Live", text.Info, t, emphasis, options...)
	case RolloutUnknown:
	default:
	}

	return text.Indicator(t.Glyph(theme.GlyphUnknown), "Unknown", text.Neutral, statusTheme(t, t.Unknown, t.DimUnknown), emphasis, options...)
}
