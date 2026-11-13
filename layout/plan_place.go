package layout

import (
	"fmt"
	"image"
)

func dockTree(n *Dock) Node {
	var horizontal []Slot
	if n.Left != nil {
		horizontal = append(horizontal, Slot{Node: n.Left, Track: Fixed(n.LeftWidth)})
	}

	horizontal = append(horizontal, Slot{Node: n.Body, Track: Grow(0)})
	if n.Right != nil {
		horizontal = append(horizontal, Slot{Node: n.Right, Track: Fixed(n.RightWidth)})
	}

	var vertical []Slot
	if n.Top != nil {
		vertical = append(vertical, Slot{Node: n.Top, Track: Fixed(n.TopHeight)})
	}

	vertical = append(vertical, Slot{Node: &Flex{Axis: Horizontal, Gap: n.Gap, Children: horizontal}, Track: Grow(0)})
	if n.Bottom != nil {
		vertical = append(vertical, Slot{Node: n.Bottom, Track: Fixed(n.BottomHeight)})
	}

	return &Flex{Axis: Vertical, Gap: n.Gap, Children: vertical}
}

func (p *planner) place(node Node, rect image.Rectangle) error {
	_, err := p.placeNode(node, rect)
	return err
}

// Each placed child returns its selected subtree's minimum. That keeps design
// validation accurate through responsive wrappers without retaining mutable nodes.
func (p *planner) placeNode(node Node, rect image.Rectangle) (Size, error) {
	need := minimumAt(node, rect.Dx())
	if rect.Dx() < need.Width || rect.Dy() < need.Height {
		return Size{}, fmt.Errorf("%T needs at least %dx%d cells; allocated %dx%d", node, need.Width, need.Height, rect.Dx(), rect.Dy())
	}

	switch n := node.(type) {
	case *Region:
		if rect.Empty() {
			return Size{}, fmt.Errorf("region %q has no cells; explicitly hide it instead", n.ID)
		}

		p.placements = append(p.placements, Placement{ID: n.ID, Bounds: rect, ClippingException: exceptionReason(n.Exceptions, ContentClipped), singleException: exceptionReason(n.Exceptions, SingleRegion), role: n.Role})
		if rect.Dx() < n.Preferred.Width || rect.Dy() < n.Preferred.Height {
			p.diagnostics = append(p.diagnostics, Diagnostic{Rule: BelowPreferred, Scope: n.ID, Exception: exceptionReason(n.Exceptions, BelowPreferred), Regions: []string{n.ID}, Message: fmt.Sprintf("%s prefers %dx%d; allocated %dx%d", n.ID, n.Preferred.Width, n.Preferred.Height, rect.Dx(), rect.Dy())})
		}

		return need, nil
	case *Flex:
		return p.placeFlex(n, rect)
	case *Grid:
		return p.placeGrid(n, rect)
	case *Flow:
		return p.placeFlow(n, rect)
	case *Dock:
		return p.placeNode(dockTree(n), rect)
	case *SplitPane:
		return p.placeSplit(n, rect)
	case *Responsive:
		for _, variant := range n.Variants {
			if rect.Dx() >= variant.Min.Width && rect.Dy() >= variant.Min.Height {
				if p.variant == "" {
					p.variant = variant.Name
				} else {
					p.variant += "/" + variant.Name
				}

				p.hidden = append(p.hidden, variant.Hidden...)
				p.selections = append(p.selections, Selection{Layout: n.Name, Variant: variant.Name})

				return p.placeNode(variant.Root, rect)
			}
		}

		return Size{}, fmt.Errorf("no responsive variant fits %dx%d", rect.Dx(), rect.Dy())
	default:
		return Size{}, fmt.Errorf("unsupported layout node %T", node)
	}
}

