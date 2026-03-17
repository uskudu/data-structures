package main

import (
	sll "awesomeDataStructures/linked-lists/singly-linked-list"
	"fmt"
)

func main() {
	n := sll.Node{
		Value: 1,
		Next:  nil,
	}

	l := sll.NewLinkedList(n)

	fmt.Println(l.Head)

	l.PushFront(2)
	fmt.Println(l.Head)

}
