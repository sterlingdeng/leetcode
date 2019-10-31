package reverse_linked_list

import (
	. "hack_hour/data_structures"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	node := &Node{
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
	}
	tests := []struct {
		name string
		in   *Node
		want []int
	}{
		{
			name: "test1",
			in:   node,
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseLinkedList(tt.in)
			for _, val := range tt.want {
				if got.Value != val {
					t.Errorf("ReverseLinkedList() = %v, want %v", got.Value, val)
				}
				got = got.Next
			}
		})
	}
}
