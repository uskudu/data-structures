package stack

import (
	"strconv"
	"strings"
)

type Stack struct {
	Capacity int
	Ar       []int
}

func NewStack(capacity int) (*Stack, bool) {
	if capacity < 1 {
		return nil, false
	}
	ar := make([]int, 0, capacity)
	return &Stack{
		Capacity: capacity,
		Ar:       ar,
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
	b.WriteString(strconv.Itoa(len(s.Ar) - 1))

	return b.String()
}

func (s *Stack) Push(elem int) bool {
	if len(s.Ar) == s.Capacity {
		return false
	}

	s.Ar = append(s.Ar, elem)
	return true
}

func (s *Stack) Pop() (int, bool) {
	if len(s.Ar) == 0 {
		return -1, false
	}
	val := s.Ar[len(s.Ar)-1]
	s.Ar = s.Ar[:len(s.Ar)-1]
	return val, true
}

func (s *Stack) GetTop() (int, bool) {
	if len(s.Ar) == 0 {
		return -1, false
	}
	return s.Ar[len(s.Ar)-1], true
}

func (s *Stack) IsEmpty() bool {
	if len(s.Ar) == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	if len(s.Ar) == 0 {
		return 0
	}
	return len(s.Ar)
}
