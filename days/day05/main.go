package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type Command struct {
	count, from, to int
}

type StrStack = utils.Stack[string]

func (c *Command) apply(stacks []StrStack, is9001 bool) {
	items := stacks[c.from].PopN(c.count)
	if is9001 {
		stacks[c.to].PushN(items)
	} else {
		for i := 0; i < len(items); i++ {
			stacks[c.to].Push(items[len(items)-1-i])
		}
	}
}

func apply(commands []Command, stacks []StrStack, is9001 bool) {
	for _, cmd := range commands {
		cmd.apply(stacks, is9001)
	}
}

func getTop(stacks []StrStack) string {
	return utils.SumValue(stacks, func(s StrStack) string {
		return s.Top()
	})
}

func prepare(lines []string) ([]StrStack, []Command) {
	i_empty := utils.FindIndexIf(lines, func(line string) bool {
		return len(line) == 0
	})

	ids_str := strings.Trim(lines[i_empty-1], " ")
	ids := strings.Split(ids_str, "   ")

	stacks := make([]StrStack, len(ids))
	for j := i_empty - 2; j >= 0; j-- {
		for c := 0; c < utils.Min(len(ids), (len(lines[j])+1)/4); c += 1 {
			col := lines[j][4*c+1]
			if col != ' ' {
				stacks[c].Push(string(col))
			}
		}
	}

	commands := make([]Command, 0)
	for ci := i_empty + 1; ci < len(lines); ci++ {
		parts := strings.Split(lines[ci], " ")
		count, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])
		cmd := Command{count, from - 1, to - 1}
		commands = append(commands, cmd)
	}

	return stacks, commands
}

func solve(lines []string, is9001 bool) string {
	stacks, commands := prepare(lines)
	apply(commands, stacks, is9001)
	return getTop(stacks)
}

func part_1(lines []string) {
	ans := solve(lines, false)
	utils.CheckTask(1, ans, "SHQWSRBDL")
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(lines []string) {
	ans := solve(lines, true)
	utils.CheckTask(2, ans, "CDTQZHBRS")
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day05/test1.txt"
	inputFile := "inputs/day05/input.txt"
	lines := utils.ReadFile(inputFile)

	part_1(lines)
	part_2(lines)
}
