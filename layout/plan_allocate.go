package layout

import (
	"fmt"
	"image"
	"math/bits"
	"sort"
)

const maximumCells = int(^uint(0) >> 1)

func plus(a, b int) int {
	if a > maximumCells-b {
		return maximumCells
	}

	return a + b
}

func times(a, b int) int {
	if a != 0 && b > maximumCells/a {
		return maximumCells
	}

	return a * b
}

func dimension(size Size, axis Axis) int {
	if axis == Horizontal {
		return size.Width
	}

	return size.Height
}
func rectangleSize(rect image.Rectangle) Size { return Size{Width: rect.Dx(), Height: rect.Dy()} }
func offset(extra int, align Alignment) int {
	switch align {
	case Start:
		return 0
	case Center:
		return extra / 2
	case End:
		return extra
	default:
		return 0
	}
}

func allocate(tracks []Track, available int) ([]int, error) {
	if available < 0 {
		return nil, fmt.Errorf("gaps exceed allocation")
	}

	out := make([]int, len(tracks))

	remaining := available
	for i, track := range tracks {
		if track.Min > remaining {
			return nil, fmt.Errorf("track minima exceed allocation %d", available)
		}

		out[i] = track.Min
		remaining -= track.Min
	}

	for remaining > 0 {
		weight := 0

		for i, track := range tracks {
			if track.Weight > 0 && (track.Max == 0 || out[i] < track.Max) {
				if track.Weight > maximumCells-weight {
					return nil, fmt.Errorf("total track weight overflows")
				}

				weight += track.Weight
			}
		}

		if weight == 0 {
			break
		}

		budget := remaining
		capped := false

		type remainder struct {
			index    int
			fraction uint64
		}

		var remainders []remainder

		for i, track := range tracks {
			if track.Weight == 0 || (track.Max > 0 && out[i] >= track.Max) {
				continue
			}

			if budget < 0 || track.Weight < 0 || weight <= 0 {
				return nil, fmt.Errorf("invalid weighted allocation")
			}

			hi, lo := bits.Mul64(uint64(budget), uint64(track.Weight))

			quotient, fraction := bits.Div64(hi, lo, uint64(weight))
			if quotient > uint64(maximumCells) {
				return nil, fmt.Errorf("weighted allocation overflows")
			}

			share := min(remaining, int(quotient))
			if track.Max > 0 && share > track.Max-out[i] {
				share = track.Max - out[i]
				capped = true
			}

			out[i] += share
			remaining -= share

			remainders = append(remainders, remainder{index: i, fraction: fraction})
		}

		if capped {
			continue
		}

		sort.SliceStable(remainders, func(i, j int) bool { return remainders[i].fraction > remainders[j].fraction })

		for _, remainder := range remainders {
			if remaining == 0 {
				break
			}

			track := tracks[remainder.index]
			if track.Max == 0 || out[remainder.index] < track.Max {
				out[remainder.index]++
				remaining--
			}
		}
	}

	return out, nil
}
