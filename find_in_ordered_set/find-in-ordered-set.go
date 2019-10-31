package main

import (
	"fmt"
	"math"
)

/*
How quickly could we check if a given integer is in the array?
Assume the array is sorted.
You can do better than O(n) time

var nums = [1, 4, 6, 7, 9, 17, 45] len 7 / 2 = 3.5 rnd up to 4
findInOrderedSet(nums, 4);  -> true
findInOrderedSet(nums, 2);  -> false

*/

func findInOrderedSet(n []int, target int) bool {
	/*
		set lower to 0 and upper bound to len(n) - 1
		get the mid point

		for lower < upper
			if value at mid == target return true
			else if target is less than, set upper to value at idx
			else set lower at value at idx
		end for
		return false
	*/
	l, r := 0, len(n) - 1
	// needs to be <= so it can evaluate when the pointers point to the last value, this makes the
	// boundaries inclusive
	for l <= r {
		mid := int(math.Ceil((float64(r)+float64(l))/float64(2))) // doesn't matter if its ceil or floor
		if target == n[mid] {
			return true
		}
		if target > n[mid] {
			// since we know the val at middle is not the target, we can move the boundaries +1 or -1
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}

func main() {
	fmt.Println(findInOrderedSet([]int{1, 4, 6, 7, 9, 17, 45}, 6))
}
