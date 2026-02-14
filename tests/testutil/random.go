package testutil

import (
	"math/rand/v2"
)

var seed = [32]byte{
	0x4e, 0xf6, 0xf5, 0x5d,
	0xca, 0xf4, 0x8b, 0x30,
	0x1a, 0x0b, 0x03, 0xe5,
	0x45, 0xe7, 0xef, 0x6f,
	0x67, 0x82, 0xcd, 0x89,
	0x64, 0x10, 0x82, 0x9e,
	0x45, 0x91, 0xa2, 0x73,
	0x32, 0xa5, 0xb7, 0x81,
}

func MoreBytes(arrays, array_length int) [][]byte {
	r := Randomizer()

	bs := make([][]byte, arrays)
	for i := range bs {
		b := make([]byte, array_length)
		for j := range b {
			b[j] = byte(r.IntN(256))
		}

		bs[i] = b
	}

	return bs
}

func Randomizer() *rand.Rand {
	src := rand.NewChaCha8(seed)

	return rand.New(src) // nolint:gosec // This is not used for cryptographic purposes.
}
