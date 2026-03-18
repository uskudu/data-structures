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
		n := &node{
			Value: tt.value,
		}
		l := NewLinkedList(n)

		got := l.String()

		if got != tt.want {
			t.Fatalf("expected %s, got %s", tt.want, got)
		}
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  string
	}{
		{"valid value (-1)", -1, "[{-1, 5}, {5, nil}]"},
		{"valid value (0)", 0, "[{0, 5}, {5, nil}]"},
		{"valid value (3)", 3, "[{3, 5}, {5, nil}]"},
	}

	for _, tt := range tests {
		n := &node{
			Value: 5,
		}
		l := NewLinkedList(n)
		l.PushFront(tt.value)

		got := l.String()

		if got != tt.want {
			t.Fatalf("expected %s, got %s", tt.want, got)
		}
	}

}
