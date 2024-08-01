package fifo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func requireSuccessfulValuePop[T any](t *testing.T, f *Fifo[T], expected T) {
	t.Helper()

	value, err := f.Next()
	require.NoError(t, err)
	require.NotNil(t, value)

	assert.Equal(t, expected, value)
}

func Test_FifoPrinciple(t *testing.T) {
	f := NewFifo[int]()

	f.Add(1)
	f.Add(2)
	f.Add(3, 4, 5)

	requireSuccessfulValuePop(t, f, 1)
	requireSuccessfulValuePop(t, f, 2)
	requireSuccessfulValuePop(t, f, 3)
	requireSuccessfulValuePop(t, f, 4)
	requireSuccessfulValuePop(t, f, 5)
}

func Test_FifoPrinciple_WithoutLocking(t *testing.T) {
	f := NewFifo[int]().WithoutLocking()

	f.Add(1)
	f.Add(2)
	f.Add(3, 4, 5)

	requireSuccessfulValuePop(t, f, 1)
	requireSuccessfulValuePop(t, f, 2)
	requireSuccessfulValuePop(t, f, 3)
	requireSuccessfulValuePop(t, f, 4)
	requireSuccessfulValuePop(t, f, 5)
}

func Test_InitialEmpty(t *testing.T) {
	f := NewFifo[int]()
	_, err := f.Next()
	assert.ErrorIs(t, err, ErrEmptyFifo)
}

func Test_DetectsEmpty(t *testing.T) {
	f := NewFifo[int]()
	f.Add(1)
	requireSuccessfulValuePop(t, f, 1)

	_, err := f.Next()
	assert.ErrorIs(t, err, ErrEmptyFifo)
}

func Test_RespectsLimit(t *testing.T) {
	f := NewFifo[int]().WithLimit(2)

	_, err := f.Add(1)
	require.NoError(t, err)
	_, err = f.Add(2)
	require.NoError(t, err)

	_, err = f.Add(3)
	require.ErrorIs(t, err, NewFifoLimitReachedError(
		int64(2),
	))
}

func Test_EmptyRestart(t *testing.T) {
	f := NewFifo[int]().WithLimit(2)

	_, err := f.Add(1)
	require.NoError(t, err)
	value, err := f.Next()
	require.NoError(t, err)
	assert.Equal(t, 1, value)

	_, err = f.Add(2)
	require.NoError(t, err)
	value, err = f.Next()
	require.NoError(t, err)
	assert.Equal(t, 2, value)
}
