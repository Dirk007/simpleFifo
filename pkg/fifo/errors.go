package fifo

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyFifo       = errors.New("empty fifo")
	_            error = FifoLimitReachedError{}
)

type FifoLimitReachedError struct {
	Limit int64
}

func NewFifoLimitReachedError(limit int64) FifoLimitReachedError {
	return FifoLimitReachedError{Limit: limit}
}

func (e FifoLimitReachedError) Error() string {
	return fmt.Sprintf("fifo limit reached, max elements: %d", e.Limit)
}
