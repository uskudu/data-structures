package binary_heap

import "math"

type minHeap struct {
	ar []int
}

func NewMinHeap() *minHeap {
	var ar []int
	return &minHeap{ar: ar}
}

func (h *minHeap) Parent(i int) int {
	return h.ar[(i-1)/2]
}
func (h *minHeap) Left(i int) int {
	return h.ar[i*2+1]
}
func (h *minHeap) Right(i int) int {
	return h.ar[i*2+2]
}

func (h *minHeap) GetMin() (int, bool) {
	if len(h.ar) == 0 {
		return 0, false
	}
	return h.ar[0], true
}

func (h *minHeap) Insert(v int) {
	h.ar = append(h.ar, v)

	for i := len(h.ar) - 1; i > 0; {
		p := (i - 1) / 2
		if h.ar[p] <= h.ar[i] {
			break
		}
		h.ar[i], h.ar[p] = h.ar[p], h.ar[i]
		i = p
	}
}

func (h *minHeap) DecreaseKey(i, newVal int) {
	if newVal > h.ar[i] {
		panic("new value is greater than current value")
	}

	h.ar[i] = newVal

	for i != 0 {
		p := (i - 1) / 2
		if h.ar[p] <= h.ar[i] {
			break
		}
		h.ar[i], h.ar[p] = h.ar[p], h.ar[i]
		i = p
	}
}

func (h *minHeap) ExtractMin() (int, bool) {
	if len(h.ar) == 0 {
		return 0, false
	}
	if len(h.ar) == 1 {
		x := h.ar[0]
		clear(h.ar)
		return x, true
	}
	x := h.ar[0]

	h.ar[0] = h.ar[len(h.ar)-1]
	h.ar = h.ar[:len(h.ar)-1]
	h.MinHeapify(0)
	return x, true
}

func (h *minHeap) DeleteKey(i int) bool {
	h.DecreaseKey(i, int(math.Inf(-1)))
	_, ok := h.ExtractMin()
	return ok
}

func (h *minHeap) MinHeapify(i int) {
	l, r, n := h.Left(i), h.Right(i), len(h.ar)
	smallest := i

	if l < n && h.ar[l] < h.ar[smallest] {
		smallest = l
	}
	if r < n && h.ar[r] < h.ar[smallest] {
		smallest = l
	}

	if i != smallest {
		h.ar[i], h.ar[smallest] = h.ar[smallest], h.ar[i]
		h.MinHeapify(smallest)
	}
}
