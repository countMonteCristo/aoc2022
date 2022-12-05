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

// Segment AB on X-axis where x(A) = Begin, x(B) = End
// It's assumed that Begin <= End
type Segmnet struct {
	Begin, End int
}

// Check if segment r2 is fully inside segment r1
func (r1 *Segmnet) Contains(r2 *Segmnet) bool {
	return r1.Begin <= r2.Begin && r1.End >= r2.End
}

// Check if r1 and r2 have any common part
func (r1 *Segmnet) Intersects(r2 *Segmnet) bool {
	return r1.End >= r2.Begin && r2.End >= r1.Begin
}
