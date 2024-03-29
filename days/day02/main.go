package main

import (
	"fmt"
	"strings"

	"aoc2022/utils"
)

type ScoreType = map[string]int
type StrategyType = map[string]string

var score = map[string]ScoreType{
	"A": {"X": 3 + 1, "Y": 6 + 2, "Z": 0 + 3},
	"B": {"X": 0 + 1, "Y": 3 + 2, "Z": 6 + 3},
	"C": {"X": 6 + 1, "Y": 0 + 2, "Z": 3 + 3},
}

var rule = map[string]StrategyType{
	"A": {"X": "Z", "Y": "X", "Z": "Y"},
	"B": {"X": "X", "Y": "Y", "Z": "Z"},
	"C": {"X": "Y", "Y": "Z", "Z": "X"},
}

func getScore(lines []string, fn func([]string) int) int {
	return utils.SumValue(lines, func(line string)int{
		round := strings.Split(line, " ")
		return fn(round)
	})
}

func part_1(input []string) {
	ans := getScore(
		input,
		func(round []string) int {
			you := round[1]
			return score[round[0]][you]
		},
	)
	utils.CheckTask(1, ans, 15422)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(input []string) {
	ans := getScore(
		input,
		func(round []string) int {
			you := rule[round[0]][round[1]]
			return score[round[0]][you]
		},
	)
	utils.CheckTask(2, ans, 15442)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day02/test1.txt"
	inputFile := "inputs/day02/input.txt"
	data := utils.ReadFile(inputFile)
	part_1(data)
	part_2(data)
}
