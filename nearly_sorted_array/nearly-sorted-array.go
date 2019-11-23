package nearly_sorted_array

/*

Given an array of n elements, where each element is at most k away from its target position, devise an algorithm that sorts in O(n log k) time.
For example, let us consider k is 2, an element at index 7 in the sorted array, can be at indexes 5, 6, 7, 8, 9 in the given array.

Examples:

Input : arr[] = {6, 5, 3, 2, 8, 10, 9}
            k = 3
Output : arr[] = {2, 3, 5, 6, 8, 9, 10}

Input : arr[] = {10, 9, 8, 7, 4, 70, 60, 50}
         k = 4
Output : arr[] = {4, 7, 8, 9, 10, 50, 60, 70}

*/

import (
	. "hack_hour/data_structures"
)

func NearlySortedArray(arr []int, k int) []int {
	// already sorted or doesn't need to be sorted
	if k == 0 || len(arr) < 2 {
		return arr
	}

	if k > len(arr) {
		panic("if k > len(arr) it won't work.")
	}

	minHeap := NewHeap(func(a, b int) bool {
		return a < b
	})

	for i := 0; i <= k && i < len(arr); i++ {
		minHeap.Add(arr[i])
	}

	idx := 0
	for i := k + 1; i < len(arr); i++ {
		popped, _ := minHeap.Remove()
		arr[idx] = popped
		minHeap.Add(arr[i])
		idx++
	}

	for !minHeap.IsEmpty() {
		popped, _ := minHeap.Remove()
		arr[idx] = popped
		idx++
	}

	return arr
}
