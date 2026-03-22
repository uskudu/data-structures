package main

import (
	bt "awesomeDataStructures/binary-tree"
	"fmt"
)

func main() {

	t := bt.NewBinaryTree(2)

	t.Insert(1)
	t.Insert(3)
	fmt.Println(t)
}
