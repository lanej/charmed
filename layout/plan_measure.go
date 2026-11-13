package layout

func minimum(node Node) Size { return minimumAt(node, 0) }

// Width-sensitive measurement accounts for wrapping before height is shared.
func minimumAt(node Node, width int) Size {
	switch n := node.(type) {
	case *Region:
		return Size{Width: max(1, n.Min.Width), Height: max(1, n.Min.Height)}
	case *Flex:
		tracks := make([]Track, len(n.Children))
		for i, child := range n.Children {
			tracks[i] = child.Track

			childWidth := 0
			if n.Axis == Vertical {
				childWidth = width
				if child.Cross > 0 {
					childWidth = child.Cross
				}
			}

			tracks[i].Min = max(tracks[i].Min, dimension(minimumAt(child.Node, childWidth), n.Axis))
		}

		var widths []int

		if n.Axis == Horizontal && width > 0 {
			allocated, allocationErr := allocate(tracks, width-times(n.Gap, len(tracks)-1))
			if allocationErr == nil {
				widths = allocated
			}
		}

		main, cross := times(n.Gap, len(n.Children)-1), 0
		for i, child := range n.Children {
			childWidth := width
			if n.Axis == Vertical && child.Cross > 0 {
				childWidth = child.Cross
			}

			if n.Axis == Horizontal {
				childWidth = 0
				if len(widths) > 0 {
					childWidth = widths[i]
				}
			}

			need := minimumAt(child.Node, childWidth)
			main = plus(main, max(child.Track.Min, dimension(need, n.Axis)))

			other := need.Height
			if n.Axis == Vertical {
				other = need.Width
			}

			cross = max(cross, other, child.Cross)
		}

		if n.Axis == Horizontal {
			return Size{Width: main, Height: cross}
		}

		return Size{Width: cross, Height: main}
	case *Grid:
		columns, err := gridRequirements(n, Horizontal, nil)
		if err != nil {
			return Size{}
		}

		var widths []int

		if width > 0 {
			allocated, allocationErr := allocate(columns, width-times(n.Gap.Width, len(columns)-1))
			if allocationErr == nil {
				widths = allocated
			}
		}

		rows, err := gridRequirements(n, Vertical, widths)
		if err != nil {
			return Size{}
		}

		size := Size{Width: times(n.Gap.Width, len(columns)-1), Height: times(n.Gap.Height, len(rows)-1)}
		for _, track := range columns {
			size.Width = plus(size.Width, track.Min)
		}

		for _, track := range rows {
			size.Height = plus(size.Height, track.Min)
		}

		return size
	case *Flow:
		size := Size{}
		x, y, rowHeight := 0, 0, 0

		for _, child := range n.Children {
			size.Width = max(size.Width, child.Size.Width)

			size.Height = max(size.Height, child.Size.Height)
			if width > 0 {
				if x > 0 && child.Size.Width > width-x {
					y = plus(y, plus(rowHeight, n.Gap.Height))
					x = 0
					rowHeight = 0
				}

				x = plus(x, plus(child.Size.Width, n.Gap.Width))
				rowHeight = max(rowHeight, child.Size.Height)
			}
		}

		if width > 0 {
			size.Height = plus(y, rowHeight)
		}

		return size
	case *Dock:
		return minimumAt(dockTree(n), width)
	case *SplitPane:
		a, b := minimum(n.First), minimum(n.Second)
		if n.Axis == Horizontal {
			if width >= plus(plus(a.Width, b.Width), n.Gap) {
				space := width - n.Gap

				ratio := n.Ratio
				if ratio == 0 {
					ratio = .5
				}

				first := max(a.Width, min(int(float64(space)*ratio), space-b.Width))
				a, b = minimumAt(n.First, first), minimumAt(n.Second, space-first)
			}

			return Size{Width: plus(plus(a.Width, b.Width), n.Gap), Height: max(a.Height, b.Height)}
		}

		a, b = minimumAt(n.First, width), minimumAt(n.Second, width)

		return Size{Width: max(a.Width, b.Width), Height: plus(plus(a.Height, b.Height), n.Gap)}
	case *Responsive:
		return Size{}
	default:
		return Size{}
	}
}
