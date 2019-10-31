package main

import (
	"math"
	"sort"
)

/**
 * Given an array of integers, find the highest product you can get from three of the integers.
 */

// time n*logn space 1
func highestProducts(n []int) int {
	sort.Ints(n)
	positives := n[len(n)-1] * n[len(n)-2] * n[len(n)-3]
	withNegatives := n[0] * n[1] * n[len(n)-1]
	if positives > withNegatives {
		return positives
	}
	return withNegatives
}

// time n space 1
func highestProductOptimal(n []int) int {
	// scan to get largest three
	// scan to get smallest two

	largest := math.Inf(-1)
	second_largest := math.Inf(-1)
	third_largest := math.Inf(-1)

	smallest := math.Inf(1)
	second_smallest := math.Inf(1)

	for _, intval := range n {
		val := float64(intval)
		if val > largest {
			third_largest = second_largest
			second_largest = largest
			largest = val
		} else if val <= largest && val > second_largest {
			third_largest = second_largest
			second_largest = val
		} else if val <= second_largest && val > third_largest {
			third_largest = val
		}

		if val < smallest {
			second_smallest = smallest
			smallest = val
		} else if val > smallest && val < second_smallest {
			second_smallest = val
		}
	}

	candidate1 := largest * second_largest * third_largest
	candidate2 := smallest * second_smallest * largest

	return int(math.Max(candidate1, candidate2))

}

