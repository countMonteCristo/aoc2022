package utils_test

import (
	"testing"

	"aoc2022/utils"

	"github.com/stretchr/testify/assert"
)

func TestCountIf(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	count_got := utils.CountIf(data, func(x int) bool { return x%2 == 0 })
	assert.Equal(t, 3, count_got, "CountIf test failed")
}

func TestSumIf(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	sumif_got := utils.SumIf(data, func(x int) bool { return x%2 == 0 })
	assert.Equal(t, 12, sumif_got, "SumIf test failed")

	sum_got := utils.Sum(data)
	assert.Equal(t, 28, sum_got, "Sum test failed")

	sumv_got := utils.SumValue(data, func(x int) int { return x * x })
	assert.Equal(t, 140, sumv_got, "SumValue test failed")
}

func TestTransform(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	res_want := []int{1, 4, 9, 16, 25, 36, 49}

	res_got := utils.Transform(data, func(x int) int { return x * x })
	assert.EqualValues(t, res_want, res_got, "Transform test failed")
}

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	even := []int{2, 4, 6}

	res_got := utils.Filter(data, func(x int) bool { return x%2 == 0})
	assert.EqualValues(t, even, res_got)
}

func TestPreds(t *testing.T) {
	assert.Equal(t, 5, utils.Id(5))
	assert.Equal(t, "hello", utils.Id("hello"))

	assert.Equal(t, true, utils.True(5))
	assert.Equal(t, true, utils.True("hello"))

	assert.Equal(t, false, utils.False(5))
	assert.Equal(t, false, utils.False("hello"))
}

func TestUpdateMap(t *testing.T) {
	m1 := map[string]int{
		"a": 1, "b": 2, "c": 3,
	}
	m2 := map[string]int{
		"b":2, "c":4, "d": 5,
	}

	utils.UpdateMap(m1, m2)
	xa, exists := m1["a"]
	assert.True(t, exists)
	assert.Equal(t, xa, 1)
	xb, exists := m1["b"]
	assert.True(t, exists)
	assert.Equal(t, xb, 2)
	xc, exists := m1["c"]
	assert.True(t, exists)
	assert.Equal(t, xc, 4)
	xd, exists := m1["d"]
	assert.True(t, exists)
	assert.Equal(t, xd, 5)
}
