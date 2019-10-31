package reverse_linked_list

import (
	"fmt"
	. "hack_hour/data_structures"
)

func ReverseLinkedList(n *Node) *Node {
	if n.Next == nil {
		return n
	}
	var a *Node
	b := n
	c := n.Next
	for c != nil {
		b.Next = a
		a = b
		b = c
		c = c.Next
	}
	b.Next = a
	return b
}

func main() {
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
	node = ReverseLinkedList(node)
	fmt.Println(node.Value, node.Next.Value, node.Next.Next.Value, node.Next.Next.Next.Value)

}
