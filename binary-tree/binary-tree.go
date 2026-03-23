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

func (t *binaryTree) Sliced() [][]int {
	if t.root == nil {
		return nil
	}
	var st [][]int
	queue := []*node{t.root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var level []int

		for i := 0; i < levelSize; i++ {
			n := queue[0]
			queue = queue[1:]

			level = append(level, n.val)

			if n.left != nil {
				queue = append(queue, n.left)
			}
			if n.right != nil {
				queue = append(queue, n.right)
			}
		}

		st = append(st, level)
	}

	return st
}

func (t *binaryTree) Insert(val int) {
	if t.root == nil {
		t.root = &node{val: val}
		return
	}

	var inner func(cur *node)

	inner = func(cur *node) {
		if val < cur.val {
			if cur.left == nil {
				cur.left = &node{val: val}
				return
			}
			inner(cur.left)
		} else {
			if cur.right == nil {
				cur.right = &node{val: val}
				return
			}
			inner(cur.right)
		}
	}
	inner(t.root)
}

//func (t *binaryTree) Delete(n *node) bool {
//	if t.root == nil {
//		return false
//	}
//	var inner func(cur *node) bool
//
//	inner = func(cur *node) bool {
//		if cur == nil {
//			return false
//		}
//
//		if cur == n {
//
//			return true
//		}
//
//		inner(cur.left)
//		inner(cur.right)
//		return false
//	}
//}
