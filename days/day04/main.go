package main

import (
	"fmt"
	"strings"

	"aoc2022/utils"
)

type PairSeg utils.Pair[utils.Segment]

func pairFromStr(str string) *PairSeg {
	pair := strings.Split(str, ",")
	return &PairSeg{
		First: segmentFromStr(pair[0]), Second: segmentFromStr(pair[1]),
	}
}

func segmentFromStr(str string) (i utils.Segment) {
	rangeStr := strings.Split(str, "-")
	i.Begin, i.End = StrToInt(rangeStr[0]), StrToInt(rangeStr[1])
	return
}

func prepare(lines []string) []*PairSeg {
	return utils.Transform(lines, pairFromStr)
}

func part_1(pairs []*PairSeg) {
	ans := utils.CountIf(pairs, func(p *PairSeg) bool { return p.First.Contains(&p.Second) || p.Second.Contains(&p.First) })
	utils.CheckTask(1, ans, 431)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(pairs []*PairSeg) {
	ans := utils.CountIf(pairs, func(p *PairSeg) bool { return p.First.Intersects(&p.Second) })
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
