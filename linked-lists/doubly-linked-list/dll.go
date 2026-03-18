package doubly_linked_list

import (
	"fmt"
	"strings"
)

type node struct {
	Value    int
	Next     *node
	Previous *node
}

type LinkedList struct {
	Head *node
	Tail *node
}

func NewNode(val int) *node {
	return &node{Value: val}
}

func NewLinkedList(head *node) *LinkedList {
	if head == nil {
		return &LinkedList{}
	}
	if head.Next != nil {
		return nil
	}
	return &LinkedList{Head: head, Tail: head}
}

// String returns string of nodes as if they were slice of struct {val int, next int}
// example: [{1, 2}, {2, 5}, {5, 7}, {7, nil}]
func (l *LinkedList) String() string {
	if l.Head == nil {
		return "[]"
	}
	b := strings.Builder{}

	b.WriteString("[")

	cur := l.Head
	for cur.Next != nil {
		fmt.Fprintf(&b, "{%d, %d} <-> ", cur.Value, cur.Next.Value)
		cur = cur.Next
	}
	fmt.Fprintf(&b, "{%d, nil}", cur.Value)

	b.WriteString("]")
	return b.String()
}
