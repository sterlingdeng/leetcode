package main

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "test1",
			args: []int{1, 3, 4},
			want: 2,
		},
		{
			name: "test2",
			args: []int{1, 3},
			want: 2,
		},
		{
			name: "test3",
			args: []int{1, 3, 4, 5, 6, 7},
			want: 2,
		},
		{
			name: "test4",
			args: []int{4, 3, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}