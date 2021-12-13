package main

import (
	"sync"
)

type Stack struct {
	stack []int
	lock  sync.RWMutex
}

func (s *Stack) push(val int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stack = append(s.stack, val)
}

func (s *Stack) pop() int {
	len := len(s.stack)
	if len > 0 {
		s.lock.Lock()
		defer s.lock.Unlock()
		v := s.stack[len-1]
		s.stack = s.stack[:len-1]
		return v
	}
	GorthError("Stack underflow")
	return -1
}
