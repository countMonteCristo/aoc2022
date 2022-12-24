package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type Rope struct {
	Body []IntPoint
}

func NewRope(body_size int) *Rope {
	return &Rope{
		Body: make([]IntPoint, body_size),
	}
}

func (r *Rope) Head() *IntPoint {
	return &r.Body[0]
}

func (r *Rope) Tail() *IntPoint {
	return &r.Body[r.Len()-1]
}

func (r *Rope) Len() int {
	return len(r.Body)
}

func (r *Rope) apply(cmd *Cmd) IpSet {
	dp := cmd.GetDirection()

	visited := NewIpSet()
	for i := 0; i < cmd.Count; i++ {
		r.Head().Add(dp)

		for j := 1; j < r.Len(); j++ {
			diff := r.Body[j-1].Minus(r.Body[j])

			d := utils.Max(utils.Abs(diff.X), utils.Abs(diff.Y))
			if d <= 1 { // do not move tail, because current knot is not moving
				break
			}

			dq := IntPoint{
				X: utils.Sign(diff.X), Y: utils.Sign(diff.Y),
			}
			r.Body[j].Add(dq)
		}

		visited.Add(*r.Tail())
	}
	return visited
}

func (rope *Rope) print(topLeft, size IntPoint) {
	for r := topLeft.Y; r <= topLeft.Y+size.Y; r++ {
		for c := topLeft.X; c <= topLeft.X+size.X; c++ {
			p := IntPoint{X: c, Y: r}
			j := "*"
			for k, q := range rope.Body {
				if p == q {
					j = "H"
					if k > 0 {
						j = strconv.Itoa(k)
					}

					break
				}
			}
			fmt.Print(j)
		}
		fmt.Println()
	}
	fmt.Println()
}

type Cmd struct {
	Dir   string
	Count int
}

var CmdDirMap = map[string]IntPoint{
	"U": {X: 0, Y: -1}, "D": {X: 0, Y: 1},
	"L": {X: -1, Y: 0}, "R": {X: 1, Y: 0},
}

func (cmd *Cmd) GetDirection() (d IntPoint) {
	d, exists := CmdDirMap[cmd.Dir]
	if !exists {
		panic("Unknown direction: " + cmd.Dir)
	}
	return
}

type DataType []Cmd

func prepare(lines []string) (data DataType) {
	data = utils.Transform(lines, func(line string) Cmd {
		parts := strings.Split(line, " ")
		count := StrToInt(parts[1])
		return Cmd{Dir: parts[0], Count: count}
	})
	return
}

func solve(data DataType, rope_len int) int {
	r := NewRope(rope_len)
	visited := NewIpSet()
	visited.Add(*r.Tail())

	for _, cmd := range data {
		visited.Update(r.apply(&cmd))
	}
	return visited.Len()
}

func part_1(data DataType) {
	ans := solve(data, 2)
	utils.CheckTask(1, ans, 5735)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data DataType) {
	ans := solve(data, 10)
	utils.CheckTask(2, ans, 2478)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day09/test1.txt"
	// inputFile := "inputs/day09/test2.txt"
	// inputFile := "inputs/day09/test3.txt"
	inputFile := "inputs/day09/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
