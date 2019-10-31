package main

import (
	"reflect"
	"testing"
)

func Test_mergeRanges(t *testing.T) {
	type args struct {
		in [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				in: [][]int{
					{
						0, 1,
					},
					{
						3, 5,
					},
					{
						4, 8,
					},
					{
						9, 10,
					},
					{
						10, 12,
					},
					{
						11, 14,
					},
				},
			},
			want: [][]int{
				{
					0, 1,
				},
				{
					3, 8,
				},
				{
					9, 14,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeRanges(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
