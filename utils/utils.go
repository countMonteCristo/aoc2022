package utils

import (
	"bufio"
	"log"
	"os"
)

func CheckTask(part_id int, ans, correct_ans int) {
	if ans != correct_ans {
		log.Fatal("Wrong answer at part_", part_id, ": ", ans, " (correct: ", correct_ans, ")")
	}
}

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
