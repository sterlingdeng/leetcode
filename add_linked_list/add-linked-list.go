package add_linked_list

/* You have two numbers represented by linked lists. Each node contains a single digit. The digits
 * are stored in reverse order, such that the 1's digit is at the head of the list. Write
 * a function that adds the two numbers and returns the sum as a linked list with the same
 * structure.
 * EXAMPLE
 * Input: (2 -> 1 -> 5) + (5 -> 9 -> 2)
 * Output: 7 -> 0 -> 8
 *
 * The 7 is the ones digit (2 + 5).
 * The 0 is the tens digit (1 + 9, carry the 1).
 * The 8 is the hundreds digit (1 carried over + 5 + 2).
 *
 */

type node struct {
	value int
	next  *node
}

func addLinkedList(l1, l2 *node) *node {
	var trailingl1 *node
	head := l1
	carry := 0

	for l1 != nil && l2 != nil {
		sum := l1.value + l2.value + carry
		if sum < 10 {
			l1.value = sum
		} else {
			l1.value = sum % 10
			carry = (sum - (sum % 10)) / 10
		}
		trailingl1 = l1
		l1 = l1.next
		l2 = l2.next
	}

	if carry != 0 {
		if l1 == nil && l2 == nil {
			trailingl1 = &node{
				value: carry,
				next:  nil,
			}
		} else if l1 != nil {
			l1.value += carry
		} else if l2 != nil {
			l2.value += carry
		}
	}

	if l2 == nil {
		return head
	}

	for l2 != nil {
		l1 = trailingl1
		l1.next = l2
		l1 = l1.next
		l2 = l2.next
	}

	return head
}
