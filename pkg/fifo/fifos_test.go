package fifo_test

import (
	"testing"

	"github.com/Dirk007/simpleFifo/pkg/fifo"
	"github.com/Dirk007/simpleFifo/pkg/fifo/implementations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func requireSuccessfulValuePop[T any](t *testing.T, f *fifo.Fifo[T], expected T) {
	t.Helper()

	value, err := f.Next()
	require.NoError(t, err)
	require.NotNil(t, value)

	assert.Equal(t, expected, value)
}

func testFifoPrincipleWorks(t *testing.T, f *fifo.Fifo[int]) {
	f.Add(1)
	f.Add(2)
	f.Add(3, 4, 5)

	requireSuccessfulValuePop(t, f, 1)
	requireSuccessfulValuePop(t, f, 2)
	requireSuccessfulValuePop(t, f, 3)
	requireSuccessfulValuePop(t, f, 4)
	requireSuccessfulValuePop(t, f, 5)
}

func testFifoPrincipleWithoutLocking(t *testing.T, f *fifo.Fifo[int]) {
	f = f.WithoutLocking()

	f.Add(1)
	f.Add(2)
	f.Add(3, 4, 5)

	requireSuccessfulValuePop(t, f, 1)
	requireSuccessfulValuePop(t, f, 2)
	requireSuccessfulValuePop(t, f, 3)
	requireSuccessfulValuePop(t, f, 4)
	requireSuccessfulValuePop(t, f, 5)
}

func testInitialEmpty(t *testing.T, f *fifo.Fifo[int]) {
	_, err := f.Next()
	assert.ErrorIs(t, err, fifo.ErrEmptyFifo)
}

func testDetectsEmpty(t *testing.T, f *fifo.Fifo[int]) {
	f.Add(1)
	requireSuccessfulValuePop(t, f, 1)

	_, err := f.Next()
	assert.ErrorIs(t, err, fifo.ErrEmptyFifo)
}

func testRespectsLimit(t *testing.T, f *fifo.Fifo[int]) {
	f = f.Clear().WithLimit(2)

	_, err := f.Add(1)
	require.NoError(t, err)
	_, err = f.Add(2)
	require.NoError(t, err)

	_, err = f.Add(3)
	require.Error(t, err)
	require.ErrorIs(t, err, fifo.NewFifoLimitReachedError(
		int64(2),
	))
}

func testEmptyRestart(t *testing.T, f *fifo.Fifo[int]) {
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

func testClear(t *testing.T, f *fifo.Fifo[int]) {
	_, err := f.Add(1)
	require.NoError(t, err)
	f.Clear()

	assert.Equal(t, 0, f.Count())
}

func TestAllFifos(t *testing.T) {
	var fifoImpls = []struct {
		name    string
		newImpl func() implementations.FifoStrategy[int]
	}{
		{
			name:    "Double Linked List Fifo",
			newImpl: implementations.NewDoubleLinkedFifo[int],
		},
		{
			name:    "Sliced Fifo",
			newImpl: implementations.NewSliceFifo[int],
		},
	}

	var tests = []struct {
		name string
		fn   func(t *testing.T, f *fifo.Fifo[int])
	}{
		{name: "Fifo Principle Works", fn: testFifoPrincipleWorks},
		{name: "Fifo Principle Without Locking", fn: testFifoPrincipleWithoutLocking},
		{name: "Initial Empty", fn: testInitialEmpty},
		{name: "Detects Empty", fn: testDetectsEmpty},
		{name: "Respects Limit", fn: testRespectsLimit},
		{name: "Empty Restart", fn: testEmptyRestart},
		{name: "Clear really clears", fn: testClear},
	}

	for _, impl := range fifoImpls {
		t.Run(impl.name, func(t *testing.T) {
			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					fifoImpl := impl.newImpl()
					testFifo := fifo.NewFifo[int]().WithImplementation(fifoImpl)
					test.fn(t, testFifo)
				})
			}
		})
	}
}
