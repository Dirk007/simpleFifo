package locking

var _ Lock = &NopLock{}

type NopLock struct{}

func NewNopLock() *NopLock {
	return &NopLock{}
}

func (*NopLock) Lock()   {}
func (*NopLock) RLock()  {}
func (*NopLock) Unlock() {}
