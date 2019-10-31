package hack_hour


/* You are given an array of integers and a target number. Write a function that returns true if
 * there is a subset of the array that sums up to the target and returns false otherwise. A subset
 * can be any size and the elements do not have to appear consecutively in the array.
 *
 * subsetSum([3, 7, 4, 2], 5) - > true, 3 + 2 = 5
 * subsetSum([3, 34, 4, 12, 5, 12], 32) -> true, 3 + 12 + 5 + 12 = 32
 * subsetSum([8, 2, 4, 12], 13) -> false
 */

// assume numbers are positive

func SubsetSum(slice []int, target int) bool {
	return subsetSum_helper(slice, target, 0, 0)
}

func subsetSum_helper(slice []int, target, sum, idx int) bool {
	if sum == target {
		return true
	}
	for i := idx; i < len(slice); i++ {
		newSum := sum + slice[i]
		if subsetSum_helper(slice, target, newSum, i+1) {
			return true
		}
	}
	return false
}
