package layout

import "slices"

// Compiled owns a validated declaration independent of subsequent caller edits.
type Compiled struct{ root Node }

func Compile(root Node) (*Compiled, error) {
	if _, err := Variants(root); err != nil {
		return nil, err
	}

	return &Compiled{root: cloneNode(root)}, nil
}

func (c *Compiled) Resolve(size Size) (Plan, error) {
	return resolveValidated(c.root, size)
}

func cloneNode(node Node) Node {
	if node == nil {
		return nil
	}

	switch n := node.(type) {
	case *Region:
		v := *n
		v.Exceptions = slices.Clone(n.Exceptions)

		return &v
	case *Flex:
		v := *n

		v.Exceptions, v.Children = slices.Clone(n.Exceptions), slices.Clone(n.Children)
		for i := range v.Children {
			v.Children[i].Node = cloneNode(v.Children[i].Node)
		}

		return &v
	case *Grid:
		v := *n

		v.Columns, v.Rows, v.Cells = slices.Clone(n.Columns), slices.Clone(n.Rows), slices.Clone(n.Cells)
		for i := range v.Cells {
			v.Cells[i].Node = cloneNode(v.Cells[i].Node)
		}

		return &v
	case *Flow:
		v := *n

		v.Children = slices.Clone(n.Children)
		for i := range v.Children {
			v.Children[i].Node = cloneNode(v.Children[i].Node)
		}

		return &v
	case *Dock:
		v := *n
		v.Top, v.Bottom, v.Left, v.Right, v.Body = cloneNode(n.Top), cloneNode(n.Bottom), cloneNode(n.Left), cloneNode(n.Right), cloneNode(n.Body)

		return &v
	case *SplitPane:
		v := *n
		v.Exceptions = slices.Clone(n.Exceptions)
		v.First, v.Second = cloneNode(n.First), cloneNode(n.Second)

		return &v
	case *Responsive:
		v := *n

		v.Variants = slices.Clone(n.Variants)
		for i := range v.Variants {
			v.Variants[i].Root = cloneNode(v.Variants[i].Root)
			v.Variants[i].Hidden = slices.Clone(v.Variants[i].Hidden)
		}

		return &v
	default:
		panic("unknown layout node")
	}
}
