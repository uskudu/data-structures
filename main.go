package main

import (
	bt "awesomeDataStructures/binary-tree"
	"fmt"
)

func main() {

	t := bt.NewBinaryTree(2)

	t.Insert(1)
	t.Insert(3)
	t.Insert(776)
	t.Insert(6)
	t.Insert(24)
	t.Insert(8)
	t.Insert(8)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(-1)
	t.Insert(30)
	t.Insert(340)
	t.Insert(55550)
	t.Insert(344)
	fmt.Println(t.Sliced())
}
