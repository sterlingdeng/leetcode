package main

import "fmt"

/**
 * We have our lists of orders sorted numerically already, in arrays.
 * Write a functin to merge our arrays of orders into one sorted array.
 * These may be of different lengths.
 *
 * var my_array = [3,4,6,10,11,15,21];
 * var another_array = [1,5,8,12,14,19];
 *
 * mergeArrays(my_array, another_array); -> [1, 3, 4, 5, 6, 8, 10, 11, 12, 14, 15, 19, 21]
 *
 * BONUS:
 * Complete in O(n) time
 *
 */

func mergeArrays(x, y []int) []int {
	merged := make([]int, len(x) + len(y))
	xPtr := 0
	yPtr := 0
	mergedPtr := 0
	for xPtr < len(x) && yPtr <len(y) {
		if x[xPtr] < y[yPtr] {
			merged[mergedPtr] = x[xPtr]
			xPtr++
		} else {
			merged[mergedPtr] = y[yPtr]
			yPtr++
		}
		mergedPtr++
	}
	for xPtr < len(x) {
		merged[mergedPtr] = x[xPtr]
		mergedPtr++
		xPtr++
	}
	for yPtr < len(y) {
		merged[mergedPtr] = y[yPtr]
		mergedPtr++
		yPtr++
	}
	return merged
}

func main() {
	ar1 := []int{3,4,6,10,11,15,21}
	ar2 := []int{1,5,8,12,14,19}

	fmt.Println(mergeArrays(ar1, ar2))
}
