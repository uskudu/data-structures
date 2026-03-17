package main

import (
	sll "awesomeDataStructures/linked-lists/singly-linked-list"
	"fmt"
)

func main() {
	n := &sll.Node{
		Value: 2,
		Next:  nil,
	}

	l := sll.NewLinkedList(n)

	l.PushFront(1)
	l.PushBack(3)

	fmt.Println(l.String())
	l.Push(, 4)
	fmt.Println(l.String())
}
