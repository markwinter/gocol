package search

import "testing"

func Test_BinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	cmp := func(a, b int) int {
		if a == b {
			return 0
		}

		if a > b {
			return 1
		}

		return -1
	}

	if _, err := BinarySearch(4, data, cmp); err != nil {
		t.Errorf("%v", err)
	}
}

func Test_BinarySearchNotFound(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	cmp := func(a, b int) int {
		if a == b {
			return 0
		}

		if a > b {
			return 1
		}

		return -1
	}

	if _, err := BinarySearch(10, data, cmp); err == nil {
		t.Error("expected err but it was nil")
	}
}
