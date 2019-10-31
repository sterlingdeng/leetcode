package main

import (
	"fmt"
	"testing"
)
import . "hack_hour/data_structures"

func Test_deleteDups(t *testing.T) {
	node1 := &Node{
		Value: 1,
		Next: &Node{
			Value: 2,
			Next: &Node{
				Value: 3,
				Next: &Node{
					Value: 3,
					Next:  nil,
				},
			},
		},
	}
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				head: node1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteDups(tt.args.head)
			fmt.Printf("%+v\n", tt.args.head)
			fmt.Printf("%+v\n", tt.args.head.Next)
			fmt.Printf("%+v\n", tt.args.head.Next.Next)
		})
	}
}
