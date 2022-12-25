package utils_test

import (
	"testing"

	"aoc2022/utils"

	"github.com/stretchr/testify/assert"
)

type Fruit struct {
	name     string
	priority int
}

func (i Fruit) LessThan(j utils.PQItem) bool {
	return i.priority < j.(Fruit).priority
}

func TestPriorityQueue(t *testing.T) {
	pq := utils.NewPq[Fruit]()
	assert.True(t, pq.Empty())

	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	for name, pr := range items {
		pq.Push(&Fruit{
			name: name, priority: pr,
		})
	}

	assert.Equal(t, 3, pq.Len())

	item := pq.Pop()
	assert.Equal(t, "apple", item.name)
	assert.Equal(t, 2, item.priority)
}
