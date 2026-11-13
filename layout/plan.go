package layout

import (
	"fmt"
	"image"
	"strings"
)

// Node declares terminal allocation, independently of component models and rendering.
//
//sumtype:decl
type Node interface{ layoutNode() }

type Axis uint8

const (
	Horizontal Axis = iota
	Vertical
)

type Alignment uint8

const (
	Start Alignment = iota
	Center
	End
)

type RegionRole uint8

const (
	ContentRegion RegionRole = iota
	ChromeRegion
	SupportingRegion
)

// Region is a named allocation. Min is a hard geometry requirement; Preferred
// is a design rule. Exceptions are scoped to this region and require a reason.
type Region struct {
	ID             string
	Min, Preferred Size
	Role           RegionRole
	Exceptions     []Exception
}

func (*Region) layoutNode() {}

// Track reserves Min cells, then shares remaining cells by Weight up to Max.
// Max zero is unbounded for weighted tracks. A zero-weight track is fixed at Min.
type Track struct{ Min, Max, Weight int }

func Fixed(cells int) Track  { return Track{Min: cells, Max: cells} }
func Grow(minimum int) Track { return Track{Min: minimum, Weight: 1} }

type Slot struct {
	Node  Node
	Track Track
	// Cross zero stretches; positive Cross uses the container's cross-axis alignment.
	Cross int
}

// Flex allocates weighted rows or columns. Exceptions apply only to this container.
type Flex struct {
	Name       string
	Axis       Axis
	Gap        int
	Align      Alignment
	Justify    Alignment
	Children   []Slot
	Exceptions []Exception
}

func (*Flex) layoutNode() {}

type Cell struct {
	Node                Node
	Column, Row         int
	ColumnSpan, RowSpan int
}

type Grid struct {
	Columns, Rows []Track
	Gap           Size
	Cells         []Cell
}

func (*Grid) layoutNode() {}

type Box struct {
	Node Node
	Size Size
}

// Flow wraps whole, fixed-size boxes. It fails if even a wrapped row cannot fit.
type Flow struct {
	Gap      Size
	Children []Box
}

func (*Flow) layoutNode() {}

// Dock reserves chrome around Body; absent edges reserve no cells or gap.
type Dock struct {
	Top, Bottom, Left, Right, Body                 Node
	TopHeight, BottomHeight, LeftWidth, RightWidth int
	Gap                                            int
}

func (*Dock) layoutNode() {}

// SplitPane preserves both panes' minima. Ratio zero means an equal split.
// Responsive chooses a different layout when the two panes cannot fit.
type SplitPane struct {
	Name          string
	Axis          Axis
	First, Second Node
	Ratio         float64
	Gap           int
	Exceptions    []Exception
}

func (*SplitPane) layoutNode() {}

type Omission struct{ ID, Reason string }

// Variant ordering is explicit: the first minimum width AND height that fits wins.
// Hidden names regions intentionally absent in this variant, with their reasons.
type (
	Variant struct {
		Name   string
		Min    Size
		Root   Node
		Hidden []Omission
	}
	Responsive struct {
		Name     string
		Variants []Variant
	}
)

func (*Responsive) layoutNode() {}

type Rule string

const (
	BelowPreferred   Rule = "below-preferred"
	StackOpportunity Rule = "stack-opportunity"
	SingleRegion     Rule = "single-region"
	ContentClipped   Rule = "content-clipped"
)

type Exception struct {
	Rule   Rule
	Reason string
}

// Diagnostic records a design violation and, when accepted, its local exception.
type Diagnostic struct {
	Rule      Rule
	Scope     string
	Regions   []string
	Message   string
	Exception string
}

type RuleError struct{ Diagnostics []Diagnostic }

func (e *RuleError) Error() string {
	if len(e.Diagnostics) == 0 {
		return "layout design rules failed"
	}

	first := e.Diagnostics[0]

	return fmt.Sprintf("layout rule %s at %s: %s", first.Rule, first.Scope, first.Message)
}

// Enforce accepts only diagnostics carrying a scoped exception. Geometry errors
// never enter this path and cannot be exempted.
func Enforce(diagnostics []Diagnostic) error {
	var failures []Diagnostic

	for _, diagnostic := range diagnostics {
		if strings.TrimSpace(diagnostic.Exception) == "" {
			failures = append(failures, diagnostic)
		}
	}

	if len(failures) > 0 {
		return &RuleError{Diagnostics: failures}
	}

	return nil
}

func exceptionReason(exceptions []Exception, rule Rule) string {
	for _, exception := range exceptions {
		if exception.Rule == rule {
			return exception.Reason
		}
	}

	return ""
}

type (
	Selection struct{ Layout, Variant string }
	Placement struct {
		ID                string
		Bounds            image.Rectangle
		ClippingException string
		singleException   string
		role              RegionRole
	}
)

// Plan is an immutable allocation snapshot. Accessors return independent slices.
type planner struct{ Plan }

type Plan struct {
	size        Size
	placements  []Placement
	hidden      []Omission
	diagnostics []Diagnostic
	variant     string
	selections  []Selection
}

func (p Plan) Size() Size              { return p.size }
func (p Plan) Variant() string         { return p.variant }
func (p Plan) Selections() []Selection { return append([]Selection(nil), p.selections...) }
func (p Plan) Regions() []Placement    { return append([]Placement(nil), p.placements...) }
func (p Plan) Hidden() []Omission      { return append([]Omission(nil), p.hidden...) }
func (p Plan) Diagnostics() []Diagnostic {
	out := append([]Diagnostic(nil), p.diagnostics...)
	for i := range out {
		out[i].Regions = append([]string(nil), out[i].Regions...)
	}

	return out
}

func (p Plan) Bounds(id string) (image.Rectangle, bool) {
	for _, region := range p.placements {
		if region.ID == id {
			return region.Bounds, true
		}
	}

	return image.Rectangle{}, false
}

// Resolve rejects malformed declarations and impossible allocations without a partial plan.
// Node shapes and cycles are validated in every branch. Selected identities are
// checked here; tui.Screen.CheckLayouts exercises all declared responsive variants.
func Resolve(root Node, size Size) (Plan, error) {
	if size.Width <= 0 || size.Height <= 0 {
		return Plan{}, fmt.Errorf("layout requires a positive terminal size")
	}

	if _, err := Variants(root); err != nil {
		return Plan{}, err
	}

	builder := planner{Plan: Plan{size: size}}
	if err := builder.place(root, image.Rect(0, 0, size.Width, size.Height)); err != nil {
		return Plan{}, err
	}

	plan := builder.Plan

	seen := map[string]bool{}
	for _, region := range plan.placements {
		if seen[region.ID] {
			return Plan{}, fmt.Errorf("duplicate region %q", region.ID)
		}

		seen[region.ID] = true
	}

	for _, hidden := range plan.hidden {
		if seen[hidden.ID] {
			return Plan{}, fmt.Errorf("region %q is placed or hidden more than once", hidden.ID)
		}

		seen[hidden.ID] = true
	}

	if len(plan.placements) == 1 {
		region := plan.placements[0]
		plan.diagnostics = append(plan.diagnostics, Diagnostic{Rule: SingleRegion, Scope: region.ID, Regions: []string{region.ID}, Message: "The screen has one visible region; declare internal regions or document the intended single surface.", Exception: region.singleException})
	}

	if err := Enforce(plan.diagnostics); err != nil {
		return Plan{}, err
	}

	return plan, nil
}
