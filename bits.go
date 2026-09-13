package bloomfilters

import (
	"math/bits"
	"slices"
)

// Bits is a structure that represents a bit array for use in bloom filters. It provides methods to set and get bits, as well as to marshal and unmarshal the data for storage or transmission.
type Bits struct {
	data []uint64
}

// NewBits creates a new Bits structure with the specified size in bits.
// It calculates the necessary number of uint64 words to accommodate the specified size.
func NewBits(size uint64) Bits {
	words := (size + 63) / 64

	return Bits{
		data: make([]uint64, max(words, 1)), // Calculate the number of uint64 needed to store 'size' bits
	}
}

// Size returns the total number of bits this storage can hold.
func (b *Bits) Size() uint64 {
	return uint64(len(b.data)) * 64
}

// SetHash sets the bit at the index corresponding to the given hash value to 1.
func (b *Bits) SetHash(hash uint64) {
	// Lemire's "fastrange"
	index, _ := bits.Mul64(hash, b.Size())

	b.Setbit(index)
}

// GetHash checks if the bit at the index corresponding to the given hash value is set to 1.
func (b *Bits) GetHash(hash uint64) bool {
	// Lemire's "fastrange"
	index, _ := bits.Mul64(hash, b.Size())

	return b.Getbit(index)
}

// Setbit sets the bit at the specified index to 1. If the index is out of bounds (greater than or equal to the size of the Bits), it will not set any bit.
func (b *Bits) Setbit(index uint64) {
	word, bit := b.calcaluteIndex(index)

	// Index is out of bounds?, do not set any bit
	bitmask := uint64(1 << bit)
	if word < uint64(len(b.data)) {
		b.data[word] |= bitmask
	}
}

// Getbit returns the value of the bit at the specified index (true if set, false otherwise).
func (b *Bits) Getbit(index uint64) (result bool) {
	word, bit := b.calcaluteIndex(index)

	bitmask := uint64(1 << bit)
	if word < uint64(len(b.data)) {
		result = (b.data[word] & bitmask) != 0
	}

	return
}

// Equals returns true if this Bits structure is equal to the other Bits structure.
func (b *Bits) Equals(other *Bits) bool {
	return slices.Equal(b.data, other.data)
}

// Words returns a copy slice of uint64 words representing the bits in the bloom filter.
func (b *Bits) Words() []uint64 {
	return slices.Clone(b.data)
}

// Copy returns a deep copy of the Bits structure.
func (b *Bits) Copy() Bits {
	return Bits{
		data: slices.Clone(b.data),
	}
}

// BitsCount returns the total number of bits that are set to 1 in the bloom filter.
func (b *Bits) BitsCount() uint64 {
	var count uint64
	for _, word := range b.data {
		count += uint64(bits.OnesCount64(word))
	}

	return count
}

// calcaluteIndex calculates the word and bit position for a given index in the bloom filter.
func (b *Bits) calcaluteIndex(index uint64) (word, bit uint64) {
	word = index / 64
	bit = index % 64

	return
}
