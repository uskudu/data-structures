package binary_tree

import (
	"strconv"
	"strings"
)

type node struct {
	val   int
	left  *node
	right *node
}

const (
	leaf = iota
	unary
	binary
)

// Relationships todo
//func (n *node) Relationships() int {
//	return -1
//}

func (n *node) CountDescendants() int {
	c := 0
	if n == nil {
		return c
	}
	c++
	if n.left != nil {
		n.left.CountDescendants()
	}
	if n.right != nil {
		n.right.CountDescendants()
	}
	return c
}

var inLeft bool

func printer(n *node, b *strings.Builder) {
	if n == nil {
		b.WriteString("")
		return
	}

	b.WriteString(strings.Repeat(" ", n.CountDescendants()))
	b.WriteString(strconv.Itoa(n.val))
	if inLeft {
		b.WriteString("\n/")
	} else {
		b.WriteString("\n\\")
	}

	if n.left != nil {
		inLeft = true
		printer(n.left, b)
	}
	if n.right != nil {
		inLeft = false
		printer(n.right, b)
	}
}
