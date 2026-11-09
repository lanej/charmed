package text

// Presentation controls an atom's semantic adornment, never its underlying data.
type Presentation string

const (
	DefaultPresentation Presentation = ""
	GlyphAndLabel       Presentation = "both"
	GlyphOnly           Presentation = "glyph"
	LabelOnly           Presentation = "label"
)

// Format escapes caller labels and keeps absent parts from adding stray spaces.
func (p Presentation) Format(glyph, label string) string {
	glyph, label = Inline(glyph), Inline(label)

	switch p {
	case GlyphOnly:
		return glyph
	case LabelOnly:
		return label
	case DefaultPresentation, GlyphAndLabel:
	default:
	}

	if glyph == "" {
		return label
	}

	if label == "" {
		return glyph
	}

	return glyph + " " + label
}

func presentation(options []Presentation) Presentation {
	if len(options) > 0 {
		return options[0]
	}

	return DefaultPresentation
}
