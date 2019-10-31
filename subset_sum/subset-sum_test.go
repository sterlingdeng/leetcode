package main

import "testing"

func Test_subsetSum(t *testing.T) {
	type args struct {
		arr []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				arr: []int{3, 7, 4, 2},
				target: 5,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				arr: []int{3, 34, 4, 12, 5, 12},
				target: 32,
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				arr: []int{8,2,4,12},
				target: 13,
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				arr: []int{8,-2,1,-3},
				target: 6,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubsetSum(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("subsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
