package main

import (
	sll "awesomeDataStructures/linked-lists/singly-linked-list"
	"fmt"
)

func main() {
	n := &sll.Node{
		Value: 9,
	}

	l := sll.NewLinkedList(n)
	l.PushFront(8)
	l.PushFront(7)
	l.PushFront(6)
	l.PushFront(5)
	l.PushFront(4)
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)
	fmt.Println(l)
	fmt.Println(l.GetNodeByIndex(9))

	//
	//l.PushFront(1)
	//l.PushBack(3)
	//l.PushBack(4)
	//
	//fmt.Println(l.String())
	//
	//l.Del(4)
	//fmt.Println(l.String())
}
