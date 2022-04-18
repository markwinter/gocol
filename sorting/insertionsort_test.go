package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_InsertionSortAsc(t *testing.T) {
	data := []int{3, 8, 4, 9, 1, 5, 0}

	lt := func(a, b int) bool { return a < b }
	InsertionSort(data, lt)

	if !cmp.Equal(data, []int{0, 1, 3, 4, 5, 8, 9}) {
		t.Errorf("array was not sorted: %v", data)
	}
}

func Test_InsertionSortDesc(t *testing.T) {
	data := []int{3, 8, 4, 9, 1, 5, 0}

	gt := func(a, b int) bool { return a > b }
	InsertionSort(data, gt)

	if !cmp.Equal(data, []int{9, 8, 5, 4, 3, 1, 0}) {
		t.Errorf("array was not sorted: %v", data)
	}
}
