package theme

import "image/color"

// DiffColors distinguish changed lines and their inline edits from lexical colors.
type DiffColors struct{ Added, Deleted, AddedEmphasis, DeletedEmphasis color.Color }

func diffColors(mode Mode) DiffColors {
	switch mode {
	case Dark:
		return DiffColors{hex("#142b24"), hex("#302027"), hex("#1d3d30"), hex("#492630")}
	case Light:
		return DiffColors{hex("#e8f3eb"), hex("#faeaea"), hex("#d5e9da"), hex("#f9dede")}
	case Contrast:
		return DiffColors{hex("#0e2118"), hex("#271519"), hex("#214331"), hex("#4d2630")}
	case Mono:
		return DiffColors{}
	default:
		return DiffColors{}
	}
}
