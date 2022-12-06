package main

import (
	"fmt"

	"aoc2022/utils"
)

func prepare(lines []string) (data string) {
	data = lines[0]
	return
}

func countDiff(str string) int {
	count := map[rune]int{}
	for _, c := range str {
		count[c] += 1
	}
	return len(count)
}

func solve(str string, count int) int {
	for i := count; i < len(str); i++ {
		if countDiff(str[i-count:i]) == count {
			return i
		}
	}
	return -1
}

func part_1(data string) {
	ans := solve(data, 4)
	utils.CheckTask(1, ans, 1198)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data string) {
	ans := solve(data, 14)
	utils.CheckTask(2, ans, 3120)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day06/test1.txt"
	// inputFile := "inputs/day06/test2.txt"
	// inputFile := "inputs/day06/test3.txt"
	// inputFile := "inputs/day06/test4.txt"
	// inputFile := "inputs/day06/test5.txt"
	inputFile := "inputs/day06/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
