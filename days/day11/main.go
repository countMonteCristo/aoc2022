package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type Unary = utils.Unary[int]
type Binary = utils.Binary[int]

var Funcs = map[string]Binary{
	"+": func(x, y int) int { return x + y },
	"*": func(x, y int) int { return x * y },
}

func parseArg(a string) (int, bool) {
	arg := 0
	is_old := (a == "old")
	if !is_old {
		arg, _ = strconv.Atoi(a)
	}
	return arg, is_old
}

func getOr(x, y int, flag bool) int {
	res := x
	if flag {
		res = y
	}
	return res
}

func parseFunc(line string) Unary {
	parts := strings.Split(line, " ")
	a1, a1_old := parseArg(parts[2])
	f := Funcs[parts[3]]
	a2, a2_old := parseArg(parts[4])

	return func(old int) int {
		arg1 := getOr(a1, old, a1_old)
		arg2 := getOr(a2, old, a2_old)
		return f(arg1, arg2)
	}
}

type Monkey struct {
	Id          int
	Items       []int
	ThrowTo     Unary
	Operation   Unary
	Inspected   int
	DivisibleBy int
}

func NewMonkey(lines []string) *Monkey {
	id, _ := strconv.Atoi(strings.Split(strings.Trim(lines[0], ":"), " ")[1])

	items := utils.Transform(
		strings.Split(strings.Split(lines[1], ": ")[1], ", "),
		func(i string) int {
			x, _ := strconv.Atoi(strings.Trim(i, " "))
			return x
		},
	)

	op := parseFunc(strings.Split(lines[2], ": ")[1])

	div_by, _ := strconv.Atoi(strings.Split(lines[3], " ")[5])
	m_true, _ := strconv.Atoi(strings.Split(lines[4], " ")[9])
	m_false, _ := strconv.Atoi(strings.Split(lines[5], " ")[9])
	throw_to := func(x int) int {
		if x%div_by == 0 {
			return m_true
		}
		return m_false
	}

	return &Monkey{
		Id: id, Items: items, ThrowTo: throw_to,
		Operation: op, Inspected: 0,
		DivisibleBy: div_by,
	}
}

func (m *Monkey) Print() {
	fmt.Printf("Monkey(id=%d items=%v inspected=%d)\n", m.Id, m.Items, m.Inspected)
}

type Monkeys []*Monkey

func prepare(lines []string) (monkeys Monkeys) {
	monkey_lines := make([]string, 6)
	n := 0
	for _, line := range lines {
		if len(line) == 0 {
			monkeys = append(monkeys, NewMonkey(monkey_lines))
			n = 0
		} else {
			monkey_lines[n] = line
			n += 1
		}
	}
	if n > 0 {
		monkeys = append(monkeys, NewMonkey(monkey_lines))
	}
	return
}

func solve(monkeys Monkeys, N int, f Unary) int {
	for i := 0; i < N; i++ {
		for _, m := range monkeys {
			for _, item := range m.Items {
				value := f(m.Operation(item))
				next_m_id := m.ThrowTo(value)
				monkeys[next_m_id].Items = append(monkeys[next_m_id].Items, value)
			}
			m.Inspected += len(m.Items)
			m.Items = make([]int, 0)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})
	return monkeys[0].Inspected * monkeys[1].Inspected
}

func part_1(lines []string) {
	monkeys := prepare(lines)
	ans := solve(monkeys, 20, func(x int) int { return x / 3 })
	utils.CheckTask(1, ans, 120056)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(lines []string) {
	monkeys := prepare(lines)
	var p int = 1
	for _, m := range monkeys {
		p *= m.DivisibleBy
	}
	ans := solve(monkeys, 10000, func(x int) int { return x % p })
	utils.CheckTask(2, ans, 21816744824)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day11/test1.txt"
	inputFile := "inputs/day11/input.txt"
	lines := utils.ReadFile(inputFile)
	part_1(lines)
	part_2(lines)
}
