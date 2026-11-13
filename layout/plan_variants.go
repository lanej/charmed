package layout

import "fmt"

// Variants lists every declared responsive branch, including nested branches.
// Container names identify coverage independently of terminal size or branch order.
func Variants(root Node) ([]Selection, error) {
	if err := validateNode(root, map[Node]bool{}); err != nil {
		return nil, err
	}

	seen := map[Node]bool{}
	names := map[string]*Responsive{}

	var (
		selections []Selection
		visit      func(Node) error
	)

	visit = func(node Node) error {
		if seen[node] {
			return nil
		}

		seen[node] = true

		var children []Node

		switch n := node.(type) {
		case *Region:
		case *Flex:
			for _, child := range n.Children {
				children = append(children, child.Node)
			}
		case *Grid:
			for _, cell := range n.Cells {
				children = append(children, cell.Node)
			}
		case *Flow:
			for _, child := range n.Children {
				children = append(children, child.Node)
			}
		case *Dock:
			for _, child := range []Node{n.Top, n.Bottom, n.Left, n.Right, n.Body} {
				if child != nil {
					children = append(children, child)
				}
			}
		case *SplitPane:
			children = append(children, n.First, n.Second)
		case *Responsive:
			if previous, exists := names[n.Name]; exists && previous != n {
				return fmt.Errorf("duplicate responsive container name %q", n.Name)
			}

			names[n.Name] = n
			for _, variant := range n.Variants {
				selections = append(selections, Selection{Layout: n.Name, Variant: variant.Name})
				children = append(children, variant.Root)
			}
		default:
			return fmt.Errorf("unsupported layout node %T", node)
		}

		for _, child := range children {
			if err := visit(child); err != nil {
				return err
			}
		}

		return nil
	}
	if err := visit(root); err != nil {
		return nil, err
	}

	return selections, nil
}
