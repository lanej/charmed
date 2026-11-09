package diagram

import "github.com/charmbracelet/x/ansi"

type (
	point   struct{ x, y int }
	nodeBox struct {
		state          State
		label, markers string
		x, y, w, h     int
	}
)

type (
	link struct {
		from, to int
		label    string
	}
	edgePath struct {
		points  []point
		label   string
		labelAt point
	}
	drawing struct {
		nodes         []nodeBox
		edges         []edgePath
		width, height int
	}
)

type layoutInfo struct {
	loops, remote        []int
	loopWidths           [][]int
	pairCount, pairWidth []int
}

func classify(nodes []nodeBox, edges []link) layoutInfo {
	info := layoutInfo{loops: make([]int, len(nodes)), remote: make([]int, len(nodes)), loopWidths: make([][]int, len(nodes)), pairCount: make([]int, max(0, len(nodes)-1)), pairWidth: make([]int, max(0, len(nodes)-1))}
	for _, edge := range edges {
		if edge.from == edge.to {
			info.loops[edge.from]++
			info.loopWidths[edge.from] = append(info.loopWidths[edge.from], ansi.StringWidth(edge.label)+4)
		} else if max(edge.from, edge.to)-min(edge.from, edge.to) == 1 {
			pair := min(edge.from, edge.to)
			info.pairCount[pair]++
			info.pairWidth[pair] += ansi.StringWidth(edge.label) + 4
		} else {
			info.remote[edge.from]++
			info.remote[edge.to]++
		}
	}

	return info
}

