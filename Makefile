$(V).SILENT:

GO=go

SOURCES := $(wildcard days/day*)
TARGETS := $(basename ${SOURCES})

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
