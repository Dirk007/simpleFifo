package implementations

import "errors"

var ErrEmptyFifoInternal = errors.New("empty FIFO (internal)")

type FifoStrategy[T any] interface {
	Add(values ...T) (FifoStrategy[T], error)
	Next() (T, error)
	Clear() FifoStrategy[T]

	Count() int
}
