package ds

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_MinHeap(t *testing.T) {
	d := []int{5, 4, 7, 3, 1}

	lt := func(a, b int) bool { return a < b }
	heap := MakeHeap(d, lt)

	heap.Push(0)
	heap.Push(6)

	var sorted []int
	for !heap.Empty() {
		sorted = append(sorted, heap.Pop())
	}

	if !cmp.Equal(sorted, []int{0, 1, 3, 4, 5, 6, 7}) {
		t.Errorf("array was not sorted: %v", sorted)
	}
}

func Test_MaxHeap(t *testing.T) {
	d := []int{3, 5, 7, 4, 1}

	gt := func(a, b int) bool { return a > b }
	heap := MakeHeap(d, gt)

	heap.Push(0)
	heap.Push(6)

	var sorted []int
	for !heap.Empty() {
		sorted = append(sorted, heap.Pop())
	}

	if !cmp.Equal(sorted, []int{7, 6, 5, 4, 3, 1, 0}) {
		t.Errorf("array was not sorted: %v", sorted)
	}
}
