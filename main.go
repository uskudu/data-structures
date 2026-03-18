package main

import (
	dll "awesomeDataStructures/linked-lists/doubly-linked-list"
	"fmt"
)

func main() {

	l := dll.NewLinkedList(1)

	fmt.Println(l)

	l.PushBack(2)
	fmt.Println(l)
	l.DelFront()
	fmt.Println(l)
}
