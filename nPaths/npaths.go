package main

/**
 *
 *  How many unique ways can one move from one corner of a n x n board to the other corner.
 *  You can move up, down, left, and right. You cannot visit spots that you have visited already.
 *  Make your solution work for a grid of any size.
 *
 */

func npaths(n int) int {
	if n < 2 {
		return n
	}
	board, visited := newBoard(n)

	return helper(board, 0, 0, visited)
}

func newBoard(size int) ([][]int, [][]bool) {
	// to create a size x size matrix
	board := make([][]int, size)
	visited := make([][]bool, size)
	for i := range board {
		board[i] = make([]int, size)
		visited[i] = make([]bool, size)
	}
	return board, visited
}

func helper(board [][]int, r, c int, memo [][]bool) int {
	// base case: if end of board return 1
	count := 0
	if r == len(board) - 1 && c == len(board) - 1 {
		return 1
	}
	// did we visit already?
	if memo[r][c] {
		return count
	}
	memo[r][c] = true
	// left is c - 1, right is c + 1, up is r - 1, down is r + 1
	// move left first
	if c != 0 {
		count += helper(board, r, c - 1, memo)
	}
	// move right
	if c != len(board) - 1 {
		count += helper(board, r, c + 1, memo)
	}
	// move up
	if r != 0 {
		count += helper(board, r - 1, c, memo)
	}

	if r != len(board) - 1 {
		count += helper(board, r + 1, c, memo)
	}
	memo[r][c] = false
	return count
}
