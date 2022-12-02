package main

import (
	"fmt"
	_ "log"

	"aoc2022/utils"
)

func prepare(lines []string) (input []int) {
	return
}

func part_1(input []int) {
	ans := 0
	// solve here

	//if ans != 0 {
	//	log.Fatal("Wrong answer at part 1: ", ans, " (correct: <UNKNOWN>)")
	//}
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(input []int) {
	ans := 0
	// solve here

	//if ans != 0 {
	//	log.Fatal("Wrong answer at part 2: ", ans, " (correct: <UNKNOWN>)")
	//}
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/@@dayXX@@/test1.txt"
	inputFile := "inputs/@@dayXX@@/input.txt"
	data := utils.ReadFile(inputFile)
	input := prepare(data)
	part_1(input)
	part_2(input)
}
