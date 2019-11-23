package nearly_sorted_array

import (
	"reflect"
	"testing"
)

func TestNearlySortedArray(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				arr: []int{6, 5, 3, 2, 8, 10, 9},
				k:   3,
			},
			want: []int{2, 3, 5, 6, 8, 9, 10},
		},
		{
			name: "2",
			args: args{
				arr: []int{10, 9, 8, 7, 4, 70, 60, 50},
				k:   4,
			},
			want: []int{4, 7, 8, 9, 10, 50, 60, 70},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NearlySortedArray(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NearlySortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
