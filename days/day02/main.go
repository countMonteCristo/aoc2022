package main

import (
	"fmt"
	"log"
	"strings"

	"aoc2022/utils"
)

type ScoreType = map[string]int
type StrategyType = map[string]string

var score = map[string]ScoreType {
	"A": {"X": 3+1, "Y": 6+2, "Z": 0+3},
	"B": {"X": 0+1, "Y": 3+2, "Z": 6+3},
	"C": {"X": 6+1, "Y": 0+2, "Z": 3+3},
}

var rule = map[string]StrategyType {
	"A": {"X": "Z", "Y": "X", "Z": "Y"},
	"B": {"X": "X", "Y": "Y", "Z": "Z"},
	"C": {"X": "Y", "Y": "Z", "Z": "X"},
}

func getScore(lines []string, fn func([]string) int) (total int) {
	for _, line := range lines {
		round := strings.Split(line, " ")
		total += fn(round)
	}
	return
}

func part_1(input []string) {
	ans := getScore(
		input,
		func(round []string) int {
			you := round[1]
			return score[round[0]][you]
		},
	)
	if ans != 15422 {
		log.Fatal("Wrong answer at part 1: ", ans, " (correct: 15422)")
	}
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
	if ans != 15442 {
		log.Fatal("Wrong answer at part 2: ", ans, " (correct: 15442)")
	}
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day02/test1.txt"
	inputFile := "inputs/day02/input.txt"
	data := utils.ReadFile(inputFile)
	part_1(data)
	part_2(data)
}
