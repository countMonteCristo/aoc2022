package utils

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](args ...T) T {
	return MinSlice(args)
}

func MinSlice[T constraints.Ordered](args []T) T {
	if len(args) == 0 {
		panic("No arguments provided for Min function")
	}
	min := args[0]
	for _, arg := range args[1:] {
		if arg < min {
			min = arg
		}
	}
	return min
}

func Max[T constraints.Ordered](args ...T) T {
	return MaxSlice(args)
}

func MaxSlice[T constraints.Ordered](args []T) T {
	if len(args) == 0 {
		panic("No arguments provided for Max function")
	}
	max := args[0]
	for _, arg := range args[1:] {
		if arg > max {
			max = arg
		}
	}
	return max
}
