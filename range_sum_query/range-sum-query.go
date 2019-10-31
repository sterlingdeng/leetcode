package main

import "fmt"

/*

Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.

The update(i, val) function modifies nums by updating the element at index i to val.

Example:

Given nums = [1, 3, 5]

sumRange(0, 2) -> 9
update(1, 2)
sumRange(0, 2) -> 8
Note:

The array is only modifiable by the update function.
You may assume the number of calls to update and sumRange function is distributed evenly.

*/

type NumArray struct {
	data   []int
	cumSum []int
}

func NewNumArray(nums []int) NumArray {
	cumSum := make([]int, len(nums))
	cumSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		cumSum[i] = nums[i] + cumSum[i-1]
	}
	return NumArray{
		cumSum: cumSum,
		data: nums,
	}
}

func (n *NumArray) Update(i, val int) {
	n.data[i] = val
	for j := i; j < len(n.cumSum); j++ {
		n.cumSum[j] = n.cumSum[j] - n.data[i] + val
	}
}

func (n *NumArray) SumRange(i, j int) int {
	if i == 0 {
		return n.cumSum[j]
	}
	return n.cumSum[j] - n.cumSum[i-1]
}

func main() {
	n := NewNumArray([]int{1, 3, 5, 1})
	fmt.Printf("%+v\n", n)
	n.Update(0, 2)
	fmt.Printf("%+v\n", n)
	fmt.Println(n.SumRange(2, 3))
}
