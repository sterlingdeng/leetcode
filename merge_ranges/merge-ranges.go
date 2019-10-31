package main

import (
	"sort"
)

/**
 * Write a function mergeRanges that takes an array of meeting time ranges and returns an array of condensed ranges.
 *
 * Example:
 * var times = [[0, 1], [3, 5], [4, 8], [9, 10], [10, 12], [11, 14]]
 *
 * mergeRanges(times); -> [[0, 1], [3, 8], [9, 12]]
 *
 * Do not assume the ranges are in order
 */

func mergeRanges(in [][]int) [][]int {
	var merged [][]int
	sort.Slice(in, func(i, j int) bool {
		return in[i][0] < in[j][0]
	})
	for i := 0; i < len(in)-1; i++ {
		starting := in[i][0]
		ending := in[i][1]
		for j := i + 1; j < len(in); j++ {
			if ending >= in[j][0] {
				ending = in[j][1]
				i = j
			} else {
				break
			}
		}
		merged = append(merged, []int{starting, ending})
	}
	return merged
}
