package bloomhashes

import (
	"hash/fnv"
)

const FNV1_64_HASHES = 1 // The number of uint64 values that can be extracted from a FNV-1 hash.

func Fnv1_64(data []byte, hashes []uint64) int {
	hasher := fnv.New64()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum64()
	if len(hashes) > 0 {
		hashes[0] = sum

		return 1
	}

	return 0
}

const FNV1_64a_HASHES = 1 // The number of uint64 values that can be extracted from a FNV-1a hash.

func Fnv1_64a(data []byte, hashes []uint64) int {
	hasher := fnv.New64a()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum64()
	if len(hashes) > 0 {
		hashes[0] = sum

		return 1
	}

	return 0
}

const FNV1_128_HASHES = 2 // The number of uint64 values that can be extracted from a FNV-1 hash.

func Fnv1_128(data []byte, hashes []uint64) int {
	hasher := fnv.New128()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum(nil)
	
	return PutUint64(sum, hashes)
}

const FNV1_128a_HASHES = 2 // The number of uint64 values that can be extracted from a FNV-1a hash.

func Fnv1_128a(data []byte, hashes []uint64) int {
	hasher := fnv.New128a()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum(nil)
	
	return PutUint64(sum, hashes)
}