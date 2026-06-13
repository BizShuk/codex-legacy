package stack

import "errors"

type stack struct {
	stack []interface{}
}

func (s *stack) Push(e interface{}) error {
	s.stack = append(s.stack, e)
	return nil
}

func (s *stack) Pop() (error, interface{}) {
	if s.Size() <= 0 {
		return errors.New("no size..."), nil
	}
	len := len(s.stack) - 1
	popd := s.stack[len]
	s.stack = s.stack[0:len]
	return nil, popd
}

func (s *stack) Size() int {
	return len(s.stack)
}

func (s *stack) Peek() interface{} {
	var last interface{}
	last = s.stack[len(s.stack)-1]

	return last
}

func (s *stack) Reverse() {
	_, last := s.Pop()
	s.Reverse()
	s.Push(last)
}
