package main

import (
	. "hack_hour/data_structures"
)

/**
 * Write a function deleteDups that removes duplicates from an unsorted linked list.
 *
 * Example:
 * 1->2->3->3
 *
 * deleteDups(head); -> 1->2->3
 *
 * Extra:
 * How would you solve this problem if a temporary buffer is not allowed?
 */

func deleteDups(head *Node) {
	// use a map to see if this node has been seen before
	seen := make(map[int]bool)
	var back *Node
	var curr = head
	for curr != nil {
		if _, exists := seen[curr.Value]; exists {
			back.Next = curr.Next
		}
		seen[curr.Value] = true
		back = curr
		curr = curr.Next
	}
}
