package main

import (
	sll "awesomeDataStructures/linked-lists/singly-linked-list"
	"fmt"
)

func main() {
	n := sll.NewNode(2)

	l := sll.NewLinkedList(n)

	fmt.Println(l)

	l.PushFront(1)
	l.PushBack(3)
	l.PushBack(4)

	fmt.Println(l.String())

	//l := sll.NewLinkedList()
}
