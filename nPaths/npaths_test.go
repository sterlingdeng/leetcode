package main

import "testing"

func Test_npaths(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"1", 1, 1},
		{"2", 2, 2},
		{"3", 3, 12},
		{"4", 4, 184},
		{"5", 5, 8512},
		{"6", 6, 1262816},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := npaths(tt.size); got != tt.want {
				t.Errorf("npaths() = %v, want %v", got, tt.want)
			}
		})
	}
}