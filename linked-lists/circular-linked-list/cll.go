package circular_linked_list

import (
	"fmt"
	"strings"
)

type node struct {
	Value int
	Next  *node
}

type LinkedList struct {
	Head *node
	Tail *node
}

func NewLinkedList(headVal int) *LinkedList {
	head := &node{
		Value: headVal,
	}
	head.Next = head
	return &LinkedList{Head: head, Tail: head}
}

// String returns string of nodes as if they were slice of struct {val int, next int}
// example: [{1, 2} -> {2, 5} -> {5, 7} -> {7, nil}]
func (l *LinkedList) String() string {
	if l.Head == nil {
		return "[]"
	}
	b := strings.Builder{}

	b.WriteString("[")

	cur := l.Head
	for {
		fmt.Fprintf(&b, "{%d, %d}", cur.Value, cur.Next.Value)

		cur = cur.Next
		if cur == l.Head {
			break
		}
		b.WriteString(" -> ")
	}

	b.WriteString("]")
	return b.String()
}

func (l *LinkedList) PushFront(val int) {
	n := &node{Value: val}

	n.Next = l.Head
	l.Head = n
	l.Tail.Next = n

	if l.Tail == nil {
		l.Tail = n
	}
}
