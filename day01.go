package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
	"sort"
)

func readFile(fn string) []int {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	current := 0
	elves := make([]int, 0)
	for scanner.Scan() {
		calories_str := scanner.Text()
		if (len(calories_str) == 0) {
			elves = append(elves, current)
			current = 0
		} else {
			calories, _ := strconv.Atoi(calories_str)
			current += calories
		}
	}
	elves = append(elves, current)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return elves
}


func part_1(fn string) {
	data := readFile(fn)
	ans := 0
	for _, q := range(data) {
		if (q >= ans) {
			ans = q
		}
	}
	fmt.Println("[Part 1] Answer:", ans)
}


func part_2(fn string) {
	data := readFile(fn)
	sort.Slice(data, func(i, j int) bool {
		return data[i] > data[j]
	})
	fmt.Println("[Part 2] Answer:", data[0] + data[1] + data[2])
}


func main() {
	// inputFile := "input/day01/test1.txt"
	inputFile := "input/day01/input.txt"
	part_1(inputFile)
	part_2(inputFile)
}
