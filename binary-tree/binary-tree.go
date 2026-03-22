package binary_tree

type binaryTree struct {
	root *node
}

func NewBinaryTree(rootNodeValue int) *binaryTree {
	t := &node{
		val:   rootNodeValue,
		left:  nil,
		right: nil,
	}
	return &binaryTree{t}
}

func (t *binaryTree) Insert(val int) {
	n := &node{
		val: val,
	}

	if t.root == nil {
		t.root = n
		return
	}

	var inner func(cur *node)

	inner = func(cur *node) {
		if val > cur.val {
			if cur.left == nil {
				cur.left = n
				return
			}
			inner(cur.left)
		} else {
			if cur.right == nil {
				cur.right = n
				return
			}
			inner(cur.left)
		}
	}
	inner(t.root)
}
