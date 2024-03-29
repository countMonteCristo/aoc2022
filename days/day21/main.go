package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

var Funcs = map[string]utils.Binary[int64]{
	"+": func(x, y int64) int64 { return x + y },
	"*": func(x, y int64) int64 { return x * y },
	"-": func(x, y int64) int64 { return x - y },
	"/": func(x, y int64) int64 { return x / y },
}

type Monkey struct {
	Name      string
	Number    int64
	Args      [2]string
	Operation string
	Known     bool
}
type Monkeys = map[string]*Monkey
type Path = []string

func (m *Monkey) String() string {
	if m.Known {
		return fmt.Sprintf("%s: %d", m.Name, m.Number)
	} else {
		return fmt.Sprintf("%s: %s %s %s", m.Name, m.Args[0], m.Operation, m.Args[1])
	}
}

func (m *Monkey) Resolve(monkeys Monkeys, val int64) {
	if m.Name == "humn" {
		m.Known = true
		m.Number = val
		return
	}
	m1, m2 := monkeys[m.Args[0]], monkeys[m.Args[1]]
	m1.Eval(monkeys, true)
	m2.Eval(monkeys, true)

	if !m1.Known && !m2.Known {
		panic(fmt.Sprintf(
			"Cannot resolve for monkey %s: %s and %s are unknown",
			m.Name, m1.Name, m2.Name,
		))
	}

	m_unknown, arg := m2, m1.Number
	if m2.Known {
		m_unknown, arg = m1, m2.Number
	}
	var new_val int64

	switch m.Operation {
	case "+":
		new_val = val - arg
	case "*":
		new_val = val / arg
	case "-":
		new_val = arg + val
		if m1.Known {
			new_val = arg - val
		}
	case "/":
		new_val = arg * val
		if m1.Known {
			new_val = arg / val
		}
	default:
		panic(fmt.Sprintf("Unknown operation for monkey %s: '%s'", m.Name, m.Operation))
	}
	m_unknown.Resolve(monkeys, new_val)
}

func (m *Monkey) Eval(monkeys Monkeys, skip_humn bool) {
	if m.Known || (m.Name == "humn" && skip_humn) {
		return
	}
	m1, m2 := monkeys[m.Args[0]], monkeys[m.Args[1]]
	m1.Eval(monkeys, skip_humn)
	m2.Eval(monkeys, skip_humn)
	if m1.Known && m2.Known {
		m.Number = Funcs[m.Operation](
			m1.Number, m2.Number,
		)
		m.Known = true
	}
}

func prepare(lines []string) (mokeys Monkeys) {
	mokeys = make(Monkeys)
	for _, line := range lines {
		m := Monkey{Known: false}
		parts := strings.Split(line, " ")
		m.Name = strings.Trim(parts[0], ":")
		x, is_num := strconv.ParseInt(parts[1], 10, 64)
		if is_num == nil {
			m.Known = true
			m.Number = x
		} else {
			m.Args[0], m.Args[1] = parts[1], parts[3]
			m.Operation = parts[2]
		}
		mokeys[m.Name] = &m
	}
	return
}

func solve_1(monkeys Monkeys) int64 {
	root := monkeys["root"]
	root.Eval(monkeys, false)
	return root.Number
}

func solve_2(monkeys Monkeys) int64 {
	humn, root := monkeys["humn"], monkeys["root"]
	humn.Known = false
	first, second := monkeys[root.Args[0]], monkeys[root.Args[1]]
	first.Eval(monkeys, true)
	second.Eval(monkeys, true)
	if first.Known {
		second.Resolve(monkeys, first.Number)
	} else {
		first.Resolve(monkeys, second.Number)
	}
	return humn.Number
}

func part_1(lines []string) {
	data := prepare(lines)
	ans := solve_1(data)
	utils.CheckTask(1, ans, 169525884255464)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(lines []string) {
	data := prepare(lines)
	ans := solve_2(data)
	utils.CheckTask(2, ans, 3247317268284)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day21/test1.txt"
	inputFile := "inputs/day21/input.txt"
	lines := utils.ReadFile(inputFile)
	part_1(lines)
	part_2(lines)
}
