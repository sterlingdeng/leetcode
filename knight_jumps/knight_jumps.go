package main

// Have the function knightjumps(str) read str which will be a
// string consisting of the location of a knight on a standard 8x8 chess board with no other pieces on the board.

// The structure of str will be the following string: "(x y)" which represents the position of the knight with x and y ranging
// from 1 to 8.

// Your function should determine the number of spaces the knight can move to from a given location.
// For example: if str is "(4 5)" then your program should output 8 because the knight can move to 8 different spaces
// from position x = 4 and y = 5.
//  example input:
// var str = "(4 5)"

func knightsJump(x, y int) int {
	jumps := 8
	// 8 x 8 board
	x_moves := []int{-2, -2, -1, -1, 1, 1, 2, 2}
	y_moves := []int{-1, 1, 2, -2, 2, -2, 1, -1}

	for i := range x_moves {
		newx := x + x_moves[i]
		newy := y + y_moves[i]
		if newx < 1 || newx > 8 || newy < 1 || newy > 8 {
			jumps -= 1
		}
	}
	return jumps
}
