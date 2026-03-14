package static_array

import (
	"errors"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type StaticArray[T any] struct {
	data []T
}

func (s *StaticArray[T]) Len() int {
	return len(s.data)
}

func (s *StaticArray[T]) Cap() int {
	return cap(s.data)
}

func (s *StaticArray[T]) Get(idx int) (T, error) {
	if idx < 0 || idx > len(s.data)-1 {
		var zero T
		return zero, ErrIndexOutOfRange
	}
	return s.data[idx], nil
}

func (s *StaticArray[T]) Set(idx int, v T) error {
	if idx < 0 || idx > len(s.data)-1 {
		return ErrIndexOutOfRange
	}
	s.data[idx] = v
	return nil
}

func NewStaticArray[T any]() *StaticArray[T] {
	return &StaticArray[T]{}
}
