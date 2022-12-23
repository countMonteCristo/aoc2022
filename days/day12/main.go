package main

import (
	"fmt"

	"aoc2022/utils"
)

// ---------------------------------------------------------------------------------------------------------------------

func prepare(lines []string) (field *Field) {
	field = &Field{
		data: lines, size: IntPoint{X: len(lines[0]), Y: len(lines)},
		elev: make(map[byte]int),
	}
	for i, line := range field.data {
		for j, c := range line {
			ec := c
			if c == 'E' {
				field.e = IntPoint{X: j, Y: i}
				ec = 'z'
			}
			if c == 'S' {
				field.s = IntPoint{X: j, Y: i}
				ec = 'a'
			}
			field.elev[line[j]] = int(ec) - int('a')
		}
	}
	return
}

type IntPoint = utils.Point2d[int]

var D = []IntPoint{
	{X: -1, Y: 0}, {X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1},
}

type Field struct {
	data []string
	s    IntPoint
	e    IntPoint
	size IntPoint
	elev map[byte]int
}

func (f *Field) elevation(p IntPoint) int {
	return f.elev[f.data[p.Y][p.X]]
}

func (f *Field) contains(p IntPoint) bool {
	return p.X >= 0 && p.X < f.size.X && p.Y >= 0 && p.Y < f.size.Y
}

func nbrs(p IntPoint, field *Field) []IntPoint {
	neighbours := make([]IntPoint, 0, len(D))
	e := field.elevation(p)
	for _, dp := range D {
		np := p.Plus(dp)
		if field.contains(np) {
			ne := field.elevation(np)
			if ne-e <= 1 {
				neighbours = append(neighbours, np)
			}
		}
	}
	return neighbours
}

type Path struct {
	points    []IntPoint // all path points
	heuristic int        // minimum distance to the destination point
}

func (p Path) LessThan(j utils.PQItem) bool {
	q := j.(Path)
	return len(p.points)+p.heuristic < len(q.points)+q.heuristic
}

// A* has been stolen from https://ru.wikipedia.org/wiki/A*
func astar(field *Field, S, E IntPoint) int {
	visited := utils.NewSet[IntPoint]()

	queue := utils.NewPq[Path]()
	queue.Push(&Path{
		points: []IntPoint{S}, heuristic: utils.Manhattan(S, E),
	})

	for !queue.Empty() {
		item := queue.Pop()

		last := item.points[len(item.points)-1]
		if visited.Contains(last) {
			continue
		}

		if last == E {
			return len(item.points) - 1
		}
		visited.Add(last)

		for _, np := range nbrs(last, field) {
			temp_path := make([]IntPoint, len(item.points), len(item.points)+1)
			copy(temp_path, item.points)
			new_item := &Path{
				points:    append(temp_path, np),
				heuristic: utils.Manhattan(np, E),
			}
			queue.Push(new_item)
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
				x := astar(field, IntPoint{X: j, Y: i}, field.e)
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
