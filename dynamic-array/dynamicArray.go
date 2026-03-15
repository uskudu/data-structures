package dynamic_array

import (
	"fmt"
	"strings"
)

type DynamicArray[T any] struct {
	data           []T
	EvacuatedCount int
}

func (d *DynamicArray[T]) String() string {
	var b strings.Builder
	b.WriteByte('[')
	for i, v := range d.data {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprint(&b, v)
	}
	b.WriteByte(']')
	return b.String()
}

func (d *DynamicArray[T]) Len() int {
	return len(d.data)
}

func (d *DynamicArray[T]) Cap() int {
	return cap(d.data)
}

func (d *DynamicArray[T]) Get(idx int) (T, bool) {
	if idx < 0 || idx >= len(d.data) {
		var zero T
		return zero, false
	}
	return d.data[idx], true
}

func (d *DynamicArray[T]) evacuate() {
	newCap := 1
	if cap(d.data) > 256 {
		newCap = cap(d.data) * 2
	}
	if cap(d.data) > 64 {
		newCap = cap(d.data) * 3
	}
	if cap(d.data) > 0 {
		newCap = cap(d.data) * 4
	}
	newSlice := make([]T, len(d.data), newCap)
	copy(newSlice, d.data)
	d.data = newSlice

	d.EvacuatedCount++
	fmt.Printf("evacuation #%d\n", d.EvacuatedCount)
}

func (d *DynamicArray[T]) Set(idx int, v T) bool {
	if idx < 0 || idx > len(d.data)-1 {
		return false
	}

	d.data[idx] = v
	return true
}

func (d *DynamicArray[T]) Push(v T) {
	if len(d.data) == cap(d.data) {
		d.evacuate()
	}
	d.data = append(d.data, v)
}

func NewDynamicArray[T any]() *DynamicArray[T] {
	var d []T
	return &DynamicArray[T]{
		data: d,
	}
}
