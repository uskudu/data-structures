package stack

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
