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

func (t *binaryTree) Delete(val int) {
	t.root = deleteNode(t.root, val)
}

func deleteNode(n *node, val int) *node {
	if n == nil {
		return nil
	}

	if val < n.val {
		n.left = deleteNode(n.left, val)
	} else if val > n.val {
		n.right = deleteNode(n.right, val)
	} else {
		// no children
		if n.left == nil && n.right == nil {
			return nil
		}
		// one child
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		// two children: replace with in-order successor (smallest in right subtree)
		successor := n.right
		for successor.left != nil {
			successor = successor.left
		}
		n.val = successor.val
		n.right = deleteNode(n.right, successor.val)
	}

	return n
}
