package utils_test

import (
	// "fmt"
	"testing"

	"aoc2022/utils"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := utils.NewSet[int]()

	assert.Panics(t, func() {s.Pop()})

	s.Add(5)
	assert.Panics(t, func(){ s.Remove(1)})

	assert.True(t, s.Contains(5))
	assert.True(t, !s.Empty())
	assert.Equal(t, 1, s.Len())
	x := s.Pop()
	assert.Equal(t, 5, x)
	assert.True(t, s.Empty())
	s.Add(6)
	s.Remove(6)
	assert.True(t, s.Empty())

	s.Add(2)
	s.Add(3)
	sum := 0
	for k := range s.Iter() {
		sum += k
	}
	assert.Equal(t, 5, sum)

	q := utils.NewSet[int]()
	q.Add(4)
	s.Update(q)
	assert.True(t, s.Contains(4))
}
