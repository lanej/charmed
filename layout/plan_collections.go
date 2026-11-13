package layout

import (
	"fmt"
	"image"
	"strings"
)

func trackStarts(sizes []int, gap int) []int {
	starts := make([]int, len(sizes)+1)
	for i, size := range sizes {
		starts[i+1] = plus(starts[i], plus(size, gap))
	}

	return starts
}

func (p *planner) placeGrid(n *Grid, rect image.Rectangle) (Size, error) {
	columnTracks, err := gridRequirements(n, Horizontal, nil)
	if err != nil {
		return Size{}, err
	}

	columns, err := allocate(columnTracks, rect.Dx()-times(n.Gap.Width, len(n.Columns)-1))
	if err != nil {
		return Size{}, err
	}

	rowTracks, err := gridRequirements(n, Vertical, columns)
	if err != nil {
		return Size{}, err
	}

	rows, err := allocate(rowTracks, rect.Dy()-times(n.Gap.Height, len(n.Rows)-1))
	if err != nil {
		return Size{}, err
	}

	xs, ys := trackStarts(columns, n.Gap.Width), trackStarts(rows, n.Gap.Height)

	var measured []Size

	for _, cell := range n.Cells {
		lastX, lastY := cell.Column+max(1, cell.ColumnSpan)-1, cell.Row+max(1, cell.RowSpan)-1

		bounds := image.Rect(rect.Min.X+xs[cell.Column], rect.Min.Y+ys[cell.Row], rect.Min.X+xs[lastX]+columns[lastX], rect.Min.Y+ys[lastY]+rows[lastY])

		need, placeErr := p.placeNode(cell.Node, bounds)
		if placeErr != nil {
			return Size{}, placeErr
		}

		measured = append(measured, need)
	}

	columnMin, err := gridMeasuredRequirements(n, Horizontal, nil, measured)
	if err != nil {
		return Size{}, err
	}

	rowMin, err := gridMeasuredRequirements(n, Vertical, nil, measured)
	if err != nil {
		return Size{}, err
	}

	need := Size{Width: times(n.Gap.Width, len(columnMin)-1), Height: times(n.Gap.Height, len(rowMin)-1)}
	for _, track := range columnMin {
		need.Width = plus(need.Width, track.Min)
	}

	for _, track := range rowMin {
		need.Height = plus(need.Height, track.Min)
	}

	return need, nil
}

func (p *planner) placeFlow(n *Flow, rect image.Rectangle) (Size, error) {
	x, y, rowHeight := 0, 0, 0

	for _, child := range n.Children {
		if child.Size.Width > rect.Dx() {
			return Size{}, fmt.Errorf("flow box width %d exceeds allocation %d", child.Size.Width, rect.Dx())
		}

		if x > 0 && child.Size.Width > rect.Dx()-x {
			y = plus(y, plus(rowHeight, n.Gap.Height))
			x = 0
			rowHeight = 0
		}

		if y > rect.Dy() || child.Size.Height > rect.Dy()-y {
			return Size{}, fmt.Errorf("wrapped flow rows exceed height %d", rect.Dy())
		}

		bounds := image.Rect(rect.Min.X+x, rect.Min.Y+y, rect.Min.X+x+child.Size.Width, rect.Min.Y+y+child.Size.Height)
		if err := p.place(child.Node, bounds); err != nil {
			return Size{}, err
		}

		x = plus(x, plus(child.Size.Width, n.Gap.Width))
		rowHeight = max(rowHeight, child.Size.Height)
	}

	return minimumAt(n, rect.Dx()), nil
}

// The stacking rule uses declared minima and content roles; it does not infer
// aesthetics from arbitrary terminal strings.
func (p *planner) warnStack(n *Flex, rect image.Rectangle, groups [][]string, widths []int) {
	if n.Axis != Vertical || len(groups) < 2 {
		return
	}

	width := times(n.Gap, len(groups)-1)

	var ids []string
	for i, group := range groups {
		ids = append(ids, group...)
		width = plus(width, widths[i])
	}

	if width <= rect.Dx() {
		scope := n.Name
		if scope == "" {
			scope = "column[" + strings.Join(ids, ",") + "]"
		}

		p.diagnostics = append(p.diagnostics, Diagnostic{Rule: StackOpportunity, Scope: scope, Exception: exceptionReason(n.Exceptions, StackOpportunity), Regions: ids, Message: fmt.Sprintf("%d content groups are stacked at %d columns although their effective minimum widths fit a row; consider columns or declare a scoped exception", len(groups), rect.Dx())})
	}
}
