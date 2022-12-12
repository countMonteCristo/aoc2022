package main

import (
	"container/heap"
	"fmt"

	"aoc2022/utils"
)

//
// PriorityQueue has been stolen from https://pkg.go.dev/container/heap
//

// An Item is something we manage in a priority queue.
type Item struct {
	path     []IntPos // The value of the item; arbitrary.
	priority int      // The priority of the item in the queue.
	index    int      // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority+len(pq[i].path) < pq[j].priority+len(pq[j].path)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// ---------------------------------------------------------------------------------------------------------------------

func prepare(lines []string) (field *Field) {
	field = &Field{
		data: lines, size: IntPos{X: len(lines[0]), Y: len(lines)},
		elev: make(map[byte]int),
	}
	for i, line := range field.data {
		for j, c := range line {
			ec := c
			if c == 'E' {
				field.e = IntPos{X: j, Y: i}
				ec = 'z'
			}
			if c == 'S' {
				field.s = IntPos{X: j, Y: i}
				ec = 'a'
			}
			field.elev[line[j]] = int(ec) - int('a')
		}
	}
	return
}

type IntPos = utils.Pos[int]

var D = []IntPos{
	{X: -1, Y: 0}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1},
}

type Field struct {
	data []string
	s    IntPos
	e    IntPos
	size IntPos
	elev map[byte]int
}

func (f *Field) elevation(p IntPos) int {
	return f.elev[f.data[p.Y][p.X]]
}

func (f *Field) contains(p IntPos) bool {
	return p.X >= 0 && p.X < f.size.X && p.Y >= 0 && p.Y < f.size.Y
}

func nbrs(p IntPos, field *Field) []IntPos {
	neighbours := make([]IntPos, 0, 4)
	e := field.elevation(p)
	for _, dp := range D {
		np := p.Plus(&dp)
		if field.contains(np) {
			ne := field.elevation(np)
			if ne-e <= 1 {
				neighbours = append(neighbours, np)
			}
		}
	}
	return neighbours
}

// A* has been stolen from https://ru.wikipedia.org/wiki/A*
func astar(field *Field, S, E IntPos) int {
	closed := make(map[IntPos]bool)
	open := make(PriorityQueue, 0)
	heap.Init(&open)

	path_start := &Item{
		path: []IntPos{S}, priority: utils.Manhattan(S, E),
	}
	heap.Push(&open, path_start)

	for open.Len() > 0 {
		p := heap.Pop(&open).(*Item)
		p_last := p.path[len(p.path)-1]

		_, exists := closed[p_last]
		if exists {
			continue
		}

		if p_last == E {
			return len(p.path) - 1
		}
		closed[p_last] = true

		for _, np := range nbrs(p_last, field) {
			tp := make([]IntPos, len(p.path))
			copy(tp, p.path)
			ni := &Item{
				path:     append(tp, np),
				priority: utils.Manhattan(np, E),
			}
			heap.Push(&open, ni)
		}
	}
	return -100
}

func solve_1(field *Field) (ans int) {
	ans = astar(field, field.s, field.e)
	return
}

func solve_2(field *Field) (ans int) {
	ans = 99999
	for i, line := range field.data {
		for j, c := range line {
			if c == 'S' || c == 'a' {
				x := astar(field, IntPos{X: j, Y: i}, field.e)
				if x < 0 {
					continue
				}
				ans = utils.Min(x, ans)
			}
		}
	}
	return
}

func part_1(field *Field) {
	ans := solve_1(field)
	utils.CheckTask(1, ans, 437)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(field *Field) {
	ans := solve_2(field)
	utils.CheckTask(2, ans, 430)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day12/test1.txt"
	inputFile := "inputs/day12/input.txt"
	lines := utils.ReadFile(inputFile)
	field := prepare(lines)
	part_1(field)
	part_2(field)
}
