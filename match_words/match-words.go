package main

import (
	. "hack_hour/data_structures"
	"regexp"
	"strings"
)

// Some languages have "if" statements that are closed by "fi" instead of curly brackets. Or they close a "case" with "esac",
//i.e. the same keyword backwards. for this problem we'll check that all words in a string are "closed". Write a function that
//takes a string and returns true if every word is closed by its backwards counterpart. Words must be separated by space or
//punctuation.

// matchWord('__END_DNE-----');  -> true
// matchWord('__ENDDNE__');  -> false       (not separated by a space)
// matchWord('IF()()fi[]');  -> true        (should be case-insensitive)
// matchWord('for__if__rof__fi');  -> false     not properly closed. like ( [) ]
// matchWord('%%$@$while  try ! yrt  for if_fi rof #*#  elihw');  -> true
// matchWord('');  -> true

func matchWords(s string) bool {
	re := regexp.MustCompile(`[a-zA-Z]+`)
	words := re.FindAllString(s, -1)
	stack := Stack{}
	for _, word := range words {
		word = strings.ToLower(word)
		peeked, ok := stack.Peek().(string) // need to type assert to string, because no generics
		if ok {
			if isReverse(peeked, word) {
				stack.Pop()
				continue
			}
		}
		stack.Push(word)
	}
	// return bool based on len of stack
	return stack.Size() == 0
}

func isReverse(first, second string) bool {
	var sb strings.Builder
	for i := len(first) - 1; i >= 0; i-- {
		sb.WriteRune(rune(first[i])) // accessing chars using s[i] notation returns a uint8 that needs to be converted to a rune
	}
	reversedFirst := sb.String()
	if reversedFirst == second {
		return true
	}
	return false
}

func main() {
	matchWords("__END_DNE-----")
}

