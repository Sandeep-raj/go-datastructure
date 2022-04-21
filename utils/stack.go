package utils

import "errors"

type Stack struct {
	Count   int
	Element []interface{}
}

func InitStack() *Stack {
	return &Stack{
		Count:   0,
		Element: make([]interface{}, 0),
	}
}

func (s *Stack) Push(val interface{}) {
	s.Element = append(s.Element, val)
	s.Count++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Count == 0 {
		return nil, errors.New("Stack is Empty")
	} else {
		currVal := s.Element[len(s.Element)-1]
		s.Element = s.Element[:len(s.Element)-1]
		s.Count--
		return currVal, nil
	}
}

func (s *Stack) Size() int {
	return s.Count
}

func (s *Stack) Peek() (interface{}, error) {
	if s.Count == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.Element[len(s.Element)-1], nil
}
