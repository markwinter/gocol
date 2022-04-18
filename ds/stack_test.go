package ds

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Stack(t *testing.T) {
	stack := Stack[int]{}

	stack.Push(0)
	stack.Push(1)
	stack.Push(2)

	var data []int
	for e, err := stack.Pop(); err == nil; e, err = stack.Pop() {
		data = append(data, e)
	}

	if !cmp.Equal(data, []int{2, 1, 0}) {
		t.Errorf("stack was not correct order: %v", data)
	}
}
