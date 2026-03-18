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
	return &LinkedList{Head: head, Tail: head}
}
