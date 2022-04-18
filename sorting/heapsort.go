package sorting

import "github.com/markwinter/gocol/ds"

func HeapSort[T any](data []T, cmp func(a, b T) bool) []T {
	heap := ds.MakeHeap(data, cmp)

	sorted := make([]T, 0, len(data))
	for e, err := heap.Pop(); err == nil; e, err = heap.Pop() {
		sorted = append(sorted, e)
	}

	return sorted
}
