package bloomfilters

import (
	"encoding"
	"encoding/binary"
	"slices"
)

var (
	_ encoding.BinaryMarshaler   = (*Bits)(nil)
	_ encoding.BinaryUnmarshaler = (*Bits)(nil)
)

// Bits is a structure that represents a bit array for use in bloom filters. It provides methods to set and get bits, as well as to marshal and unmarshal the data for storage or transmission.
type Bits struct {
	data []uint64
}

func NewBits(size uint64) Bits {
	bits := (size + 63) / 64

	return Bits{
		data: make([]uint64, max(bits, 1)), // Calculate the number of uint64 needed to store 'size' bits
	}
}

// Size returns the total number of bits this storage can hold.
func (b *Bits) Size() uint64 {
	return uint64(len(b.data)) * 64
}

// Setbit sets the bit at the specified index to 1. If the index is out of bounds (greater than or equal to the size of the Bits), it will not set any bit.
func (b *Bits) Setbit(index uint64) {
	word, bit := b.calcaluteIndex(index)

	if word >= uint64(len(b.data)) {
		return // Index is out of bounds, do not set any bit
	}
	b.data[word] |= 1 << bit
}

func (b *Bits) Getbit(index uint64) bool {
	word, bit := b.calcaluteIndex(index)

	return (b.data[word] & (1 << bit)) != 0
}

// UnmarshalBinary implements [encoding.BinaryUnmarshaler].
func (b *Bits) UnmarshalBinary(data []byte) error {
	neededcap := (len(data) + 7) / 8
	if b.data == nil {
		b.data = make([]uint64, 0, neededcap)
	} else {
		// Grow and reset the existing slice to accommodate the new data
		b.data = slices.Grow(b.data, neededcap)[0:0]
	}

	for i := 0; i < len(data); i += 8 {
		word := binary.LittleEndian.Uint64(data[i:])
		b.data = append(b.data, word)
	}

	return nil
}

// MarshalBinary implements [encoding.BinaryMarshaler].
func (b *Bits) MarshalBinary() (data []byte, err error) {
	data = make([]byte, len(b.data)*8)
	for i, word := range b.data {
		binary.LittleEndian.PutUint64(data[i*8:], word)
	}

	return data, nil
}

// Words returns the underlying slice of uint64 words representing the bits in the bloom filter.
// WARNING: Modifying the returned slice will affect the internal state of the Bits struct. Use with caution.
func (b *Bits) Words() []uint64 {
	return b.data
}

func (b *Bits) Equals(other *Bits) bool {
	return slices.Equal(b.data, other.data)
}

// calcaluteIndex calculates the word and bit position for a given index in the bloom filter.
func (b *Bits) calcaluteIndex(index uint64) (word, bit uint64) {
	word = index / 64
	bit = index % 64

	return
}
