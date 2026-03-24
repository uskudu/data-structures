package main

import (
	bh "awesomeDataStructures/binary-heap"
	"fmt"
)

func main() {
	q := bh.NewMinHeap()

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
	m, ok := q.ExtractMin()
	fmt.Println(m, ok)
	fmt.Println(q)
}
