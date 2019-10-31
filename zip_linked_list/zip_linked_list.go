package zip_linked_list

/* Merge two linked lists so that their nodes alternate. Let the first node of the zipped list be
 * the first node of the first argument, if it exists.
 * Implement the linked list using only a Node class. No need for a wrapper LinkedList class
 *
 * BONUS: Do this in place
 */

import (
	. "hack_hour/data_structures"
)

func zip(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l1curr := l1.Head
	l2curr := l2.Head
	for l1curr.Next != nil && l2curr.Next != nil {
		l1front := l1curr.Next
		l1curr.Next = l2curr
		l2curr = l2curr.Next
		l1curr.Next.Next = l1front
		l1curr = l1front
	}

	if l1curr.Next == nil && l2curr.Next == nil {
		l1curr.Next = l2curr
		return l1
	}

	if l1curr.Next == nil && l2curr.Next != nil {
		l1curr.Next = l2curr
		return l1
	}

	l2curr.Next = l1curr.Next
	l1curr.Next = l2curr
	return l1
}
