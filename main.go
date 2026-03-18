package main

import (
	sll "awesomeDataStructures/linked-lists/singly-linked-list"
	"fmt"
)

func main() {
	n := sll.NewNode(1)

	l := sll.NewLinkedList(n)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

}
