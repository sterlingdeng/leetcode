package data_structures

import "errors"

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("Nothing to pop")
	}
	popped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return popped, nil
}

func (s *Stack) Peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data) - 1]
}

func (s *Stack) Size() int {
	return len(s.data)
}
