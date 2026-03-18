package singly_linked_list

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
		fmt.Fprintf(&b, "{%d, %d} -> ", cur.Value, cur.Next.Value)
		cur = cur.Next
	}
	fmt.Fprintf(&b, "{%d, nil}", cur.Value)

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

func (l *LinkedList) PushBack(val int) {
	n := &node{Value: val}

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
// returns true if pushed, else false
func (l *LinkedList) Push(zeroBasedIdx, val int) bool {
	if zeroBasedIdx < 0 {
		return false
	}
	// use PushFront for inserting front
	if zeroBasedIdx == 0 {
		l.PushFront(val)
		return true
	}

	n := &node{Value: val}

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

// DelIndex deletes node with index zeroBasedIdx
// returns true if deleted, else false
func (l *LinkedList) DelIndex(zeroBasedIdx int) bool {
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

// GetNodeByIndex returns node with index zeroBasedIdx
func (l *LinkedList) GetNodeByIndex(zeroBasedIdx int) *node {
	if zeroBasedIdx < 0 || l.Head == nil {
		return nil
	}
	if l.Head == l.Tail {
		return l.Head
	}

	curIdx := 0
	cur := l.Head
	for curIdx != zeroBasedIdx && cur != nil {
		if curIdx > zeroBasedIdx {
			return nil
		}
		cur = cur.Next
		curIdx++
	}
	return cur
}

// BulkFront pushes multiple values into front of LinkedList in provided order
//func (l *LinkedList) BulkFront(values []int) {
//
//}

// BulkBack pushes multiple values into back of LinkedList in provided order
//func (l *LinkedList) BulkBack(values []int) {
//
//}

//func (l *LinkedList) Reverse() {
//
//}

// ToSlice turns LinkedList into slice
func (l *LinkedList) ToSlice() []int {
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