func planDiagram(input []nodeBox, edges []link, vertical bool) drawing {
	info := classify(input, edges)
	d := drawing{nodes: append([]nodeBox(nil), input...)}

	gaps := make([]int, max(0, len(input)-1))
	for i := range gaps {
		gaps[i] = 3
	}

	for _, edge := range edges {
		if !vertical && edge.from != edge.to {
			pair := min(edge.from, edge.to)
			if edge.label != "" {
				gaps[pair] = max(gaps[pair], ansi.StringWidth(edge.label)+6)
			}
		}
	}

	commonW, commonH, maxLoops, loopBand := 0, 4, 0, 0
	for _, count := range info.pairCount {
		commonH = max(commonH, count+2)
	}

	for i := range d.nodes {
		n := &d.nodes[i]

		possible := "active"
		if n.state.Initial {
			possible = "initial, active"
		}

		if n.state.Final {
			possible = "final, active"
		}

		if n.state.Initial && n.state.Final {
			possible = "initial, final, active"
		}

		n.w = max(7, ansi.StringWidth(n.label)+4, ansi.StringWidth(possible)+4)
		n.h = 4

		loopLabel, band := 0, 0
		for _, w := range info.loopWidths[i] {
			loopLabel = max(loopLabel, w-4)
			band += w
		}

		loopBand = max(loopBand, band)
		maxLoops = max(maxLoops, info.loops[i])

		space := max(2*info.remote[i]+3, 2*info.remote[i]+4*info.loops[i]+loopLabel+4)
		if vertical {
			n.h = max(n.h, 2*info.remote[i]+4*info.loops[i]+3)
		} else {
			n.w = max(n.w, space)
			n.h = commonH
		}

		commonW = max(commonW, n.w)
	}

	if vertical {
		for _, pairWidth := range info.pairWidth {
			commonW = max(commonW, pairWidth+2)
		}
	}

	gapStarts := make([]int, len(gaps))

	for i := range d.nodes {
		n := &d.nodes[i]
		if vertical {
			n.w = commonW
			n.y = d.height

			d.height += n.h
			if i < len(gaps) {
				gapStarts[i] = d.height
				d.height += gaps[i]
			}
		} else {
			n.x = d.width

			d.width += n.w
			if i < len(gaps) {
				gapStarts[i] = d.width
				d.width += gaps[i]
			}
		}
	}

	if vertical {
		d.width = commonW
	} else {
		d.height = commonH
	}

	loopIndex, portIndex, pairIndex, pairOffset := make([]int, len(input)), make([]int, len(input)), make([]int, len(gaps)), make([]int, len(gaps))
	remoteIndex, remoteX := 0, commonW+loopBand+2

	for _, edge := range edges {
		from, to := d.nodes[edge.from], d.nodes[edge.to]

		path := edgePath{label: edge.label}
		if edge.from == edge.to {
			i := loopIndex[edge.from]
			loopIndex[edge.from]++

			if vertical {
				lane := commonW + 1
				for _, w := range info.loopWidths[edge.from][i+1:] {
					lane += w
				}

				a, b := point{from.w - 1, from.y + 1 + 2*info.remote[edge.from] + 2*i}, point{from.w - 1, from.y + from.h - 2 - 2*i}
				path.points = []point{a, {lane, a.y}, {lane, b.y}, b}
				path.labelAt = point{lane + 2, a.y + 1}
				d.width = max(d.width, lane+2+ansi.StringWidth(edge.label))
			} else {
				lane := commonH + 1 + 2*(info.loops[edge.from]-1-i)
				a, b := point{from.x + 1 + 2*info.remote[edge.from] + 2*i, from.h - 1}, point{from.x + from.w - 2 - 2*i, from.h - 1}
				path.points = []point{a, {a.x, lane}, {b.x, lane}, b}
				path.labelAt = point{(a.x + b.x - ansi.StringWidth(edge.label) + 1) / 2, lane}
				d.height = max(d.height, lane+1)
			}
		} else if max(edge.from, edge.to)-min(edge.from, edge.to) == 1 {
			pair := min(edge.from, edge.to)
			if vertical {
				x := 1 + pairOffset[pair]
				pairOffset[pair] += ansi.StringWidth(edge.label) + 4

				a, b := point{x, from.y + from.h - 1}, point{x, to.y}
				if edge.from > edge.to {
					a.y = from.y
					b.y = to.y + to.h - 1
				}

				path.points = []point{a, b}
				path.labelAt = point{x + 2, gapStarts[pair] + 1}
			} else {
				y := 1 + pairIndex[pair]
				pairIndex[pair]++

				a, b := point{from.x + from.w - 1, y}, point{to.x, y}
				if edge.from > edge.to {
					a.x = from.x
					b.x = to.x + to.w - 1
				}

				path.points = []point{a, b}
				path.labelAt = point{gapStarts[pair] + (gaps[pair]-ansi.StringWidth(edge.label))/2, y}
			}
		} else {
			if vertical {
				lane := remoteX
				remoteX += ansi.StringWidth(edge.label) + 4
				a, b := point{from.w - 1, from.y + 1 + 2*portIndex[edge.from]}, point{to.w - 1, to.y + 1 + 2*portIndex[edge.to]}
				path.points = []point{a, {lane, a.y}, {lane, b.y}, b}
				path.labelAt = point{lane + 2, gapStarts[min(edge.from, edge.to)] + 1}
				d.width = max(d.width, lane+max(1, 2+ansi.StringWidth(edge.label)))
			} else {
				lane := commonH + 2*maxLoops + 1 + 2*remoteIndex
				a, b := point{from.x + 1 + 2*portIndex[edge.from], from.h - 1}, point{to.x + 1 + 2*portIndex[edge.to], to.h - 1}
				path.points = []point{a, {a.x, lane}, {b.x, lane}, b}
				pair := min(edge.from, edge.to)
				path.labelAt = point{gapStarts[pair] + (gaps[pair]-ansi.StringWidth(edge.label))/2, lane}
				d.height = max(d.height, lane+1)
			}

			portIndex[edge.from]++
			portIndex[edge.to]++
			remoteIndex++
		}

		d.edges = append(d.edges, path)
	}

	return d
}
