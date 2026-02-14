package bloomfilters

type GenericBloomFilter[T any] struct {
	base     IBloomFilter
	getBytes func(T) []byte
}

func NewGenericBloomFilter[T any](base IBloomFilter, getBytes func(T) []byte) *GenericBloomFilter[T] {
	return &GenericBloomFilter[T]{
		base:     base,
		getBytes: getBytes,
	}
}

// Add implements [IBloomFilter].
func (g *GenericBloomFilter[T]) Add(data T) {
	d := g.getBytes(data)
	g.base.Add(d)
}

// Test implements [IBloomFilter].
func (g *GenericBloomFilter[T]) Test(data T) bool {
	d := g.getBytes(data)

	return g.base.Test(d)
}

// Bits returns a copy of the Bits struct representing the bit array of the bloom filter.
// Modifying the returned Bits will not affect the internal state of the bloom filter.
func (bf *GenericBloomFilter[T]) Bits() Bits {
	return bf.base.Bits()
}

// GetHash implements [IBloomFilter].
func (g *GenericBloomFilter[T]) GetHash(hash uint64) bool {
	return g.base.GetHash(hash)
}

// SetHash implements [IBloomFilter].
func (g *GenericBloomFilter[T]) SetHash(hash uint64) {
	g.base.SetHash(hash)
}

func (bf *GenericBloomFilter[T]) BitsCount() uint64 {
	return bf.base.BitsCount()
}
