package stack_queue

import "testing"

func TestQueue_Dequeue(t *testing.T) {
}

func Test_StackFunctionality(t *testing.T) {
	s := stack{}
	isEmpty := s.IsEmpty()
	if isEmpty != true {
		t.Error("isEmpty() should return true, but returned false")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	expect := []int{3,2,1}
	for _, val := range expect {
		if got, err := s.Pop(); got != val || err != nil  {
			t.Error("stack did not pop properly or returned an error")
		}
	}
}

func Test_QueueFunctionality(t *testing.T) {
	q := queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	expect := []int{1,2,3}
	for _, val := range expect {
		if got, err := q.Dequeue(); got != val || err != nil {
			t.Error("queue did not dequeue properly or returned an error")

		}
	}
}
