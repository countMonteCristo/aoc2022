package main

import (
	"fmt"

	"aoc2022/utils"
)


func prepare(lines []string) (elves IpSet) {
	elves = NewIpSet()
	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				elves.Add(IntPoint{X: j, Y: i})
			}
		}
	}
	return
}

var DP = []IntPoint{
	{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
	{X: -1, Y: 0}, {X: 1, Y: 0},
	{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
}

var ChoicesIds = [][3]int {
	{0, 1, 2},
	{5, 6, 7},
	{0, 3, 5},
	{2, 4, 7},
}

func HasNeighbours(elves IpSet, elf IntPoint) bool {
	for _, d := range DP {
		if elves.Contains(elf.Plus(d)) {
			return true
		}
	}
	return false
}

func HasFreePlaceAtIds(elves IpSet, elf IntPoint, ids [3]int) bool {
	for _, i := range ids {
		if elves.Contains(elf.Plus(DP[i])) {
			return false
		}
	}
	return true
}

func diffuseIpSet(elves IpSet, max_steps int) (IpSet, int) {
	choice_id := 0
	step := 1

	for ;;step++ {
		next := NewIpSet()
		proposes := make(map[IntPoint]IntPoint)
		for e := range elves.Iter() {
			if !HasNeighbours(elves, e) {
				next.Add(e)
				continue
			}

			intent := e
			for i := choice_id; i<choice_id+len(ChoicesIds); i++ {
				ids := ChoicesIds[i % len(ChoicesIds)]
				if HasFreePlaceAtIds(elves, e, ids) {
					intent = e.Plus(DP[ids[1]])
					break
				}
			}

			pe, has := proposes[intent]
			if has {
				delete(proposes, intent)
				next.Add(pe)
				next.Add(e)
			} else {
				proposes[intent] = e
			}
		}

		if next.Len() == elves.Len() || step > max_steps {
			break
		}

		for intent := range proposes {
			next.Add(intent)
		}

		elves = next
		choice_id = (choice_id + 1)%len(ChoicesIds)
	}
	return elves, step
}

func solve(elves IpSet, max_steps int, part1 bool) (ans int) {
	elves, steps := diffuseIpSet(elves, max_steps)

	if part1 {
		var xmin, xmax, ymin, ymax int
		for e := range elves.Iter() {
			xmin, xmax, ymin, ymax = e.X, e.X, e.Y, e.Y
			break
		}
		for e := range elves.Iter() {
			xmin, xmax = utils.Min(xmin, e.X), utils.Max(xmax, e.X)
			ymin, ymax = utils.Min(ymin, e.Y), utils.Max(ymax, e.Y)
		}
		ans = (xmax-xmin+1)*(ymax-ymin+1) - elves.Len()
	} else {
		ans = steps
	}

	return
}

func part_1(elves IpSet) {
	ans := solve(elves, 10, true)
	utils.CheckTask(1, ans, 3906)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(elves IpSet) {
	ans := solve(elves, 9999, false)
	utils.CheckTask(2, ans, 895)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day23/test1.txt"
	// inputFile := "inputs/day23/test2.txt"
	inputFile := "inputs/day23/input.txt"
	lines := utils.ReadFile(inputFile)
	elves := prepare(lines)
	part_1(elves)
	part_2(elves)
}
