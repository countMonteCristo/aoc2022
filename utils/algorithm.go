package utils

import "errors"

type Summable interface {
	Number | string
}

// Identity function
func Id[T any](x T) T {
	return x
}

// Predicate which is true for any input arguments
func True[T any](x T) bool {
	return true
}

// Predicate which is false for any input arguments
func False[T any](x T) bool {
	return false
}

// Count the number of elements in `items` for which pred(Item) == true
func CountIf[T any](items []T, pred func(T) bool) (count int) {
	for _, item := range items {
		if pred(item) {
			count += 1
		}
	}
	return
}

// Sum the numbers in `items` for which pred(number) == true
func SumIf[T Summable](items []T, pred func(T) bool) (sum T) {
	for _, item := range items {
		if pred(item) {
			sum += item
		}
	}
	return
}

// Sum the values for each element in `items`. Value of element is computed as value = f(element)
func SumValue[T any, V Summable](items []T, f func(T) V) V {
	return Sum(Transform(items, f))
}

// Sums all numbers in array
func Sum[V Summable](items []V) V {
	return SumIf(items, True[V])
}

// Transform each element from `items` array using function f: g = f(element)
func Transform[T, V any](items []T, f func(T) V) []V {
	res := make([]V, len(items))
	for index, item := range items {
		res[index] = f(item)
	}
	return res
}

// Get all elements from `items` for which pred(element) == true
func Filter[T any](items []T, pred func(T) bool) []T {
	res := make([]T, 0)
	for _, item := range items {
		if pred(item) {
			res = append(res, item)
		}
	}
	return res
}

// Get first item from slice such that pred(item) == true or -1
func FindIndexIf[T any](items []T, pred func(T) bool) (int, error) {
	for index, item := range items {
		if pred(item) {
			return index, nil
		}
	}
	return -1, errors.New("no such element")
}

func FindIndex[T comparable](items []T, item T) (int, error) {
	return FindIndexIf(items, func(i T) bool { return i == item })
}

// Adds all key-value pairs from second map to first
func UpdateMap[T comparable, V any](first, second map[T]V) {
	for key, value := range second {
		first[key] = value
	}
}

// Ideas:
//
// 1. Apply
// 2. Map* algorithms
// 3. Str* algorithms
// ...
