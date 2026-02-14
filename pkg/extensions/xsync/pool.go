package xsync

import "sync"

// Pool is a generic type-safe wrapper around sync.Pool.
// It provides strongly typed Get and Put methods for a specific type T.
type Pool[T any] struct {
	pool   sync.Pool
	create func() T
}

// NewPool creates a new type-safe pool with the given factory function.
// The factory function is called to create new instances when the pool is empty.
func NewPool[T any](newFunc func() T) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{
			New: func() any {
				return newFunc()
			},
		},
		create: newFunc,
	}
}

// Get retrieves an item from the pool or creates a new one if the pool is empty.
// The returned item is strongly typed as T.
func (p *Pool[T]) Get() T {
	item := p.pool.Get()
	if item == nil {
		return p.create()
	}

	return item.(T)
}

// Put adds an item back to the pool for reuse.
// The item should be in a reusable state before being put back.
func (p *Pool[T]) Put(item T) {
	p.pool.Put(item)
}
