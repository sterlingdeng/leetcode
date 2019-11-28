package merge_arrays

import (
	"reflect"
	"testing"
)

func Test_mergeArrays(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				x: []int{3, 4, 6, 10, 11, 15, 21},
				y: []int{1, 5, 8, 12, 14, 19},
			},
			want: []int{1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeArrays(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}