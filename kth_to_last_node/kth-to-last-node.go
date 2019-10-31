package main

import (
	. "hack_hour/data_structures"
)

/**
 * Write a function that takes two parameters, an integer and the head of a
 * singly linked list, and returns the VALUE on the kth to last node in the list.
 *
 * const a = new Node('A');
 * const b = new Node('B');
 * const c = new Node('C');
 * const d = new Node('D');
 * const e = new Node('E');
 *
 * a.next = b;
 * b.next = c;
 * c.next = d;
 * d.next = e;
 *
 * kthToLastNode(2, a); -> returns 'D' (the value on the second to last node)
 */

func kthToLastNode(k int, node *Node) *Node {
	var backPtr *Node
	curr := node
	for i := 1; i < k+1; i++ {
		if curr == nil {
			return nil
		}
		curr = curr.Next
	}
	backPtr = node
	for curr != nil {
		backPtr = backPtr.Next
		curr = curr.Next
	}
	return backPtr
}
