package main

import "errors"

/*
 * Create a stack with the push, pop, and getMax methods.
 * push should return the new length of the stack.
 * pop should return the element that was just removed.
 * getMax should return the largest value currently in the stack.
 * BONUS: The getMax method should retrieve the maximum value from the stack in O(1) time.
 */

var ErrEmptyStack = errors.New("empty stack")

type stack struct {
	data []int
}

func (s *stack) push(n int) {
	s.data = append(s.data, n)
}

func (s *stack) pop() (int, error) {
	if len(s.data) == 0 {
		return 0, ErrEmptyStack
	}
	popped := s.data[len(s.data) - 1]
	s.data = s.data[:len(s.data) - 1]
	return popped, nil
}

func (s *stack) peek() (int, error) {
	if len(s.data) == 0 {
		return -1, ErrEmptyStack
	}
	return s.data[len(s.data) - 1], nil
}

type MaxStack struct {
	data stack
	max stack
}

func (ms *MaxStack) Push(n int) {
	ms.data.push(n)
	currentMax, err := ms.max.peek()
	if err == ErrEmptyStack {
		ms.max.push(n)
	} else {
		if currentMax < n {
			ms.max.push(n)
		} else {
			ms.max.push(currentMax)
		}
	}
}

func (ms *MaxStack) Pop() int {
	num, _ := ms.data.pop()
	ms.max.pop()
	return num
}

func (ms *MaxStack) GetMax() int {
	max, err := ms.max.peek()
	if err != nil {
		return 0
	}
	return max
}
