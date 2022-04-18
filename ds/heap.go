package ds

import "errors"

type Heap[T any] struct {
	data []T
	cmp  func(a, b T) bool
}

func MakeHeap[T any](data []T, cmp func(a, b T) bool) Heap[T] {
	h := Heap[T]{
		data: make([]T, len(data)+1),
		cmp:  cmp,
	}

	for i := range data {
		h.data[i+1] = data[i]
	}

	for i := len(data) / 2; i >= 1; i-- {
		h.bubbleDown(i)
	}

	return h
}

func (h *Heap[T]) Push(t T) {
	h.data = append(h.data, t)
	h.bubbleUp(len(h.data) - 1)
}

func (h *Heap[T]) Pop() (T, error) {
	var r T

	if h.Empty() {
		return r, errors.New("heap is empty")
	}

	r = h.data[1]

	h.data[1], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[1]

	h.data = h.data[:len(h.data)-1]
	h.bubbleDown(1)

	return r, nil
}

func (h *Heap[T]) Empty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Size() int {
	return len(h.data) - 1
}

func (h *Heap[T]) bubbleDown(i int) {
	minIndex := i
	child := firstChild(i)

	for i := 0; i < 2; i++ {
		if child+i < len(h.data) {
			if h.cmp(h.data[child+i], h.data[minIndex]) {
				minIndex = child + i
			}
		}
	}

	if minIndex != i {
		h.data[i], h.data[minIndex] = h.data[minIndex], h.data[i]
		h.bubbleDown(minIndex)
	}
}

func (h *Heap[T]) bubbleUp(i int) {
	if parent(i) == -1 {
		return
	}

	if h.cmp(h.data[i], h.data[parent(i)]) {
		h.data[i], h.data[parent(i)] = h.data[parent(i)], h.data[i]
		h.bubbleUp(parent(i))
	}
}

func parent(i int) int {
	if i == 1 {
		return -1
	}
	return i / 2
}

func firstChild(i int) int { return 2 * i }
