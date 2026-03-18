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

func NewLinkedList(headVal int) *LinkedList {
	head := &node{
		Value: headVal,
	}
	return &LinkedList{Head: head, Tail: head}
}

// String returns string of nodes as if they were slice of struct {Previous int, Value int, Next int}
// example: [{nil, 1, 2} <-> {1, 2, 5} <-> {2, 5, 7} <-> {5, 7, nil}]
func (l *LinkedList) String() string {
	if l.Head == nil {
		return "[]"
	}
	b := strings.Builder{}

	b.WriteString("[")

	cur := l.Head
	for cur.Next != nil {
		if cur.Previous != nil {
			fmt.Fprintf(&b, "{%d, %d, %d} <-> ", cur.Previous.Value, cur.Value, cur.Next.Value)
			cur = cur.Next
		}
		fmt.Fprintf(&b, "{nil, %d, %d} <-> ", cur.Value, cur.Next.Value)
		cur = cur.Next
	}
	fmt.Fprintf(&b, "{nil, %d, nil}", cur.Value)

	b.WriteString("]")
	return b.String()
}

func (l *LinkedList) PushFront(val int) {
	n := &node{Value: val}

	n.Next = l.Head
	l.Head = n

	if l.Tail == nil {
		l.Tail = n
	}
}
