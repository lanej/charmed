package tui

import (
	"fmt"
	"slices"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
)

// RegionView opts into caching. Advance Revision whenever any visible input,
// including native styles, focus or scroll position, changes. Zero is valid.
// Callback identity is not a cache key. Calls follow Screen's serialized lifecycle.
type RegionView struct {
	Revision uint64
	Render   func(layout.Size) string
}

type retainedRegion struct {
	revision      uint64
	size          layout.Size
	content       string
	width, height int
}

// ArrangeCompiled reuses unchanged allocations. Failure invalidates rendering.
func (s *Screen) ArrangeCompiled(c *layout.Compiled, size layout.Size) error {
	if c != nil && s.arranged && s.compiled == c && s.plan.Size() == size {
		return nil
	}

	s.arranged, s.retainedFrame = false, nil

	if c == nil {
		return fmt.Errorf("nil compiled layout")
	}

	plan, err := c.Resolve(size)
	if err != nil {
		return err
	}

	if err = s.acceptPlan(plan); err != nil {
		return err
	}

	s.compiled = c

	return nil
}

// RenderRetained retains only fitted output for currently visible regions.
func (s *Screen) RenderRetained(inputs map[string]RegionView) (Frame, error) {
	if !s.arranged {
		return Frame{}, fmt.Errorf("screen has no successful arrangement")
	}

	regions := s.plan.Regions()
	if len(inputs) != len(regions) {
		return Frame{}, fmt.Errorf("screen needs exactly %d visible regions", len(regions))
	}

	for _, region := range regions {
		input, ok := inputs[region.ID]
		if !ok || input.Render == nil {
			return Frame{}, fmt.Errorf("missing renderer for %s", region.ID)
		}
	}

	if s.retainedOwner != s {
		s.retained, s.retainedFrame = nil, nil
		s.retainedOwner = s
	}

	next := make(map[string]retainedRegion, len(regions))
	frame := Frame{Diagnostics: s.plan.Diagnostics()}
	changed := s.retainedFrame == nil

	var layers []*lipgloss.Layer

	for _, region := range regions {
		input := inputs[region.ID]
		size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}

		cached, ok := s.retained[region.ID]
		if !ok || cached.revision != input.Revision || cached.size != size {
			content := input.Render(size)
			cached = retainedRegion{revision: input.Revision, size: size, content: strings.Clone(layout.Fit(content, size)), width: lipgloss.Width(content), height: lipgloss.Height(content)}
			changed = true
		}

		if cached.width > size.Width || cached.height > size.Height {
			frame.Diagnostics = append(frame.Diagnostics, layout.Diagnostic{Rule: layout.ContentClipped, Scope: region.ID, Exception: region.ClippingException, Regions: []string{region.ID}, Message: fmt.Sprintf("%s rendered %dx%d into %dx%d cells", region.ID, cached.width, cached.height, size.Width, size.Height)})
		}

		next[region.ID] = cached
		layers = append(layers, lipgloss.NewLayer(cached.content).X(region.Bounds.Min.X).Y(region.Bounds.Min.Y))
	}

	if err := layout.Enforce(frame.Diagnostics); err != nil {
		s.retained, s.retainedFrame = nil, nil
		return Frame{}, err
	}

	if changed {
		frame.Content = s.compose(layers)
	} else {
		frame.Content = s.retainedFrame.Content
	}

	s.retained = next
	saved := frame
	saved.Diagnostics = cloneDiagnostics(frame.Diagnostics)
	s.retainedFrame = &saved

	return frame, nil
}

func cloneDiagnostics(source []layout.Diagnostic) []layout.Diagnostic {
	result := slices.Clone(source)
	for i := range result {
		result[i].Regions = slices.Clone(result[i].Regions)
	}

	return result
}
