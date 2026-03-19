package main

import (
	"awesomeDataStructures/stack"
	"fmt"
)

func main() {

	l, _ := stack.NewStack(2)

	fmt.Println(l)

	l.Push(1)
	l.Push(4)
	fmt.Println(l)
	fmt.Println(l.Pop())
	fmt.Println(l)
	fmt.Println(l.GetTop())
}
