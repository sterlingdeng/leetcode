package kth_largest_element

/*
Note: Avoid using built-in std::nth_element (or analogous built-in functions in languages other than C++) when solving this challenge. Implement them yourself, since this is what you would be asked to do during a real interview.

Find the kth largest element in an unsorted array. This will be the kth largest element in sorted order, not the kth distinct element.

Example

For nums = [7, 6, 5, 4, 3, 2, 1] and k = 2, the output should be
kthLargestElement(nums, k) = 6;
For nums = [99, 99] and k = 1, the output should be
kthLargestElement(nums, k) = 99.
Input/Output

[execution time limit] 4 seconds (go)

[input] array.integer nums

Guaranteed constraints:
1 ≤ nums.length ≤ 105,
-105 ≤ nums[i] ≤ 105.

[input] integer k

Guaranteed constraints:
1 ≤ k ≤ nums.length.

[output] integer

*/

func KthLargestElement(nums []int, k int) int {
	k = len(nums) - k
	helper(nums, 0, len(nums)-1, k)
	return nums[k]
}

func helper(nums []int, low, high, k int) {
	if low < high {
		partitionIdx := partition(nums, low, high, k)
		if partitionIdx == k {
			return
		}
		if partitionIdx < k {
			helper(nums, partitionIdx+1, high, k)
		} else {
			helper(nums, low, partitionIdx-1, k)
		}
	}
}

func partition(arr []int, low, high, k int) int {
	partitionValue := arr[high]
	partitionIdx := low
	for i := low; i < high; i++ {
		if arr[i] < partitionValue {
			swap(arr, partitionIdx, i)
			partitionIdx++
		}
	}
	swap(arr, partitionIdx, high)
	return partitionIdx
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
