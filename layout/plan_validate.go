package layout

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"unicode"
)

func validSize(size Size) bool { return size.Width >= 0 && size.Height >= 0 }
func validAxis(axis Axis) bool {
	switch axis {
	case Horizontal, Vertical:
		return true
	default:
		return false
	}
}

func validAlignment(align Alignment) bool {
	switch align {
	case Start, Center, End:
		return true
	default:
		return false
	}
}

func validRole(role RegionRole) bool {
	switch role {
	case ContentRegion, ChromeRegion, SupportingRegion:
		return true
	default:
		return false
	}
}

func validateTrack(track Track) error {
	if track.Min < 0 || track.Max < 0 || track.Weight < 0 || (track.Max > 0 && track.Max < track.Min) || (track.Weight == 0 && track.Max != track.Min) {
		return fmt.Errorf("invalid track: %+v", track)
	}

	return nil
}

func validateNode(node Node, path map[Node]bool) error {
	if node == nil {
		return fmt.Errorf("layout node is nil")
	}

	value := reflect.ValueOf(node)
	if value.Kind() != reflect.Pointer {
		return fmt.Errorf("unsupported layout node %T", node)
	}

	if value.IsNil() {
		return fmt.Errorf("layout node is nil")
	}

	if path[node] {
		return fmt.Errorf("layout contains a cycle")
	}

	path[node] = true
	defer delete(path, node)

	var children []Node

	switch n := node.(type) {
	case *Region:
		if (n.Preferred.Width > 0 && n.Preferred.Width < n.Min.Width) || (n.Preferred.Height > 0 && n.Preferred.Height < n.Min.Height) {
			return fmt.Errorf("preferred size is below minimum for %q", n.ID)
		}

		if err := validateExceptions(n.Exceptions, n.ID, SingleRegion, BelowPreferred, ContentClipped); err != nil {
			return err
		}

		if strings.TrimSpace(n.ID) == "" || strings.IndexFunc(n.ID, unicode.IsControl) >= 0 || !validSize(n.Min) || !validSize(n.Preferred) || !validRole(n.Role) {
			return fmt.Errorf("invalid region %q", n.ID)
		}
	case *Flex:
		if err := validateExceptions(n.Exceptions, n.Name, StackOpportunity); err != nil {
			return err
		}

		if !validAxis(n.Axis) || !validAlignment(n.Align) || !validAlignment(n.Justify) || n.Gap < 0 || len(n.Children) == 0 {
			return fmt.Errorf("invalid flex declaration")
		}

		for _, child := range n.Children {
			if err := validateTrack(child.Track); err != nil {
				return err
			}

			if child.Cross < 0 {
				return fmt.Errorf("negative flex cross size")
			}

			children = append(children, child.Node)
		}
	case *Grid:
		if len(n.Columns) == 0 || len(n.Rows) == 0 || len(n.Cells) == 0 || !validSize(n.Gap) {
			return fmt.Errorf("invalid grid declaration")
		}

		for _, tracks := range [][]Track{n.Columns, n.Rows} {
			for _, track := range tracks {
				if err := validateTrack(track); err != nil {
					return err
				}
			}
		}

		occupied := map[[2]int]bool{}

		for _, cell := range n.Cells {
			if cell.Column < 0 || cell.Row < 0 || cell.ColumnSpan < 0 || cell.RowSpan < 0 || cell.Column >= len(n.Columns) || cell.Row >= len(n.Rows) || max(1, cell.ColumnSpan) > len(n.Columns)-cell.Column || max(1, cell.RowSpan) > len(n.Rows)-cell.Row {
				return fmt.Errorf("grid cell outside declared tracks")
			}

			for y := cell.Row; y < cell.Row+max(1, cell.RowSpan); y++ {
				for x := cell.Column; x < cell.Column+max(1, cell.ColumnSpan); x++ {
					if occupied[[2]int{x, y}] {
						return fmt.Errorf("grid cells overlap at %d,%d", x, y)
					}

					occupied[[2]int{x, y}] = true
				}
			}

			children = append(children, cell.Node)
		}
	case *Flow:
		if !validSize(n.Gap) || len(n.Children) == 0 {
			return fmt.Errorf("invalid flow declaration")
		}

		for _, child := range n.Children {
			if child.Size.Width <= 0 || child.Size.Height <= 0 {
				return fmt.Errorf("flow boxes require positive sizes")
			}

			children = append(children, child.Node)
		}
	case *Dock:
		if n.Gap < 0 {
			return fmt.Errorf("negative dock gap")
		}

		for _, edge := range []struct {
			node Node
			size int
		}{{n.Top, n.TopHeight}, {n.Bottom, n.BottomHeight}, {n.Left, n.LeftWidth}, {n.Right, n.RightWidth}} {
			if edge.node != nil {
				if edge.size <= 0 {
					return fmt.Errorf("dock edge requires a positive size")
				}

				children = append(children, edge.node)
			} else if edge.size != 0 {
				return fmt.Errorf("absent dock edge has a size")
			}
		}

		children = append(children, n.Body)
	case *SplitPane:
		if err := validateExceptions(n.Exceptions, n.Name, StackOpportunity); err != nil {
			return err
		}

		if !validAxis(n.Axis) || n.Gap < 0 || math.IsNaN(n.Ratio) || math.IsInf(n.Ratio, 0) || n.Ratio < 0 || n.Ratio >= 1 {
			return fmt.Errorf("invalid split declaration")
		}

		children = append(children, n.First, n.Second)
	case *Responsive:
		if strings.TrimSpace(n.Name) == "" || len(n.Variants) == 0 {
			return fmt.Errorf("responsive layout has no variants")
		}

		names := map[string]bool{}
		for _, variant := range n.Variants {
			if strings.TrimSpace(variant.Name) == "" || names[variant.Name] || !validSize(variant.Min) {
				return fmt.Errorf("invalid responsive variant %q", variant.Name)
			}

			names[variant.Name] = true
			for _, hidden := range variant.Hidden {
				if strings.TrimSpace(hidden.ID) == "" || strings.TrimSpace(hidden.Reason) == "" {
					return fmt.Errorf("hidden regions require identity and reason")
				}
			}

			children = append(children, variant.Root)
		}
	default:
		return fmt.Errorf("unsupported layout node %T", node)
	}

	for _, child := range children {
		if err := validateNode(child, path); err != nil {
			return err
		}
	}

	return nil
}

func validateExceptions(exceptions []Exception, scope string, allowed ...Rule) error {
	seen := map[Rule]bool{}

	for _, exception := range exceptions {
		permitted := false

		for _, rule := range allowed {
			if exception.Rule == rule {
				permitted = true
			}
		}

		if !permitted || seen[exception.Rule] || strings.TrimSpace(exception.Reason) == "" || strings.TrimSpace(scope) == "" {
			return fmt.Errorf("invalid or unscoped exception for %s", exception.Rule)
		}

		seen[exception.Rule] = true
	}

	return nil
}
