package locking

type Lock interface {
	Lock()
	RLock()
	Unlock()
}
