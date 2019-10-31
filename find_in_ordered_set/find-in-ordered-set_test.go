package main

import "testing"

func Test_findInOrderedSet(t *testing.T) {
	testSlice := []int{1, 4, 6, 7, 9, 17, 45}
	tests := []struct {
		name string
		target int
		want bool
	}{
		{
			name: "1",
			target: 1,
			want: true,
		},
		{
			name: "2",
			target: 4,
			want: true,
		},
		{
			name: "3",
			target: 6,
			want: true,
		},
		{
			name: "4",
			target: 7,
			want: true,
		},
		{
			name: "5",
			target: 9,
			want: true,
		},
		{
			name: "6",
			target: 17,
			want: true,
		},
		{
			name: "7",
			target: 45,
			want: true,
		},
		{
			name: "8",
			target: 5,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInOrderedSet(testSlice, tt.target); got != tt.want {
				t.Errorf("findInOrderedSet() = %v, want %v", got, tt.want)
			}
		})
	}
}