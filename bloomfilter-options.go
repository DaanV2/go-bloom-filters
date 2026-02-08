package bloomfilters

import "github.com/daanv2/go-bloom-filters/pkg/bloomhashes"

type BloomFilterOptions func(*BloomFilter)

// WithSize sets the size of the bloom filter in bits.
// It calculates the necessary number of uint64 words to accommodate the specified size and initializes the Bits structure accordingly.
func WithSize(size uint64) BloomFilterOptions {
	return func(bf *BloomFilter) {
		bf.bits = NewBits(size)
	}
}

// WithHashFunctions sets the hash functions to be used by the bloom filter. It replaces any existing hash functions with the provided slice of HashFunction.
func WithHashFunctions(hashFunctions []bloomhashes.HashFunction) BloomFilterOptions {
	return func(bf *BloomFilter) {
		bf.hashes = hashFunctions
	}
}

// WithAppendHashFunctions appends additional hash functions to the existing list of hash functions used by the bloom filter. It takes a slice of HashFunction and appends it to the current list of hash functions.
func WithAppendHashFunctions(hashFunctions []bloomhashes.HashFunction) BloomFilterOptions {
	return func(bf *BloomFilter) {
		bf.hashes = append(bf.hashes, hashFunctions...)
	}
}
