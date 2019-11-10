package levenshtein_distance

import (
	"math"
)

/*

Write a function that takes in two strings and returns the minimum number of edit operations that need to be performed on the first string to obtain the second string.
There are three edit operations: insertion of a character, deletion of a character, and substitution of a character for another.

Sample input: "abc", "yabd"

	_ a b c
y
a
b
d
*/

func LevenshteinDistance(a, b string) int {
	// setup
	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
	}
	for i := 0; i < len(matrix[0]); i++ {
		matrix[0][i] = i
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i][0] = i
	}
	// work
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			// if letters are the same
			if a[r-1] == b[c-1] {
				matrix[r][c] = matrix[r-1][c-1]
			} else {
				if matrix[r-1][c] == matrix[r][c-1] {
					matrix[r][c] = matrix[r-1][c]
				} else {
					matrix[r][c] = int(math.Min(float64(matrix[r-1][c]+1), float64(matrix[r][c-1]+1)))
				}
			}
		}
	}
	return matrix[len(matrix)-1][len(matrix[0])-1]
}

