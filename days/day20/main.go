package main

import (
	"fmt"
	"strconv"

	"aoc2022/utils"
)

type DataType []int

func prepare(lines []string) (data DataType) {
	return utils.Transform(lines, func(line string) int {
		x, _ := strconv.Atoi(line)
		return x
	})
}

func PosMod(x, m int) int {
	return ((x%m) + m)%m
}

func solve(data DataType, M, N int ) (ans int) {
	indicies := make([]int, len(data))
	zero_id := 0
	for i := 0; i< len(data); i++ {
		indicies[i] = i
		data[i] *= M
		if data[i] == 0 {
			zero_id = i
		}
	}

	for n:=0; n<N; n++ {
		for i:=0; i< len(indicies); i++ {
			index := indicies[i]
			new_index := PosMod(index+data[i], len(data)-1)

			d := -1
			if new_index >= index {
				d = 1
			}

			begin, end := utils.MinMax(index, new_index)
			for j:=0; j<len(indicies); j++ {
				if indicies[j] >= begin && indicies[j] <= end {
					indicies[j] = PosMod(indicies[j] - d, len(data))
				}
			}
			indicies[i] = new_index
		}
	}

	for _, d := range []int{1000, 2000, 3000} {
		t := (indicies[zero_id] + d) % len(data)
		for i, x := range indicies {
			if x == t {
				ans += data[i]
				break
			}
		}
	}

	return
}

func part_1(data DataType) {
	ans := solve(data, 1, 1)
	// utils.CheckTask(1, ans, 23321)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data DataType) {
	ans := solve(data, 811589153, 10)
	utils.CheckTask(2, ans, 1428396909280)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day20/test1.txt"
	inputFile := "inputs/day20/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
