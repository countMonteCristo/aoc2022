package main

import (
	"fmt"
	"sort"
	"strconv"

	"aoc2022/utils"
)

func prepare(lines []string) (calories []int) {
	current := 0
	for _, energy_str := range lines {
		if len(energy_str) == 0 {
			calories = append(calories, current)
			current = 0
		} else {
			energy, _ := strconv.Atoi(energy_str)
			current += energy
		}
	}
	calories = append(calories, current)
	return
}

func part_1(calories []int) {
	ans := 0
	for _, q := range calories {
		if q >= ans {
			ans = q
		}
	}
	utils.CheckTask(1, ans, 72070)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(calories []int) {
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	ans := calories[0] + calories[1] + calories[2]
	utils.CheckTask(2, ans, 211805)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day01/test1.txt"
	inputFile := "inputs/day01/input.txt"
	data := utils.ReadFile(inputFile)
	calories := prepare(data)
	part_1(calories)
	part_2(calories)
}
