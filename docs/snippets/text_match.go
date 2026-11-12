package snippets

import (
	"github.com/lanej/charmed/text"
	"github.com/lanej/charmed/theme"
)

func ExampleTextMatch() string {
	normal := theme.New(theme.Dark).Text()
	// Ranges are half-open rune offsets into this original source string.
	return text.HighlightMatches("Layout: timeline",
		[]text.MatchRange{{Start: 0, End: 3}, {Start: 8, End: 9}},
		normal, normal.Underline(true))
}
