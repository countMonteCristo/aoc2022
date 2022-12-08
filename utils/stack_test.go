package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"aoc2022/utils"
)

func TestStack(t *testing.T) {
	stack := utils.Stack[int]{}

	assert.Panics(t, func() {stack.Pop()})
	assert.Panics(t, func() {stack.Top()})
	assert.Panics(t, func() {stack.PopN(3)})

	stack.Push(5)
	assert.Equal(t, 5, stack.Top())
	stack.Push(6)
	assert.Equal(t, 6, stack.Top())

	x := stack.Pop()
	assert.Equal(t, 6, x)
	assert.Equal(t, 5, stack.Top())

	stack.PushN([]int{3, 4})
	assert.Equal(t, 4, stack.Top())
	assert.EqualValues(t, []int{3, 4}, stack.PopN(2))
}
