package tui

import (
	"fmt"
	"regexp"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/lanej/charmed/layout"
)

// StaticMaxHeight bounds the canvas height RenderStatic will search.
const StaticMaxHeight = 2048

// staticPasses bounds how often widths may change while heights are measured.
const staticPasses = 4

// StaticHeights reports the rendered height of each region measured so far.
type StaticHeights map[string]int

// Height returns the measured height of a region, or one row before the region
// has been measured. Declarations use it as the region's minimum height.
func (h StaticHeights) Height(id string) int { return max(1, h[id]) }

// StaticDeclare declares the layout for a canvas. RenderStatic calls it for a
// tall canvas to learn region widths and again for each candidate canvas, so it
// must be deterministic and free of effects.
type StaticDeclare func(canvas layout.Size, heights StaticHeights) layout.Node

// StaticRender renders one visible region at the given width. The text may have
// any height but no line may be wider than the width. Hidden regions are never
// rendered.
type StaticRender func(id string, width int) string

// StaticFrame is a composed static render and the plan it was composed from.
type StaticFrame struct {
	// Content has no padding: trailing blanks on each line and trailing blank
	// rows are removed, and styling sequences are preserved.
	Content     string
	Plan        layout.Plan
	Diagnostics []layout.Diagnostic
}

// RenderStatic composes a declared layout for output that is printed rather
// than run in a full-screen terminal, where a finite document has no fixed
// height. It declares the layout once to learn region widths, renders each
// visible region at its width, redeclares the layout with the measured heights
// and uses the smallest canvas height that arranges. It is deterministic and
// retains no state.
//
// Any arrangement, design-rule or content failure is returned as an error so
// that the caller can choose a stacked fallback. If heights change which
// recipe is selected so that widths keep changing, RenderStatic fails rather
// than guessing.
func RenderStatic(ids []string, width int, declare StaticDeclare, render StaticRender) (StaticFrame, error) {
	if width <= 0 {
		return StaticFrame{}, fmt.Errorf("static render requires a positive width")
	}

	if declare == nil || render == nil {
		return StaticFrame{}, fmt.Errorf("static render requires declare and render functions")
	}

	screen, err := NewScreen(ids...)
	if err != nil {
		return StaticFrame{}, err
	}

	heights := StaticHeights{}
	tall := layout.Size{Width: width, Height: StaticMaxHeight}

	if err = screen.Arrange(declare(tall, heights), tall); err != nil {
		return StaticFrame{}, fmt.Errorf("static layout at width %d: %w", width, err)
	}

	for range staticPasses {
		rendered := screen.Plan()

		views, measured, renderErr := renderVisible(rendered, render)
		if renderErr != nil {
			return StaticFrame{}, renderErr
		}

		heights = measured

		if arrangeErr := smallestCanvas(screen, width, declare, heights); arrangeErr != nil {
			return StaticFrame{}, arrangeErr
		}

		if sameWidths(rendered, screen.Plan()) {
			frame, composeErr := screen.Render(views)
			if composeErr != nil {
				return StaticFrame{}, composeErr
			}

			return StaticFrame{Content: trimPadding(frame.Content), Plan: screen.Plan(), Diagnostics: frame.Diagnostics}, nil
		}
	}

	return StaticFrame{}, fmt.Errorf("static layout at width %d did not settle after %d passes", width, staticPasses)
}

func renderVisible(plan layout.Plan, render StaticRender) (map[string]string, StaticHeights, error) {
	views, heights := map[string]string{}, StaticHeights{}

	for _, region := range plan.Regions() {
		content := render(region.ID, region.Bounds.Dx())
		if got := lipgloss.Width(content); got > region.Bounds.Dx() {
			return nil, nil, fmt.Errorf("static region %q rendered %d cells into width %d", region.ID, got, region.Bounds.Dx())
		}

		views[region.ID], heights[region.ID] = content, lipgloss.Height(content)
	}

	return views, heights, nil
}

// smallestCanvas arranges screen on the least canvas height that satisfies the
// declaration given the measured heights.
func smallestCanvas(screen *Screen, width int, declare StaticDeclare, heights StaticHeights) error {
	lowest := 1
	for _, h := range heights {
		lowest = max(lowest, h)
	}

	var first error

	for height := lowest; height <= StaticMaxHeight; height++ {
		canvas := layout.Size{Width: width, Height: height}

		err := screen.Arrange(declare(canvas, heights), canvas)
		if err == nil {
			return nil
		}

		if first == nil {
			first = err
		}
	}

	return fmt.Errorf("static layout at width %d has no arrangement up to height %d: %w", width, StaticMaxHeight, first)
}

// sameWidths reports whether both plans place the same regions at the same widths.
func sameWidths(a, b layout.Plan) bool {
	ra, rb := a.Regions(), b.Regions()
	if len(ra) != len(rb) {
		return false
	}

	for i := range ra {
		if ra[i].ID != rb[i].ID || ra[i].Bounds.Dx() != rb[i].Bounds.Dx() {
			return false
		}
	}

	return true
}

var trailingBlank = regexp.MustCompile(`(?:[ \t]|\x1b\[[0-9;:]*m)+$`)

var sgr = regexp.MustCompile(`\x1b\[[0-9;:]*m`)

// trimPadding removes layout padding while keeping styling sequences.
func trimPadding(content string) string {
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		lines[i] = trailingBlank.ReplaceAllStringFunc(line, func(run string) string {
			return strings.Join(sgr.FindAllString(run, -1), "")
		})
	}

	for len(lines) > 0 && ansi.Strip(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	return strings.Join(lines, "\n")
}
