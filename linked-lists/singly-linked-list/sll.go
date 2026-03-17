package singly_linked_list

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head Node
}

func NewLinkedList(head Node) *LinkedList {
	return &LinkedList{Head: head}
}

func (l *LinkedList) PushFront(val int) {
	n := Node{Value: val}

	prev := l.Head
	n.Next = &prev
	l.Head = n
}
