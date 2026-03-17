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

	fmt.Println(l.Head)

	l.PushFront(1)
	fmt.Println(l.Head)

	l.PushBack(3)
	fmt.Println(l.Tail)

	fmt.Println(l.String())
}
