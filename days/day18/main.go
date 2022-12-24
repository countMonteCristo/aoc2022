package main

import (
	"fmt"
	"strings"

	"aoc2022/utils"
)

var dp = [][]int{
	{0, 0, 1}, {0, 0, -1}, {0, 1, 0}, {0, -1, 0}, {1, 0, 0}, {-1, 0, 0},
}

type Cube struct {
	X, Y, Z int
}

func NewCube(x, y, z int) Cube {
	return Cube{X: x, Y: y, Z: z}
}

func (c *Cube) IsInRange(cmin, cmax int) bool {
	return cmin <= c.X && c.X <= cmax && cmin <= c.Y && c.Y <= cmax && cmin <= c.Z && c.Z <= cmax
}

func (c *Cube) GetNbrs(cmin, cmax int) *utils.Set[Cube] {
	nbrs := utils.NewSet[Cube]()
	for _, q := range dp {
		n := NewCube(c.X+q[0], c.Y+q[1], c.Z+q[2])
		if n.IsInRange(cmin, cmax) {
			nbrs.Add(n)
		}
	}
	return nbrs
}

type Cubes []Cube

func prepare(lines []string) Cubes {
	return utils.Transform(lines, func(line string) Cube {
		coords := utils.Transform(strings.Split(line, ","), StrToInt)
		return Cube{X: coords[0], Y: coords[1], Z: coords[2]}
	})
}

func solve(cubes Cubes) (ans int) {
	ans = 6 * len(cubes)
	for i := 0; i < len(cubes); i++ {
		c1 := cubes[i]
		for j := i + 1; j < len(cubes); j++ {
			c2 := cubes[j]
			if utils.Abs(c1.X-c2.X)+utils.Abs(c1.Y-c2.Y)+utils.Abs(c1.Z-c2.Z) == 1 {
				ans -= 2
			}
		}
	}
	return
}

func solve2(cubes Cubes) (ans int) {
	cmin := utils.Min(cubes[0].X, cubes[0].Y, cubes[0].Z)
	cmax := utils.Max(cubes[0].X, cubes[0].Y, cubes[0].Z)
	all := utils.NewSet[Cube]()
	for _, c := range cubes {
		all.Add(c)
		cmin = utils.Min(c.X, c.Y, c.Z, cmin)
		cmax = utils.Max(c.X, c.Y, c.Z, cmax)
	}
	cmin--
	cmax++

	start := Cube{X: cmin, Y: cmin, Z: cmin}
	out := utils.NewSet[Cube]()
	edge := utils.NewSet[Cube]()
	edge.Add(start)

	for edge.Len() > 0 {
		e := edge.Pop()
		nbrs := e.GetNbrs(cmin, cmax)
		for n := range nbrs.Iter() {
			if out.Contains(n) {
				continue
			}
			if all.Contains(n) {
				ans += 1
			} else {
				edge.Add(n)
			}
		}
		out.Add(e)
	}
	return
}

func part_1(data Cubes) {
	ans := solve(data)
	utils.CheckTask(1, ans, 3494)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data Cubes) {
	ans := solve2(data)
	utils.CheckTask(2, ans, 2062)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day18/test1.txt"
	inputFile := "inputs/day18/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
