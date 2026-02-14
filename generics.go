package bloomfilters

// GenericBloomFilter is a wrapper around IBloomFilter that allows using custom types with a serializer function.
// It converts custom types to byte slices before passing them to the underlying bloom filter.
type GenericBloomFilter[T any] struct {
	base     IBloomFilter
	getBytes func(T) []byte
}

// NewGenericBloomFilter creates a new generic bloom filter that wraps the provided bloom filter and uses the given serializer function.
// The serializer function converts custom types to byte slices for hashing.
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

// BitsCount returns the total number of bits that are set to 1 in the bloom filter.
func (bf *GenericBloomFilter[T]) BitsCount() uint64 {
	return bf.base.BitsCount()
}
