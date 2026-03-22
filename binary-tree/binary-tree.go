package binary_tree

import (
	"strconv"
	"strings"
)

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

// String prints in BFS
func (t *binaryTree) String() string {
	if t.root == nil {
		return "empty"
	}
	b := &strings.Builder{}

	printer(t.root, b)

	return b.String()
}


func ()