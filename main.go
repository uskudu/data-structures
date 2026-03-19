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
	l.Push(23)
	fmt.Println(l)

}
