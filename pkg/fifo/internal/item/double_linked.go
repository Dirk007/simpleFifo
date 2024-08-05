package item

type DoubleLinkedItem[T any] struct {
	content  T
	previous *DoubleLinkedItem[T]
	next     *DoubleLinkedItem[T]
}

func NewUnbound[T any](content T) *DoubleLinkedItem[T] {
	return &DoubleLinkedItem[T]{content: content}
}

func (e *DoubleLinkedItem[T]) Value() T {
	return e.content
}

// Remove removes the current item from the list and returns the new tail and the removed value.
// If the current item is the only one in the list, the returned tail will be nil.
func (e *DoubleLinkedItem[T]) Remove() (*DoubleLinkedItem[T], T) {
	if e.previous == nil {
		return nil, e.content
	}
	e.previous.next = nil
	return e.previous, e.content
}

// Prepend adds a new item to the beginning of the list and returns the new head.
func (e *DoubleLinkedItem[T]) Prepend(value T) *DoubleLinkedItem[T] {
	e.previous = &DoubleLinkedItem[T]{
		content:  value,
		previous: nil,
		next:     e,
	}

	return e.previous
}
