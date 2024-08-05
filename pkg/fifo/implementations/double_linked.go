package implementations

import (
	"github.com/Dirk007/simpleFifo/pkg/fifo/internal/item"
)

var _ FifoStrategy[any] = &DoubleLinkedFifo[any]{}

type DoubleLinkedFifo[T any] struct {
	count int
	first *item.DoubleLinkedItem[T]
	last  *item.DoubleLinkedItem[T]
}

func NewDoubleLinkedFifo[T any]() FifoStrategy[T] {
	return &DoubleLinkedFifo[T]{}
}

func (f *DoubleLinkedFifo[T]) Add(values ...T) (FifoStrategy[T], error) {
	for _, itemValue := range values {
		f.count++
		if f.count == 1 {
			// Very first entry
			entry := item.NewUnbound(itemValue)
			f.first = entry
			f.last = entry
			continue
		}

		f.first = f.first.Prepend(itemValue)
	}

	return f, nil
}

// Next pops the next value from the end of the FIFO.
// Returns an error if the FIFO is empty.
func (f *DoubleLinkedFifo[T]) Next() (T, error) {
	var value T

	if f.count == 0 {
		return value, ErrEmptyFifoInternal
	}

	f.count--
	f.last, value = f.last.Remove()

	return value, nil
}

func (f *DoubleLinkedFifo[T]) Count() int {
	return f.count
}

func (f *DoubleLinkedFifo[T]) Clear() FifoStrategy[T] {
	f.count = 0
	f.first = nil
	f.last = nil

	return f
}
