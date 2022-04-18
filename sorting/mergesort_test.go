package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_MergeSortAsc(t *testing.T) {
	data := []int{5, 4, 3, 1, 9, 0, 2}

	sorted := MergeSort(data, func(a, b int) bool { return a < b })

	if !cmp.Equal(sorted, []int{0, 1, 2, 3, 4, 5, 9}) {
		t.Errorf("array was not sorted: %v", sorted)
	}
}

func Test_MergeSortDesc(t *testing.T) {
	data := []int{5, 4, 3, 1, 9, 0, 2}

	sorted := MergeSort(data, func(a, b int) bool { return a > b })

	if !cmp.Equal(sorted, []int{9, 5, 4, 3, 2, 1, 0}) {
		t.Errorf("array was not sorted: %v", sorted)
	}
}
