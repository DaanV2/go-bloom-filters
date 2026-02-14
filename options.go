package bloomfilters

import (
	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

type BloomFilterOptions interface {
	applyBF(*BloomFilter)
	applyCBF(*ConcurrentBloomFilter)
}

type withSize struct {
	size uint64
}

func (w withSize) applyBF(bf *BloomFilter)            { bf.bits = NewBits(w.size) }
func (w withSize) applyCBF(bf *ConcurrentBloomFilter) { bf.bits = NewBits(w.size) }

// WithSize sets the size of the bloom filter in bits.
// It calculates the necessary number of uint64 words to accommodate the specified size and initializes the Bits structure accordingly.
func WithSize(size uint64) BloomFilterOptions {
	return withSize{size: size}
}

type withHashFunctions struct {
	hashFunctions []bloomhashes.HashFunction
}

func (w withHashFunctions) applyBF(bf *BloomFilter)            { bf.hashes = w.hashFunctions }
func (w withHashFunctions) applyCBF(bf *ConcurrentBloomFilter) { bf.hashes = w.hashFunctions }

// WithDefaultHashFunctions sets the default hash functions to be used by the bloom filter.
// It retrieves the default hash functions from the bloomhashes package and assigns them to the bloom filter's hashes field.
func WithDefaultHashFunctions() BloomFilterOptions {
	return WithHashFunctions(bloomhashes.DefaultHashFunctions())
}

// WithAllHashFunctions sets all available hash functions to be used by the bloom filter.
// It retrieves all hash functions from the bloomhashes package and assigns them to the bloom filter's hashes field.
func WithAllHashFunctions() BloomFilterOptions {
	return WithHashFunctions(bloomhashes.AllHashFunctions())
}

// WithHashFunctions sets the hash functions to be used by the bloom filter. It replaces any existing hash functions with the provided slice of HashFunction.
func WithHashFunctions(hashFunctions []bloomhashes.HashFunction) BloomFilterOptions {
	return withHashFunctions{
		hashFunctions: hashFunctions,
	}
}

type withAppendHashFunctions struct {
	hashFunctions []bloomhashes.HashFunction
}

func (w withAppendHashFunctions) applyBF(bf *BloomFilter) {
	bf.hashes = append(bf.hashes, w.hashFunctions...)
}
func (w withAppendHashFunctions) applyCBF(bf *ConcurrentBloomFilter) {
	bf.hashes = append(bf.hashes, w.hashFunctions...)
}

// WithAppendHashFunctions appends additional hash functions to the existing list of hash functions used by the bloom filter. It takes a slice of HashFunction and appends it to the current list of hash functions.
func WithAppendHashFunctions(hashFunctions []bloomhashes.HashFunction) BloomFilterOptions {
	return withAppendHashFunctions{
		hashFunctions: hashFunctions,
	}
}

type withWords struct {
	words []uint64
}

func (w withWords) applyBF(bf *BloomFilter)            { bf.bits = Bits{data: w.words} }
func (w withWords) applyCBF(bf *ConcurrentBloomFilter) { bf.bits = Bits{data: w.words} }

// WithWords sets the bits of the bloom filter using a slice of uint64 words. It initializes the Bits structure with the provided words, allowing for direct manipulation of the bloom filter's bit array.
func WithWords(words []uint64) BloomFilterOptions {
	return withWords{
		words: words,
	}
}

type withBits struct {
	bits Bits
}

func (w withBits) applyBF(bf *BloomFilter)            { bf.bits = w.bits }
func (w withBits) applyCBF(bf *ConcurrentBloomFilter) { bf.bits = w.bits }

// WithBits sets the bits of the bloom filter using a Bits structure. It directly assigns the provided Bits to the bloom filter, allowing for more flexible manipulation of the bit array.
func WithBits(bits Bits) BloomFilterOptions {
	return withBits{
		bits: bits,
	}
}
