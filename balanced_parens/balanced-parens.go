package main

import "github.com/pkg/errors"

/*
 * write a function that takes a string of text and returns true if
 * the parentheses are balanced and false otherwise.
 *
 * Example:
 *   balancedParens('(');  // false
 *   balancedParens('()'); // true
 *   balancedParens(')(');  // false
 *   balancedParens('(())');  // true
 *
 * Step 2:
 *   make your solution work for all types of brackets
 *
 * Example:
 *  balancedParens('[](){}'); // true
 *  balancedParens('[({})]');   // true
 *  balancedParens('[(]{)}'); // false
 *
 * Step 3:
 * ignore non-bracket characters
 * balancedParens(' var wow  = { yo: thisIsAwesome() }'); // true
 * balancedParens(' var hubble = function() { telescopes.awesome();'); // false
 *
 *
 */

type Stack struct {
	d []rune
}

func (s *Stack) Push(r rune) {
	s.d = append(s.d, r)
}

func (s *Stack) Pop() (rune, error) {
	if len(s.d) == 0 {
		return 0, errors.New("nothing to pop")
	}
	popped := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return popped, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.d) == 0
}

func balanced_parens(s string) bool {
	// if (, [, { .. it will go onto the stack
	// if we encounter ), ], }, we will pop off the stack and verify the compliment
	// if its not the compliment, we can return false
	// if length of the stack if not 0, then return true
	var complimentMap = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var parens = map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
	}
	stack := Stack{}
	// iterate through the characters in s
	for _, char := range s {
		// check if its (, [, or {
		if _, exists := parens[char]; exists {
			stack.Push(char)
			continue
		}
		if compliment, exists := complimentMap[char]; exists {
			popped, err := stack.Pop()
			if err != nil {
				return false
			}
			if popped != compliment {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
