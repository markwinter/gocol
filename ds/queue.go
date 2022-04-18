package ds

import (
	"errors"
)

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Push(t T) {
	q.data = append(q.data, t)
}

func (q *Queue[T]) Pop() (T, error) {
	var ret T

	if q.Empty() {
		return ret, errors.New("queue is empty")
	}

	ret = q.data[0]

	q.data = q.data[1:]

	return ret, nil
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Empty() bool {
	return q.Size() == 0
}
