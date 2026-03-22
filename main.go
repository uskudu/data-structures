package main

import (
	q "awesomeDataStructures/queue"
	"fmt"
)

func main() {

	q := q.NewQueue(1)

	fmt.Println(q)

	q.Enq(2)
	q.Enq(3)
	fmt.Println(q)
	q.Peek()
	fmt.Println(q.Peek())
	fmt.Println(q)
}
