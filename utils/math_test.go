package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"aoc2022/utils"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, utils.Abs(5))
	assert.Equal(t, 5, utils.Abs(-5))
	assert.Equal(t, 0, utils.Abs(0))
	assert.Equal(t, 5.2, utils.Abs(5.2))
	assert.Equal(t, 5.2, utils.Abs(-5.2))
}

func TestSign(t *testing.T) {
	assert.Equal(t, 1, utils.Sign(4))
	assert.Equal(t, 1, utils.Sign(4.5))
	assert.Equal(t, -1, utils.Sign(-4))
	assert.Equal(t, -1, utils.Sign(-4.5))
	assert.Equal(t, 0, utils.Sign(0))
	assert.Equal(t, 0, utils.Sign(0.0))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, utils.MinSlice([]int{2, 5, 2, 1, 7, 4}))
	assert.Panics(t, func() {utils.MinSlice([]int{})})
	assert.Equal(t, 1, utils.Min(2, 5, 2, 1, 7, 4))
	assert.Panics(t, func() {utils.Min[int]()})
}

func TestMax(t *testing.T) {
	assert.Equal(t, 7, utils.MaxSlice([]int{2, 5, 2, 1, 7, 4}))
	assert.Panics(t, func() {utils.MaxSlice([]int{})})
	assert.Equal(t, 7, utils.Max(2, 5, 2, 1, 7, 4))
	assert.Panics(t, func() {utils.Max[int]()})
}

func TestSegs(t *testing.T) {
	s1 := utils.Segment{Begin: 0, End: 10}
	s2 := utils.Segment{Begin: 1, End: 9}
	s3 := utils.Segment{Begin: -4, End: 5}
	s4 := utils.Segment{Begin: 5, End: 15}
	s5 := utils.Segment{Begin: -10, End: -5}

	assert.True(t, s1.Contains(&s2))
	assert.True(t, s1.Intersects(&s2))
	assert.False(t, s1.Contains(&s3))
	assert.True(t, s1.Intersects(&s3))
	assert.False(t, s1.Contains(&s4))
	assert.True(t, s1.Intersects(&s4))
	assert.True(t, s3.Intersects(&s4))
	assert.True(t, s4.Intersects(&s3))
	assert.False(t, s3.Contains(&s4))
	assert.False(t, s4.Contains(&s3))
	assert.False(t, s1.Contains(&s5))
	assert.False(t, s5.Contains(&s1))
	assert.False(t, s1.Intersects(&s5))
	assert.False(t, s5.Intersects(&s1))
}

func TestPosMethods(t *testing.T) {
	p1 := utils.Pos[int]{X: 1, Y: 2}
	q1 := p1
	p2 := utils.Pos[int]{X: 4, Y: 6}

	ps := utils.Pos[int]{5, 8}

	assert.Equal(t, ps, p2.Plus(&p1))
	assert.Equal(t, ps, p1.Plus(&p2))

	assert.Equal(t, utils.Pos[int]{3, 4}, p2.Minus(&p1))
	assert.Equal(t, utils.Pos[int]{-3, -4}, p1.Minus(&p2))

	assert.Equal(t, utils.Pos[int]{3, 6}, p1.Prod(3))

	p1.Add(&p2)
	assert.Equal(t, ps, p1)

	p1.Sub(&p2)
	assert.Equal(t, q1, p1)

	p1.Mul(3)
	assert.Equal(t, utils.Pos[int]{3, 6}, p1)
}
