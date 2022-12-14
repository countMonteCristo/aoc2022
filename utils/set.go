package utils

import (
	"fmt"
)

// Set is implemented as map[T]bool
type Set[T comparable] struct {
	items map[T]bool
}

// Create new Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

// Add item to Set
func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

// Check if item is in Set
func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

// Remove item from Set if it exists
func (s *Set[T]) Remove(item T) {
	if !s.Contains(item) {
		msg := fmt.Sprint("Attempt to remove item which is not in Set:", item)
		panic(msg)
	}
	delete(s.items, item)
}

// Pop random element from the Set
func (s *Set[T]) Pop() T {
	for k := range s.items {
		delete(s.items, k)
		return k
	}
	panic("Pop from empty Set")
}

// Size of the Set
func (s *Set[T]) Len() int {
	return len(s.items)
}

// Check if Set is empty
func (s *Set[T]) Empty() bool {
	return s.Len() == 0
}

// Iterator for Set to be used in range expressions
func (s *Set[T]) Iter() map[T]bool {
	return s.items
}

// Remove all elements from the Set
func (s *Set[T]) Clear() {
	for k := range s.items {
		delete(s.items, k)
	}
}

// Add all elements from Set `q`
func (s *Set[T]) Update(q *Set[T]) {
	UpdateMap(s.items, q.items)
}
