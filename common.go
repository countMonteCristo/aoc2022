package main

import (
	"aoc2022/utils"
	"strconv"
)

type IntPoint = utils.Point2d[int]
type I64Point = utils.Point2d[int64]
type IntPoint3d = utils.Point3d[int]

var IpZero = IntPoint{X: 0, Y: 0}
var IpZero3d = IntPoint3d{X: 0, Y: 0, Z: 0}

var DD = []IntPoint{
	{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1},
}


type IntSet = *utils.Set[int]
type I64Set = *utils.Set[int64]
type IpSet = *utils.Set[IntPoint]
type I64pSet = *utils.Set[I64Point]
type StrSet = *utils.Set[string]

func NewIntSet() IntSet{
	return utils.NewSet[int]()
}
func NewI64Set() I64Set{
	return utils.NewSet[int64]()
}

func NewIpSet() IpSet {
	return utils.NewSet[IntPoint]()
}
func NewI64pSet() I64pSet {
	return utils.NewSet[I64Point]()
}
func NewStrSet() StrSet {
	return utils.NewSet[string]()
}


func StrToInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}
