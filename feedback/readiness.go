package feedback

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Readiness says whether a target may accept work; it does not assert health or traffic.
type Readiness string

const (
	ReadinessUnknown Readiness = ""
	Ready            Readiness = "Ready"
	NotReady         Readiness = "Not ready"
)

// View preserves the supplied dimension; zero and unsupported values display Unknown.
func (s Readiness) View(emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	switch s {
	case Ready:
		return text.Indicator(t.Glyph(theme.GlyphReady), "Ready", text.Success, t, emphasis, options...)
	case NotReady:
		return text.Indicator(t.Glyph(theme.GlyphNotReady), "Not ready", text.Neutral, t, emphasis, options...)
	case ReadinessUnknown:
	default:
	}

	return text.Indicator(t.Glyph(theme.GlyphUnknown), "Unknown", text.Neutral, statusTheme(t, t.Unknown, t.DimUnknown), emphasis, options...)
}
