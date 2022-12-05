# aoc2022
[Advent of Code 2022](https://adventofcode.com/2022) solved (probably) in Golang

## Requirements
* go >= 1.19
* golang.org/x/exp

## Commands
### Make commands
All make commands can get an optional V=1 argument for verbosity
* Build all tasks to binaries located inside **bin** directory:
```console
make V=1 all
```
* Check if all tasks are still working correctly:
```console
make test
```
* Build specified day task:
```console
make day12
```
* Run specified day task:
```console
make run_day12
```
* Clean
```console
make clean
```
### Other commands
* Add new day:
```console
./scripts/add_day.sh 15
```
* Run all tests manually:
```console
./scripts/check.sh
```
