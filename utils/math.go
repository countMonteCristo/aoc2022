package utils

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	int8 | int16 | int32 | int64 | int | float32 | float64
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

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
type Segment struct {
	Begin, End int
}

// Check if segment s2 is fully inside segment s1
func (s1 *Segment) Contains(s2 *Segment) bool {
	return s1.Begin <= s2.Begin && s1.End >= s2.End
}

// Check if s1 and s2 have any common part
func (s1 *Segment) Intersects(s2 *Segment) bool {
	return s1.End >= s2.Begin && s2.End >= s1.Begin
}

type Pair[T any] struct {
	First, Second T
}
