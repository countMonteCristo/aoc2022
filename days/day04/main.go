package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type Pair struct {
	first, second utils.Segmnet
}

func pairFromStr(str string) (p Pair) {
	pair := strings.Split(str, ",")
	p.first = segmentFromStr(pair[0])
	p.second = segmentFromStr(pair[1])
	return
}

func segmentFromStr(str string) (i utils.Segmnet) {
	rangeStr := strings.Split(str, "-")
	i.Begin, _ = strconv.Atoi(rangeStr[0])
	i.End, _ = strconv.Atoi(rangeStr[1])
	return
}

func prepare(lines []string) (pairs []Pair) {
	for _, line := range lines {
		p := pairFromStr(line)
		pairs = append(pairs, p)
	}
	return
}

func part_1(pairs []Pair) {
	ans := utils.CountIf(pairs, func(p *Pair) bool { return p.first.Contains(&p.second) || p.second.Contains(&p.first) })
	utils.CheckTask(1, ans, 431)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(pairs []Pair) {
	ans := utils.CountIf(pairs, func(p *Pair) bool { return p.first.Intersects(&p.second) })
	utils.CheckTask(2, ans, 823)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day04/test1.txt"
	inputFile := "inputs/day04/input.txt"
	lines := utils.ReadFile(inputFile)
	pairs := prepare(lines)
	part_1(pairs)
	part_2(pairs)
}
