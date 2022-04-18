package ds

import "errors"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(t T) {
	s.data = append(s.data, t)
}

func (s *Stack[T]) Pop() (T, error) {
	var ret T

	if s.Empty() {
		return ret, errors.New("stack is empty")
	}

	l := len(s.data) - 1

	ret = s.data[l]

	s.data = s.data[:l]

	return ret, nil
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Empty() bool {
	return s.Size() == 0
}
