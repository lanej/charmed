package feedback

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// CanaryMembership is supplied routing membership for a canary target.
// On describes the new-version target in the serving pool; off describes its exclusion.
// Membership is independent of the target's health, readiness and reachability.
type CanaryMembership string

const (
	CanaryMembershipUnknown CanaryMembership = ""
	CanaryOn                CanaryMembership = "Canary on"
	CanaryOff               CanaryMembership = "Canary off"
)

// View renders pool membership as a labeled pill or a compact glyph.
// Zero and unsupported values remain Unknown rather than implying removal.
func (s CanaryMembership) View(emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	switch s {
	case CanaryOn:
		return text.Indicator(canaryMark(theme.GlyphCanaryOn, t, options...), "Canary on", text.Info, t, emphasis, options...)
	case CanaryOff:
		return text.Indicator(canaryMark(theme.GlyphCanaryOff, t, options...), "Canary off", text.Neutral, statusTheme(t, t.Inactive, t.DimInactive), emphasis, options...)
	case CanaryMembershipUnknown:
	default:
	}

	return text.Indicator(t.Glyph(theme.GlyphUnknown), "Unknown", text.Neutral, statusTheme(t, t.Unknown, t.DimUnknown), emphasis, options...)
}

func canaryMark(glyph theme.Glyph, t theme.Theme, options ...text.Presentation) string {
	if len(options) > 0 && options[0] == text.Compact {
		return t.Glyph(glyph)
	}

	return t.FullGlyph(glyph)
}
