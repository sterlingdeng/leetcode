package hack_hour

import (
	"errors"
)

/*
 * Create a stack. Then create a queue using two stacks.
 * assume a stack of ints to make it simple
 */

type stack struct {
	values []int
}

// use pointer receivers because we will be mutating the struct
func (s *stack) Push(v int) {
	s.values = append(s.values, v)
}

func (s *stack) Pop() (int, error) {
	if len(s.values) == 0 {
		return -1, errors.New("nothing to pop")
	}
	popped := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return popped, nil
}

func (s *stack) IsEmpty() bool {
	return len(s.values) == 0
}

// queue struct represents the queue data structure using two stacks
type queue struct {
	s1 stack
	s2 stack
}

func (q *queue) Enqueue(v int) {
	q.s1.Push(v)
}

func (q *queue) Dequeue() (int, error) {
	for !q.s1.IsEmpty() {
		val, _ := q.s1.Pop()
		q.s2.Push(val)
	}
	dequeued, err := q.s2.Pop()
	if err != nil {
		return -1, errors.New("nothing to dequeue")
	}
	for !q.s2.IsEmpty() {
		val, _ := q.s2.Pop()
		q.s1.Push(val)
	}
	return dequeued, nil
}
