package singly_linked_list

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList(head *Node) *LinkedList {
	if head == nil {
		return &LinkedList{}
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
		fmt.Fprintf(&b, "{%d, %d}, ", cur.Value, cur.Next.Value)
		cur = cur.Next
	}
	fmt.Fprintf(&b, "{%d, nil}", cur.Value)

	b.WriteString("]")
	return b.String()
}

func (l *LinkedList) PushFront(val int) {
	n := &Node{Value: val}

	n.Next = l.Head
	l.Head = n

	if l.Tail == nil {
		l.Tail = n
	}
}

func (l *LinkedList) PushBack(val int) {
	n := &Node{Value: val}

	if l.Tail == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n
}

func (l *LinkedList) DelFront() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next

	if l.Head == nil {
		l.Tail = nil
	}
}

func (l *LinkedList) DelBack() {
	if l.Tail == nil {
		return
	}
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		return
	}

	cur := l.Head
	for cur.Next != l.Tail {
		cur = cur.Next
	}

	cur.Next = nil
	l.Tail = cur
}

// Push sets node with Value val at index zeroBasedIdx
func (l *LinkedList) Push(zeroBasedIdx, val int) bool {
	if zeroBasedIdx < 0 {
		return false
	}
	// use PushFront for inserting front
	if zeroBasedIdx == 0 {
		l.PushFront(val)
		return true
	}

	n := &Node{Value: val}

	// if no elems in LinkedList
	if l.Head == nil {
		l.Head = n
		return true
	}

	// if one elem in LinkedList
	if l.Head == l.Tail {
		l.Head.Next = n
		return true
	}

	// if more than one elem in LinkedList
	cur := l.Head
	prev := l.Head
	curIdx := 0

	for curIdx != zeroBasedIdx && cur != nil {
		prev = cur
		cur = cur.Next
		curIdx++
	}

	prev.Next = n
	n.Next = cur
	cur = n
	return true
}

// Del deletes node with index zeroBasedIdx
func (l *LinkedList) Del(zeroBasedIdx int) bool {
	if zeroBasedIdx < 0 {
		return false
	}
	// use DelFront for inserting front
	if zeroBasedIdx == 0 {
		l.DelFront()
		return true
	}

	// if no elems in LinkedList
	if l.Head == nil {
		return false
	}

	// if one elem in LinkedList
	if l.Head == l.Tail {
		l.Head = nil
		return true
	}

	// if more than one elem in LinkedList
	prev := l.Head
	cur := l.Head
	curIdx := 0

	for curIdx != zeroBasedIdx && cur.Next != nil {
		prev = cur
		cur = cur.Next
		curIdx++
	}

	prev.Next = cur.Next
	cur.Next = nil
	return true
}

func (l *LinkedList) Len() int {
	var res int

	cur := l.Head
	for cur != nil {
		cur = cur.Next
		res++
	}
	return res
}

// GetNodeByIndex returns Node with index zeroBasedIdx
//func (l *LinkedList) GetNodeByIndex(zeroBasedIdx int) Node {
//	if zeroBasedIdx
//}

//func (l *LinkedList) DelNode(node Node) bool {
//
//}

// BulkFront pushes multiple values into front of LinkedList in provided order
//func (l *LinkedList) BulkFront(values []int) {
//
//}

// BulkBack pushes multiple values into back of LinkedList in provided order
//func (l *LinkedList) BulkBack(values []int) {
//
//}
