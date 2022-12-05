package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type Stack struct {
	data []string
}

func (s *Stack) Push(e string) {
	s.data = append(s.data, e)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		panic("Pop from empty stack")
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top
}

func (s *Stack) Top() string {
	return s.data[len(s.data)-1]
}

type Command struct {
	count, from, to int
}

func (c *Command) apply(stacks []Stack) {
	for n := 0; n < c.count; n++ {
		item := stacks[c.from].Pop()
		stacks[c.to].Push(item)
	}
}

func (c *Command) apply2(stacks []Stack) {
	ts := Stack{}
	for n := 0; n < c.count; n++ {
		item := stacks[c.from].Pop()
		ts.Push(item)
	}
	for n := 0; n < c.count; n++ {
		stacks[c.to].Push(ts.Pop())
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func prepare(lines []string) ([]Stack, []Command) {
	i_empty := 0
	for i, line := range lines {
		if len(line) == 0 {
			i_empty = i
			break
		}
	}
	ids_str := strings.Trim(lines[i_empty-1], " ")
	ids := strings.Split(ids_str, "   ")

	stacks := make([]Stack, len(ids))
	for j := i_empty - 2; j >= 0; j-- {
		for c := 0; c < min(len(ids), (len(lines[j])+1)/4); c += 1 {
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

func part_1(lines []string) {
	stacks, commands := prepare(lines)
	ans := ""

	for _, cmd := range commands {
		cmd.apply(stacks)
	}

	for _, s := range stacks {
		ans += s.Top()
	}

	utils.CheckTask(1, ans, "SHQWSRBDL")
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(lines []string) {
	stacks, commands := prepare(lines)
	ans := ""

	for _, cmd := range commands {
		cmd.apply2(stacks)
	}

	for _, s := range stacks {
		ans += s.Top()
	}

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
