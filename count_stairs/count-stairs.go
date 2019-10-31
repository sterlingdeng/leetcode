package main

/* You are faced by a staircase that is N steps high. You can take 1 or 2 steps at a time. Write
 * a function to calculate how many different ways you can go up the flight of stairs.
 *
 * Example: n === 5. You are 5 steps away from the top. You can take these different ways to the top:
 * 1 + 1 + 1 + 1 + 1
 * 1 + 1 + 1 + 2
 * 1 + 1 + 2 + 1
 * 1 + 2 + 1 + 1
 * 1 + 2 + 2
 * 2 + 1 + 1 + 1
 * 2 + 1 + 2
 * 2 + 2 + 1
 *
 * That is a total of 8 different ways to take 5 steps, given that you can take 1 or 2 steps at a time.
 */


func countStairs(n int) int {
	matrix := make([]int, n + 1) // create memo table
	matrix[0], matrix[1], matrix[2] = 0, 1, 2 // seed the table with initial conditions
	for i := 3; i < len(matrix); i++ {
		matrix[i] = matrix[i-1] + matrix[i-2] // build up results
	}
	return matrix[len(matrix)-1] // return final result
}
