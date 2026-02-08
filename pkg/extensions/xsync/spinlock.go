package xsync

import (
	"sync/atomic"
)

// SpinLock is a simple spin lock using atomic operations.
// It avoids OS scheduler overhead for very short critical sections.
type SpinLock struct {
	locked atomic.Uint32
}

func NewSpinLock() *SpinLock {
	return &SpinLock{}
}

// Lock acquires the spin lock, spinning until successful.
func (s *SpinLock) Lock() {
	for !s.locked.CompareAndSwap(0, 1) {
		// runtime.Gosched() // yield to avoid starving other goroutines
	}
}

// Unlock releases the spin lock.
func (s *SpinLock) Unlock() {
	s.locked.Store(0)
}
