package main

import (
	sll "awesomeDataStructures/linked-lists/singly-linked-list"
	"fmt"
)

func main() {

	l := sll.NewLinkedList(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	fmt.Println(l)
	l.Reverse()
	fmt.Println(l)
	fmt.Println(l.Head)
	fmt.Println(l.Tail)
}
