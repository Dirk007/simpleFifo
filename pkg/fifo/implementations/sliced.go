package implementations

var _ FifoStrategy[any] = &SliceFifo[any]{}

type SliceFifo[T any] struct {
	inner []T
}

func NewSliceFifo[T any]() FifoStrategy[T] {
	return &SliceFifo[T]{
		inner: make([]T, 0),
	}
}

func (s *SliceFifo[T]) Add(values ...T) (FifoStrategy[T], error) {
	s.inner = append(s.inner, values...)

	return s, nil
}

func (s *SliceFifo[T]) Next() (T, error) {
	var value T

	if len(s.inner) == 0 {
		return value, ErrEmptyFifoInternal
	}

	nextValue := s.inner[0]
	s.inner = s.inner[1:]

	return nextValue, nil
}

func (s *SliceFifo[T]) Count() int {
	return len(s.inner)
}

func (s *SliceFifo[T]) Clear() FifoStrategy[T] {
	s.inner = make([]T, 0)
	return s
}
