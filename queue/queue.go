package queue

import (
	"fmt"
	"strings"
)

// using linked list for implementation of queue

type node struct {
	val  int
	next *node
}

type Queue struct {
	head *node
	tail *node
}

func (q *Queue) String() string {
	if q.head == nil {
		return "[]"
	}
	b := strings.Builder{}

	b.WriteString("[")

	cur := q.head
	for cur.next != nil {
		fmt.Fprintf(&b, "%d ", cur.val)
		cur = cur.next
	}
	fmt.Fprintf(&b, "%d", cur.val)

	b.WriteString("]")
	return b.String()
}

func NewQueue(headVal int) *Queue {
	head := &node{
		val: headVal,
	}
	return &Queue{head: head, tail: head}
}

func (q *Queue) Length() int {
	cur := q.head
	var c int
	for cur != nil {
		c++
		cur = cur.next
	}
	return c
}

func (q *Queue) Enq(v int) {
	n := &node{
		val: v,
	}

	q.tail.next = n
	q.tail = n
}

func (q *Queue) Deq() (int, bool) {
	if q.Length() == 0 {
		return 0, false
	}

	v := q.head.val
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return v, true
}

func (q *Queue) Peek() (int, bool) {
	if q.Length() == 0 {
		return 0, false
	}
	return q.head.val, true
}

// for deque just add push front and del back funcs
