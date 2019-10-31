package main

import (
	"reflect"
	"testing"
)

func Test_merge_sort(t *testing.T) {
	tests := []struct {
		name string
		in []int
		want []int
	}{
		{
			name: "1",
			in: []int{5, 3, 8, 20, 19, 1, 2},
			want: []int{1, 2, 3, 5, 8, 19, 20},
		},
		{
			name: "2",
			in: []int{3},
			want: []int{3},
		},
		{
			name: "3",
			in: []int{-1, -10},
			want: []int{-10, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge_sort(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge_sort() = %v, want %v", got, tt.want)
			}
		})
	}
}