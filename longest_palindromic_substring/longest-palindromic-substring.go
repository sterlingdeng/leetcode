package main

import "fmt"

/*
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
*/

type substring struct {
	l int
	r int
}

func (s substring) Dist() int {
	return s.r - s.l
}

func LongestPalindromicSubstring(s string) string {
	res := substring{0, 1}
	for i := 1; i < len(s); i++ {
		// even
		var largest substring
		even := getDistance(s, i-1, i)
		odd := getDistance(s, i, i)

		if even.Dist() > odd.Dist() {
			largest = even
		} else {
			largest = odd
		}
		if largest.Dist() > res.Dist() {
			res = largest
		}
	}
	return s[res.l:res.r]
}

func getDistance(s string, l, r int) substring {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return substring{l+1, r}
}

func main() {
	fmt.Println(LongestPalindromicSubstring("abaxyzzyxf"))

	fmt.Println(LongestPalindromicSubstring("it's highnoon"))
}
