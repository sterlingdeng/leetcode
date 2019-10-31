package main

import (
	"reflect"
	"testing"
)

func Test_getAllProducts(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "1",
			in: []int{1},
			want: []int{1},
		},
		{
			name: "2",
			in: []int{4, 5},
			want: []int{5, 4},
		},
		{
			name: "3",
			in: []int{1, 7, 3, 4},
			want: []int{84, 12, 28, 21},
		},
		{
			name: "4",
			in: []int{},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllProducts(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
