package anagram_count

import (
	"sort"
)

func anagram_count(s []string) int {
	m := make(map[string]int)
	for i := 0; i< len(s); i++ {
		rn := []rune(s[i])
		sort.Slice(rn, func(i, j int) bool {
			return rn[i] < rn[j]
		})
		s[i] = string(rn)
		if _, has := m[s[i]]; has {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	count := 0
	for _, v := range m {
		count += v - 1
	}
	return count
}

