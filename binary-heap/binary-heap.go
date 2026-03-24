package binary_heap

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

func (h *minHeap) Decrease(i, newVal int) {
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
