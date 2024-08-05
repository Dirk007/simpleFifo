package locking

import "sync"

var _ Lock = &MutexLock{}

type MutexLock struct {
	inner sync.RWMutex
}

func NewMutexLock() *MutexLock {
	return &MutexLock{}
}

func (m *MutexLock) Lock() {
	m.inner.Lock()
}

func (m *MutexLock) Unlock() {
	m.inner.Unlock()
}

func (m *MutexLock) RLock() {
	m.inner.RLock()
}

func (m *MutexLock) RUnlock() {
	m.inner.RUnlock()
}
