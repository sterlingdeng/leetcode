package main

import (
	"fmt"
	"math"
)

/* You are given a string that represents a binary number
 *
 * Write a function that converts the binary string to a decimal number
 *
 * Example:
 * 	binToDec('0')   -> 0
 * 	binToDec('11')  -> 3
 * 	binToDec('100') -> 4
 * 	binToDec('101') -> 5
 *  binToDec('0101') -> 5
 *
 * Extension:
 * Write a function that converts a decimal number to binary (then maybe hexadecimal)
 */

// 10 9 8 7 6
// 0 1 2 3 4
func binToDec(s string) int {
	var rtn int
	for i := len(s) - 1; i >= 0; i-- {
		num := int(s[i] - '0')
		if num > 0 {
			rtn += int(math.Pow(2, float64(len(s) - 1 - i)))
		}
	}
	return rtn
}

func main() {
	fmt.Println(binToDec("110"))
}
