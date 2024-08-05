package locking

type Lock interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}
