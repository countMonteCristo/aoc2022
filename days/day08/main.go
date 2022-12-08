package main

import (
	"fmt"

	"aoc2022/utils"
)

type Forest []string

func TreeHeight(data Forest, r, c int) int {
	return int(data[r][c]) - int('0')
}

func scanRows(data Forest, vis_map [][]bool, w, h, dc int) {
	cmin, cmax := 1, w-1
	if dc < 0 {
		cmin, cmax = w-2, 0
	}

	for r := 1; r < h-1; r++ {
		max_height := TreeHeight(data, r, cmin-dc)
		for c := cmin; c != cmax; c += dc {
			height := TreeHeight(data, r, c)
			if height > max_height {
				vis_map[r][c] = true
				max_height = height
			}
		}
	}
}

func scanCols(data Forest, vis_map [][]bool, w, h, dr int) {
	rmin, rmax := 1, h-1
	if dr < 0 {
		rmin, rmax = h-2, 0
	}

	for c := 1; c < w-1; c++ {
		max_height := TreeHeight(data, rmin-dr, c)
		for r := rmin; r != rmax; r += dr {
			height := TreeHeight(data, r, c)
			if height > max_height {
				vis_map[r][c] = true
				max_height = height
			}
		}
	}
}

func getRowClosest(data Forest, r, c, w, dc int) int {
	th := TreeHeight(data, r, c)
	end, val := -1, c
	if dc > 0 {
		end, val = w, w-1-c
	}

	for cc := c + dc; cc != end; cc += dc {
		hh := TreeHeight(data, r, cc)
		if hh >= th {
			return utils.Abs(c - cc)
		}
	}
	return val
}

func getColClosest(data Forest, r, c, h, dr int) int {
	th := TreeHeight(data, r, c)
	end, val := -1, r
	if dr > 0 {
		end, val = h, h-1-r
	}

	for rr := r + dr; rr != end; rr += dr {
		hh := TreeHeight(data, rr, c)
		if hh >= th {
			return utils.Abs(r - rr)
		}
	}
	return val
}

func solve(data Forest, part int) (ans int) {
	w, h := len(data[0]), len(data)

	if part == 1 {
		vis_map := make([][]bool, h)
		for r := range data {
			vis_map[r] = make([]bool, w)
		}
		// set edges to be visible
		for r := 0; r < h; r++ {
			vis_map[r][0], vis_map[r][w-1] = true, true
		}
		for c := 0; c < w; c++ {
			vis_map[0][c], vis_map[h-1][c] = true, true
		}

		// left and right edges
		scanRows(data, vis_map, w, h, 1)
		scanRows(data, vis_map, w, h, -1)
		// top and bottom edges
		scanCols(data, vis_map, w, h, 1)
		scanCols(data, vis_map, w, h, -1)

		// count visible trees
		ans = utils.Sum(
			utils.Transform(vis_map, func(row []bool) int {
				return utils.CountIf(row, utils.Id[bool])
			}),
		)
	} else {
		for r := 0; r < h; r++ {
			for c := 0; c < w; c++ {
				dl := getRowClosest(data, r, c, w, -1)
				dr := getRowClosest(data, r, c, w, 1)
				dt := getColClosest(data, r, c, h, -1)
				db := getColClosest(data, r, c, h, 1)

				ans = utils.Max(ans, dl*dr*dt*db)
			}
		}
	}

	return
}

func part_1(data Forest) {
	ans := solve(data, 1)
	utils.CheckTask(1, ans, 1859)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data Forest) {
	ans := solve(data, 2)
	utils.CheckTask(2, ans, 332640)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day08/test1.txt"
	inputFile := "inputs/day08/input.txt"
	data := utils.ReadFile(inputFile)
	part_1(data)
	part_2(data)
}
