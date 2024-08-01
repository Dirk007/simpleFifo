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

func (e *FifoItem[T]) Remove() (*FifoItem[T], T) {
	if e.previous == nil {
		return nil, e.content
	}
	e.previous.next = nil
	return e.previous, e.content
}

func (e *FifoItem[T]) Prepend(value T) *FifoItem[T] {
	e.previous = &FifoItem[T]{
		content:  value,
		previous: nil,
		next:     e,
	}

	return e.previous
}
