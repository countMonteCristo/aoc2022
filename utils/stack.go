package utils

import (
	"fmt"
)

// Simple Stack implementation based on slice
type Stack[T any] struct {
	data []T
}

// Push new value to the Stack
func (s *Stack[T]) Push(e T) {
	s.data = append(s.data, e)
}

// Pushes all elements from array `elems` to stack:
// [1, 2, 3].PushN([4, 5]) -> [1, 2, 3, 4, 5]
func (s *Stack[T]) PushN(elems []T) {
	s.data = append(s.data, elems...)
}

// Return and remove top element from the stack
func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		panic("Pop from empty stack")
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

// Return and remove top N elements from the stack
// [1, 2, 3, 4, 5].PopN(2) -> [4, 5], stack=[1, 2, 3]
func (s *Stack[T]) PopN(count int) []T {
	if len(s.data) < count {
		msg := fmt.Sprintf("Not enough elemnts in stack to pop: have %d, need %d", len(s.data), count)
		panic(msg)
	}
	top := s.data[len(s.data)-count:]
	s.data = s.data[:len(s.data)-count]
	return top
}

// Return (but not remove!) top element from the stack
func (s *Stack[T]) Top() T {
	if len(s.data) == 0 {
		panic("Top from empty stack")
	}
	return s.data[len(s.data)-1]
}
