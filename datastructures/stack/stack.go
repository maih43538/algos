package stack

import "sync"

// stack is an ordered list of similar data type
// last in first out

type Element struct {
	value string
	next  *Element
}

type Stack struct {
	lock     sync.Mutex
	top      *Element
	size     int
	capacity int
}

func New(n int) *Stack {
	if n <= 0 {
		n = 100
	}
	return &Stack{
		lock:     sync.Mutex{},
		size:     0,
		capacity: n,
	}
}

func (s *Stack) Length() int {
	return s.size
}
