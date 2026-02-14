package bloomhashes

import (
	"hash/fnv"
)

// Fnv1_64 computes a FNV-1 64-bit hash.
// It returns a 64-bit hash value suitable for use in bloom filters.
func Fnv1_64(data []byte) uint64 {
	hasher := fnv.New64()
	hasher.Reset()
	_, _ = hasher.Write(data)

	return hasher.Sum64()
}

// Fnv1_64a computes a FNV-1a 64-bit hash.
// FNV-1a is a variant of FNV-1 with improved avalanche properties.
func Fnv1_64a(data []byte) uint64 {
	hasher := fnv.New64a()
	hasher.Reset()
	_, _ = hasher.Write(data)

	return hasher.Sum64()
}

// Fnv1_128 computes a FNV-1 128-bit hash and converts it to a uint64.
// It returns the first 8 bytes of the 128-bit hash as a 64-bit value.
func Fnv1_128(data []byte) uint64 {
	hasher := fnv.New128()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum(nil)

	return bytesToUint64(sum)
}

// Fnv1_128a computes a FNV-1a 128-bit hash and converts it to a uint64.
// FNV-1a is a variant of FNV-1 with improved avalanche properties.
func Fnv1_128a(data []byte) uint64 {
	hasher := fnv.New128a()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum(nil)

	return bytesToUint64(sum)
}
