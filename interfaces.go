package bloomfilters

type IBloomFilter interface {
	SetHash(hash uint64)
	GetHash(hash uint64) bool
	Add(data []byte)
	Test(data []byte) bool
	BitsCount() uint64
	Bits() Bits
}
