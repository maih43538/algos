package stack

import "errors"

var ErrMaxCapReached = errors.New("max_capacity_reached")

func (s *Stack) Push(in string) error {
	// if stack is at capacity, return error
	if s.Length() >= s.capacity {
		return ErrMaxCapReached
	}

	// replace stack top with new incoming data
	s.lock.Lock()
	e := &Element{
		value: in,
		next:  s.top,
	}
	s.top = e
	s.size++
	s.lock.Unlock()
	return nil
}
