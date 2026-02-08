package bloomfilters

import (
	"errors"

	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
)

var (
	ErrInvalidSize = errors.New("invalid size")
)

type BloomFilter struct {
	bits   Bits
	hashes []bloomhashes.HashFunction
}

func NewBloomFilter(opts ...BloomFilterOptions) (*BloomFilter, error) {
	r := &BloomFilter{}
	for _, opt := range opts {
		opt(r)
	}
	if r.bits.Size() == 0 {
		return nil, ErrInvalidSize
	}
	if len(r.hashes) == 0 {
		return nil, errors.New("at least one hash function is required")
	}
	for _, hashFunc := range r.hashes {
		if hashFunc == nil {
			return nil, errors.New("hash functions cannot be nil")
		}
	}

	return r, nil
}

// Add adds the given data to the bloom filter by applying each hash function to the data and setting the corresponding bits in the filter.
func (bf *BloomFilter) Add(data []byte) {
	var buf [bloomhashes.MAX_HASHES]uint64

	for _, hashFunc := range bf.hashes {
		if hashFunc == nil {
			continue
		}
		n := hashFunc(data, buf[:])
		for _, hash := range buf[:n] {
			bf.Set(hash)
		}
	}
}

// Test checks if the given data is likely to be in the bloom filter by applying each hash function to the data and checking if the corresponding bits in the filter are set. It returns true if all bits are set, indicating that the data is likely to be in the filter, and false otherwise.
func (bf *BloomFilter) Test(data []byte) bool {
	var buf [bloomhashes.MAX_HASHES]uint64

	for _, hashFunc := range bf.hashes {
		if hashFunc == nil {
			continue
		}
		n := hashFunc(data, buf[:])
		for _, hash := range buf[:n] {
			if !bf.Get(hash) {
				return false
			}
		}
	}

	return true
}

// Set sets the bit at the index corresponding to the given hash value to 1.
func (bf *BloomFilter) Set(hash uint64) {
	bf.bits.Setbit(bf.index(hash))
}

// Get checks if the bit at the index corresponding to the given hash value is set to 1.
func (bf *BloomFilter) Get(hash uint64) bool {
	return bf.bits.Getbit(bf.index(hash))
}

func (bf *BloomFilter) index(hash uint64) uint64 {
	return hash % bf.bits.Size()
}
