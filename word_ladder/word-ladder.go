package word_ladder

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
func word_ladder(beginWord string, endWord string, wordList []string) int {
	if !isSolvable(wordList, endWord) {
		return 0
	}
	// instantiate necessary data structures
	wordWild, wildWord := makeAdjList(append(wordList, beginWord))
	visited := make(map[string]bool)
	depthqueue := newSlice(1, 1)
	// instantiate initial variables
	queue := []string{beginWord}
	for len(queue) != 0 {
		// unshift
		currNode := queue[0]
		queue = queue[1:]
		// unshift
		currDepth := depthqueue[0]
		depthqueue = depthqueue[1:]
		// if we've traversed this node already, skip it
		if _, has := visited[currNode]; has {
			continue
		}
		visited[currNode] = true
		if currNode == endWord {
			return currDepth
		}
		// add children to the queue
		wildCards := wordWild[currNode]
		for _, wildCard := range wildCards {
			queue = append(queue, wildWord[wildCard]...)
			depthqueue = append(depthqueue, newSlice(len(wildWord[wildCard]), currDepth+1)...)
		}
	}
	return 0
}

func isSolvable(wordList []string, endWord string) bool {
	for _, word := range wordList {
		if word == endWord {
			return true
		}
	}
	return false
}

func newSlice(size, val int) []int {
	r := make([]int, size)
	for i := range r {
		r[i] = val
	}
	return r
}

// rtn 1: map word -> wild cards
// rtn 2: map wild cards -> words
func makeAdjList(wordList []string) (map[string][]string, map[string][]string) {
	wordWild := make(map[string][]string)
	wildWord := make(map[string][]string)

	// for each word
	for _, word := range wordList {
		var withWildcards []string
		// convert word to rune
		// for each index in the word
		for i := range []rune(word) {
			t := []rune(word)
			t[i] = '*'
			if list, has := wildWord[string(t)]; !has {
				wildWord[string(t)] = []string{word}
			} else {
				wildWord[string(t)] = append(list, word)
			}
			withWildcards = append(withWildcards, string(t))
		}
		wordWild[word] = withWildcards
	}
	return wordWild, wildWord

	//aj := make(map[string][]string)
	//for i := 0; i < len(wordList); i++ {
	//	currWord := wordList[i]
	//	if _, has := aj[currWord]; !has {
	//		aj[currWord] = []string{}
	//	}
	//	for j := 0; j < len(wordList); j++ {
	//		if i == j {
	//			continue
	//		}
	//		if offByLessThanOne(currWord, wordList[j]) {
	//			aj[currWord] = append(aj[currWord], wordList[j])
	//		}
	//	}
	//}
	//return aj
}

func offByLessThanOne(s1, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	count := 0
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			if count == 1 {
				return false
			}
			count++
		}
	}
	return true
}
