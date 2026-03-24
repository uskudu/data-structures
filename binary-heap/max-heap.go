package binary_heap

import "math"

type maxHeap struct {
	ar []int
}

func NewMaxHeap() *maxHeap {
	var ar []int
	return &maxHeap{ar: ar}
}

func (h *maxHeap) Parent(i int) int {
	return h.ar[(i-1)/2]
}
func (h *maxHeap) Left(i int) int {
	return h.ar[i*2+1]
}
func (h *maxHeap) Right(i int) int {
	return h.ar[i*2+2]
}

func (h *maxHeap) GetMax() (int, bool) {
	if len(h.ar) == 0 {
		return 0, false
	}
	return h.ar[0], true
}

func (h *maxHeap) Insert(v int) {
	h.ar = append(h.ar, v)

	for i := len(h.ar) - 1; i > 0; {
		p := (i - 1) / 2
		if h.ar[p] >= h.ar[i] {
			break
		}
		h.ar[i], h.ar[p] = h.ar[p], h.ar[i]
		i = p
	}
}
