package sorting

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_QuickSortAsc(t *testing.T) {
	data := []int{2, 9, 3, 8, 4, 5, 7, 6, 1, 0}

	lt := func(a, b int) bool { return a < b }

	QuickSort(data, lt)

	log.Print(data)

	if !cmp.Equal(data, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("data was not sorted: %v", data)
	}
}

func Test_QuickSortDesc(t *testing.T) {
	data := []int{2, 9, 3, 8, 4, 5, 7, 6, 1, 0}

	lt := func(a, b int) bool { return a > b }

	QuickSort(data, lt)

	log.Print(data)

	if !cmp.Equal(data, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) {
		t.Errorf("data was not sorted: %v", data)
	}
}
