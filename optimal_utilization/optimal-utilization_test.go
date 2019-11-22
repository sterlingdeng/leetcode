package optimal_utilization

import (
	"reflect"
	"testing"
)

func Test_optimalUtilizationBruteForce(t *testing.T) {
	type args struct {
		a      [][]int
		b      [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				a: [][]int{
					{1, 2},
					{2, 4},
					{3, 6},
				},
				b: [][]int{
					{1, 2},
				},
				target: 7,
			},
			want: [][]int{
				{2, 1},
			},
		},
		{
			name: "test2",
			args: args{
				a: [][]int{
					{1, 3},
					{2, 5},
					{3, 7},
					{4, 10},
				},
				b: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 5},
				},
				target: 10,
			},
			want: [][]int{
				{2, 4},
				{3, 2},
			},
		},
		{
			name: "test3",
			args: args{
				a: [][]int{
					{1, 8},
					{2, 7},
					{3, 14},
				},
				b: [][]int{
					{1, 5},
					{2, 10},
					{3, 14},
				},
				target: 20,
			},
			want: [][]int{
				{3, 1},
			},
		},
		{
			name: "test4",
			args: args{
				a: [][]int{
					{1, 8},
					{2, 15},
					{3, 9},
				},
				b: [][]int{
					{1, 8},
					{2, 11},
					{3, 12},
				},
				target: 20,
			},
			want: [][]int{
				{1, 3},
				{3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalUtilizationBruteForce(tt.args.a, tt.args.b, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optimalUtilizationBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_optimalUtilization(t *testing.T) {
	type args struct {
		a      [][]int
		b      [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				a: [][]int{
					{1, 2},
					{2, 4},
					{3, 6},
				},
				b: [][]int{
					{1, 2},
				},
				target: 7,
			},
			want: [][]int{
				{2, 1},
			},
		},
		{
			name: "test2",
			args: args{
				a: [][]int{
					{1, 3},
					{2, 5},
					{3, 7},
					{4, 10},
				},
				b: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
					{4, 5},
				},
				target: 10,
			},
			want: [][]int{
				{2, 4},
				{3, 2},
			},
		},
		{
			name: "test3",
			args: args{
				a: [][]int{
					{1, 8},
					{2, 7},
					{3, 14},
				},
				b: [][]int{
					{1, 5},
					{2, 10},
					{3, 14},
				},
				target: 20,
			},
			want: [][]int{
				{3, 1},
			},
		},
		{
			name: "test4",
			args: args{
				a: [][]int{
					{1, 8},
					{2, 15},
					{3, 9},
				},
				b: [][]int{
					{1, 8},
					{2, 11},
					{3, 12},
				},
				target: 20,
			},
			want: [][]int{
				{1, 3},
				{3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalUtilization(tt.args.a, tt.args.b, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optimalUtilizationBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
