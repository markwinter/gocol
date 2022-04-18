package ds

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Queue(t *testing.T) {
	queue := Queue[int]{}

	queue.Push(0)
	queue.Push(1)
	queue.Push(3)

	var data []int
	for e, err := queue.Pop(); err == nil; e, err = queue.Pop() {
		data = append(data, e)
	}

	if !cmp.Equal(data, []int{0, 1, 3}) {
		t.Errorf("data was not correct order: %v", data)
	}
}

func Test_EmptyQueue(t *testing.T) {
	queue := Queue[int]{}

	if _, err := queue.Pop(); err == nil {
		t.Errorf("expected error but it was nil")
	}
}
