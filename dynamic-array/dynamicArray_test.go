package dynamic_array

import "testing"

func TestDynamicArray_PushAndGet(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		idx   int
		want  int
		ok    bool
	}{
		{
			name:  "get first element",
			input: []int{1, 2, 3},
			idx:   0,
			want:  1,
			ok:    true,
		},
		{
			name:  "get middle element",
			input: []int{1, 2, 3},
			idx:   1,
			want:  2,
			ok:    true,
		},
		{
			name:  "index out of range",
			input: []int{1, 2, 3},
			idx:   5,
			want:  0,
			ok:    false,
		},
		{
			name:  "negative index",
			input: []int{1, 2, 3},
			idx:   -1,
			want:  0,
			ok:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray[int]()

			for _, v := range tt.input {
				arr.Push(v)
			}

			got, ok := arr.Get(tt.idx)

			if ok != tt.ok {
				t.Fatalf("expected ok=%v got %v", tt.ok, ok)
			}

			if got != tt.want {
				t.Fatalf("expected %v got %v", tt.want, got)
			}
		})
	}
}

func TestDynamicArray_Set(t *testing.T) {
	tests := []struct {
		name      string
		initial   []int
		idx       int
		value     int
		wantOk    bool
		wantValue int
	}{
		{
			name:      "set valid index",
			initial:   []int{1, 2, 3},
			idx:       1,
			value:     10,
			wantOk:    true,
			wantValue: 10,
		},
		{
			name:    "index out of range",
			initial: []int{1, 2, 3},
			idx:     5,
			value:   10,
			wantOk:  false,
		},
		{
			name:    "negative index",
			initial: []int{1, 2, 3},
			idx:     -1,
			value:   10,
			wantOk:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray[int]()

			for _, v := range tt.initial {
				arr.Push(v)
			}

			ok := arr.Set(tt.idx, tt.value)

			if ok != tt.wantOk {
				t.Fatalf("expected ok=%v got %v", tt.wantOk, ok)
			}

			if tt.wantOk {
				got, _ := arr.Get(tt.idx)
				if got != tt.wantValue {
					t.Fatalf("expected %v got %v", tt.wantValue, got)
				}
			}
		})
	}
}

func TestDynamicArray_Len(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "empty",
			input: []int{},
			want:  0,
		},
		{
			name:  "three elements",
			input: []int{1, 2, 3},
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray[int]()

			for _, v := range tt.input {
				arr.Push(v)
			}

			if arr.Len() != tt.want {
				t.Fatalf("expected %d got %d", tt.want, arr.Len())
			}
		})
	}
}

func TestDynamicArray_String(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  string
	}{
		{
			name:  "empty",
			input: []int{},
			want:  "[]",
		},
		{
			name:  "three elements",
			input: []int{1, 2, 3},
			want:  "[1, 2, 3]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := NewDynamicArray[int]()

			for _, v := range tt.input {
				arr.Push(v)
			}

			if arr.String() != tt.want {
				t.Fatalf("expected %s got %s", tt.want, arr.String())
			}
		})
	}
}
