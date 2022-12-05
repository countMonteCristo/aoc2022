package main

import (
	"fmt"

	"aoc2022/utils"
)

func prepare(lines []string) (data []int) {
	return
}

func part_1(input []int) {
	ans := 0
	// solve here

	utils.CheckTask(1, ans, 0)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(input []int) {
	ans := 0
	// solve here

	utils.CheckTask(2, ans, 0)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/@@dayXX@@/test1.txt"
	inputFile := "inputs/@@dayXX@@/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
