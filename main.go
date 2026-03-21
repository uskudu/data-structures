package main

import (
	q "awesomeDataStructures/queue"
	"fmt"
)

func main() {

	q := q.NewQueue(1)

	fmt.Println(q)

	q.Enq(3)
	q.Enq(3)
	q.Enq(3)
	fmt.Println(q)
	p, _ := q.Deq()
	fmt.Println(p)
	fmt.Println(q)
}
