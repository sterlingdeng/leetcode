package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "test1",
			args: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "test2",
			args: []int{2, 0, 2, 1, 0, 1, 0, 2, 1, 0},
			want: []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			name: "test3",
			args: []int{1},
			want: []int{1},
		},
		{
			name: "test4",
			args: []int{},
			want: []int{},
		},
		{
			name: "test4",
			args: []int{2, 1, 0},
			want: []int{0, 1, 2},
		},
		{
			name: "test5",
			args: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("Got sortColors() = %v. Want: %v", tt.args, tt.want)
			}
		})
	}
}
