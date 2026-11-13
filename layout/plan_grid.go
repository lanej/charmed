package layout

import (
	"fmt"
	"slices"
	"sort"
)

type spanMinimum struct{ start, end, need int }

// Interval minima are satisfied in end order by filling the rightmost available
// tracks first. Shared spans reuse that allocation, avoiding unnecessary growth.
func gridRequirements(n *Grid, axis Axis, widths []int) ([]Track, error) {
	return gridMeasuredRequirements(n, axis, widths, nil)
}

func gridMeasuredRequirements(n *Grid, axis Axis, widths []int, measured []Size) ([]Track, error) {
	tracks := slices.Clone(n.Columns)

	gap := n.Gap.Width
	if axis == Vertical {
		tracks = slices.Clone(n.Rows)
		gap = n.Gap.Height
	}

	var spans []spanMinimum

	for index, cell := range n.Cells {
		start, span, need := cell.Column, max(1, cell.ColumnSpan), minimum(cell.Node).Width
		if axis == Vertical {
			start, span = cell.Row, max(1, cell.RowSpan)

			width := 0
			if len(widths) > 0 {
				width = times(n.Gap.Width, max(1, cell.ColumnSpan)-1)
				for i := cell.Column; i < cell.Column+max(1, cell.ColumnSpan); i++ {
					width = plus(width, widths[i])
				}
			}

			need = minimumAt(cell.Node, width).Height
		}

		if len(measured) > index {
			need = dimension(measured[index], axis)
		}

		spans = append(spans, spanMinimum{start: start, end: start + span, need: max(0, need-times(gap, span-1))})
	}

	sort.SliceStable(spans, func(i, j int) bool {
		if spans[i].end == spans[j].end {
			return spans[i].start > spans[j].start
		}

		return spans[i].end < spans[j].end
	})

	for _, span := range spans {
		present := 0
		for i := span.start; i < span.end; i++ {
			present = plus(present, tracks[i].Min)
		}

		deficit := max(0, span.need-present)
		for i := span.end - 1; i >= span.start && deficit > 0; i-- {
			if tracks[i].Weight == 0 {
				continue
			}

			growth := deficit
			if tracks[i].Max > 0 {
				growth = min(growth, tracks[i].Max-tracks[i].Min)
			}

			tracks[i].Min = plus(tracks[i].Min, growth)
			deficit -= growth
		}

		if deficit > 0 {
			return nil, fmt.Errorf("grid tracks cannot meet child minimum %d", span.need)
		}
	}

	return tracks, nil
}
