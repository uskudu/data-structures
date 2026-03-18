package main

import (
	sll "awesomeDataStructures/linked-lists/singly-linked-list"
	"fmt"
)

func main() {
	n := &sll.Node{
		Value: 1,
	}

	l := sll.NewLinkedList(n)
	l.PushFront(1)
	l.PushFront(1)
	l.PushFront(1)
	l.PushFront(1)
	l.PushFront(1)
	l.PushFront(1)
	l.PushFront(1)
	l.PushFront(1)
	fmt.Println(l)
	fmt.Println(l.Len())
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
