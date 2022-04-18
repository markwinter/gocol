package sorting

func InsertionSort[T any](data []T, cmp func(a, b T) bool) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && cmp(data[j], data[j-1]); j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}
