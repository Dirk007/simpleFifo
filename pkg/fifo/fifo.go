package fifo

import (
	"fifo/pkg/fifo/internal/item"
	"sync"
)

// Fifo that can solely be used to push items to the end and pop items from the beginning.
type Fifo[T any] struct {
	lock  sync.RWMutex
	count uint64
	limit uint64
	first *item.FifoItem[T]
	last  *item.FifoItem[T]
}

func NewFifo[T any]() *Fifo[T] {
	return &Fifo[T]{}
}

func (f *Fifo[T]) WithLimit(limit uint64) *Fifo[T] {
	f.limit = limit
	return f
}

func (f *Fifo[T]) Add(values ...T) (*Fifo[T], error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.limit > 0 && f.count >= f.limit {
		return f, NewFifoLimitReachedError(int64(f.limit))
	}

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

func (f *Fifo[T]) Next() (T, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	var value T

	if f.count == 0 {
		return value, ErrEmptyFifo
	}

	f.count--
	f.last, value = f.last.Remove()

	return value, nil
}

func (f *Fifo[T]) Len() uint64 {
	f.lock.RLock()
	defer f.lock.Unlock()

	return f.count
}

func (f *Fifo[T]) IsEmpty() bool {
	return f.Len() == 0
}

func (f *Fifo[T]) Clear() *Fifo[T] {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.count = 0
	f.first = nil
	f.last = nil

	return f
}
