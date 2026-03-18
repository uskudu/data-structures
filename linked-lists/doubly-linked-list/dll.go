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
	for cur != nil {
		prev := "nil"
		next := "nil"

		if cur.Previous != nil {
			prev = fmt.Sprintf("%d", cur.Previous.Value)
		}
		if cur.Next != nil {
			next = fmt.Sprintf("%d", cur.Next.Value)
		}

		fmt.Fprintf(&b, "{%s, %d, %s}", prev, cur.Value, next)

		if cur.Next != nil {
			b.WriteString(" <-> ")
		}

		cur = cur.Next
	}

	b.WriteString("]")
	return b.String()
}

func (l *LinkedList) PushFront(val int) {
	n := &node{Value: val}

	n.Next = l.Head
	n.Next.Previous = n
	l.Head = n

	if l.Tail == nil {
		l.Tail = n
	}
}

func (l *LinkedList) PushBack(val int) {
	n := &node{Value: val}

	if l.Tail == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	n.Previous = l.Tail
	l.Tail = n
}

// Sliced turns LinkedList into slice
func (l *LinkedList) Sliced() []int {
	var res []int
	if l.Head == nil {
		return res
	}

	cur := l.Head
	for cur != nil {
		res = append(res, cur.Value)
		cur = cur.Next
	}
	return res
}
