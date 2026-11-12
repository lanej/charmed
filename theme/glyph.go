package theme

// Glyph names meaning independently of its Unicode or ASCII representation.
type Glyph string

const (
	GlyphOK           Glyph = "ok"
	GlyphWarning      Glyph = "warning"
	GlyphError        Glyph = "error"
	GlyphUnknown      Glyph = "unknown"
	GlyphOff          Glyph = "off"
	GlyphPending      Glyph = "pending"
	GlyphRunning      Glyph = "running"
	GlyphPaused       Glyph = "paused"
	GlyphDisabled     Glyph = "disabled"
	GlyphConnected    Glyph = "connected"
	GlyphDisconnected Glyph = "disconnected"
	GlyphSearch       Glyph = "search"
	GlyphFilter       Glyph = "filter"
	GlyphSort         Glyph = "sort"
	GlyphRefresh      Glyph = "refresh"
	GlyphLink         Glyph = "link"
	GlyphExpand       Glyph = "expand"
	GlyphInfo         Glyph = "info"
	GlyphSelected     Glyph = "selected"
	GlyphUnselected   Glyph = "unselected"
	GlyphExpanded     Glyph = "expanded"
	GlyphCollapsed    Glyph = "collapsed"
	GlyphUp           Glyph = "up"
	GlyphDown         Glyph = "down"
	GlyphLeft         Glyph = "left"
	GlyphRight        Glyph = "right"
	GlyphLoading      Glyph = "loading"
	GlyphUnavailable  Glyph = "unavailable"
	GlyphLog          Glyph = "log"
	GlyphTrace        Glyph = "trace"
	GlyphSpan         Glyph = "span"
	GlyphUpstream     Glyph = "upstream"
	GlyphDownstream   Glyph = "downstream"
	GlyphScrollTrack  Glyph = "scroll-track"
	GlyphScrollThumb  Glyph = "scroll-thumb"
	GlyphChecked      Glyph = "checked"
	GlyphUnchecked    Glyph = "unchecked"
	GlyphMixed        Glyph = "mixed"
	GlyphAction       Glyph = "action"
	GlyphBusy         Glyph = "busy"
	GlyphDelete       Glyph = "delete"
	GlyphInactive     Glyph = "inactive"
	GlyphAdded        Glyph = "added"
	GlyphModified     Glyph = "modified"
	GlyphRemoved      Glyph = "removed"
	GlyphUnchanged    Glyph = "unchanged"
	GlyphMissing      Glyph = "missing"
	GlyphRedacted     Glyph = "redacted"
	GlyphClock        Glyph = "clock"
	GlyphTruncated    Glyph = "truncated"
	GlyphAttribution  Glyph = "attribution-gap"
)

// Glyphs returns the semantic catalog in display order.
func Glyphs() []Glyph {
	return []Glyph{GlyphOK, GlyphWarning, GlyphError, GlyphUnknown, GlyphOff, GlyphPending, GlyphRunning, GlyphPaused, GlyphDisabled, GlyphConnected, GlyphDisconnected, GlyphInfo, GlyphSearch, GlyphFilter, GlyphSort, GlyphRefresh, GlyphLink, GlyphExpand, GlyphSelected, GlyphUnselected, GlyphExpanded, GlyphCollapsed, GlyphUp, GlyphDown, GlyphLeft, GlyphRight, GlyphLoading, GlyphUnavailable, GlyphLog, GlyphTrace, GlyphSpan, GlyphUpstream, GlyphDownstream, GlyphScrollTrack, GlyphScrollThumb, GlyphChecked, GlyphUnchecked, GlyphMixed, GlyphAction, GlyphBusy, GlyphDelete, GlyphInactive, GlyphAdded, GlyphModified, GlyphRemoved, GlyphUnchanged, GlyphMissing, GlyphRedacted, GlyphClock, GlyphTruncated, GlyphAttribution}
}

