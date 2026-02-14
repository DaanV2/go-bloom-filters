package bloomhashes_test

import (
	"testing"

	"github.com/daanv2/go-bloom-filters/pkg/bloomhashes"
	"github.com/daanv2/go-bloom-filters/tests/testutil"
)

func Benchmark_Hashes_Cost(b *testing.B) {
	type testHash struct {
		name string
		hash func([]byte) uint64
	}

	tests := []testHash{
		{name: "MD5", hash: bloomhashes.MD5},
		{name: "Sha1", hash: bloomhashes.Sha1},
		{name: "Sha224", hash: bloomhashes.Sha224},
		{name: "Sha3_384", hash: bloomhashes.Sha3_384},
		{name: "Sha256", hash: bloomhashes.Sha256},
		{name: "Sha512", hash: bloomhashes.Sha512},
		{name: "Crc64_ISO", hash: bloomhashes.Crc64_ISO},
		{name: "Crc64_ECMA", hash: bloomhashes.Crc64_ECMA},
		{name: "Fnv1_64", hash: bloomhashes.Fnv1_64},
		{name: "Fnv1_64a", hash: bloomhashes.Fnv1_64a},
		{name: "Fnv1_128", hash: bloomhashes.Fnv1_128},
		{name: "Fnv1_128a", hash: bloomhashes.Fnv1_128a},
	}

	d := testutil.MoreBytes(1000, 32)

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for b.Loop() {
				for i := range d {
					v := test.hash(d[i])
					if v == 0 {
						b.Fatalf("unexpected zero value")
					}
				}
			}
		})
	}
}
