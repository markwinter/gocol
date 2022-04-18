package sorting

func MergeSort[T any](data []T, cmp func(a, b T) bool) []T {
	if len(data) < 2 {
		return data
	}

	middle := len(data) / 2
	leftReturn := MergeSort(data[:middle], cmp)
	rightReturn := MergeSort(data[middle:], cmp)

	output := make([]T, 0, len(leftReturn)+len(rightReturn))

	i, j := 0, 0
	for i < len(leftReturn) && j < len(rightReturn) {
		if cmp(leftReturn[i], rightReturn[j]) {
			output = append(output, leftReturn[i])
			i++
		} else {
			output = append(output, rightReturn[j])
			j++
		}
	}

	output = append(output, leftReturn[i:]...)
	output = append(output, rightReturn[j:]...)

	return output
}
