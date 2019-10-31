package main

import (
	. "hack_hour/data_structures"
	"reflect"
	"testing"
)

func Test_kthToLastNode(t *testing.T) {
	n := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next: &Node{
					Value: 4,
					Next: &Node{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	}
	type args struct {
		k    int
		node *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "0th to the last node",
			args: args{
				k:    0,
				node: n,
			},
			want: nil,
		},
		{
			name: "last node",
			args: args{1, n},
			want: n.Next.Next.Next.Next,
		},
		{
			name: "2nd to the last node",
			args: args{2, n},
			want: n.Next.Next.Next,
		},
		{
			name: "out of bounds",
			args: args{99, n},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthToLastNode(tt.args.k, tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kthToLastNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
