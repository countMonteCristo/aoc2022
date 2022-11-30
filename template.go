package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func readFile(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


func part_1(fn string) {
	// data := readFile(fn)
	// process
	// fmt.Println("Answer:", 3)
}


func part_2(fn string) {
	// data := readFile(fn)
	// process
	// fmt.Println("Answer:", 3)
}


func main() {
	inputFile := "path/to/file"
	part_1(inputFile)
	part_2(inputFile)
}
