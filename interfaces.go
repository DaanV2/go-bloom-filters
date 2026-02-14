package bloomfilters

// IBloomFilter defines the interface for bloom filter implementations.
// It provides methods for adding elements, testing membership, and managing hash values.
type IBloomFilter interface {
	SetHash(hash uint64)
	GetHash(hash uint64) bool
	Add(data []byte)
	Test(data []byte) bool
	BitsCount() uint64
	Bits() Bits
}
