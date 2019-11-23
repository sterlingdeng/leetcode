package sort_colors

/*
https://leetcode.com/problems/sort-colors/
LC: medium

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:

Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Follow up:

A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with a one-pass algorithm using only constant space?

 */

// 2 0 1
// modify in place
func sortColors(nums []int) {
	zeroPtr := 0
	twoPtr := len(nums) - 1
	for i := 0; i <= twoPtr; i++ {
		currVal := nums[i]
		for currVal == 2 && i <= twoPtr {
			swap(nums, i, twoPtr)
			twoPtr--
			currVal = nums[i]
		}
		if currVal == 0 {
			if i != zeroPtr {
				swap(nums, i, zeroPtr)
			}
			zeroPtr++
		}
	}
}

func swap(arr []int, idx1, idx2 int) {
	arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
}
