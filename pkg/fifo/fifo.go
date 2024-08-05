package fifo

import (
	"github.com/Dirk007/simpleFifo/pkg/fifo/implementations"
	"github.com/Dirk007/simpleFifo/pkg/fifo/internal/locking"
)

// Fifo that can solely be used to push items to the beginning and pop items from the end.
type Fifo[T any] struct {
	lock  locking.Lock
	limit int
	inner implementations.FifoStrategy[T]
}

func (f *Fifo[T]) wrap(inner implementations.FifoStrategy[T], err error) (*Fifo[T], error) {
	if err != nil {
		return nil, err
	}

	return &Fifo[T]{
		lock:  f.lock,
		inner: inner,
	}, nil
}

func NewFifo[T any]() *Fifo[T] {
	return &Fifo[T]{
		lock:  locking.NewMutexLock(),
		inner: implementations.NewDoubleLinkedFifo[T](),
	}
}

func (f *Fifo[T]) WithImplementation(impl implementations.FifoStrategy[T]) *Fifo[T] {
	f.inner = impl
	return f
}

func (f *Fifo[T]) WithoutLocking() *Fifo[T] {
	f.lock = locking.NewNopLock()
	return f
}

func (f *Fifo[T]) WithLimit(limit int) *Fifo[T] {
	f.limit = limit
	return f
}

func (f *Fifo[T]) Add(values ...T) (*Fifo[T], error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.limit != 0 && f.inner.Count()+len(values) > f.limit {
		return nil, NewFifoLimitReachedError(int64(f.limit))
	}

	return f.wrap(f.inner.Add(values...))
}

func (f *Fifo[T]) Next() (T, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.inner.Count() == 0 {
		var zero T
		return zero, ErrEmptyFifo
	}

	return f.inner.Next()
}

func (f *Fifo[T]) Count() int {
	f.lock.RLock()
	defer f.lock.RUnlock()

	return f.inner.Count()
}

func (f *Fifo[T]) IsEmpty() bool {
	return f.Count() == 0
}

func (f *Fifo[T]) IsFull() bool {
	if f.limit == 0 {
		return false
	}

	return f.Count() >= f.limit
}

func (f *Fifo[T]) Clear() *Fifo[T] {
	f.lock.Lock()
	defer f.lock.Unlock()

	f.inner.Clear()

	return f
}
