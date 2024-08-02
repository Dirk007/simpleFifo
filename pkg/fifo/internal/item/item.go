package item

type FifoItem[T any] struct {
	content  T
	previous *FifoItem[T]
	next     *FifoItem[T]
}

func NewUnbound[T any](content T) *FifoItem[T] {
	return &FifoItem[T]{content: content}
}

func (e *FifoItem[T]) Value() T {
	return e.content
}

// Remove removes the current item from the list and returns the new tail and the removed value.
// If the current item is the only one in the list, the returned tail will be nil.
func (e *FifoItem[T]) Remove() (*FifoItem[T], T) {
	if e.previous == nil {
		return nil, e.content
	}
	e.previous.next = nil
	return e.previous, e.content
}

// Prepend adds a new item to the beginning of the list and returns the new head.
func (e *FifoItem[T]) Prepend(value T) *FifoItem[T] {
	e.previous = &FifoItem[T]{
		content:  value,
		previous: nil,
		next:     e,
	}

	return e.previous
}
