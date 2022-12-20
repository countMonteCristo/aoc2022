package utils

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	int8 | int16 | int32 | int64 | int | float32 | float64 | uint8 | uint16 | uint32 | uint64 | uint
}

// Absolute value of number
func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Sign of number
func Sign[T Number](x T) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

// Get minimum number from arguments `args`
func Min[T constraints.Ordered](args ...T) T {
	return MinSlice(args)
}

// Get minimum number from slice `args`
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

// Get maximum number from arguments `args`
func Max[T constraints.Ordered](args ...T) T {
	return MaxSlice(args)
}

// Get maximum number from slice `args`
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

func MinMax[T constraints.Ordered](x, y T) (T, T) {
	return Min(x, y), Max(x, y)
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

// Point on 2D plane
type Point2d[T Number] struct {
	X, Y T
}

// p += q
func (p *Point2d[T]) Add(q *Point2d[T]) {
	p.Y += q.Y
	p.X += q.X
}

// x = p + q
func (p *Point2d[T]) Plus(q *Point2d[T]) Point2d[T] {
	return Point2d[T]{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

// p -= q
func (p *Point2d[T]) Sub(q *Point2d[T]) {
	p.Y -= q.Y
	p.X -= q.X
}

// x = p - q
func (p *Point2d[T]) Minus(q *Point2d[T]) Point2d[T] {
	return Point2d[T]{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

// p *= k
func (p *Point2d[T]) Mul(k T) {
	p.Y *= k
	p.X *= k
}

// x = p * k
func (p *Point2d[T]) Prod(k T) Point2d[T] {
	return Point2d[T]{
		X: k * p.X,
		Y: k * p.Y,
	}
}

// |p| = |p.x - q.x| + |p.y - q.y|
func Manhattan[T Number](p, q Point2d[T]) T {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}
