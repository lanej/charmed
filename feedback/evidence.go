package feedback

import (
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

// Evidence is the state a measurement is presented under. Only EvidenceHealthy is healthy: missing, thin, aged,
// cut-short and unattributed evidence each keep their own word and glyph, and none uses the OK status, glyph or
// success tone, so an absent measurement cannot pass as good. The zero value renders as unknown.
type Evidence string

const (
	EvidenceHealthy        Evidence = "HEALTHY"
	EvidenceDegraded       Evidence = "DEGRADED"
	EvidenceDown           Evidence = "DOWN"
	EvidenceUnknown        Evidence = "UNKNOWN"
	EvidenceStale          Evidence = "STALE"
	EvidenceInsufficient   Evidence = "INSUFFICIENT"
	EvidenceTruncated      Evidence = "TRUNCATED"
	EvidenceAttributionGap Evidence = "ATTRIBUTION GAP"
)

// Evidences lists every state in learning order.
func Evidences() []Evidence {
	return []Evidence{EvidenceHealthy, EvidenceDegraded, EvidenceDown, EvidenceUnknown, EvidenceStale, EvidenceInsufficient, EvidenceTruncated, EvidenceAttributionGap}
}

// severity orders states worst first for summaries: failure, missing evidence, aged or cut-short evidence, healthy.
func severity() []Evidence {
	return []Evidence{EvidenceDown, EvidenceDegraded, EvidenceUnknown, EvidenceInsufficient, EvidenceAttributionGap, EvidenceStale, EvidenceTruncated, EvidenceHealthy}
}

// Worst is the least trustworthy state among counts with a nonzero N, so a summary never reads healthier than its
// weakest evidence. A state this package does not know counts as unknown, and no evidence at all is unknown.
func Worst(counts []EvidenceCount) Evidence {
	for _, s := range severity() {
		for _, c := range counts {
			if c.N > 0 && (c.State == s || (s == EvidenceUnknown && !known(c.State))) {
				return s
			}
		}
	}

	return EvidenceUnknown
}

// ParseEvidence reads a state word once at a boundary; anything else is not a state.
func ParseEvidence(value string) (Evidence, bool) {
	e := Evidence(strings.ToUpper(strings.TrimSpace(value)))

	for _, known := range Evidences() {
		if e == known {
			return e, true
		}
	}

	return EvidenceUnknown, false
}

// Status maps onto the display statuses. Only healthy is OK; evidence gaps are UNKNOWN rather than OFF so they never
// read as a deliberate off switch, and aged or cut-short evidence is WARN.
func (e Evidence) Status() Status {
	switch e {
	case EvidenceHealthy:
		return OK
	case EvidenceDown:
		return Error
	case EvidenceDegraded, EvidenceStale, EvidenceTruncated:
		return Warn
	case EvidenceUnknown, EvidenceInsufficient, EvidenceAttributionGap:
		return Unknown
	default:
		return Unknown
	}
}

// Label is the lowercase word that accompanies the glyph, so meaning survives without color.
func (e Evidence) Label() string {
	switch e {
	case EvidenceHealthy, EvidenceDegraded, EvidenceDown, EvidenceUnknown, EvidenceStale, EvidenceInsufficient, EvidenceTruncated, EvidenceAttributionGap:
		return strings.ToLower(string(e))
	default:
		return "unknown"
	}
}

func (e Evidence) mark() (text.Tone, theme.Glyph) {
	switch e {
	case EvidenceHealthy:
		return text.Success, theme.GlyphOK
	case EvidenceDegraded:
		return text.Warning, theme.GlyphWarning
	case EvidenceDown:
		return text.Danger, theme.GlyphError
	case EvidenceStale:
		return text.Warning, theme.GlyphClock
	case EvidenceTruncated:
		return text.Warning, theme.GlyphTruncated
	case EvidenceInsufficient:
		return text.Neutral, theme.GlyphMissing
	case EvidenceAttributionGap:
		return text.Neutral, theme.GlyphAttribution
	case EvidenceUnknown:
		return text.Neutral, theme.GlyphUnknown
	default:
		return text.Neutral, theme.GlyphUnknown
	}
}

// View is the state's pill. Gaps use the muted unknown palette, as the UNKNOWN status does.
func (e Evidence) View(emphasis text.Emphasis, t theme.Theme, options ...text.Presentation) string {
	tone, glyph := e.mark()
	if tone == text.Neutral {
		t = statusTheme(t, t.Unknown, t.DimUnknown)
	}

	mode := text.DefaultPresentation
	if len(options) > 0 {
		mode = options[0]
	}

	return text.Pill(mode.Format(t.Glyph(glyph), strings.ToUpper(e.Label())), tone, t, emphasis)
}

// EvidenceCount is the number of items in one evidence state.
type EvidenceCount struct {
	State Evidence
	N     uint64
}

// SummaryLine reads "× 1 down · ◌ 2 unknown · ✓ 4 healthy", worst state first, wrapped at width. An empty tally
// says "no evidence" in the unknown tone; it is never zero healthy. Repeated states add up.
func SummaryLine(counts []EvidenceCount, width int, t theme.Theme) string {
	var items []string

	for _, s := range severity() {
		var n uint64

		for _, c := range counts {
			if c.State == s || (s == EvidenceUnknown && !known(c.State)) {
				n += c.N
			}
		}

		if n == 0 {
			continue
		}

		tone, glyph := s.mark()
		items = append(items, text.ToneStyle(tone, text.Normal, evidenceTheme(s, t)).Render(t.Glyph(glyph)+" "+count(n)+" "+s.Label()))
	}

	if len(items) == 0 {
		return text.ToneStyle(text.Neutral, text.Normal, evidenceTheme(EvidenceUnknown, t)).Render(t.Glyph(theme.GlyphUnknown) + " no evidence")
	}

	return joinParts(items, width, t)
}

func known(e Evidence) bool {
	for _, k := range Evidences() {
		if e == k {
			return true
		}
	}

	return false
}

func evidenceTheme(e Evidence, t theme.Theme) theme.Theme {
	if tone, _ := e.mark(); tone == text.Neutral {
		return statusTheme(t, t.Unknown, t.DimUnknown)
	}

	return t
}

// count is exact below one thousand and compact above, through the shared quantity scaling.
func count(n uint64) string {
	if n < 1_000 {
		return strconv.FormatUint(n, 10)
	}

	return text.Quantity{Value: float64(n), Scale: text.DecimalScale, Precision: 1}.Label()
}

// EvidenceFooter says where a view's data came from and how complete it is.
type EvidenceFooter struct {
	Source, Scope string
	State         Evidence
	// ObservedAt and Now are explicit, as in Freshness: a zero value is unavailable and a future observation is
	// labeled as such.
	ObservedAt, Now time.Time
	Omission        *text.Omission // nil means nothing was omitted; a zero Omission reads "more hidden"
}

// View is "source · scope · state age · omission", wrapped at width. A missing source or scope reads missing, and
// the state keeps its word, so aged or partial evidence stays visible.
func (f EvidenceFooter) View(width int, t theme.Theme) string {
	parts := []string{
		t.Subtle().Render("source: ") + footerValue(f.Source, t),
		t.Subtle().Render("scope: ") + footerValue(f.Scope, t),
		f.State.View(text.Normal, t) + " " + t.Subtle().Render(f.age()),
	}

	if f.Omission != nil {
		if omitted := f.Omission.View(t); ansi.StringWidth(omitted) > 0 {
			parts = append(parts, omitted)
		}
	}

	return joinParts(parts, width, t)
}

func footerValue(value string, t theme.Theme) string {
	state := text.ValuePresent
	if strings.TrimSpace(value) == "" {
		state = text.ValueMissing
	}

	return state.View(value, text.Normal, t)
}

func (f EvidenceFooter) age() string {
	switch {
	case f.ObservedAt.IsZero():
		return "age unavailable"
	case f.Now.IsZero():
		return "clock unavailable"
	case f.Now.Before(f.ObservedAt):
		return "in " + text.Age(f.ObservedAt.Sub(f.Now)) + " (future)"
	default:
		return text.Age(f.Now.Sub(f.ObservedAt)) + " ago"
	}
}

// joinParts fills lines with parts separated by " · ", starting a new line before a part that would overflow.
func joinParts(parts []string, width int, t theme.Theme) string {
	sep := t.Subtle().Render(" · ")
	sepWidth := ansi.StringWidth(sep)
	lines, line, used := []string{}, "", 0

	for _, part := range parts {
		// A part wider than the allocation wraps at the cell boundary, so no source or scope text is lost or cut.
		if width > 0 && ansi.StringWidth(part) > width {
			wrapped := strings.Split(ansi.Hardwrap(part, width, true), "\n")

			if line != "" {
				lines = append(lines, line)
			}

			lines = append(lines, wrapped[:len(wrapped)-1]...)
			part = wrapped[len(wrapped)-1]
			line, used = "", 0
		}

		w := ansi.StringWidth(part)

		switch {
		case line == "":
			line, used = part, w
		case width > 0 && used+sepWidth+w > width:
			lines, line, used = append(lines, line), part, w
		default:
			line, used = line+sep+part, used+sepWidth+w
		}
	}

	rows := strings.Split(layout.Column(append(lines, line)...), "\n")
	for i, row := range rows {
		rows[i] = strings.TrimRight(row, " ")
	}

	return strings.Join(rows, "\n")
}
