package locking

var _ Lock = &NopLock{}

type NopLock struct{}

func NewNopLock() *NopLock {
	return &NopLock{}
}

func (*NopLock) Lock()    {}
func (*NopLock) Unlock()  {}
func (*NopLock) RLock()   {}
func (*NopLock) RUnlock() {}
