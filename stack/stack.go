package stack

import (
	"strconv"
	"strings"
)

type Stack struct {
	Capacity int
	Ar       *[]int
	Top      int
}

func NewStack(capacity int) (*Stack, bool) {
	if capacity < 1 {
		return nil, false
	}
	ar := make([]int, 0, capacity)
	return &Stack{
		Capacity: capacity,
		Ar:       &ar,
		Top:      -1,
	}, true
}

func (s *Stack) String() string {
	b := strings.Builder{}

	b.WriteString(strconv.Itoa(s.Capacity))
	b.WriteString(" [")

	for _, i := range *s.Ar {
		b.WriteString(strconv.Itoa(i))
	}

	b.WriteString("] ")
	b.WriteString(strconv.Itoa(s.Top))

	return b.String()
}
