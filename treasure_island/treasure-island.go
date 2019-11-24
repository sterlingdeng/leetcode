package treasure_island

/*
https://leetcode.com/discuss/interview-question/347457
backtracking
*/

import "math"

const (
	danger   = "D"
	treasure = "X"
)

func TreasureIsland(board [][]string) int {
	if len(board) < 1 {
		panic("not possible")
	}
	visitedBoard := make([][]bool, len(board))
	for i := range visitedBoard {
		visitedBoard[i] = make([]bool, len(board[0]))
	}
	shortest := walker(0, 0, 0, math.MaxInt64, board, visitedBoard)
	return shortest
}

func walker(r, c, currDist, currShortest int, board [][]string, visited [][]bool) int {
	// check if we are outside
	if isOutside(r, c, len(board), len(board[0])) {
		return currShortest
	}
	// check if we visited before
	if visited[r][c] {
		return currShortest
	}
	// check if its dangerous
	if board[r][c] == danger {
		return currShortest
	}
	// if we are at the treasure, return currDist
	if board[r][c] == treasure {
		return currDist
	}
	// if currDist is the currShortest and we aren't at the treasure, exit
	if currDist == currShortest {
		return currShortest
	}
	visited[r][c] = true

	var pathDist int
	pathDist = walker(r+1, c, currDist+1, currShortest, board, visited)
	currShortest = min(pathDist, currShortest)

	pathDist = walker(r-1, c, currDist+1, currShortest, board, visited)
	currShortest = min(pathDist, currShortest)

	pathDist = walker(r, c+1, currDist+1, currShortest, board, visited)
	currShortest = min(pathDist, currShortest)

	pathDist = walker(r, c-1, currDist+1, currShortest, board, visited)
	currShortest = min(pathDist, currShortest)

	visited[r][c] = false
	return currShortest
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isOutside(r, c, rLength, cLength int) bool {
	return r < 0 || c < 0 || r >= rLength || c >= cLength
}
