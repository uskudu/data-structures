package main

import (
	bh "awesomeDataStructures/binary-heap"
	"fmt"
)

func main() {
	q := bh.NewMinHeap()

	q.Insert(1)
	q.Insert(9)
	q.Insert(2)
	q.Insert(8)
	q.Insert(3)
	q.Insert(7)
	q.Insert(4)
	q.Insert(6)
	q.Insert(5)

	fmt.Println(q)
}
