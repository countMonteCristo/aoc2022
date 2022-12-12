package utils

import (
	"fmt"
)

// Set is implemented as map[T]bool
type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]bool)}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set[T]) Remove(item T) {
	if !s.Contains(item) {
		msg := fmt.Sprint("Attempt to remove item which is not in Set:", item)
		panic(msg)
	}
	delete(s.items, item)
}

func (s *Set[T]) Pop() T {
	for k := range s.items {
		delete(s.items, k)
		return k
	}
	panic("Pop from empty Set")
}

func (s *Set[T]) Len() int {
	return len(s.items)
}

func (s *Set[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Set[T]) Iter() map[T]bool {
	return s.items
}

func (s *Set[T]) Update(q *Set[T]) {
	UpdateMap(s.items, q.items)
}
