# aoc2022
[Advent of Code 2022](https://adventofcode.com/2022) solved (probably) in Golang

## Commands
### Make commands
All make commands can get an optional V=1 argument for verbosity
* Build all tasks to binaries located inside **bin** directory:
```bash
make V=1 all
```
* Check if all tasks are still working correctly:
```bash
make test
```
* Build specified day task:
```bash
make day12
```
* Run specified day task:
```bash
make run_day12
```
* Clean
```bash
make clean
```
### Other commands
* Add new day:
```bash
./scripts/add_day.sh 15
```
* Run all tests manually:
```bash
./scripts/check.sh
```
