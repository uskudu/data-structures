package static_array

import (
	"fmt"
	"strings"
)

type StaticArray[T any] struct {
	data []T
}

func (s *StaticArray[T]) String() string {
	var b strings.Builder
	b.WriteByte('[')
	for i, v := range s.data {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprint(&b, v)
	}
	b.WriteByte(']')
	return b.String()
}

func (s *StaticArray[T]) Len() int {
	return len(s.data)
}

func (s *StaticArray[T]) Cap() int {
	return cap(s.data)
}

func (s *StaticArray[T]) Get(idx int) (T, bool) {
	if idx < 0 || idx >= len(s.data) {
		var zero T
		return zero, false
	}
	return s.data[idx], true
}

func (s *StaticArray[T]) Set(idx int, v T) bool {
	if idx < 0 || idx > len(s.data)-1 {
		return false
	}
	s.data[idx] = v
	return true
}

func NewStaticArray[T any](size int) *StaticArray[T] {
	return &StaticArray[T]{
		data: make([]T, size),
	}
}
