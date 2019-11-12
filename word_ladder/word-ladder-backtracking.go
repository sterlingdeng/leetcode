package word_ladder

import (
	"math"
)

/*
https://leetcode.com/problems/word-ladder/
medium

Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

*/

// not an optimal solution. try again
func WordLadder_backtracking(start, end string, list []string) int {
	if !isThere(end, list) {
		return 0
	}
	memo := make([]bool, len(list))
	min := helper(start, end, list, 0, 1, math.MaxInt64, memo)
	if min == math.MaxInt64 {
		return 0
	}
	return min
}

func helper(currWord, end string, list []string, idx, count, min int, memo []bool) int {
	if currWord == end {
		return int(math.Min(float64(count), float64(min)))
	}
	for i := 0; i < len(list); i++ {
		nextWord := list[i]
		if differsByAtMostOne(currWord, nextWord) && !memo[i] {
			memo[i] = true
			min = helper(nextWord, end, list, i, count+1, min, memo)
			memo[i] = false
		}
	}
	return min
}

func differsByAtMostOne(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	diffCount := 0
	r1 := []rune(s1)
	r2 := []rune(s2)
	for i := 0; i < len(s1); i++ {
		if r1[i] != r2[i] {
			if diffCount == 1 {
				return false
			}
			diffCount++
		}
	}
	return true
}

func isThere(end string, list []string) bool {
	found := false
	for i := 0; i < len(list); i++ {
		if list[i] == end {
			found = true
			break
		}
	}
	return found
}
