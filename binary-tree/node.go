package binary_tree

type node struct {
	val   int
	left  *node
	right *node
}

func (n *node) CountDescendants() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.CountDescendants() + n.right.CountDescendants()
}
