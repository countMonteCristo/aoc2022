$(V).SILENT:

GO=go

.PHONY: clean all

all:
	bash -c 'for f in days/*; do tgt=$$(basename $$f); make -B $$tgt; done'

clean:
	rm bin/day[0-2][0-9]

test:
	./scripts/check.sh -v

day01: days/day01/main.go
	${GO} build -o bin/$@ days/day01/main.go

run_day01: day01
	./bin/day01

day02: days/day02/main.go
	${GO} build -o bin/$@ days/day02/main.go

run_day02: day02
	./bin/day02

day03: days/day03/main.go
	${GO} build -o bin/$@ days/day03/main.go

run_day03: day03
	./bin/day03

day04: days/day04/main.go
	${GO} build -o bin/$@ days/day04/main.go

run_day04: day04
	./bin/day04

day05: days/day05/main.go
	${GO} build -o bin/$@ days/day05/main.go

run_day05: day05
	./bin/day05

day06: days/day06/main.go
	${GO} build -o bin/$@ days/day06/main.go

run_day06: day06
	./bin/day06

day07: days/day07/main.go
	${GO} build -o bin/$@ days/day07/main.go

run_day07: day07
	./bin/day07

day08: days/day08/main.go
	${GO} build -o bin/$@ days/day08/main.go

run_day08: day08
	./bin/day08

day09: days/day09/main.go
	${GO} build -o bin/$@ days/day09/main.go

run_day09: day09
	./bin/day09

day10: days/day10/main.go
	${GO} build -o bin/$@ days/day10/main.go

run_day10: day10
	./bin/day10

day11: days/day11/main.go
	${GO} build -o bin/$@ days/day11/main.go

run_day11: day11
	./bin/day11

day12: days/day12/main.go
	${GO} build -o bin/$@ days/day12/main.go

run_day12: day12
	./bin/day12

day13: days/day13/main.go
	${GO} build -o bin/$@ days/day13/main.go

run_day13: day13
	./bin/day13

day14: days/day14/main.go
	${GO} build -o bin/$@ days/day14/main.go

run_day14: day14
	./bin/day14

day15: days/day15/main.go
	${GO} build -o bin/$@ days/day15/main.go

run_day15: day15
	./bin/day15

day16: days/day16/main.go
	${GO} build -o bin/$@ days/day16/main.go

run_day16: day16
	./bin/day16
