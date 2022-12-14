package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type IntPoint = utils.Pos[int]

type Cave struct {
	Rocks      *utils.Set[IntPoint]
	Sand       *utils.Set[IntPoint]
	MaxY       int
	Floor      int
	SandSource IntPoint
	SandMoves  []IntPoint
}

func NewCave() *Cave {
	return &Cave{
		Rocks:      utils.NewSet[IntPoint](),
		Sand:       utils.NewSet[IntPoint](),
		MaxY:       0,
		Floor:      0,
		SandSource: IntPoint{X: 500, Y: 0},
		SandMoves: []IntPoint{
			{X: 0, Y: 1}, {X: -1, Y: 1}, {X: 1, Y: 1},
		},
	}
}

func (c *Cave) ClearSand() {
	c.Sand.Clear()
}

func (c *Cave) AddRocks(p1, p2 IntPoint) {
	delta := p2.Minus(&p1)
	dp := IntPoint{X: utils.Sign(delta.X), Y: utils.Sign(delta.Y)}
	for p := p1; p != p2; p.Add(&dp) {
		c.Rocks.Add(p)
	}
	c.Rocks.Add(p2)
	c.MaxY = utils.Max(c.MaxY, p1.Y, p2.Y)
}

func prepare(lines []string) (cave *Cave) {
	cave = NewCave()
	for _, line := range lines {
		points := utils.Transform(strings.Split(line, " -> "), func(s string) (p IntPoint) {
			c := strings.Split(s, ",")
			p.X, _ = strconv.Atoi(c[0])
			p.Y, _ = strconv.Atoi(c[1])
			return
		})
		for i := 0; i < len(points)-1; i++ {
			cave.AddRocks(points[i], points[i+1])
		}
	}
	cave.Floor = cave.MaxY + 1
	return
}

func solve(cave *Cave, maxy, start int, stop func(p IntPoint) bool) (ans int) {
	cave.Sand.Clear()
	for ans = start; ; {
		p := cave.SandSource
		for np := p; ; {
			for _, dp := range cave.SandMoves {
				if q := p.Plus(&dp); !cave.Rocks.Contains(q) && !cave.Sand.Contains(q) && q.Y <= maxy {
					np = q
					break
				}
			}
			if np == p {
				break
			}
			p = np
		}
		ans += 1
		cave.Sand.Add(p)
		if stop(p) {
			break
		}
	}
	return
}

func part_1(cave *Cave) {
	ans := solve(cave, cave.MaxY, -1, func(p IntPoint) bool { return p.Y == cave.MaxY })
	utils.CheckTask(1, ans, 763)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(cave *Cave) {
	ans := solve(cave, cave.Floor, 0, func(p IntPoint) bool { return p == cave.SandSource })
	utils.CheckTask(2, ans, 23921)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day14/test1.txt"
	inputFile := "inputs/day14/input.txt"
	lines := utils.ReadFile(inputFile)
	cave := prepare(lines)
	part_1(cave)
	part_2(cave)
}
