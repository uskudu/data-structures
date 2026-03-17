package singly_linked_list

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
