package singly_linked_list

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head Node
	Tail Node
}

func NewLinkedList(head Node) *LinkedList {
	return &LinkedList{Head: head}
}

func (l *LinkedList) PushFront(val int) {
	n := &Node{Value: val}

	n.Next = &l.Head
	l.Head = *n

	if &l.Tail == nil {
		l.Tail = *n
	}
}

func (l *LinkedList) PushBack(val int) {
	n := &Node{Value: val}

	if &l.Tail == nil {
		l.Head = *n
		l.Tail = *n
		return
	}
	l.Tail.Next = n
	l.Tail = *n
}
