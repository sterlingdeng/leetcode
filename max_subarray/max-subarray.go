package main

import (
	"math"
)

/* You are given an array of integers with both positive and negative numbers. Write a function to
 * find the maximum sum of all subarrays. A subarray is a section of consecutive elements from the
 * original array. The whole array can be considered a sub array of itself.
 *
 * For example: maxSubarray([1, -2, 3, 10, -4, 7, 2, -5]) -> 18 from [3, 10, -4, 7, 2]
 *              maxSubarray([15,20,-5,10])  -> 40
 *
 */

func maxSubarray(n []int) int {
	// naive solution, but dp can be used if the numbers are positive
	subarraymax := math.Inf(-1)
	globalMax := math.Inf(-1)
	for _, number := range n {
			fltnumber := float64(number)
			subarraymax = math.Max(fltnumber, subarraymax + fltnumber)
			globalMax = math.Max(globalMax, subarraymax)
	}
	return int(globalMax)
}
