package bloomhashes

import (
	"hash"
)

// HashFunction defines the type for hash functions used in the bloom filter.
// It takes a byte slice as input and adds the resulting uint64 hashes to the provided slice of uint64 values, starting at the beginning of the slice.
// The function returns the number of hashes added to the slice.
// Disregarding any bytes that fall outside of the length of the provided slice of uint64 values.
type HashFunction func(data []byte) uint64

// WrapFunction wraps a function that returns a byte slice into a HashFunction.
// It converts the byte slice output to a uint64 hash value.
func WrapFunction(f func(data []byte) []byte) HashFunction {
	return func(data []byte) uint64 {
		b := f(data)

		return bytesToUint64(b)
	}
}

// WrapHasher64 wraps a hash.Hash64 factory function into a HashFunction.
// It creates a new hasher, writes the data, and returns the 64-bit hash sum.
func WrapHasher64(h func() hash.Hash64) HashFunction {
	return func(data []byte) uint64 {
		hasher := h()
		_, _ = hasher.Write(data)

		return hasher.Sum64()
	}
}

// WrapHasher wraps a hash.Hash factory function into a HashFunction.
// It creates a new hasher, writes the data, computes the hash sum, and converts it to a uint64.
func WrapHasher(h func() hash.Hash) HashFunction {
	return func(data []byte) uint64 {
		hasher := h()
		_, _ = hasher.Write(data)
		sum := hasher.Sum(nil)

		return bytesToUint64(sum)
	}
}
