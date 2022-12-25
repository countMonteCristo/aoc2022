package main

import (
	"fmt"
	"strings"
	"unicode"

	"aoc2022/utils"
)

func priority(char rune) (pr int) {
	if unicode.IsUpper(char) {
		pr = int(char) - int('A') + 27
	} else {
		pr = int(char) - int('a') + 1
	}
	return
}

func common(rucks []string) rune {
	for _, c := range rucks[0] {
		char := string(c)
		count := utils.SumValue(rucks, func (r string)int{
			if r != rucks[0] && strings.Contains(r, char){
				return 1
			}
			return 0
		})
		if count == len(rucks)-1 {
			return c
		}
	}
	panic("No common chars found")
}

func process_1(lines []string) int {
	return utils.SumValue(lines, func(line string) int {
		return priority(common([]string{
			line[:len(line)/2],
			line[len(line)/2:],
		}))
	})
}

func process_2(lines []string, ngroups int) (total int) {
	group := make([]string, ngroups)
	for index, line := range lines {
		n := index % ngroups
		group[n] = line
		if n == ngroups-1 {
			c := common(group)
			total += priority(c)
		}
	}
	return
}

func part_1(data []string) {
	ans := process_1(data)
	utils.CheckTask(1, ans, 8233)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data []string) {
	ans := process_2(data, 3)
	utils.CheckTask(2, ans, 2821)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day03/test1.txt"
	inputFile := "inputs/day03/input.txt"
	data := utils.ReadFile(inputFile)
	part_1(data)
	part_2(data)
}
