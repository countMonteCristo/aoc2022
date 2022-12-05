package utils

import (
	"log"
)

// Simple Stack implementation based on slice
type Stack struct {
	data []interface{}
}

func (s *Stack) Push(e interface{}) {
	s.data = append(s.data, e)
}

// Pushes all elements from array `elems` to stack:
// [1, 2, 3].PushN([4, 5]) -> [1, 2, 3, 4, 5]
func (s *Stack) PushN(elems []interface{}) {
	s.data = append(s.data, elems...)
}

// Return and remove top element from the stack
func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		panic("Pop from empty stack")
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

// Return and remove top N elements from the stack
// [1, 2, 3, 4, 5].PopN(2) -> [4, 5], stack=[1, 2, 3]
func (s *Stack) PopN(count int) []interface{} {
	if len(s.data) < count {
		log.Fatal("Not enough elemnts in stack to pop: have ", len(s.data), ", need ", count)
	}
	top := s.data[len(s.data)-count:]
	s.data = s.data[:len(s.data)-count]
	return top
}

// Return (but not remove!) top element from the stack
func (s *Stack) Top() interface{} {
	if len(s.data) == 0 {
		panic("Top from empty stack")
	}
	return s.data[len(s.data)-1]
}
