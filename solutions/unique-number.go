package hack_hour

/**
 * Write a function that will take an array of integers, all of which will appear exactly twice,
 * except for one integer that will appear exactly once. Return the integer that appears once.
 *
 * uniqueNumber([1,2,1,3,3]); -> 2
 *
 * BONUS:
 * Complete the challenge in O(n) time
 * Complete the challenge in O(1) space
 *
 */

// for hint: https://hackernoon.com/a-javascript-interview-question-and-a-digression-into-xor-3f88bb5ab3be

func uniqueNumber(slice []int) int {
	unique := slice[0]
	for i := 1; i < len(slice); i++ {
		unique ^= slice[i]
	}
	return unique
}

func uniqueNumber_naive(slice []int) int {
	occurence := make(map[int]int)
	for _, val := range slice {
		count, exists := occurence[val]
		if !exists {
			occurence[val] = 1
		} else {
			occurence[val] = count + 1
		}
	}

	for k, v := range occurence {
		if v == 1 {
			return k
		}
	}
	return -1
}
