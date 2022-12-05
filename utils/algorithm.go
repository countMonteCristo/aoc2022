package utils

func CountIf[T any](items []T, predicate func(*T) bool) (count int) {
	for _, pair := range items {
		if predicate(&pair) {
			count += 1
		}
	}
	return
}
