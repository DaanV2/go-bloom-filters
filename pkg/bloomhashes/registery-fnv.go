package bloomhashes

import (
	"hash/fnv"
)

func Fnv1_64(data []byte) uint64 {
	hasher := fnv.New64()
	hasher.Reset()
	_, _ = hasher.Write(data)

	return hasher.Sum64()
}

func Fnv1_64a(data []byte) uint64 {
	hasher := fnv.New64a()
	hasher.Reset()
	_, _ = hasher.Write(data)

	return hasher.Sum64()
}

func Fnv1_128(data []byte) uint64 {
	hasher := fnv.New128()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum(nil)

	return bytesToUint64(sum)
}

func Fnv1_128a(data []byte) uint64 {
	hasher := fnv.New128a()
	hasher.Reset()
	_, _ = hasher.Write(data)
	sum := hasher.Sum(nil)

	return bytesToUint64(sum)
}
