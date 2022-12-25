package main

import (
	"fmt"

	"aoc2022/utils"
)


var SnafuDigits = map[rune]int{
	'-': -1, '=': -2, '0': 0, '1': 1, '2': 2,
}

var Snafu = []string{
	"0", "1", "2", "=", "-",
}

func SnafuToBase10(s string) int {
	n := 0
	for _, d := range s {
		n = n*5 + SnafuDigits[d]
	}
	return n
}

func DecToBase5(x int) []int {
	res := make([]int, 0)
	for x > 0 {
		rem := x % 5
		res = append(res, rem)
		x = x / 5
	}
	return res
}

func Base5ToSnafu(x []int) string {
	snafu := ""
	carry := 0
	for _, digit := range x {
		nd := digit + carry
		carry = 0
		if nd > 2 {
			carry = 1
		}
		snafu = Snafu[nd % 5] + snafu
	}
	if carry == 1 {
		snafu = Snafu[carry] + snafu
	}
	return snafu
}

func DecToSnafu(x int) string {
	return Base5ToSnafu(DecToBase5(x))
}

func prepare(lines []string) (data []int) {
	data = utils.Transform(lines, SnafuToBase10)
	return
}

func solve(data []int) (ans string) {
	ans = DecToSnafu(utils.Sum(data))
	return
}

func part_1(data []int) {
	ans := solve(data)
	utils.CheckTask(1, ans, "2-0==21--=0==2201==2")
	fmt.Println("[Part 1] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day25/test1.txt"
	inputFile := "inputs/day25/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
}
