package main

import "testing"
import . "hack_hour/data_structures"

func Test_hasCycle(t *testing.T) {

	a := &Node{
		Value: 1,
		Next:  nil,
	}
	b := &Node{
		Value: 2,
		Next:  nil,
	}
	c := &Node{
		Value: 3,
		Next:  nil,
	}
	d := &Node{
		Value: 4,
		Next:  nil,
	}
	e := &Node{
		Value: 5,
		Next:  nil,
	}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = a


	f := &Node{
		Value: 1,
		Next: nil,
	}

	g := &Node{
		Value: 1,
		Next: nil,
	}
	h := &Node{
		Value: 2,
		Next:  nil,
	}

	g.Next = h
	h.Next = g


	tests := []struct {
		name string
		args *Node
		want bool
	}{
		{
			name: "has cycle",
			args: a,
			want: true,
		},
		{
			name: "only 1 node",
			args: f,
			want: false,
		},
		{
			name: "circular with 2 nodes",
			args: g,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}