package gitui

import "github.com/rivo/uniseg"

type editSpan struct{ start, end int }

type diffGrapheme struct {
	text       string
	start, end int
}

func diffGraphemes(source string) []diffGrapheme {
	var out []diffGrapheme

	iterator := uniseg.NewGraphemes(source)
	for iterator.Next() {
		start, end := iterator.Positions()
		out = append(out, diffGrapheme{text: iterator.Str(), start: start, end: end})
	}

	return out
}

// Large replacements use a single changed middle rather than quadratic work.
func inlineEdits(before, after string) ([]editSpan, []editSpan) {
	a, b := diffGraphemes(before), diffGraphemes(after)
	for len(a) > 0 && len(b) > 0 && a[0].text == b[0].text {
		a, b = a[1:], b[1:]
	}

	for len(a) > 0 && len(b) > 0 && a[len(a)-1].text == b[len(b)-1].text {
		a, b = a[:len(a)-1], b[:len(b)-1]
	}

	if len(a) > 256 || len(b) > 256 {
		return changedMiddle(a), changedMiddle(b)
	}

	columns := len(b) + 1

	lcs := make([]int, (len(a)+1)*columns)
	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if a[i].text == b[j].text {
				lcs[i*columns+j] = 1 + lcs[(i+1)*columns+j+1]
			} else {
				lcs[i*columns+j] = max(lcs[(i+1)*columns+j], lcs[i*columns+j+1])
			}
		}
	}

	var removed, added []editSpan

	for i, j := 0, 0; i < len(a) || j < len(b); {
		switch {
		case i < len(a) && j < len(b) && a[i].text == b[j].text:
			i++
			j++
		case i < len(a) && (j == len(b) || lcs[(i+1)*columns+j] >= lcs[i*columns+j+1]):
			removed = appendEdit(removed, a[i])
			i++
		default:
			added = appendEdit(added, b[j])
			j++
		}
	}

	return removed, added
}

func changedMiddle(graphemes []diffGrapheme) []editSpan {
	if len(graphemes) == 0 {
		return nil
	}

	return []editSpan{{start: graphemes[0].start, end: graphemes[len(graphemes)-1].end}}
}

func appendEdit(spans []editSpan, grapheme diffGrapheme) []editSpan {
	if len(spans) > 0 && spans[len(spans)-1].end == grapheme.start {
		spans[len(spans)-1].end = grapheme.end
		return spans
	}

	return append(spans, editSpan{start: grapheme.start, end: grapheme.end})
}

// Pair replacement lines by position within each contiguous change block.
func hunkPairs(hunk Hunk) [][2]int {
	var pairs [][2]int

	var removed, added []int

	flush := func() {
		for i := range min(len(removed), len(added)) {
			pairs = append(pairs, [2]int{removed[i], added[i]})
		}

		removed, added = nil, nil
	}

	for i, line := range hunk.Lines {
		switch line.Kind {
		case DeletedLine:
			removed = append(removed, i)
		case AddedLine:
			added = append(added, i)
		case ContextLine:
			flush()
		default:
			flush()
		}
	}

	flush()

	return pairs
}

func hunkEdits(hunk Hunk, lines []syntaxLine) [][]editSpan {
	edits := make([][]editSpan, len(hunk.Lines))
	for _, pair := range hunkPairs(hunk) {
		oldIndex, newIndex := pair[0], pair[1]
		edits[oldIndex], edits[newIndex] = inlineEdits(lines[oldIndex].text, lines[newIndex].text)
	}

	return edits
}

func shiftedLineNumbers(hunk Hunk) bool {
	for _, line := range hunk.Lines {
		if line.OldLine > 0 && line.NewLine > 0 && line.OldLine != line.NewLine {
			return true
		}
	}

	for _, pair := range hunkPairs(hunk) {
		oldNumber, newNumber := hunk.Lines[pair[0]].OldLine, hunk.Lines[pair[1]].NewLine
		if oldNumber > 0 && newNumber > 0 && oldNumber != newNumber {
			return true
		}
	}

	return false
}
