package tui

import (
	"fmt"

	"github.com/lanej/charmed/layout"
)

type LayoutCase struct {
	Size layout.Size
	Plan layout.Plan
}

// CheckLayouts validates registry coverage, geometry and design rules at supplied
// terminal sizes, and requires every declared responsive variant to be exercised.
// It does not mutate this screen or execute component rendering and effects.
func (s *Screen) CheckLayouts(root layout.Node, sizes ...layout.Size) ([]LayoutCase, error) {
	if len(sizes) == 0 {
		return nil, fmt.Errorf("layout validation requires terminal sizes")
	}

	variants, err := layout.Variants(root)
	if err != nil {
		return nil, err
	}

	candidate, err := NewScreen(s.ids...)
	if err != nil {
		return nil, err
	}

	covered := map[layout.Selection]bool{}

	var cases []LayoutCase

	for _, size := range sizes {
		if err := candidate.Arrange(root, size); err != nil {
			return nil, fmt.Errorf("layout at %dx%d: %w", size.Width, size.Height, err)
		}

		plan := candidate.Plan()

		cases = append(cases, LayoutCase{Size: size, Plan: plan})
		for _, selection := range plan.Selections() {
			covered[selection] = true
		}
	}

	for _, variant := range variants {
		if !covered[variant] {
			return nil, fmt.Errorf("responsive variant %s/%s has no validation case", variant.Layout, variant.Variant)
		}
	}

	return cases, nil
}
