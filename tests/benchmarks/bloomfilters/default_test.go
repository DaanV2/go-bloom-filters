package bloomfilters_test

import (
	"testing"

	bloomfilters "github.com/daanv2/go-bloom-filters"
	"github.com/daanv2/go-bloom-filters/tests/testutil"
	"github.com/stretchr/testify/require"
)

func Benchmark_Default(b *testing.B) {
	const bf_size = 100_000
	const arrays = 1000
	const array_length = 32

	data := testutil.MoreBytes(arrays, array_length)

	b.Run("Add", func(b *testing.B) {
		bg, err := bloomfilters.NewBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			for i := range data {
				bg.Add(data[i])
			}
		}
	})

	b.Run("Add_Test", func(b *testing.B) {
		bg, err := bloomfilters.NewBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			for i := range data {
				bg.Add(data[i])
			}
			for i := range data {
				v := bg.Test(data[i])
				if !v {
					b.Fatalf("expected to find %v", data[i])
				}
			}
		}
	})

	b.Run("Test_Nothing", func(b *testing.B) {
		bg, err := bloomfilters.NewBloomFilter(
			bloomfilters.WithDefaultHashFunctions(),
			bloomfilters.WithSize(bf_size),
		)
		require.NoError(b, err)

		for b.Loop() {
			for i := range data {
				v := bg.Test(data[i])
				if v {
					b.Fatalf("found %v", data[i])
				}
			}
		}
	})
}
