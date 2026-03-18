package singly_linked_list

import (
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"valid value (-10)", -10, "[{-10, nil}]"},
		{"valid value (0)", 0, "[{0, nil}]"},
		{"valid value (3)", 3, "[{3, nil}]"},
	}

	for _, tt := range tests {
		n := &Node{
			Value: tt.value,
		}
		l := NewLinkedList(n)

		got := l.String()

		if got != tt.want {
			t.Fatalf("expected %s, got %s", tt.want, got)
		}

	}

}
