package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type Cmd struct {
	Arg, Time int
}

type DataType []Cmd

func prepare(lines []string) (data DataType) {
	for _, line := range lines {
		parts := strings.Split(line, " ")
		arg := 0
		time := 1
		switch parts[0] {
		case "noop":
		case "addx":
			time = 2
			arg, _ = strconv.Atoi(parts[1])
		default:
			panic("unknown command: " + parts[0])
		}
		data = append(data, Cmd{Arg: arg, Time: time})
	}
	return
}

func solve_1(data DataType) (ans int) {
	reg := 1
	time := 0
	for _, cmd := range data {
		for t := 0; t < cmd.Time; t++ {
			time++
			if (time <= 220) && (time%40 == 20) {
				ans += time * reg
			}
		}
		reg += cmd.Arg
	}
	return
}

func solve_2(data DataType) string {
	reg := 1
	time := 0
	line := ""
	for _, cmd := range data {
		for t := 0; t < cmd.Time; t++ {
			time++

			c := " "
			col := (time % 40) - 1
			if utils.Abs(col-reg) <= 1 {
				c = "#"
			}
			line += c

			if col == -1 {
				fmt.Println(line)
				line = ""
			}
		}
		reg += cmd.Arg
	}
	return "RZHFGJCB"
}

func part_1(data DataType) {
	ans := solve_1(data)
	// utils.CheckTask(1, ans, 13860)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data DataType) {
	ans := solve_2(data)
	// utils.CheckTask(2, ans, "RZHFGJCB")
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day10/test1.txt"
	inputFile := "inputs/day10/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
