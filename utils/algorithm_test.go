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

	countp_got := utils.CountIfP(data, func(x *int) bool { return *x%2 == 1 })
	assert.Equal(t, 4, countp_got, "CountIf test failed")
}

func TestSumIf(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}

	sumif_got := utils.SumIf(data, func(x int) bool { return x%2 == 0 })
	assert.Equal(t, 12, sumif_got, "SumIf test failed")

	sumifp_got := utils.SumIfP(data, func(x *int) bool { return *x%2 == 1 })
	assert.Equal(t, 16, sumifp_got, "SumIfP test failed")

	sum_got := utils.Sum(data)
	assert.Equal(t, 28, sum_got, "Sum test failed")

	sumv_got := utils.SumValue(data, func(x int) int { return x * x })
	assert.Equal(t, 140, sumv_got, "SumValue test failed")

	sumvp_got := utils.SumValueP(data, func(x *int) int { return *x * *x })
	assert.Equal(t, 140, sumvp_got, "SumValueP test failed")
}

func TestTransform(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	res_want := []int{1, 4, 9, 16, 25, 36, 49}

	res_got := utils.Transform(data, func(x int) int { return x * x })
	assert.EqualValues(t, res_want, res_got, "Transform test failed")

	resp_got := utils.TransformP(data, func(x *int) int { return *x * *x })
	assert.EqualValues(t, resp_got, res_got, "TransformP test failed")
}

func TestPreds(t *testing.T) {
	assert.Equal(t, 5, utils.Id(5))
	assert.Equal(t, "hello", utils.Id("hello"))

	assert.Equal(t, true, utils.True(5))
	assert.Equal(t, true, utils.True("hello"))

	assert.Equal(t, false, utils.False(5))
	assert.Equal(t, false, utils.False("hello"))
}
