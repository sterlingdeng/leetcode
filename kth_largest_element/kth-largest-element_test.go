package kth_largest_element

import "testing"

func TestKthLargestElement(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{7, 6, 5, 4, 3, 2, 1},
				k:    3,
			},
			want: 5,
		},
		{
			name: "smaller edge of constraint",
			args: args{
				nums: []int{7},
				k:    1,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthLargestElement(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("KthLargestElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
