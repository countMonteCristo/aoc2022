package utils

import (
	"bufio"
	"log"
	"os"
)

// Check task `part_id` answer for correctness, aborting if `ans“ does not equal to `correct_ans“
func CheckTask[AnsType Number | string](part_id int, ans, correct_ans AnsType) {
	if ans != correct_ans {
		log.Fatal("Wrong answer at part_", part_id, ": ", ans, " (correct: ", correct_ans, ")")
	}
}

// Returns array of lines from file name `fn`
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
