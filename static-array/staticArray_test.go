package static_array

import "testing"

func TestNewStaticArray(t *testing.T) {
	var size = 3
	arr := NewStaticArray[int](size)
	tests := []struct {
		name   string
		index  int
		value  int
		wantOk bool
	}{
		{"valid index 0", 0, 10, true},
		{"valid index 1", 1, 20, true},
		{"valid index 2", 2, 30, true},
		{"negative index", -1, 0, false},
		{"index equal len", size, 0, false},
		{"index > len", size + 1, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok := arr.Set(tt.index, tt.value)
			if ok != tt.wantOk {
				t.Fatalf("Set() ok = %v, wantOk = %v", ok, tt.wantOk)
			}

			v, ok := arr.Get(tt.index)
			if tt.wantOk {
				if !ok {
					t.Fatalf("Get() ok = false; expected true")
				}
				if v != tt.value {
					t.Errorf("Get() = %d; want %d", v, tt.value)
				}
			} else {
				if ok {
					t.Fatalf("Get() ok = true; expected false")
				}
			}
		})
	}

	if arr.Len() != size {
		t.Errorf("Len() = %d; want %d", arr.Len(), size)
	}
	if arr.Cap() != size {
		t.Errorf("Cap() = %d; want %d", arr.Cap(), size)
	}
}

func TestStaticArray_Len(t *testing.T) {
	const size = 5
	arr := NewStaticArray[int](size)

	if got := arr.Len(); got != size {
		t.Errorf("Len() = %d; want %d", got, size)
	}

	ok := arr.Set(0, 42)
	if !ok {
		t.Fatalf("Set() failed for index 0")
	}
	if got := arr.Len(); got != size {
		t.Errorf("Len() after Set = %d; want %d", got, size)
	}
}
