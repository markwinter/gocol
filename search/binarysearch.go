package search

import "errors"

// BinarySearch returns the index of the target in the given slice
// The comparison function must return one of the following:
// When equal, return 0
// When a is larger, return 1
// Otherwise return return -1
func BinarySearch[T any](target T, data []T, cmp func(a, b T) int) (int, error) {
	l := 0
	r := len(data) - 1

	for l <= r {
		mid := ((r - l) / 2) + l

		v := cmp(target, data[mid])
		if v == 0 {
			return mid, nil
		} else if v == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1, errors.New("target not found")
}
