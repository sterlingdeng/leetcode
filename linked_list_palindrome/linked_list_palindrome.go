package linked_list_palindrome

import (
	. "hack_hour/data_structures"
	. "hack_hour/reverse_linked_list"
)

func isLinkedListPalindrome(l *Node) bool {
	if l == nil {
		return true
	}
	// get to the half way point first
	slow := l
	fast := l
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// now we are at the half way point
	head := l
	half := ReverseLinkedList(slow) // reverse the second half of the linked list

	for head != half && half != nil {
		if head.Value != half.Value {
			return false
		}
		head = head.Next
		half = half.Next
	}
	return true
}
