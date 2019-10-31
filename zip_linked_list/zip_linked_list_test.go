package zip_linked_list

import (
	. "hack_hour/data_structures"
	"testing"
)

func Test_zip(t *testing.T) {
	l1 := &LinkedList{Head: &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next: &Node{
					Value: 4,
					Next:  nil,
				},
			},
		},
	}}
	l2 := &LinkedList{Head: &Node{
		Value: 5,
		Next: &Node{
			Value: 6,
			Next: &Node{
				Value: 7,
				Next: &Node{
					Value: 8,
					Next: &Node{
						Value: 9,
						Next: &Node{
							Value: 10,
							Next:  nil,
						},
					},
				},
			},
		},
	}}

	l3 := &LinkedList{Head: &Node{
		Value: 1,
		Next:  nil,
	}}
	l4 := &LinkedList{Head: &Node{
		Value: 2,
		Next:  nil,
	}}

	type args struct {
		l1 *LinkedList
		l2 *LinkedList
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{l1, l2},
			want: []int{1, 5, 2, 6, 3, 7, 4, 8, 9, 10},
		},
		{
			name: "test2",
			args: args{l3, l4},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := zip(tt.args.l1, tt.args.l2)
			i := 0
			node := got.Head
			for node != nil {
				if node.Value != tt.want[i] {
					t.Errorf("value = %v, want %v", node.Value, tt.want[i])
				}
				node = node.Next
				i++
			}
		})
	}
}
