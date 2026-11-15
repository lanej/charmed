package layout

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
)

// TextRowItem contains trusted, optionally ANSI-styled display text. Escape data
// with text.Inline before styling it. Flex inserts an expanding spacer after
// Text; an empty Text with Flex is a standalone spacer. Empty items are ignored.
type TextRowItem struct {
	Text string
	Flex bool
}

// TextRow arranges text in terminal cells with a minimum Gap between nonempty
// items. Flexible spacers share remaining cells equally on each rendered line.
// When an item will not fit, it moves to the next line with any preceding spacer;
// long items wrap internally. No item is silently omitted or abbreviated.
// Stack puts every text item on its own wrapped lines and ignores spacers.
//
// View is pure and returns padded, width-bounded lines with content-driven height.
// Callers allocate that height or provide scrolling; Fit would clip overflow.
// A grapheme wider than the allocation is represented by !, rather than an
// overflowing cell. Styling, semantic meaning and input handling belong to callers.
type TextRow struct {
	Items []TextRowItem
	Gap   int
	Stack bool
}

func (r TextRow) View(width int) string {
	if width <= 0 {
		return ""
	}

	gap := min(width, max(0, r.Gap))
	lines := []string{}
	line := []TextRowItem{}
	used, texts, pending := 0, 0, 0
	flush := func() {
		if len(line) == 0 {
			return
		}

		lines = append(lines, renderTextRow(line, width, gap))
		line, used, texts = nil, 0, 0
	}

	for _, item := range r.Items {
		if item.Text != "" {
			if r.Stack {
				flush()

				pending = 0
			}

			content := strings.ReplaceAll(item.Text, "\r\n", "\n")
			if width == 1 {
				content = textRowSingleCell(content)
			}

			fragments := strings.Split(ansi.Hardwrap(ansi.Wrap(content, width, ""), width, true), "\n")

			for i, fragment := range fragments {
				if ansi.StringWidth(fragment) > width {
					fragment = ansi.Truncate(fragment, width, "!")
				}

				cells := ansi.StringWidth(fragment)
				if texts > 0 && used+gap+cells > width {
					flush()
				}

				for range pending {
					line = append(line, TextRowItem{Flex: true})
				}

				pending = 0

				if texts > 0 {
					used += gap
				}

				line = append(line, TextRowItem{Text: fragment})
				used += cells
				texts++

				if i < len(fragments)-1 {
					flush()
				}
			}
		}

		if item.Flex && !r.Stack {
			pending++
		}
	}

	if len(line) > 0 {
		for range pending {
			line = append(line, TextRowItem{Flex: true})
		}
	}

	flush()

	return strings.Join(lines, "\n")
}

func renderTextRow(items []TextRowItem, width, gap int) string {
	used, texts, spacers := 0, 0, 0

	for _, item := range items {
		if item.Flex {
			spacers++
		} else {
			used += ansi.StringWidth(item.Text)
			texts++
		}
	}

	remaining := max(0, width-used-max(0, texts-1)*gap)
	parts := []string{}

	seen, expanded := 0, 0

	for _, item := range items {
		if item.Flex {
			cells := remaining / spacers
			if expanded < remaining%spacers {
				cells++
			}

			parts = append(parts, strings.Repeat(" ", cells))

			expanded++
		} else {
			if seen > 0 {
				parts = append(parts, strings.Repeat(" ", gap))
			}

			parts = append(parts, item.Text)

			seen++
		}
	}

	return Fit(strings.Join(parts, ""), Size{Width: width, Height: 1})
}

// Replace impossible graphemes before wrapping so the upstream wrapper does not
// insert empty rows around them. Native ANSI sequences remain unchanged.
func textRowSingleCell(content string) string {
	parts := []string{}

	var state byte

	for len(content) > 0 {
		sequence, cells, read, next := ansi.DecodeSequence(content, state, nil)
		if cells > 1 {
			sequence = "!"
		}

		parts = append(parts, sequence)
		content, state = content[read:], next
	}

	return strings.Join(parts, "")
}
