package main

/*

You are given coins of different denominations and a total amount of money amount. Write a function to compute the
fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any
combination of the coins, return -1.

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
*/

func coinChange(coins []int, target int) int {
	count := helper(coins, target, 0)
	if count == 0 {
		return -1
	}
	return count
}

func helper(coins []int, target, count int) int {
	if target < 0 {
		return -1
	}
	if target == 0 {
		return count
	}
	for i := len(coins)-1; i >= 0; i-- {
		ct := helper(coins, target-coins[i], count+1)
		if ct != -1 {
			return ct
		}
	}
	return -1
}
