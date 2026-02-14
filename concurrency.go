package bloomfilters

import (
	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/daanv2/go-bloom-filters/pkg/extensions/xsync"
)

var _ IBloomFilter = &ConcurrentBloomFilter{}

type ConcurrentBloomFilter struct {
	bits   Bits
	hashes []bloomhashes.HashFunction
	lock   xsync.SpinLock
}

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

func (bf *ConcurrentBloomFilter) GetHash(hash uint64) bool {
	return bf.getHashes(bf.index(hash))
}

func (bf *ConcurrentBloomFilter) SetHash(hash uint64) {
	bf.setHashes(bf.index(hash))
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
