package main

/* You are given a tree. Write a function to check if it is a valid binary search tree. A tree is
 * a valid binary search tree if it satisfies the following constraints:
 *      at any given node, the value of all the nodes in its left tree must be < its value
 *      at any given node, the value of all the nodes in its right tree must be > its value
 * Assume that each value in the tree is unique.
 */

/*
							 15
						/			\
					10			20
				/		\		 /
			4			11  3
			 \		  \
				9		 22
				 \
					11
 */

import (
	"fmt"
	. "hack_hour/data_structures"
	"math"
)

func validBst(n *BinarySearchTree) bool {
	return validBstHelper(n, math.MinInt64, math.MaxInt64)
}

func validBstHelper(n *BinarySearchTree, min, max int) bool {
	if n == nil {
		return true
	}
	if n.Value < min || n.Value > max {
		return false
	}
	if !validBstHelper(n.Left, min, n.Value) {
		return false
	}
	if !validBstHelper(n.Right, n.Value, max) {
		return false
	}
	return true
}

func main() {
	entry := &BinarySearchTree{
		Value: 15,
		Left:  &BinarySearchTree{
			Value: 10,
			Left:  &BinarySearchTree{
				Value: 4,
				Left:  nil,
				Right: &BinarySearchTree{
					Value: 9,
					Left:  nil,
					Right: &BinarySearchTree{
						Value: 11,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &BinarySearchTree{
				Value: 11,
				Left:  nil,
				Right: &BinarySearchTree{
					Value: 22,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &BinarySearchTree{
			Value: 20,
			Left:  &BinarySearchTree{
				Value: 3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(validBst(entry))
}
