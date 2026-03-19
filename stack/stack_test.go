package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	tests := []struct {
		name      string
		capacity  int
		wantOk    bool
		wantArCap int
	}{
		{"invalid Stack with negative capacity", -10, false, -1},
		{"invalid Stack with 0 capacity", 0, false, -1},
		{"valid Stack with capacity 1", 1, true, 1},
		{"valid Stack with capacity 12", 12, true, 12},
	}
	_ = tests

}
