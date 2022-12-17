package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2022/utils"
)

func prepare(lines []string) (calories []int) {
	calories = utils.Transform(strings.Split(strings.Join(lines, "\n"), "\n\n"), func(s string)int{
		return utils.Sum(utils.Transform(strings.Split(s, "\n"), func(c string)int{
			energy, _ := strconv.Atoi(c)
			return energy
		}))
	})
	return
}

func part_1(calories []int) {
	ans := utils.MaxSlice(calories)
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
