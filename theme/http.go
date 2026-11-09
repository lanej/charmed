package theme

import "image/color"

// HTTPColors are categorical method roles, independent of response severity.
type HTTPColors struct{ Read, Create, Change, Delete, Meta, Other color.Color }

// SyntaxColors describe lexical meaning independently of interaction and status.
type SyntaxColors struct{ Keyword, String, Number, Name, Comment color.Color }

func httpColors(mode Mode) (HTTPColors, SyntaxColors) {
	var values []color.Color

	switch mode {
	case Light:
		values = colors("#365e70", "#725a20", "#65517f", "#7b4252", "#526176", "#365b78")
	case Dark:
		values = colors("#99c8d4", "#e1cd91", "#c3b2df", "#e6b6bf", "#a1acbd", "#a9c4e5")
	case Contrast:
		values = colors("#b7dfea", "#f4dfa6", "#d9c8f0", "#f3c6d0", "#b0bdcf", "#c1d9f5")
	case Mono:
		return HTTPColors{}, SyntaxColors{}
	default:
		return HTTPColors{}, SyntaxColors{}
	}

	return HTTPColors{values[0], values[1], values[2], values[3], values[4], values[5]}, SyntaxColors{values[2], values[0], values[1], values[5], values[4]}
}
