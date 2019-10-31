package stack_queue

/*
 * Create a stack. Then create a queue using two stacks.
 * assume a stack of ints to make it simple
 */

// stack struct represents the stack data structure
type stack struct {}

// use pointer receivers because we will be mutating the struct
func (s *stack) Push(v int) {}

func (s *stack) Pop() (int, error)  {
	return -1, nil
}

func (s *stack) IsEmpty() bool {
	return false
}


// queue struct represents the queue data structure using two stacks
type queue struct {}

func (q *queue) Enqueue(v int) {}

func (q *queue) Dequeue() (int, error) {
	return -1, nil
}


