package bloomhashes

import "hash"

// HashFunction defines the type for hash functions used in the bloom filter.
// It takes a byte slice as input and adds the resulting uint64 hashes to the provided slice of uint64 values, starting at the beginning of the slice.
// The function returns the number of hashes added to the slice.
// Disregarding any bytes that fall outside of the length of the provided slice of uint64 values.
type HashFunction func(data []byte, hashes []uint64) int

// DefaultHashFunctions returns a slice of default hash functions to be used in the bloom filter.
// Currently, it returns an empty slice, indicating that no default hash functions are provided.
func WrapFunction(f func(data []byte) []byte) HashFunction {
	return func(data []byte, hashes []uint64) int {
		b := f(data)
		if len(b) < 8 {
			return 0
		}

		n := min(len(b)/uint64_size, len(hashes))
		PutUint64(b[:n*uint64_size], hashes)

		return n
	}
}

func WrapHasher64(h func() hash.Hash64) HashFunction {
	return func(data []byte, hashes []uint64) int {
		hasher := h()
		hasher.Reset()
		_, _ = hasher.Write(data)
		sum := hasher.Sum64()
		if len(hashes) > 0 {
			hashes[0] = sum

			return 1
		}

		return 0
	}
}

func WrapHasher(h func() hash.Hash) HashFunction {
	return func(data []byte, hashes []uint64) int {
		hasher := h()
		hasher.Reset()
		_, _ = hasher.Write(data)
		sum := hasher.Sum(nil)

		return PutUint64(sum, hashes)
	}
}
