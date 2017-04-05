package stack

import "errors"

var ErrEmptyStack = errors.New("empty_stack")

func (s *Stack) Pop() (string, error) {
	if s.Length() == 0 {
		return "", ErrEmptyStack
	}

	// extract the last item in stack
	s.lock.Lock()
	last := s.top.value
	s.top = s.top.next

	s.size--
	s.lock.Unlock()

	return last, nil
}
