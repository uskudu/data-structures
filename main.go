package main

import (
	cll "awesomeDataStructures/linked-lists/circular-linked-list"
	"fmt"
)

func main() {

	l := cll.NewLinkedList(1)

	fmt.Println(l)

	l.PushBack(2)
	l.PushBack(3)
	fmt.Println(l)
	//l.DelFront()
	//fmt.Println(l)
}
