package utils

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

// Create func(*T)V from func(T)V
func Deref[T, V any](f func(T) V) func(*T) V {
	return func(x *T) V {
		return f(*x)
	}
}

func CountIfP[T any](items []T, pred func(*T) bool) (count int) {
	for _, item := range items {
		if pred(&item) {
			count += 1
		}
	}
	return
}

func CountIf[T any](items []T, pred func(T) bool) (count int) {
	return CountIfP(items, Deref(pred))
}

func SumIfP[T Number](items []T, pred func(*T) bool) (count T) {
	for _, item := range items {
		if pred(&item) {
			count += item
		}
	}
	return
}

func SumIf[T Number](items []T, pred func(T) bool) (count T) {
	return SumIfP(items, Deref(pred))
}

func SumValueP[T any, V Number](items []T, f func(*T) V) V {
	return Sum(TransformP(items, f))
}

func SumValue[T any, V Number](items []T, f func(T) V) (count V) {
	return Sum(Transform(items, f))
}

func Sum[V Number](items []V) (count V) {
	return SumIf(items, True[V])
}

func TransformP[T, V any](items []T, f func(*T) V) []V {
	res := make([]V, len(items))
	for index, item := range items {
		res[index] = f(&item)
	}
	return res
}

func Transform[T, V any](items []T, f func(T) V) []V {
	return TransformP(items, Deref(f))
}

func FilterP[T any](items []T, pred func(*T) bool) []T {
	res := make([]T, 0)
	for _, item := range items {
		if pred(&item) {
			res = append(res, item)
		}
	}
	return res
}

func Filter[T any](items []T, pred func(T) bool) []T {
	return FilterP(items, Deref(pred))
}

func UpdateMap[T comparable, V any](first, second map[T]V) {
	for key, value := range second {
		first[key] = value
	}
}
