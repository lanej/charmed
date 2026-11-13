// Package tui composes native terminal views through declared layout regions.
// It owns placement validation and diagnostics, while applications retain their
// Bubble Tea models, commands, focus routing and hidden component state.
package tui

import (
	"fmt"
	"image"
	"strings"
	"unicode"

	"charm.land/lipgloss/v2"
	"github.com/lanej/charmed/layout"
	"github.com/lanej/charmed/theme"
)

// Screen registers the complete set of screen regions. Arrange requires each to
// be placed exactly once or explicitly hidden with a reason. Models are not stored.
type Screen struct {
	ids      []string
	plan     layout.Plan
	arranged bool
}

type Frame struct {
	Content     string
	Diagnostics []layout.Diagnostic
}

func NewScreen(ids ...string) (*Screen, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("screen requires registered regions")
	}

	seen := map[string]bool{}
	for _, id := range ids {
		if strings.TrimSpace(id) == "" || strings.IndexFunc(id, unicode.IsControl) >= 0 || seen[id] {
			return nil, fmt.Errorf("empty, unsafe or duplicate screen region %q", id)
		}

		seen[id] = true
	}

	return &Screen{ids: append([]string(nil), ids...)}, nil
}

// Arrange requires a successful allocation for the current size. Failure invalidates
// rendering until the next success; caller-owned component state remains untouched.
func (s *Screen) Arrange(root layout.Node, size layout.Size) error {
	s.arranged = false

	plan, err := layout.Resolve(root, size)
	if err != nil {
		return err
	}

	declared := map[string]bool{}
	for _, region := range plan.Regions() {
		declared[region.ID] = true
	}

	for _, hidden := range plan.Hidden() {
		declared[hidden.ID] = true
	}

	for _, id := range s.ids {
		if !declared[id] {
			return fmt.Errorf("screen region %q is neither placed nor explicitly hidden", id)
		}

		delete(declared, id)
	}

	if len(declared) > 0 {
		for id := range declared {
			return fmt.Errorf("layout names unregistered screen region %q", id)
		}
	}

	s.plan, s.arranged = plan, true

	return nil
}

func (s *Screen) Plan() layout.Plan {
	if !s.arranged {
		return layout.Plan{}
	}

	return s.plan
}

func (s *Screen) Region(id string) (image.Rectangle, bool) {
	if !s.arranged {
		return image.Rectangle{}, false
	}

	return s.plan.Bounds(id)
}

// Render requires exactly the visible region views; hidden or unknown input is an
// error. Oversized content blocks rendering unless that region declares a
// ContentClipped exception; accepted exceptions remain visible in diagnostics.
func (s *Screen) Render(views map[string]string) (Frame, error) {
	if !s.arranged {
		return Frame{}, fmt.Errorf("screen has no successful arrangement")
	}

	visible := s.plan.Regions()
	if len(views) != len(visible) {
		return Frame{}, fmt.Errorf("screen needs exactly %d visible region views; got %d", len(visible), len(views))
	}

	frame := Frame{Diagnostics: s.plan.Diagnostics()}
	size := s.plan.Size()
	layers := []*lipgloss.Layer{lipgloss.NewLayer(layout.Fit("", size))}

	for _, region := range visible {
		content, exists := views[region.ID]
		if !exists {
			return Frame{}, fmt.Errorf("missing view for screen region %q", region.ID)
		}

		allocation := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
		if lipgloss.Width(content) > allocation.Width || lipgloss.Height(content) > allocation.Height {
			frame.Diagnostics = append(frame.Diagnostics, layout.Diagnostic{Rule: layout.ContentClipped, Scope: region.ID, Exception: region.ClippingException, Regions: []string{region.ID}, Message: fmt.Sprintf("%s rendered %dx%d into %dx%d cells", region.ID, lipgloss.Width(content), lipgloss.Height(content), allocation.Width, allocation.Height)})
		}

		layers = append(layers, lipgloss.NewLayer(layout.Fit(content, allocation)).X(region.Bounds.Min.X).Y(region.Bounds.Min.Y))
	}

	if err := layout.Enforce(frame.Diagnostics); err != nil {
		return Frame{}, err
	}

	frame.Content = layout.Fit(lipgloss.NewCompositor(layers...).Render(), size)

	return frame, nil
}

// RegionMap shows the same allocations used for rendering, including their gutters
// only when the application explicitly placed a panel in the region.
func (s *Screen) RegionMap(t theme.Theme) string {
	if !s.arranged {
		return "Screen has no arrangement"
	}

	layers := []*lipgloss.Layer{lipgloss.NewLayer(layout.Fit("", s.plan.Size()))}
	for _, region := range s.plan.Regions() {
		size := layout.Size{Width: region.Bounds.Dx(), Height: region.Bounds.Dy()}
		label := fmt.Sprintf("%s (%d,%d) %dx%d", region.ID, region.Bounds.Min.X, region.Bounds.Min.Y, size.Width, size.Height)
		content := layout.Panel(region.ID, label, size, t, false)
		layers = append(layers, lipgloss.NewLayer(content).X(region.Bounds.Min.X).Y(region.Bounds.Min.Y))
	}

	return layout.Fit(lipgloss.NewCompositor(layers...).Render(), s.plan.Size())
}

// Inspect is a textual diagnostic report; it does not consume screen allocation.
func (s *Screen) Inspect() string {
	var lines []string

	if !s.arranged {
		return "Screen has no arrangement"
	}

	lines = append(lines, fmt.Sprintf("Layout %s · %dx%d", s.plan.Variant(), s.plan.Size().Width, s.plan.Size().Height))
	for _, region := range s.plan.Regions() {
		lines = append(lines, fmt.Sprintf("%s: x=%d y=%d width=%d height=%d", region.ID, region.Bounds.Min.X, region.Bounds.Min.Y, region.Bounds.Dx(), region.Bounds.Dy()))
	}

	for _, hidden := range s.plan.Hidden() {
		lines = append(lines, fmt.Sprintf("%s: hidden — %s", hidden.ID, hidden.Reason))
	}

	for _, diagnostic := range s.plan.Diagnostics() {
		lines = append(lines, string(diagnostic.Rule)+": "+diagnostic.Message+" · exception: "+diagnostic.Exception)
	}

	return strings.Join(lines, "\n")
}
