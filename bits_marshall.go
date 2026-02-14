package bloomfilters

import (
	"encoding"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"slices"
)

var (
	_ encoding.BinaryMarshaler = (*Bits)(nil)
	_ encoding.BinaryMarshaler = (*Bits)(nil)
	_ encoding.TextUnmarshaler = (*Bits)(nil)
	_ encoding.TextMarshaler   = (*Bits)(nil)
	_ json.Marshaler           = (*Bits)(nil)
	_ json.Unmarshaler         = (*Bits)(nil)
)

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

// MarshalText implements [encoding.TextMarshaler].
func (b *Bits) MarshalText() (text []byte, err error) {
	src, err := b.MarshalBinary()
	if err != nil {
		return nil, err
	}

	dst := make([]byte, base64.RawStdEncoding.EncodedLen(len(src)))
	base64.RawStdEncoding.Encode(dst, src)

	return dst, nil
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (b *Bits) UnmarshalText(text []byte) error {
	src := make([]byte, base64.RawStdEncoding.DecodedLen(len(text)))
	n, err := base64.RawStdEncoding.Decode(src, text)
	if err != nil {
		return err
	}

	return b.UnmarshalBinary(src[:n])
}

// MarshalJSON implements [json.Marshaler].
func (b *Bits) MarshalJSON() ([]byte, error) {
	return b.MarshalText()
}

// UnmarshalJSON implements [json.Unmarshaler].
func (b *Bits) UnmarshalJSON(d []byte) error {
	return b.UnmarshalText(d)
}
