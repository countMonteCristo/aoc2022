package main

import (
	"fmt"

	"aoc2022/utils"
)

type DataType []int

func prepare(lines []string) (data DataType) {
	return
}

func solve(data DataType) (ans int) {
	// solve here
	return
}

func part_1(data DataType) {
	ans := solve(data)
	// utils.CheckTask(1, ans, 0)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data DataType) {
	ans := solve(data)
	// utils.CheckTask(2, ans, 0)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	inputFile := "inputs/@@dayXX@@/test1.txt"
	// inputFile := "inputs/@@dayXX@@/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
