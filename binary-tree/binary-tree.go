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
	newNode := &node{val: val}

	if t.root == nil {
		t.root = newNode
		return
	}

	cur := t.root
	for {
		if val < cur.val {
			if cur.left == nil {
				cur.left = newNode
				return
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = newNode
				return
			}
			cur = cur.right
		}
	}
}

//func (t *binaryTree) Delete(n *node) bool {
//	if t.root == nil {
//		return false
//	}
//	var inner func(cur *node) bool
//	prev := t.root
//
//	inner = func(cur *node) bool {
//		if cur == nil {
//			return false
//		}
//
//
//		if cur == n {
//			if prev.left == n {
//				prev.left = nil
//				prev.
//			}
//			return true
//		}
//
//		prev = cur
//
//		inner(cur.left)
//		inner(cur.right)
//		return false
//	}
//
//	return inner(t.root)
//}
