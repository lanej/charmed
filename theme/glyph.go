package theme

// Glyph names meaning independently of its Unicode or ASCII representation.
type Glyph string

const (
	GlyphOK          Glyph = "ok"
	GlyphWarning     Glyph = "warning"
	GlyphError       Glyph = "error"
	GlyphUnknown     Glyph = "unknown"
	GlyphOff         Glyph = "off"
	GlyphInfo        Glyph = "info"
	GlyphSelected    Glyph = "selected"
	GlyphUnselected  Glyph = "unselected"
	GlyphExpanded    Glyph = "expanded"
	GlyphCollapsed   Glyph = "collapsed"
	GlyphUp          Glyph = "up"
	GlyphDown        Glyph = "down"
	GlyphLeft        Glyph = "left"
	GlyphRight       Glyph = "right"
	GlyphLoading     Glyph = "loading"
	GlyphUnavailable Glyph = "unavailable"
)

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
		unicode, ascii = "○", "-"
	case GlyphInfo:
		unicode, ascii = "i", "i"
	case GlyphSelected:
		unicode, ascii = "●", "*"
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
	case GlyphUnavailable:
		unicode, ascii = "—", "-"
	default:
		unicode, ascii = "?", "?"
	}

	if t.ASCII {
		return ascii
	}

	return unicode
}
