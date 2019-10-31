package main

import "testing"

func Test_maxSubarray(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "test1",
			args: []int{1, -2, 3, 10, -4, 7, 2, -5},
			want: 18,
		},
		{
			name: "test2",
			args: []int{15, 20, -5, 10},
			want: 40,
		},
		{
			name: "test3",
			args: []int{15},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarray(tt.args); got != tt.want {
				t.Errorf("maxSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}