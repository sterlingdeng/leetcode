package each_permutation

import (
	"reflect"
	"testing"
)

func TestEachPermutation(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				slice: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EachPermutation(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EachPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