func (p *planner) placeFlex(n *Flex, rect image.Rectangle) (Size, error) {
	tracks := make([]Track, len(n.Children))
	for i, child := range n.Children {
		tracks[i] = child.Track

		childWidth := 0
		if n.Axis == Vertical {
			childWidth = rect.Dx()
			if child.Cross > 0 {
				childWidth = child.Cross
			}
		}

		need := dimension(minimumAt(child.Node, childWidth), n.Axis)
		if tracks[i].Weight == 0 && tracks[i].Min < need {
			return Size{}, fmt.Errorf("fixed track cannot meet child minimum %d", need)
		}

		tracks[i].Min = max(tracks[i].Min, need)
		if tracks[i].Max > 0 && tracks[i].Max < tracks[i].Min {
			return Size{}, fmt.Errorf("child minimum exceeds track maximum")
		}
	}

	sizes, err := allocate(tracks, dimension(rectangleSize(rect), n.Axis)-times(n.Gap, len(tracks)-1))
	if err != nil {
		return Size{}, err
	}

	used := times(n.Gap, len(sizes)-1)
	for _, size := range sizes {
		used = plus(used, size)
	}

	cursor := offset(dimension(rectangleSize(rect), n.Axis)-used, n.Justify)

	var (
		groups      [][]string
		groupWidths []int
	)

	main, cross := times(n.Gap, len(sizes)-1), 0
	for i, child := range n.Children {
		bounds := rect
		if n.Axis == Horizontal {
			bounds.Min.X += cursor
			bounds.Max.X = bounds.Min.X + sizes[i]
		} else {
			bounds.Min.Y += cursor
			bounds.Max.Y = bounds.Min.Y + sizes[i]
		}

		if child.Cross > 0 {
			available := rect.Dy()
			if n.Axis == Vertical {
				available = rect.Dx()
			}

			if child.Cross > available {
				return Size{}, fmt.Errorf("cross size exceeds allocation")
			}

			start := offset(available-child.Cross, n.Align)
			if n.Axis == Horizontal {
				bounds.Min.Y += start
				bounds.Max.Y = bounds.Min.Y + child.Cross
			} else {
				bounds.Min.X += start
				bounds.Max.X = bounds.Min.X + child.Cross
			}
		}

		start := len(p.placements)

		need, placeErr := p.placeNode(child.Node, bounds)
		if placeErr != nil {
			return Size{}, placeErr
		}

		var ids []string

		for _, placed := range p.placements[start:] {
			if placed.role == ContentRegion {
				ids = append(ids, placed.ID)
			}
		}

		if len(ids) > 0 {
			groups = append(groups, ids)
			groupWidths = append(groupWidths, need.Width)
		}

		main = plus(main, max(child.Track.Min, dimension(need, n.Axis)))

		other := need.Height
		if n.Axis == Vertical {
			other = need.Width
		}

		cross = max(cross, other, child.Cross)
		cursor = plus(cursor, plus(sizes[i], n.Gap))
	}

	p.warnStack(n, rect, groups, groupWidths)

	if n.Axis == Horizontal {
		return Size{Width: main, Height: cross}, nil
	}

	return Size{Width: cross, Height: main}, nil
}

func (p *planner) placeSplit(n *SplitPane, rect image.Rectangle) (Size, error) {
	space := dimension(rectangleSize(rect), n.Axis) - n.Gap

	ratio := n.Ratio
	if ratio == 0 {
		ratio = .5
	}

	a, b := minimumAt(n.First, rect.Dx()), minimumAt(n.Second, rect.Dx())
	first := max(dimension(a, n.Axis), min(int(float64(space)*ratio), space-dimension(b, n.Axis)))

	left, right := rect, rect
	if n.Axis == Horizontal {
		left.Max.X = left.Min.X + first
		right.Min.X = left.Max.X + n.Gap
	} else {
		left.Max.Y = left.Min.Y + first
		right.Min.Y = left.Max.Y + n.Gap
	}

	start := len(p.placements)

	a, err := p.placeNode(n.First, left)
	if err != nil {
		return Size{}, err
	}

	var idsA []string

	for _, placed := range p.placements[start:] {
		if placed.role == ContentRegion {
			idsA = append(idsA, placed.ID)
		}
	}

	start = len(p.placements)

	b, err = p.placeNode(n.Second, right)
	if err != nil {
		return Size{}, err
	}

	var idsB []string

	for _, placed := range p.placements[start:] {
		if placed.role == ContentRegion {
			idsB = append(idsB, placed.ID)
		}
	}

	if len(idsA) > 0 && len(idsB) > 0 {
		p.warnStack(&Flex{Name: n.Name, Axis: n.Axis, Gap: n.Gap, Exceptions: n.Exceptions}, rect, [][]string{idsA, idsB}, []int{a.Width, b.Width})
	}

	if n.Axis == Horizontal {
		return Size{Width: plus(plus(a.Width, b.Width), n.Gap), Height: max(a.Height, b.Height)}, nil
	}

	return Size{Width: max(a.Width, b.Width), Height: plus(plus(a.Height, b.Height), n.Gap)}, nil
}
