package sorting

import "math/rand"

func QuickSort[T any](data []T, cmp func(a, b T) bool) {
	if len(data) < 2 {
		return
	}

	p := partition(data, cmp)
	QuickSort(data[:p], cmp)
	QuickSort(data[p+1:], cmp)
}

func partition[T any](data []T, cmp func(a, b T) bool) int {
	// Lomuto partition with random pivot
	p := rand.Intn(len(data))
	data[len(data)-1], data[p] = data[p], data[len(data)-1]
	p = len(data) - 1

	firstHigh := 0

	for i := 0; i < len(data); i++ {
		if cmp(data[i], data[p]) {
			data[i], data[firstHigh] = data[firstHigh], data[i]
			firstHigh++
		}
	}

	data[p], data[firstHigh] = data[firstHigh], data[p]

	return firstHigh
}
