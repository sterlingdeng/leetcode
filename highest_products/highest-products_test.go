package main

import "testing"

func Test_highestProducts(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "1",
			args: []int{1, 2, 3},
			want: 6,
		},
		{
			name: "2",
			args: []int{1, 2, 3, 10, 10, 11},
			want: 1100,
		},
		{
			name: "3",
			args: []int{-4, -2, 1, 2, 100},
			want: 800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestProductOptimal(tt.args); got != tt.want {
				t.Errorf("highestProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}