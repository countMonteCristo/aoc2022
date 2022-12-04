package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type IdRange struct {
	begin, end int
}

type Pair struct {
	first, second IdRange
}

func pairFromStr(str string) (p Pair) {
	pair := strings.Split(str, ",")
	p.first = idsFromStr(pair[0])
	p.second = idsFromStr(pair[1])
	return
}

func idsFromStr(str string) (i IdRange) {
	rangeStr := strings.Split(str, "-")
	i.begin, _ = strconv.Atoi(rangeStr[0])
	i.end, _ = strconv.Atoi(rangeStr[1])
	return
}

func (r1 *IdRange) contains(r2 *IdRange) bool {
	return r1.begin <= r2.begin && r1.end >= r2.end
}

func (r1 *IdRange) intersects(r2 *IdRange) bool {
	return r1.end >= r2.begin && r2.end >= r1.begin
}

func prepare(lines []string) (pairs []Pair) {
	for _, line := range lines {
		p := pairFromStr(line)
		pairs = append(pairs, p)
	}
	return
}

func countIf(pairs []Pair, f func(*Pair) bool) (ans int) {
	for _, pair := range pairs {
		if f(&pair) {
			ans += 1
		}
	}
	return
}

func part_1(pairs []Pair) {
	ans := countIf(pairs, func(p *Pair) bool { return p.first.contains(&p.second) || p.second.contains(&p.first) })
	if ans != 431 {
		log.Fatal("Wrong answer at part 1: ", ans, " (correct: 431)")
	}
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(pairs []Pair) {
	ans := countIf(pairs, func(p *Pair) bool { return p.first.intersects(&p.second) })
	if ans != 823 {
		log.Fatal("Wrong answer at part 2: ", ans, " (correct: 823)")
	}
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day04/test1.txt"
	inputFile := "inputs/day04/input.txt"
	data := utils.ReadFile(inputFile)
	pairs := prepare(data)
	part_1(pairs)
	part_2(pairs)
}
