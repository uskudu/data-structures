package main

import (
	cll "awesomeDataStructures/linked-lists/circular-linked-list"
	"fmt"
)

func main() {

	l := cll.NewLinkedList(1)

	fmt.Println(l)

	l.PushFront(2)
	l.PushFront(3)
	fmt.Println(l)
	//l.DelFront()
	//fmt.Println(l)
}