// Glyph returns a compact mark. Callers keep a text label for semantic states.
func (t Theme) Glyph(glyph Glyph) string {
	var unicode, ascii string

	switch glyph {
	case GlyphOK:
		unicode, ascii = "✓", "+"
	case GlyphWarning:
		unicode, ascii = "!", "!"
	case GlyphError:
		unicode, ascii = "×", "x"
	case GlyphUnknown:
		unicode, ascii = "?", "?"
	case GlyphOff:
		unicode, ascii = "⏻", "-"
	case GlyphPending:
		unicode, ascii = "◷", "p"
	case GlyphRunning:
		unicode, ascii = "▶", ">"
	case GlyphPaused:
		unicode, ascii = "‖", "="
	case GlyphDisabled:
		unicode, ascii = "⊘", "/"
	case GlyphConnected:
		unicode, ascii = "●", "+"
	case GlyphDisconnected:
		unicode, ascii = "○", "-"
	case GlyphSearch:
		unicode, ascii = "⌕", "/"
	case GlyphFilter:
		unicode, ascii = "▽", "f"
	case GlyphSort:
		unicode, ascii = "↕", "s"
	case GlyphRefresh:
		unicode, ascii = "↻", "r"
	case GlyphLink:
		unicode, ascii = "↗", "@"
	case GlyphExpand:
		unicode, ascii = "⤢", "+"
	case GlyphInfo:
		unicode, ascii = "i", "i"
	case GlyphSelected:
		unicode, ascii = "◉", "*"
	case GlyphUnselected:
		unicode, ascii = "○", "o"
	case GlyphExpanded:
		unicode, ascii = "▾", "v"
	case GlyphCollapsed:
		unicode, ascii = "▸", ">"
	case GlyphUp:
		unicode, ascii = "↑", "^"
	case GlyphDown:
		unicode, ascii = "↓", "v"
	case GlyphLeft:
		unicode, ascii = "←", "<"
	case GlyphRight:
		unicode, ascii = "→", ">"
	case GlyphLoading:
		unicode, ascii = "…", "."
	case GlyphScrollTrack:
		unicode, ascii = "│", "|"
	case GlyphScrollThumb:
		unicode, ascii = "┃", "#"
	case GlyphLog:
		unicode, ascii = "≡", "L"
	case GlyphTrace:
		unicode, ascii = "⋈", "T"
	case GlyphSpan:
		unicode, ascii = "─", "S"
	case GlyphUpstream:
		unicode, ascii = "↑", "^"
	case GlyphDownstream:
		unicode, ascii = "↓", "v"
	case GlyphUnavailable:
		unicode, ascii = "—", "-"
	case GlyphChecked:
		unicode, ascii = "☑", "x"
	case GlyphUnchecked:
		unicode, ascii = "☐", " "
	case GlyphMixed:
		unicode, ascii = "▣", "-"
	case GlyphAction:
		unicode, ascii = "↵", ">"
	case GlyphBusy:
		unicode, ascii = "⟳", "~"
	case GlyphDelete:
		unicode, ascii = "⌫", "x"
	case GlyphInactive:
		unicode, ascii = "◌", "-"
	case GlyphAdded:
		unicode, ascii = "⊕", "+"
	case GlyphModified:
		unicode, ascii = "↺", "~"
	case GlyphRemoved:
		unicode, ascii = "⊖", "-"
	case GlyphUnchanged:
		unicode, ascii = "≡", "="
	case GlyphMissing:
		unicode, ascii = "∅", "-"
	case GlyphRedacted:
		unicode, ascii = "▧", "#"
	case GlyphClock:
		unicode, ascii = "◷", "t"
	case GlyphTruncated:
		unicode, ascii = "…", "."
	case GlyphAttribution:
		unicode, ascii = "≠", "%"
	default:
		unicode, ascii = "?", "?"
	}

	if t.ASCII {
		return ascii
	}

	return unicode
}
