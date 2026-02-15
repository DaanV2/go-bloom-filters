package bloomfilters

// IBloomFilter defines the interface for bloom filter implementations.
// It provides methods for adding elements, testing membership, and managing hash values.
type IBloomFilter interface {
	SetHash(hash uint64)      // SetHash sets the bit corresponding to the given hash value in the bloom filter.
	GetHash(hash uint64) bool // GetHash checks if the bit corresponding to the given hash value is set in the bloom filter.
	Add(data []byte)          // Add adds the given data to the bloom filter by applying the hash functions and setting the corresponding bits.
	Test(data []byte) bool    // Test checks if the given data is likely to be in the bloom filter by applying the hash functions and checking the corresponding bits.
	BitsCount() uint64        // BitsCount returns the total number of bits that are set to 1 in the bloom filter.
	Bits() Bits               // Bits returns the underlying bit array of the bloom filter.
}
