package main

import (
	bh "awesomeDataStructures/binary-heap"
	"fmt"
)

func main() {
	q := bh.NewMaxHeap()

	q.Insert(40)
	q.Insert(43)
	q.Insert(86)
	q.Insert(88)
	q.Insert(90)
	q.Insert(73)
	q.Insert(25)
	q.Insert(57)
	q.Insert(19)

	fmt.Println(q)
}
