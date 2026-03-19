package stack

import (
	"strconv"
	"strings"
)

type Stack struct {
	Capacity int
	Ar       []int
	Top      int
}

func NewStack(capacity int) (*Stack, bool) {
	if capacity < 1 {
		return nil, false
	}
	ar := make([]int, 0, capacity)
	return &Stack{
		Capacity: capacity,
		Ar:       ar,
		Top:      -1,
	}, true
}

func (s *Stack) String() string {
	b := strings.Builder{}

	b.WriteString(strconv.Itoa(s.Capacity))
	b.WriteString(" [")

	for i, elem := range s.Ar {
		if i != 0 {
			b.WriteString(" ")
		}
		b.WriteString(strconv.Itoa(elem))
	}

	b.WriteString("] ")
	b.WriteString(strconv.Itoa(s.Top))

	return b.String()
}

func (s *Stack) Push(elem int) bool {
	if len(s.Ar) == s.Capacity {
		return false
	}

	s.Ar = append(s.Ar, elem)
	s.Top++
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.Top == -1 {
		return -1, false
	}
	val := s.Ar[s.Top]
	s.Ar = s.Ar[:s.Top]
	s.Top--
	return val, true
}

func (s *Stack) GetTop() (int, bool) {
	if s.Top == -1 {
		return -1, false
	}
	return s.Ar[s.Top], true
}

func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}
