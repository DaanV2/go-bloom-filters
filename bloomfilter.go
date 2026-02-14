package bloomfilters

import (
	"errors"

	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

var (
	ErrInvalidSize          = errors.New("invalid size")
	ErrRequiredHashFunction = errors.New("at least one hash function is required")
	ErrHashIsNil            = errors.New("hash functions cannot be nil")
)

var _ IBloomFilter = &BloomFilter{}

// BloomFilter is a probabilistic data structure that tests whether an element is a member of a set.
// False positive matches are possible, but false negatives are not.
type BloomFilter struct {
	bits   Bits
	hashes []bloomhashes.HashFunction
}

// NewBloomFilter creates a new bloom filter with the given options.
// It returns an error if the size of the bloom filter is invalid or if any of the hash functions are nil.
func NewBloomFilter(opts ...BloomFilterOptions) (*BloomFilter, error) {
	bf := &BloomFilter{}
	for _, opt := range opts {
		opt.applyBF(bf)
	}
	if bf.bits.Size() == 0 {
		return nil, ErrInvalidSize
	}
	if len(bf.hashes) == 0 {
		return nil, ErrRequiredHashFunction
	}
	for _, hashFunc := range bf.hashes {
		if hashFunc == nil {
			return nil, ErrHashIsNil
		}
	}

	return bf, nil
}

// Add adds the given data to the bloom filter by applying each hash function to the data and setting the corresponding bits in the filter.
func (bf *BloomFilter) Add(data []byte) {
	indexes := make([]uint64, 0, len(bf.hashes))

	for _, hashFunc := range bf.hashes {
		if hashFunc == nil {
			continue
		}
		indexes = append(indexes, bf.index(hashFunc(data)))
	}

	for _, hash := range indexes {
		bf.SetHash(hash)
	}
}

// Test checks if the given data is likely to be in the bloom filter by applying each hash function to the data and checking if the corresponding bits in the filter are set. It returns true if all bits are set, indicating that the data is likely to be in the filter, and false otherwise.
func (bf *BloomFilter) Test(data []byte) bool {
	for _, hashFunc := range bf.hashes {
		if hashFunc == nil {
			continue
		}
		hash := hashFunc(data)
		if !bf.GetHash(hash) {
			return false
		}
	}

	return true
}

// Set sets the bit at the index corresponding to the given hash value to 1.
func (bf *BloomFilter) SetHash(hash uint64) {
	bf.bits.Setbit(bf.index(hash))
}

// Get checks if the bit at the index corresponding to the given hash value is set to 1.
func (bf *BloomFilter) GetHash(hash uint64) bool {
	return bf.bits.Getbit(bf.index(hash))
}

// BitsCount returns the total number of bits that are set to 1 in the bloom filter.
func (bf *BloomFilter) BitsCount() uint64 {
	return bf.bits.BitsCount()
}

// Words returns a copy slice of uint64 words representing the bits in the bloom filter.
func (bf *BloomFilter) Words() []uint64 {
	return bf.bits.Words()
}

// Bits returns a copy of the Bits struct representing the bit array of the bloom filter.
// Modifying the returned Bits will not affect the internal state of the bloom filter.
func (bf *BloomFilter) Bits() Bits {
	return bf.bits.Copy()
}

func (bf *BloomFilter) index(hash uint64) uint64 {
	return hash % bf.bits.Size()
}
