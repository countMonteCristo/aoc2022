package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type Segmnet struct {
	begin, end int
}

type Pair struct {
	first, second Segmnet
}

func pairFromStr(str string) (p Pair) {
	pair := strings.Split(str, ",")
	p.first = segmentFromStr(pair[0])
	p.second = segmentFromStr(pair[1])
	return
}

func segmentFromStr(str string) (i Segmnet) {
	rangeStr := strings.Split(str, "-")
	i.begin, _ = strconv.Atoi(rangeStr[0])
	i.end, _ = strconv.Atoi(rangeStr[1])
	return
}

func (r1 *Segmnet) contains(r2 *Segmnet) bool {
	return r1.begin <= r2.begin && r1.end >= r2.end
}

func (r1 *Segmnet) intersects(r2 *Segmnet) bool {
	return r1.end >= r2.begin && r2.end >= r1.begin
}

func prepare(lines []string) (pairs []Pair) {
	for _, line := range lines {
		p := pairFromStr(line)
		pairs = append(pairs, p)
	}
	return
}

func countIf(pairs []Pair, predicate func(*Pair) bool) (ans int) {
	for _, pair := range pairs {
		if predicate(&pair) {
			ans += 1
		}
	}
	return
}

func part_1(pairs []Pair) {
	ans := countIf(pairs, func(p *Pair) bool { return p.first.contains(&p.second) || p.second.contains(&p.first) })
	utils.CheckTask(1, ans, 431)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(pairs []Pair) {
	ans := countIf(pairs, func(p *Pair) bool { return p.first.intersects(&p.second) })
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
