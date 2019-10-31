package main

/*
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded
by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all
four edges of the grid are all surrounded by water.

Input:
11110
11010
11000
00000

Output: 1

Input:
11000
11000
00100
00011

Output: 3

*/

func numberOfIslands(in [][]byte) int {
	count := 0
	for r := 0; r < len(in); r++ {
		for c := 0; c < len(in[0]); c++ {
			if in[r][c] == 0 {
				continue
			}
			clearer(in, r, c)
			count++
		}
	}
	return count
}

func clearer(in [][]byte, r, c int) {
	// remember to memo
	if in[r][c] == 0 {
		return
	}
	in[r][c] = 0
	// remember to check boundaries!
	// move left
	if c != 0 {
		clearer(in, r, c-1)
	}
	// move right
	if c != len(in[0])-1 {
		clearer(in, r, c+1)
	}
	// move down
	if r != len(in)-1 {
		clearer(in, r+1, c)
	}
	// move up
	if r != 0 {
		clearer(in, r-1, c)
	}
}

