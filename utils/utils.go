package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fn string) (lines []string) {
	file, err := os.Open(fn)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
