package main

import (
	"fmt"
	"math"
)

// implement merge sort

func merge_sort(in []int) []int {
	// splits the arrays left and right
	// calls merge sort on left and right slices
	if len(in) == 1 {
		return in
	}
	mid := int(math.Ceil(float64(len(in)) / 2.0))
	return merge(merge_sort(in[:mid]), merge_sort(in[mid:]))

	// merge to two slices
}

func merge(s1, s2 []int) []int {
	p1, p2, m1 := 0, 0, 0
	merged := make([]int, len(s1)+len(s2))
	for p1 < len(s1) && p2 < len(s2) {
		if s1[p1] <= s2[p2] {
			merged[m1] = s1[p1]
			m1++
			p1++
		} else {
			merged[m1] = s2[p2]
			m1++
			p2++
		}
	}
	for p1 < len(s1) {
		merged[m1] = s1[p1]
		p1++
		m1++
	}
	for p2 < len(s2) {
		merged[m1] = s2[p2]
		p2++
		m1++
	}
	return merged
}

func main() {
	unsorted := []int{5, 3, 8, 20, 19, 1, 2}
	fmt.Println(merge_sort(unsorted))
}
