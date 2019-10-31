package main

/*

In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:

1p piece
2p piece
5p piece
10p piece
20p piece
50p piece
1 euro (100p)
2 euro (200p)

It is possible to make £2 in the following way:

1 * £1 + 1 * 50p + 2 * 20p + 1 * 5p + 1 * 2p + 3 * 1p
How many different ways can £2 be made using any number of coins?

example usage of `coinSum`:

// aka, there's only one way to make 1p. that's with a single 1p piece
coinSum(1) === 1
// aka, there's only two ways to make 2p. that's with two, 1p pieces or with a single 2p piece
coinSum(2) === 2
*/

func coinSum(coins []int, target int) int {
		// set up
		matrix := make([][]int, len(coins) + 1)
		for i := range matrix {
			matrix[i] = make([]int, target + 1)
		}
		// seed initial conditions. (There is no way to make change for a target of 0)
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
		// seed the first column with 1's (There is 1 way to make change with a target of 0, and that is 0 ways)
		for i := range matrix {
			matrix[i][0] = 1
		}

		for r := 1; r < len(matrix); r++ {
			currentCoin := coins[r-1]
			for c := 1; c < len(matrix[0]); c++ {
				if c < currentCoin {
					matrix[r][c] = matrix[r-1][c]
					continue
				}
				matrix[r][c] = matrix[r-1][c] + matrix[r][c-currentCoin]
			}
		}
		return matrix[len(matrix)-1][len(matrix[0])-1]
}
