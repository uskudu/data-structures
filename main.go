package main

import (
	dll "awesomeDataStructures/linked-lists/doubly-linked-list"
	"fmt"
)

func main() {

	l := dll.NewLinkedList(1)

	fmt.Println(l)

	l.PushFront(2)
	l.PushFront(3)
	fmt.Println(l)
}
