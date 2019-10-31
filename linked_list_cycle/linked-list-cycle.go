package main

import (
	. "hack_hour/data_structures"
)
/**
 * We are familiar with linked lists being linear and terminating:
 *
 * A->B->C->D
 *
 * However, linked lists can also have cyclical references:
 *
 * A->B->C->D
 *    ^     |
 *    |     V
 *    G<-F<-E
 *
 * Create a function that accepts a linked list and returns true if the linked list has a cyclical reference
 *
 * var node1 = new Node('1');
 * var node2 = node1.next = new Node('2');
 * var node3 = node2.next = new Node('3');
 * var node4 = node3.next = new Node('4');
 * var node5 = node4.next = new Node('5');
 * hasCycle(node1); // => false
 * node5.next = node2;
 * hasCycle(node1); // => true
 *
 * Challenge 1: Do this in linear time
 * Challenge 2: Do this in constant space
 * Challenge 3: Do not mutate the original nodes in any way // can't mutate the original node because of static typing :)
 *
 */

func hasCycle(node *Node) bool {
	slowPtr := node
	fastPtr := node

	for fastPtr != nil && fastPtr.Next != nil {
		if slowPtr == fastPtr {
			return true
		}
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	return false
}

func hasCycleSpaceUnoptimized(node *Node) bool {
	seen := make(map[*Node]bool)  // create map to store what
	curr := node
	for curr != nil {
		if _, exists := seen[curr]; exists {
			return true
		}
		seen[curr] = true
		curr = curr.Next
	}
	return false
}