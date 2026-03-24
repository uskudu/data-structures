package binary_heap

type minHeap struct {
	ar []int
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
