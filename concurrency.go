package bloomfilters

import (
	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/daanv2/go-bloom-filters/pkg/extensions/xsync"
)

var _ IBloomFilter = &ConcurrentBloomFilter{}

// ConcurrentBloomFilter is a thread-safe bloom filter that uses a spinlock for concurrent access.
// It is safe to call Add and Test methods from multiple goroutines.
type ConcurrentBloomFilter struct {
	bits   Bits
	hashes []bloomhashes.HashFunction
	lock   xsync.SpinLock
}

// NewConcurrentBloomFilter creates a new concurrent bloom filter with the given options.
// It returns an error if the size of the bloom filter is invalid or if any of the hash functions are nil.
func NewConcurrentBloomFilter(opts ...BloomFilterOptions) (*ConcurrentBloomFilter, error) {
	bf := &ConcurrentBloomFilter{
		lock: *xsync.NewSpinLock(),
	}

	for _, opt := range opts {
		opt.applyCBF(bf)
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
// This method is thread-safe.
func (bf *ConcurrentBloomFilter) Add(data []byte) {
	indexes := make([]uint64, 0, len(bf.hashes))

	for _, hashFunc := range bf.hashes {
		if hashFunc == nil {
			continue
		}
		hash := hashFunc(data)
		indexes = append(indexes, bf.index(hash))
	}

	bf.setHashes(indexes...)
}

// Test checks if the given data is likely to be in the bloom filter by applying each hash function to the data and checking if the corresponding bits in the filter are set.
// This method is thread-safe. It returns true if all bits are set, indicating that the data is likely to be in the filter, and false otherwise.
func (bf *ConcurrentBloomFilter) Test(data []byte) bool {
	indexes := make([]uint64, 0, len(bf.hashes))

	// Spent more time on hashing, so we don't have to lock for each bit access.
	for _, hashFunc := range bf.hashes {
		if hashFunc == nil {
			continue
		}
		hash := hashFunc(data)
		indexes = append(indexes, bf.index(hash))
	}

	return bf.getHashes(indexes...)
}

// GetHash checks if the bit at the index corresponding to the given hash value is set to 1.
// This method is thread-safe.
func (bf *ConcurrentBloomFilter) GetHash(hash uint64) bool {
	return bf.getHashes(bf.index(hash))
}

// SetHash sets the bit at the index corresponding to the given hash value to 1.
// This method is thread-safe.
func (bf *ConcurrentBloomFilter) SetHash(hash uint64) {
	bf.setHashes(bf.index(hash))
}

// BitsCount returns the total number of bits that are set to 1 in the bloom filter.
func (bf *ConcurrentBloomFilter) BitsCount() uint64 {
	return bf.bits.BitsCount()
}

// Bits returns a copy of the Bits struct representing the bit array of the bloom filter.
// Modifying the returned Bits will not affect the internal state of the bloom filter.
func (bf *ConcurrentBloomFilter) Bits() Bits {
	return bf.bits.Copy()
}

func (bf *ConcurrentBloomFilter) getHashes(indexes ...uint64) bool {
	bf.lock.Lock()
	defer bf.lock.Unlock()

	for _, index := range indexes {
		if !bf.bits.Getbit(index) {
			return false
		}
	}

	return true
}

func (bf *ConcurrentBloomFilter) setHashes(indexes ...uint64) {
	bf.lock.Lock()
	defer bf.lock.Unlock()

	for _, index := range indexes {
		bf.bits.Setbit(index)
	}
}

func (bf *ConcurrentBloomFilter) index(hash uint64) uint64 {
	s := bf.bits.Size()
	if s == 0 {
		return 0
	}

	return hash % s
}
