package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type PairSeg utils.Pair[utils.Segment]

func pairFromStr(str string) (p PairSeg) {
	pair := strings.Split(str, ",")
	p.First = segmentFromStr(pair[0])
	p.Second = segmentFromStr(pair[1])
	return
}

func segmentFromStr(str string) (i utils.Segment) {
	rangeStr := strings.Split(str, "-")
	i.Begin, _ = strconv.Atoi(rangeStr[0])
	i.End, _ = strconv.Atoi(rangeStr[1])
	return
}

func prepare(lines []string) (pairs []PairSeg) {
	for _, line := range lines {
		p := pairFromStr(line)
		pairs = append(pairs, p)
	}
	return
}

func part_1(pairs []PairSeg) {
	ans := utils.CountIfP(pairs, func(p *PairSeg) bool { return p.First.Contains(&p.Second) || p.Second.Contains(&p.First) })
	utils.CheckTask(1, ans, 431)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(pairs []PairSeg) {
	ans := utils.CountIfP(pairs, func(p *PairSeg) bool { return p.First.Intersects(&p.Second) })
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
