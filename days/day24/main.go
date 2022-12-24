package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

var DP = []IntPoint{
	{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1}, {X: 0, Y: 0},
}
var DirIdxToStr = map[int]string{
	0: ">", 1: "v", 2: "<", 3: "^",
}
var StrToDirIdx = map[string]int{
	">": 0, "v": 1, "<": 2, "^": 3,
}

// ---------------------------------------------------------------------------------------------------------------------

type Valley struct {
	Blizzards     map[IntPoint][]int
	Empty         IpSet
	Start, Finish IntPoint
	W, H          int
}

func (v *Valley) Wrap(p IntPoint) IntPoint {
	if p.X == 0 {
		p.X = v.W - 2
	}
	if p.X == v.W-1 {
		p.X = 1
	}
	if p.Y == 0 {
		p.Y = v.H - 2
	}
	if p.Y == v.H-1 {
		p.Y = 1
	}
	return p
}

func (v *Valley) Next() Valley {
	nv := Valley{
		Blizzards: make(map[IntPoint][]int),
		Empty:     NewIpSet(),
		Start:     v.Start,
		Finish:    v.Finish,
		W:         v.W,
		H:         v.H,
	}

	for p, dirs := range v.Blizzards {
		for _, dir := range dirs {
			np := v.Wrap(p.Plus(DP[dir]))

			_, exists := nv.Blizzards[np]
			if !exists {
				nv.Blizzards[np] = make([]int, 0)
			}
			nv.Blizzards[np] = append(nv.Blizzards[np], dir)
		}
	}

	nv.Empty.Add(v.Start)
	nv.Empty.Add(v.Finish)
	for i := 1; i < v.H-1; i++ {
		for j := 1; j < v.W-1; j++ {
			p := IntPoint{X: j, Y: i}
			_, exists := nv.Blizzards[p]
			if !exists {
				nv.Empty.Add(p)
			}
		}
	}
	return nv
}

func (v *Valley) Print() {
	fmt.Println(strings.Repeat("#", v.W))
	for i := 1; i < v.H-1; i++ {
		fmt.Print("#")
		for j := 1; j < v.W-1; j++ {
			p := IntPoint{X: j, Y: i}
			c := "."
			blizzards, exists := v.Blizzards[p]
			if exists {
				if len(blizzards) > 1 {
					c = strconv.Itoa(len(blizzards))
				} else {
					c = DirIdxToStr[blizzards[0]]
				}
			}
			fmt.Print(c)
		}
		fmt.Println("#")
	}
	fmt.Println(strings.Repeat("#", v.W))
	fmt.Println()
}

// ---------------------------------------------------------------------------------------------------------------------

func prepare(lines []string) Valley {
	valley := Valley{
		Blizzards: make(map[IntPoint][]int),
		Empty:     NewIpSet(),
		W:         len(lines[0]),
		H:         len(lines),
	}

	for i, line := range lines {
		for j, c := range line {
			p := IntPoint{X: j, Y: i}
			if c == '#' {
				continue
			}
			if c == '.' {
				if i == 0 {
					valley.Start = p
				}
				if i == valley.H-1 {
					valley.Finish = p
				}
				continue
			}

			_, exists := valley.Blizzards[p]
			if !exists {
				valley.Blizzards[p] = make([]int, 0)
			}

			valley.Blizzards[p] = append(valley.Blizzards[p], StrToDirIdx[string(c)])
		}
	}

	valley.Empty.Add(valley.Start)
	valley.Empty.Add(valley.Finish)
	for i := 1; i < valley.H-1; i++ {
		for j := 1; j < valley.W-1; j++ {
			p := IntPoint{X: j, Y: i}
			_, exists := valley.Blizzards[p]
			if !exists {
				valley.Empty.Add(p)
			}
		}
	}

	return valley
}

// ---------------------------------------------------------------------------------------------------------------------

func GetMinTime(valley Valley, from, to IntPoint) (Valley, int) {
	accessable := NewIpSet()
	accessable.Add(from)

	for i := 1; ; i++ {
		new_accessable := NewIpSet()
		valley = valley.Next()

		for p := range accessable.Iter() {
			for _, d := range DP {
				q := p.Plus(d)
				if valley.Empty.Contains(q) {
					new_accessable.Add(q)
				}
			}
		}

		if new_accessable.Contains(to) {
			return valley, i
		}
		accessable = new_accessable
	}
}

func solve(valley Valley) int {
	_, time := GetMinTime(valley, valley.Start, valley.Finish)
	return time
}

func solve2(valley Valley) int {
	valley, time1 := GetMinTime(valley, valley.Start, valley.Finish)
	valley, time2 := GetMinTime(valley, valley.Finish, valley.Start)
	valley, time3 := GetMinTime(valley, valley.Start, valley.Finish)
	return time1 + time2 + time3
}

func part_1(valley Valley) {
	ans := solve(valley)
	utils.CheckTask(1, ans, 334)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(valley Valley) {
	ans := solve2(valley)
	utils.CheckTask(2, ans, 934)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day24/test1.txt"
	inputFile := "inputs/day24/input.txt"
	lines := utils.ReadFile(inputFile)
	valley := prepare(lines)
	part_1(valley)
	part_2(valley)
}
