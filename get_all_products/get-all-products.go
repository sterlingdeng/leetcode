package main

/* Accepts an array of integers and returns an array of all the possible products made by
 * multiplying all but one number. In other words, find all the products of multiplying any
 * array.length-1 numbers in the array.
 *
 * ex: getProducts([1, 7, 3, 4]); ->  [84, 12, 28, 21]
 * this is done via:
 * [7*3*4, 1*3*4, 1*7*4, 1*7*3]
 *
 * do not use division, because zero might be in the array and you cannot divide by zero
 */

func getAllProducts(n []int) []int {
	if len(n) == 1 {
		return n
	}
	// n = [1, 7, 3, 4]
	// left =  [  1   ,  1  , 1*7 , 1*7*3]
	// right = [7*3*4 , 3*4 ,  4  ,   0]
	left := make([]int, len(n))
	right := make([]int, len(n))
	output := make([]int, len(n))
	// prefill with 1
	for i := range left {
		left[i] = 1
		right[i] = 1
	}
	// start by calculating products of everything to the left
	for i := 1; i < len(n); i++ {
		left[i] = left[i-1] * n[i-1]
	}
	// calculate products of everything to the right of i
	for i := len(n) - 2; i >= 0; i-- {
		right[i] = right[i+1] * n[i+1]
	}
	// fill the output slice
	for i := 0; i < len(output); i++ {
		output[i] = left[i] * right[i]
	}
	return output
}

