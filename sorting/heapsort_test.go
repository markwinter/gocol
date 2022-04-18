package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_HeapSortAsc(t *testing.T) {
	data := []int{3, 8, 4, 6, 1, 9}

	lt := func(a, b int) bool { return a < b }
	sorted := HeapSort(data, lt)

	if !cmp.Equal(sorted, []int{1, 3, 4, 6, 8, 9}) {
		t.Errorf("array was not sorted: %v", sorted)
	}
}

func Test_HeapSortDesc(t *testing.T) {
	data := []int{3, 8, 4, 6, 1, 9}

	gt := func(a, b int) bool { return a > b }
	sorted := HeapSort(data, gt)

	if !cmp.Equal(sorted, []int{9, 8, 6, 4, 3, 1}) {
		t.Errorf("array was not sorted: %v", sorted)
	}
}
